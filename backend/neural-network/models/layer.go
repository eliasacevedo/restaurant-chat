package models

import (
	"gonum.org/v1/gonum/mat"
)

type Layer struct {
	// Perceptrons        *list.List
	Weights            *mat.Dense
	ActivationFunction ActivationFunction
	Inputs             *mat.VecDense
	Bias               float64
	Outputs            *mat.VecDense
	PonderedSum        *mat.VecDense
	Delta              *mat.Dense
}

func NewLayer(inputQuantity, perceptronQuantity int, activationFunction ActivationFunction) *Layer {

	layer := &Layer{
		// Perceptrons:        list.New(),
		ActivationFunction: activationFunction,
		Weights:            mat.NewDense(perceptronQuantity, inputQuantity+1, nil),
		Inputs:             mat.NewVecDense(inputQuantity+1, nil),
	}
	return layer
}

func (l *Layer) GetOutputs() *mat.VecDense {
	r, _ := l.Weights.Dims()
	matrix := mat.NewDense(r, 1, nil)
	matrix.Mul(l.Weights, l.Inputs)

	l.PonderedSum = mat.NewVecDense(r, matrix.RawMatrix().Data)

	for i := 0; i < r; i++ {
		actualValue := matrix.At(i, 0)
		activatedValue := l.ActivationFunction.GetActivationValue(actualValue)

		matrix.SetRow(i, []float64{activatedValue})
	}
	outputs := matrix.ColView(0).(*mat.VecDense)
	l.Outputs = outputs
	return outputs
}

func (l *Layer) AssignDefaultValues(Inputs *mat.VecDense, Weights *mat.Dense, function ActivationFunction, Bias float64, ThresholdValue float64) {
	l.Weights = getWeightsWithBias(Weights, Bias)
	l.Inputs = Inputs // push(Inputs, 1.0)
	l.ActivationFunction = function
	l.Bias = Bias
}

func getWeightsWithBias(Weight *mat.Dense, Bias float64) *mat.Dense {
	row, column := Weight.Dims()
	newWeights := Weight.Grow(0, 1).(*mat.Dense)
	biasCol := make([]float64, row)
	for i := 0; i < row; i++ {
		biasCol[i] = Bias
	}
	newWeights.SetCol(column, biasCol)
	return newWeights
}

func (l *Layer) TrainWeights(delta *mat.Dense, learningRate float64, lossFunction LossFunction) {
	l.Delta = delta
	row, column := l.Weights.Dims()
	updatedWeights := mat.NewDense(row, column, nil)
	for i := 0; i < row; i++ {
		actualPredictValue := l.Outputs.AtVec(i)
		actualCorrectValue := delta.At(i, 0)
		for j := 0; j < column; j++ {
			actualWeight := l.Weights.At(i, j)
			lossFunctionValue := actualCorrectValue * l.ActivationFunction.GetDerivateValue(actualPredictValue)
			//lossFunction.GetError(actualPredictValue, actualCorrectValue, actualWeight, l.ActivationFunction)
			// go fmt.Println("Peso actual: ", actualWeight, ", Actualizacion: ", lossFunctionValue, ", Diferencia: ", actualWeight-lossFunctionValue)
			newWeight := updateWeights(learningRate, actualWeight, lossFunctionValue)
			// if math.IsNaN(newWeight) {
			// 	continue
			// }
			updatedWeights.Set(i, j, newWeight)
		}
	}

	l.Weights = updatedWeights
}

func updateWeights(learningRate, weight, lossFunctionValue float64) float64 {
	return weight + learningRate*(-1*lossFunctionValue)
}

func push(vector *mat.VecDense, value float64) *mat.VecDense {
	if vector == nil {
		return nil
	}
	newVector := mat.NewVecDense(vector.Len()+1, nil)
	for i := 0; i < newVector.Len(); i++ {
		if i == newVector.Len()-1 {
			newVector.SetVec(i, value)
		} else {
			newVector.SetVec(i, vector.AtVec(i))
		}
	}
	return newVector
}

func (l *Layer) AssignInput(Inputs *mat.VecDense) {
	l.Inputs = push(Inputs, 1)
}
