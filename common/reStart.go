package common

import "fmt"

func Jaiken(fingers []*OutFinger) (bool, []string) {
    countRock := 0
    countScissors := 0
    countPaper := 0

    // 各指のフィンガーを数える
    for _, finger := range fingers {
        switch finger.Finger {
        case Rock:
            countRock++
        case Scissors:
            countScissors++
        case Paper:
            countPaper++
        default:
            fmt.Printf("%s のフィンガーが不明です。\n", finger.Name)
        }
    }
    // 勝者を判定
    winners := []string{}
    if countRock > 0 && countScissors > 0 && countPaper == 0 {
        // グーが勝つ
        for _, finger := range fingers {
            if finger.Finger == Rock {
                finger.VictoryCount++
                winners = append(winners, finger.Name)
                fmt.Println(winners)
            }
        }
    } else if countRock > 0 && countScissors == 0 && countPaper > 0 {
        // パーが勝つ
        for _, finger := range fingers {
            if finger.Finger == Paper {
                finger.VictoryCount++
                winners = append(winners, finger.Name)
                fmt.Println(winners)
            }
        }
    } else if countRock == 0 && countScissors > 0 && countPaper > 0 {
        // チョキが勝つ
        for _, finger := range fingers {
            if finger.Finger == Scissors {
                finger.VictoryCount++
                winners = append(winners, finger.Name)
                fmt.Println(winners)
            }
        }
    }

    // 勝者がいない場合はfalseを返す
    if len(winners) == 0 {
        fmt.Println(winners)
        return false, winners
    }

    return true, winners
}