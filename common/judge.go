package common

func Judge(fingers []OutFinger) (string, bool) {
    // 勝者の名前と勝利数を保存するマップ
    victories := make(map[string]int)

    // 各プレイヤーの勝利数をカウント
    for _, f := range fingers {
        victories[f.Name] = f.VictoryCount
        if f.VictoryCount >= 3 {
            return f.Name, true
        }
    }

    // 3回勝利したプレイヤーがいない場合
    return "", false
}