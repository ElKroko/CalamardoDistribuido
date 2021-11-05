package main

import ("fmt"
		"math/rand"
		"time"
		"math"
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

func RemovePlayer(players []int, i int) []int {
    players[i] = players[len(players)-1]
	fmt.Printf("El jugador_%d fue eliminado\n",value)
	//avisar jugador
	//avisar a Namenode
    return players[:len(players)-1]
}

func RemoveElemArray(array []int, i int) []int {
    array[i] = array[len(array)-1]
    return array[:len(array)-1]
}

func indexOf(element int, data []int) (int) {
	for k, v := range data {
		if element == v {
			return k
		}
	}
	return -1    //not found.
 }

func Juego3(players []int) {
	rand.Seed(time.Now().UnixNano())
	fmt.Println("Bienvenido al juego número 2")
	if len(players)%2 == 1 {
		players = RemovePlayer(rand.Intn(len(players)))
	}

	gamePlayers := player
	var player1 int
	var player2 int
	var play1 int
	var play2 int
	for i := 0; i < len(players)/2; i++ {
		player1 = rand.Intn(len(gamePlayers))
		player2 = rand.Intn(len(gamePlayers))
		play1 = (rand.Intn(10)+1)
		play2 = (rand.Intn(10)+1)
		if math.Abs(playBoss - play1) > math.Abs(playBoss - play2) {
			players = RemovePlayer(players, indexOf(player1))
		} else if math.Abs(playBoss - play1) < math.Abs(playBoss - play2) {
			players = RemovePlayer(players, indexOf(player1))
		}
		gamePlayers = RemoveElemArray(indexOf(player1))
		gamePlayers = RemoveElemArray(indexOf(player2))

	}
	return players
}

func Juego2(players []int) {
	rand.Seed(time.Now().UnixNano())
	fmt.Println("Bienvenido al juego número 2")

	team1 := []int{}
	team2 := []int{}

	teamPlayers := players 
	var newPlayer int

	for len(teamPlayers) > 0 {
		if len(teamPlayers) == 1 {
			players = RemovePlayer(players, indexOf(teamplayers[0]))
		} else {
			//ingresar juegador a team 1
			newPlayer = rand.Intn(len(teamPlayers))
			team1 = append(teamPlayers[newPlayer])
			teamPlayers = RemoveElemArray[newPlayer]
			//ingresar juegador a team 2
			newPlayer = rand.Intn(len(teamPlayers))
			team2 = append(teamPlayers[newPlayer])
			teamPlayers = RemoveElemArray[newPlayer]
		}
	}


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
			players = RemovePlayer(number)
		}
	} else if bossParity != team1Parity && bossParity == team2Parity {
		for number := range team1 {
			players = RemovePlayer(number)
		}
	} else {
		if rand.Intn(2) == 0 {
			for number := range team1 {
				players = RemovePlayer(number)
			}
		} else {
			for number := range team2 {
				players = RemovePlayer(number)
			}
		}
	}
	return players

}

func juego_1(jugador bool, ronda int) int {
	fmt.Println("Hola!")
	rand.Seed(time.Now().UnixNano())

	suma := 0
	var numero int
	gano := 0

	for ronda < 4 {

		// Seleccionar numero para jugar
		if jugador {
			fmt.Println("Elija un numero del 1 al 10")
			fmt.Print("-> ")
			fmt.Scanln(&numero)

		} else {
			numero = rand.Intn(10)

		}

		// Chequear si el numero es igual o mayor al del lider

		muere := false

		// Fin

		if !muere {
			suma += numero
			fmt.Println("La suma da: ", suma)
		} else {
			gano = 0
			break
		}

		// sumar al numero anterior
		if suma >= 21 {
			gano = 1
			break
		}

		fmt.Println(numero)
		ronda = ronda + 1
	}

	return gano

}

func main() {
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

	return
}

