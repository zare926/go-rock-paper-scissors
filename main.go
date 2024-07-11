package main

import (
	"fmt"
	"math/rand"
	"projects/golesson/akita"
	"projects/golesson/common"
	"projects/golesson/france"
	"projects/golesson/tokyo"
	"time"
)

func main() {
    akita.Intro()
    tokyo.Intro()
    akita.Tukkomi()
    france.Intro()
    france.Tukkomi()

	aIn := common.InputFinger{
		Name:   "アキタ",
		VictoryCount: 0,
	}
	
	bIn := common.InputFinger{
		Name:   "トウキョウ",
		VictoryCount: 0,
	}
	
	cIn := common.InputFinger{
		Name:   "フランス",
		VictoryCount: 0,
	}
	
	
	akitaFingerNumber := common.FingerNumber(3, aIn)
	tokyoFingerNumber := common.FingerNumber(3, bIn)
	franceFingerNumber := common.FingerNumber(3, cIn)
	
    for {
		fingerArray := []*common.OutFinger{
			&akitaFingerNumber,
			&tokyoFingerNumber,
			&franceFingerNumber,
		}

		for _, finger := range fingerArray {
			seed := time.Now().UnixNano()
			fingerNumber := rand.New(rand.NewSource(seed)).Intn(3)
			fingerChoices := [3]string{common.Rock, common.Paper, common.Scissors}
			finger.Finger = fingerChoices[fingerNumber]
    	}
		for _, finger := range fingerArray {
			fmt.Printf("Name: %s, Finger: %v, VictoryCount: %d\n", finger.Name, finger.Finger, finger.VictoryCount)
		}
		
		hasWinner, winnerNames := common.Jaiken(fingerArray)
		 
		if hasWinner {
			for _, winnerName := range winnerNames {
				fmt.Printf("%sが勝ちました！\n", winnerName)
			}
			for _, outFinger := range fingerArray {
				if outFinger.VictoryCount >= 3 {
					fmt.Printf("%sが最終的な勝者です！\n", outFinger.Name)
					return
				}
			}
		} else {
			fmt.Println("アイコなので再戦！")
		}
		time.Sleep(5 * time.Second)
	}
}