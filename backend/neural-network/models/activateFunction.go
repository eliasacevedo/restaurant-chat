package models

import "math"

type ActivationFunction interface {
	GetActivationValue(x float64) float64
	GetDerivateValue(predictValue float64) float64
}

type SygmoidFunction struct{}

func (s *SygmoidFunction) GetActivationValue(x float64) float64 {
	return 1 / (1 + math.Exp(x*(-1)))
}

func (s *SygmoidFunction) GetDerivateValue(predictValue float64) float64 {
	return predictValue * (1 - predictValue)
}

type ReluFunction struct{}

func (s *ReluFunction) GetActivationValue(x float64) float64 {
	if x > 0 {
		return x
	}
	return 0
}

func (s *ReluFunction) GetDerivateValue(predictValue float64) float64 {
	if predictValue > 0 {
		return 1
	}
	return 0
}
