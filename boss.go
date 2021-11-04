package main

import ("fmt"
		"math/rand"
		"time"
)

func ShowPlayers(players []int) {
	for _, value := range players {
		fmt.Printf("jugador_%d\n",value)
	}
}

func ShowPlays(player) {
	fmt.Println("Que revise txt")
}

func EliminatePlayers(players []int, number) {
	for number := range players {
		fmt.Prnumberntln(players[number])// send false to player
	}
}

func RemoveElemArray(players []int, i int) []int {
    players[i] = players[len(players)-1]
	fmt.Printf("El jugador_%d fue eliminado\n",value)
	//avisar jugador
	//avisar a Namenode
    return players[:len(players)-1]
}

func Juego2(players []int) {
	rand.Seed(time.Now().UnixNano())
	fmt.Println("Bienvenido al juego número 2")

	// numberPlayers := (rand.Intn(8)+9)

	// if (numberPlayers % 2) == 0 {

	// } else {

	// }


	fmt.Println("Elija un número del 1 al 4")
	var bossNumber int
    fmt.Scanln(&bossNumber)

	bossParity := bossNumber % 2
	team1Parity := (rand.Intn(4)+1) % 2
	team2Parity := (rand.Intn(4)+1) % 2

	if bossParity == team1Parity && bossParity == team2Parity {
		fmt.Println("Nadie muerio en otra ronda")
	} else if bossParity == team1Parity && bossParity != team2Parity {
		for number := range team2 {
			players = RemoveElemArray(number)
		}
	} else if bossParity != team1Parity && bossParity == team2Parity {
		for number := range team1 {
			players = RemoveElemArray(number)
		}
	} else {
		if rand.Intn(2) == 0 {
			for number := range team1 {
				players = RemoveElemArray(number)
			}
		} else {
			for number := range team2 {
				players = RemoveElemArray(number)
			}
		}
	}
	return players

}

func main() {
	rand.Seed(time.Now().UnixNano())
	players := []int{1, 2, 3,4,5,6,7,8,9,10,11,12,13,14,15,16}
	var option int
	next := false
	fmt.Println("Welcome back boss")
	for next == false {
		fmt.Println("Elija que quiere hacer:")
		fmt.Println("[1] Empezar los juegos del calamar")
		fmt.Println("[2] Revisar jugadas de jugadores")
		fmt.Scanln(&option)
		if option == 1 {
			players = Juego1()
			next = true
		} else {
			fmt.Println("Aun no hay jugadas")
		}
	}

	next = false

	fmt.Println("Los jugadores que sobrevivieron a la etapa 1 son:")
	ShowPlayers(players)
	for next == false {
		fmt.Println("Elija que quiere hacer:")
		fmt.Println("[1] Comenzar la etapa 2")
		fmt.Println("[2] Revisar jugadas de jugadores")
		fmt.Scanln(&option)
		if option == 1 {
			players = Juego2()
			next = true
		} else {
			fmt.Println("[2] Revisar jugadas de jugadores")
			fmt.Scanln(&option)
			ShowPlays(option)
		}
	}

	fmt.Println("Los jugadores que sobrevivieron a la etapa 1 son:")
	ShowPlayers(players)
	for next == false {
		fmt.Println("Elija que quiere hacer:")
		fmt.Println("[1] Comenzar la etapa 3")
		fmt.Println("[2] Revisar jugadas de jugadores")
		fmt.Scanln(&option)
		if option == 1 {
			players = Juego3()
			next = true
		} else {
			fmt.Println("[2] Revisar jugadas de jugadores")
			fmt.Scanln(&option)
			ShowPlays(option)
		}
	}

	fmt.Println("Los jugadores que ganaron los juegos del calamar son:")
	ShowPlayers(players)
}

