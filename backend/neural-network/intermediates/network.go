package intermediates

import (
	"container/list"
	"fastfoodrestaurant/neuralnetwork/models"
	"fastfoodrestaurant/neuralnetwork/services"
	"fmt"
	"math"
	"os"
	"strings"
	"time"

	"gonum.org/v1/gonum/mat"
	"gonum.org/v1/gonum/stat/distuv"
)

type NetworkIntermediate struct {
	WordService  *services.Word2VecService
	TrainService *services.DataTrainService
	network      *models.Network
}

func NewTextClassificationNetwork(WordService *services.Word2VecService, TrainService *services.DataTrainService) *NetworkIntermediate {
	sygmoidFunction := &models.SygmoidFunction{}
	reluFunction := &models.ReluFunction{}
	network := models.NewNetwork()

	perceptronQuantity := 200

	hiddenLayer := models.NewLayer(services.MaxWords, perceptronQuantity, reluFunction)
	outputLayer := models.NewLayer(perceptronQuantity, services.ClassesQuantity, sygmoidFunction)

	network.AddLayer(hiddenLayer)
	network.AddLayer(outputLayer)

	weightOutputLayer := mat.NewDense(services.ClassesQuantity, perceptronQuantity, randomArray(perceptronQuantity*services.ClassesQuantity, float64(200)))
	weightHiddenLayer := mat.NewDense(perceptronQuantity, services.MaxWords, randomArray(services.MaxWords*perceptronQuantity, float64(200)))

	hiddenLayer.AssignDefaultValues(nil, weightHiddenLayer, reluFunction, 1, 0.50)
	outputLayer.AssignDefaultValues(nil, weightOutputLayer, sygmoidFunction, 1, 0.50)

	return &NetworkIntermediate{
		network:      network,
		WordService:  WordService,
		TrainService: TrainService,
	}
}

func (i *NetworkIntermediate) PredictMessage(text string) *models.Message {
	inputs := i.WordService.GetVector(text)
	predictValues := i.network.Predict(inputs)
	rawValues := predictValues.RawVector().Data
	i.network.ThresholdPredictValues(rawValues)
	classifications := i.TrainService.GetClassesString(rawValues)
	// for index := 0; index < len(rawValues); index++ {
	// 	classifications[index] = i.ClassificateValue(rawValues[index])
	// }

	return &models.Message{
		Text:           "Hemos clasificado tu consulta como: ",
		Classification: classifications,
	}
}

func (i *NetworkIntermediate) Save() {
	layerCount := 1
	weights := i.network.GetWeights()
	for e := weights.Front(); e != nil; e = e.Next() {
		filename := fmt.Sprintf("backup-layer-%d-%s", layerCount, time.Now().String())
		direction := fmt.Sprintf("backups/%s.model", filename)
		file, err := os.Create(direction)

		if err != nil {
			return
		}

		defer file.Close()
		vector := e.Value.(*mat.Dense)
		vector.MarshalBinaryTo(file)
	}
}

func (i *NetworkIntermediate) Load(path string, layerIndex int) error {
	f, err := os.Open(path)
	defer f.Close()
	if err != nil {
		return err
	}

	weights := &mat.Dense{}
	weights.UnmarshalBinaryFrom(f)
	i.network.SetWeights(weights, layerIndex)

	return nil
}

func (n *NetworkIntermediate) Train(dataPath string, epochs int) {
	data := n.TrainService.GetData(dataPath)
	sentences, _ := n.TrainService.SeparateData(data)

	if len(sentences) == 0 {
		fmt.Println("No hay informacion de entrenamiento")
		return
	}

	wordsList := n.TrainService.GetWordsArray(sentences)
	wordsArray := transformListToArray(wordsList)
	n.WordService.Train(strings.Join(wordsArray, " "))

	for e := 0; e < epochs; e++ {
		zeroX := mat.NewDense(len(sentences), services.MaxWords, nil)
		zeroX.Zero()

		zeroY := mat.NewDense(len(sentences), services.ClassesQuantity, nil)
		zeroY.Zero()

		X := *mat.NewDense(len(sentences), services.MaxWords, nil)

		Y := *mat.NewDense(len(sentences), services.ClassesQuantity, nil)

		i := 0
		for sentence, classes := range sentences {
			vectorizedSentence := n.WordService.GetVector(sentence)
			X.SetRow(i, vectorizedSentence.RawVector().Data)

			vectorizedClasses := n.TrainService.GetClassVector(classes)
			Y.SetRow(i, vectorizedClasses)

			i++
		}

		n.network.Train(&X, &Y)

		a := n.network.Test(&X, &Y)
		fmt.Println("Iteracion: ", e, "Errores: ", a.RawVector().Data)
	}

}

func randomArray(size int, v float64) (data []float64) {
	dist := distuv.Uniform{
		Min: -1 / math.Sqrt(v),
		Max: 1 / math.Sqrt(v),
	}

	data = make([]float64, size)
	for i := 0; i < size; i++ {
		data[i] = dist.Rand()
	}
	return
}

func transformListToArray(a *list.List) []string {
	i := 0
	array := make([]string, a.Len())
	for e := a.Front(); e != nil; e = e.Next() {
		array[i] = e.Value.(string)
		i++
	}
	return array
}
