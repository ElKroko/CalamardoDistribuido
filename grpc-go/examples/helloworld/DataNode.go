package main

import ("os"
		"log"
        "strconv"
        "github.com/streadway/amqp"
)

func main() {
    next := false
	for next == false {
		test := []int{1, 2, 3,4,5,6,7,8,9,10,11,12,13,14,15,16}
		player, round, plays := 1, 1, test
		fileName := ""
		fileName = "jugador_" + strconv.Itoa(player) + "__ronda_" + strconv.Itoa(round) + ".txt"
		f, err := os.Create(fileName)
		var index int64 = 0
		if err != nil {
			log.Fatal(err)
		}
		defer f.Close()
        for play := range plays {
            data := []byte(strconv.Itoa(play) + "\n")
            _, err := f.WriteAt(data, index)
            index = 1 + index
            if err != nil {
                log.Fatal(err)
            }
        }
    }
    return
}