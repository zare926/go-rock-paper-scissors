package common

import (
	"math/rand"
)

func ChoiceFingerName(fingerNumber *rand.Rand) string {
	finger  := [3]string {"グー", "チョキ", "パー"}
	selectedFinger := finger[fingerNumber.Int31() % 3]
    return selectedFinger
}