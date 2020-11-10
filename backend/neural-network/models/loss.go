package models

import (
	"math"

	"gonum.org/v1/gonum/mat"
)

type LossFunction interface {
	GetError(predictValue, correctValue, weight float64, activateFunction ActivationFunction) float64
	ErrorBase(predictValue, correctValue float64) float64
	Derivate(predictValue, correctValue float64) float64
	LossValue(X, Y *mat.VecDense) float64
}

type BinaryCrossEntropy struct{}

func (b *BinaryCrossEntropy) ErrorBase(predictValue, correctValue float64) float64 {
	return correctValue*math.Log(predictValue) + (1-correctValue)*math.Log(1-predictValue)
}

func (b *BinaryCrossEntropy) Derivate(predictValue, correctValue float64) float64 {
	return (-1 * correctValue / predictValue) + ((1-correctValue)/(1-predictValue))
}

func (b *BinaryCrossEntropy) GetError(predictValue, correctValue, weight float64, activateFunction ActivationFunction) float64 {
	return b.Derivate(predictValue, correctValue) * activateFunction.GetDerivateValue(predictValue) * weight
}

func (b *BinaryCrossEntropy) LossValue(predictsVector, correctsVector *mat.VecDense) float64 {
	predictsVectorLen := predictsVector.Len()
	if predictsVectorLen != correctsVector.Len() {
		return 0.0
	}

	lossValue := 0.0

	for i := 0; i < predictsVectorLen; i++ {
		predictValue := predictsVector.AtVec(i)
		correctValue := correctsVector.AtVec(i)
		lossValue += b.ErrorBase(predictValue, correctValue)
	}

	lossValue *= -1
	return lossValue
}
