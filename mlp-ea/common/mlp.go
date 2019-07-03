package common

import (
	"math/rand"

	mn "github.com/made2591/go-perceptron-go/model/neural"
	mu "github.com/made2591/go-perceptron-go/util"
	mv "github.com/made2591/go-perceptron-go/validation"
	"github.com/salvacorts/eaopt"
)

// MLPConfig Stores the common variables for training and evaluating all MLPs generated by the EA
type MLPConfig struct {
	Epochs        int
	Folds         int
	Classes       *[]string
	TrainingSet   *[]mn.Pattern
	ValidationSet *[]mn.Pattern
}

// MLP is a type of MultiLayerNetwork that implements Genome interface
type MLP struct {
	nn     mn.MultiLayerNetwork
	config MLPConfig
}

// Evaluate a MLP by getting its accuracy
func (mlp MLP) Evaluate() (float64, error) {
	// Train the network
	mn.MLPTrain(&mlp.nn, *mlp.config.TrainingSet, *mlp.config.Classes, mlp.config.Epochs)

	scores := mv.MLPKFoldValidation(
		&mlp.nn,
		*mlp.config.TrainingSet,
		mlp.config.Epochs,
		mlp.config.Folds,
		0,
		*mlp.config.Classes)

	accuracy, _ := mu.MaxInSlice(scores)

	return accuracy, nil
}

// Mutate modifies the weights of certain neurons, at random, depending on the application rate.
// Modifies the weights of the network after each epoch of network training,
// adding or subtracting a small random number that follows uniform distribution with the interval [-0.1, 0.1].
// The learning rate is modified by adding a small random number that follows uniform distribution
// in the interval [-0.05, 0.05]
func (mlp MLP) Mutate(rng *rand.Rand) {
	// TODO: Mutate
}

// Crossover carries out the multipoint cross-over between two chromosome nets,
// so that two networks are obtained whose hidden layer neurons are a mixture of the
// hidden layer neurons of both parents:
// some hidden neurons along with their in and out connections, from each parent
// make one offspring and the remaining hidden neurons make the other one.
// The learningrate is swapped between the two nets.
func (mlp MLP) Crossover(Y eaopt.Genome, rng *rand.Rand) {
	// TODO: Crossover
}

// AddLayer s intended to performincremental learning: it starts with a small structure
// and increments it, if neccesary, by adding new hidden units
func AddLayer(in eaopt.Genome, rng *rand.Rand) (out eaopt.Genome) {
	// TODO: AddLayer

	return out
}

// RemoveLayer eliminates one hidden neuron at random
func RemoveLayer(in eaopt.Genome, rng *rand.Rand) (out eaopt.Genome) {
	// TODO: RemoveLayer

	return out
}

// SubstituteNeuron eplaces one hiddenlayer neuron at random with a new one,
// initialized with random weights
func SubstituteNeuron(in eaopt.Genome, rng *rand.Rand) (out eaopt.Genome) {
	// TODO: SubstituteNeuron

	return out
}

// Train is used to train the individual-net for acertain number of generations, using the QP algorithm.
func Train(in eaopt.Genome, rng *rand.Rand) (out eaopt.Genome) {
	// TODO: Train

	return out
}

// Clone a MLP to produce a new one that points to a different one.
func (mlp MLP) Clone() eaopt.Genome {
	return MLP{
		nn:     mlp.nn,
		config: mlp.config,
	}
}
