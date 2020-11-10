package models

import (
	"container/list"

	"gonum.org/v1/gonum/mat"
)

type Network struct {
	Layers         *list.List
	LearningRate   float64
	LossFunction   LossFunction
	ThresholdValue float64
}

func NewNetwork() *Network {
	return &Network{
		Layers:         list.New(),
		LearningRate:   0.01,
		LossFunction:   &BinaryCrossEntropy{},
		ThresholdValue: 0.5,
	}
}

func (n *Network) Train(X *mat.Dense, Y *mat.Dense) {
	r, _ := X.Dims()
	for i := 0; i < r; i++ {
		inputs := X.RowView(i).(*mat.VecDense)
		correctValue := Y.RowView(i).(*mat.VecDense)
		n.Predict(inputs)
		actualLayer := n.Layers.Back()

		actualLayerCount := 0
		var lastWeights *mat.Dense
		for ; actualLayer != nil; actualLayer = actualLayer.Prev() {
			var delta *mat.Dense
			layer := actualLayer.Value.(*Layer)
			if actualLayerCount == 0 {
				delta = n.getLastDeltaLayer(actualLayer, correctValue)
				lastWeights = copyWeights(layer.Weights)
			} else {
				delta = n.getOthersDeltaLayer(actualLayer.Next(), lastWeights)
				lastWeights = copyWeights(layer.Weights)
			}

			layer.TrainWeights(delta, n.LearningRate, n.LossFunction)
			actualLayerCount++
		}

	}

}

func (n *Network) getLastDeltaLayer(layerNode *list.Element, correctValues *mat.VecDense) *mat.Dense {
	result := mat.NewDense(correctValues.Len(), 1, nil)
	layer := layerNode.Value.(*Layer)
	for i := 0; i < correctValues.Len(); i++ {
		lossValue := n.LossFunction.Derivate(layer.Outputs.AtVec(i), correctValues.AtVec(i))
		result.Set(i, 0, lossValue)
	}
	return result
}

func (n *Network) getOthersDeltaLayer(layer *list.Element, weights *mat.Dense) *mat.Dense {
	deltaMatrix := layer.Value.(*Layer).Delta
	_, r := weights.Dims()
	_, c := deltaMatrix.Dims()
	resultMatrix := mat.NewDense(r, c, nil)
	resultMatrix.Mul(weights.T(), deltaMatrix)
	return resultMatrix
}

func (n *Network) Predict(input *mat.VecDense) *mat.VecDense {
	vector := input
	for element := n.Layers.Front(); element != nil; element = element.Next() {
		layer := element.Value.(*Layer)
		layer.AssignInput(vector)
		vector = layer.GetOutputs()
	}
	return vector
}

func (n *Network) AddLayer(layer *Layer) {
	n.Layers.PushBack(layer)
}

func (n *Network) GetWeights() *list.List {
	weights := list.New()
	for e := n.Layers.Front(); e != nil; e = e.Next() {
		layer := e.Value.(*Layer)
		weights.PushBack(layer.Weights)
	}

	return weights
}

func (n *Network) SetWeights(weights *mat.Dense, layerIndex int) {
	layer := n.Layers.Front()
	for actualLayerIndex := 0; actualLayerIndex < layerIndex && layer != nil; actualLayerIndex++ {
		layer = layer.Next()
	}

	layer.Value.(*Layer).Weights = weights
}

func (n *Network) Test(X *mat.Dense, Y *mat.Dense) *mat.VecDense {
	r, _ := X.Dims()
	confusion_matrix := mat.NewVecDense(2, []float64{0, 0})

	for i := 0; i < r; i++ {
		inputs := X.RowView(i).(*mat.VecDense)
		correctValues := Y.RowView(i).(*mat.VecDense)
		predictValues := n.Predict(inputs)
		// rawValues := predictValues.RawVector().Data
		n.ThresholdPredictValues(predictValues.RawVector().Data)
		updateConfusionMatrix(confusion_matrix, predictValues, correctValues)
	}

	return confusion_matrix
}

func equal_arrays(a, b []float64) bool {
	if len(a) != len(b) {
		return false
	}

	equal := true
	for i := 0; i < len(a); i++ {
		if a[i] != b[i] {
			equal = false
			break
		}
	}

	return equal
}

func updateConfusionMatrix(resultVector *mat.VecDense, predictVector, correctVector *mat.VecDense) {
	rawPredictVector := predictVector.RawVector().Data
	rawCorrectVector := correctVector.RawVector().Data

	equal := equal_arrays(rawPredictVector, rawCorrectVector)

	if equal {
		value := resultVector.AtVec(0) + 1
		resultVector.SetVec(0, value)
	} else {
		value := resultVector.AtVec(1) + 1
		resultVector.SetVec(1, value)
	}
}

func copyWeights(weights *mat.Dense) *mat.Dense {
	r, c := weights.Dims()
	copy := mat.NewDense(r, c, nil)
	for i := 0; i < r; i++ {
		for j := 0; j < c; j++ {
			copy.Set(i, j, weights.At(i, j))
		}
	}
	return copy
}

func transformVecToDense(vec *mat.VecDense) *mat.Dense {
	r := vec.Len()
	matrix := mat.NewDense(r, 1, nil)
	for i := 0; i < r; i++ {
		matrix.Set(i, 0, vec.AtVec(i))
	}
	return matrix
}

func (n *Network) ThresholdPredictValues(predictValues []float64) {
	for i, value := range predictValues {
		if value >= n.ThresholdValue {
			predictValues[i] = 1
			continue
		}
		predictValues[i] = 0
	}
}
