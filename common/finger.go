package common

import (
	"math/rand"
	"time"
)

type InputFinger struct {
	Name string
	VictoryCount int
}

type OutFinger struct {
	Name string
	Finger string
	VictoryCount int
}

func FingerNumber(number int, inputFinger InputFinger) OutFinger {
	seed := time.Now().UnixNano()
	fingerNumber := rand.New(rand.NewSource(seed)).Intn(number)
	finger  := [3]string {Rock, Paper, Scissors}
	selectedFinger := OutFinger{
		Name: inputFinger.Name,
		Finger: finger[fingerNumber],
		VictoryCount: inputFinger.VictoryCount,
	}


	return selectedFinger
} 