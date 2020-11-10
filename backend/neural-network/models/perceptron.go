package models

// type Perceptron struct {
// 	Inputs         *mat.VecDense
// 	Weights        *mat.VecDense
// 	Bias           float64
// 	ThresholdValue float64
// 	LastOutput     float64
// }

// func (p *Perceptron) ThresholdFunction(output float64) bool {
// 	if output > p.ThresholdValue {
// 		return true
// 	}

// 	return false
// }

// func (p *Perceptron) activate(activationFunction ActivationFunction, value float64) float64 {
// 	return activationFunction.GetActivationValue(value)
// }

// func (p *Perceptron) Output(activationFunction ActivationFunction) float64 {
// 	inputVector := prepend(p.Inputs, 1)
// 	weightVector := prepend(p.Weights, p.Bias)
// 	sum := mat.Dot(inputVector, weightVector)
// 	output := p.activate(activationFunction, sum)

// 	if !p.ThresholdFunction(output) {
// 		return 0
// 	}

// 	p.LastOutput = output
// 	return output
// }

// func prepend(vector *mat.VecDense, value float64) *mat.VecDense {
// 	newVector := mat.NewVecDense(vector.Len()+1, nil)
// 	for i := 1; i < newVector.Len(); i++ {
// 		if i == 1 {
// 			newVector.SetVec(i, value)
// 		} else {
// 			newVector.SetVec(i, vector.AtVec(i-1))
// 		}
// 	}

// 	return newVector
// }
