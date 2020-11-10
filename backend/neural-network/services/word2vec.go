package services

import (
	"os"
	"strings"

	"github.com/ynqa/wego/pkg/embedding"
	"github.com/ynqa/wego/pkg/model"
	"github.com/ynqa/wego/pkg/model/modelutil/save"
	"github.com/ynqa/wego/pkg/model/word2vec"
	"github.com/ynqa/wego/pkg/search"

	"gonum.org/v1/gonum/mat"
)

type Word2VecService struct {
	model    model.Model
	searcher *search.Searcher
	length   int
}

var MaxWords = 100
var ClassesQuantity = 9

func NewWordService() *Word2VecService {
	service := &Word2VecService{}
	service.NewModel()
	return service
}

func (w2v *Word2VecService) GetWordVectorFloat(word string) []float64 {
	emb, found := w2v.searcher.Items.Find(word)
	if found {
		return emb.Vector
	}
	
	neighbors, _ := w2v.searcher.SearchInternal(word, w2v.length)
	if neighbors == nil {
		return []float64{}
	}
	mostSimilarWord, _ := w2v.searcher.Items.Find(neighbors[0].Word)

	return mostSimilarWord.Vector
}

func (w2v *Word2VecService) GetVector(sentence string) *mat.VecDense {
	words := strings.Split(sentence, " ")
	vectorWord := w2v.GetVectorFromArray(words)
	return RedimensionVector(MaxWords, vectorWord)
}

func (w2v *Word2VecService) GetVectorFromArray(words []string) *mat.VecDense {
	vectorDimension := 10
	vectorWords := mat.NewVecDense(len(words)*vectorDimension, nil)
	actualIndexVectorWords := 0
	for _, word := range words {
		vectorWord := w2v.GetWordVectorFloat(word)
		for _, vector := range vectorWord {
			vectorWords.SetVec(actualIndexVectorWords, vector)
			actualIndexVectorWords++
		}
	}
	return vectorWords
}

func RedimensionVector(n int, vector *mat.VecDense) *mat.VecDense {
	newVector := mat.NewVecDense(n, nil)
	newVector.Zero()

	for i := 0; i < vector.Len(); i++ {
		newVector.SetVec(i, vector.AtVec(i))
	}

	return newVector
}

// func (w2v *Word2VecService) GetString(numberText float64) string {
// neighbors, _ := w2v.searcher.SearchVector(word, w2v.length)

// return neighbors[0].Similarity
// return "a"
// }

// func (w2v *Word2VecService) GetSentence(vector *mat.VecDense) string {
// 	vectorRaw := vector.RawVector().Data
// 	words := make([]string, vector.Len())
// 	for i, value := range vectorRaw {
// 		words[i] = w2v.GetString(value)
// 	}

// 	sentence := strings.Join(words, " ")
// 	return sentence
// }

func (w2v *Word2VecService) NewModel() {
	model, err := word2vec.New(
		word2vec.WithWindow(5),
		word2vec.WithModel(word2vec.Cbow),
		word2vec.WithOptimizer(word2vec.NegativeSampling),
		word2vec.WithNegativeSampleSize(5),
		word2vec.Verbose(),
	)
	if err != nil {
		// failed to create word2vec.
	}
	w2v.model = model
}

func (w2v *Word2VecService) Train(words string) error {
	input := strings.NewReader(words)
	if err := w2v.model.Train(input); err != nil {
		return err
	}

	fileName := "backups/word2vec.model"
	saveFileWithModel(fileName, w2v.model)

	searcher, _ := getSearcherFromFile(fileName)
	w2v.searcher = searcher
	w2v.length = len(words)

	return nil
}

func saveFileWithModel(filename string, model model.Model) {
	f, _ := os.Create(filename)
	model.Save(f, save.Single)
	defer f.Close()
}

func getSearcherFromFile(filename string) (*search.Searcher, error) {
	f, _ := os.Open(filename)
	embs, err := embedding.Load(f)
	defer f.Close()

	if err != nil {
		return nil, err
	}
	searcher, _ := search.New(embs...)
	return searcher, nil
}
