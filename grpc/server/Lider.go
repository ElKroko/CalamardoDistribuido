package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"time"
	pb "lab/game/proto"
	"google.golang.org/grpc"
	"math/rand"
	"math"
)

type server struct {
	pb.UnimplementedCalamardoGameServer
}

type PlayerStruct struct {
	id    int32
	alive bool
	score int32
	jugada int32
}

// Variables
var totalPlayers int								// Jugadores actuales
var players_online []PlayerStruct
var started bool
var total int = 2 									// Aqui se define el total de jugadores maximos!
var playsBoss []int

var etapaActual int
var total_juego2 int
var total_juego3 int = 0



// Funciones de mensajeria
func (s *server) JoinGame(ctx context.Context, in *pb.JoinRequest) (*pb.JoinReply, error) {
	log.Printf("Recibido: %s", in.GetMessage())
	
	players_online = append(players_online, PlayerStruct{int32(totalPlayers), true, 0, 0})
	totalPlayers += 1

	return &pb.JoinReply{IdJugador: int32(totalPlayers-1), Alive: true, Round: 0}, nil
}

func (s *server) StartGame(ctx context.Context, in *pb.StartRequest) (*pb.StartReply, error) {

	// log.Printf("El jugador %d envio: %s \t en la etapa: %d",in.GetId(), in.GetMessage(), in.GetEtapa())
	etapa := in.GetEtapa()

	switch etapa {
	case 0:
		if len(players_online) == total {					// Cantidad de jugadores = maxio
			return &pb.StartReply{Started: true}, nil
		} else {
			return &pb.StartReply{Started: false}, nil
		}
	case 1:
		if started{
			return &pb.StartReply{Started: true}, nil
		} else {
			return &pb.StartReply{Started: false}, nil
		}
	case 2:
		if started{
			return &pb.StartReply{Started: true}, nil
		} else {
			return &pb.StartReply{Started: false}, nil
		}
	default:
		return &pb.StartReply{Started: false}, nil
	}
}

// Para el juego 1: 
// El lider debe esperar la jugada del jugador
// Al momento de recibirla, necesitamos calcular si este numero cumple las condiciones (>= y 21)
// 		Si es >=, eliminar el jugador
// 		Si es 21>, avisar que gano y cambiar su ronda



// Funcion para enviar y recibir jugadas desde el lider, para todos los juegos.

func (s *server) JuegoMsg(ctx context.Context, in *pb.JuegoRequest) (*pb.JuegoReply, error) {
	vivo := true
	etapa_jugador := in.GetEtapa()

	if etapaActual == 1 && etapa_jugador == 1{
		ronda := in.GetRound()
		id_jugador := in.GetId()
		log.Println("El Jugador ", id_jugador, "en la ronda ", ronda)
		playPlayer := in.GetJugada()
		log.Println("Ha jugado la carta", playPlayer)

		suma := in.GetScore()

		if playPlayer >= int32(playsBoss[ronda]) && playPlayer >= 1 && playPlayer <= 10{
			RemovePlayer(indexOfPlayers(id_jugador))
			vivo = false
		} else {
			suma = suma + playPlayer
			if suma >= 21 {
				etapa_jugador += 1
				log.Println("")
				log.Println("El Jugador ", id_jugador, "en la ronda ", ronda)
				log.Println("Paso a la siguiente etapa! con puntaje ", suma)
				ronda = 0														 // Reiniciamos la ronda para la etapa 2
				started = false
			} else if ronda == 3 {
				vivo = false
				RemovePlayer(indexOfPlayers(id_jugador))
			} else {
				ronda = ronda + 1
			}
		}
		return &pb.JuegoReply{Alive: vivo, Round: ronda, Score: suma, Etapa: etapa_jugador}, nil


	} else if etapa_jugador == 2{
		// Aqui va la logica del juego 2
		fmt.Println("Recibi mensaje juego 2")
		// Recibir el mensaje, y devolver todo igual... 
		players_online[indexOfPlayers(in.Id)].jugada = in.GetJugada()
		// Para decidir si el jugador vivio o no, se usara una funcion alive() que verificara si estoy vivo al final... 
		// Preguntando si el id de el jugador esta en la lista.
		total_juego2 += 1
		etapa_jugador += 1
		started = false
		return &pb.JuegoReply{Alive: true, Round: 1, Score: 0, Etapa: etapa_jugador}, nil


	} else if etapa_jugador == 3{
		// Aqui va la logica del juego 3
		fmt.Println("Recibi mensaje juego 3")
		total_juego3 += 1
		players_online[indexOfPlayers(in.Id)].jugada = in.GetJugada()
		etapa_jugador += 1
		started = false
		return &pb.JuegoReply{Alive: vivo, Round: 0, Score: 0, Etapa: etapa_jugador}, nil


	} else{
		fmt.Println("Algo extraño paso!")
		return &pb.JuegoReply{Alive: vivo, Round: 3, Score: 0, Etapa: etapa_jugador}, nil
	}
	
}


// Funciones del Juego


// Funciones Auxiliares
func indexOfPlayers(element int32) (int) {
	for k, v := range players_online {
		if element == v.id {
			return k
		}
	}
	fmt.Println("No encontre nada")
	return -1    //not found.
}

func indexOf(array []PlayerStruct, element int32) (int) {
	for k, v := range array {
		if element == v.id {
			return k
		}
	}
	fmt.Println("No encontre nada IndexOf")
	return -1    //not found.
}

func RemovePlayer(i int){
	value := players_online[i].id 
    players_online[i] = players_online[len(players_online)-1]
	fmt.Printf("El jugador_%d fue eliminado\n",value)
	//avisar jugador
	//avisar a Namenode
}

func RemoveElemArray(array []PlayerStruct, i int) []PlayerStruct {
    array[i] = array[len(array)-1]
    return array[:len(array)-1]
}



func main() {
	log.Printf("Bienvenido al Calamardo, iniciando servicios...")
	go func() {
		listner, err := net.Listen("tcp", ":8080")

		if err != nil {
			panic("cannot create tcp connection " + err.Error())
		}

		serv := grpc.NewServer()
		pb.RegisterCalamardoGameServer(serv, &server{})
		if err = serv.Serve(listner); err != nil {
			panic("cannot initialize the server" + err.Error())
		}
	}()

	

	started = false
	totalPlayers = 0
	etapaActual = 0
	// Calamardo := ""
	// var muere bool

	var empezar string
	
	

	for totalPlayers != total {
		time.Sleep(2 * time.Second)

		if totalPlayers != total {
			fmt.Println("Aun no hay suficientes calamares... Total:", totalPlayers, "\t esperamos:", total)
		}
	}

	fmt.Println("")
	fmt.Println("Todos a bordo!")
	started = false


	for i := 0; i < total; i++ {
		fmt.Println(players_online[i].id)
	}
	

	// Logica juego 1
	fmt.Println("Hola!, el lider eligira sus cartas")

	var numero int
	//gano := 0
	ronda := 0
	etapaActual = 1

	for ronda < 4 {
		fmt.Println("Elija un numero del 6 al 10")
		fmt.Print("-> ")
		fmt.Scanln(&numero)
		playsBoss = append(playsBoss, numero)
		ronda = ronda + 1
	}

	started = true


	fmt.Println("Escribe C para continuar: ")
	fmt.Scanln(&empezar)

	started = false



	// Logica Juego 2
	rand.Seed(time.Now().UnixNano())
	fmt.Println("Bienvenido al juego número 2")
	fmt.Println("Elija un número del 1 al 4")
	var bossNumber int
    fmt.Scanln(&bossNumber)
	bossParity := bossNumber % 2
	fmt.Println(bossParity)

	total_juego2 = 0
	
	started = true

	// Esperar a que los jugadores manden su jugada

	for total_juego2 < len(players_online){
		fmt.Println("Aun no han jugado todos...")
		time.Sleep(2* time.Second)
	}
	
	started = false
	// Luego de que hagan su jugada, los separamos en equipo
	// Y calculamos la paridad


	team1 := []int32{}
	team2 := []int32{}

	teamPlayers := players_online 
	var newPlayer int


		// Cambiar team_players = players_online
	for len(teamPlayers) > 0 {
		if len(teamPlayers) == 1 {
			RemovePlayer(indexOfPlayers(teamPlayers[0].id))
			teamPlayers = RemoveElemArray(teamPlayers, indexOfPlayers(teamPlayers[0].id))
			fmt.Println("Eliminamos, habia solo 1 jugador")
		} else {
			//ingresar juegador a team 1
			newPlayer = rand.Intn(len(teamPlayers))				// al azar entre todos los players online
			team1 = append(team1, teamPlayers[newPlayer].id)				// 
			teamPlayers = RemoveElemArray(teamPlayers, newPlayer)
			//ingresar juegador a team 2
			newPlayer = rand.Intn(len(teamPlayers))
			team2 = append(team2, teamPlayers[newPlayer].id)
			teamPlayers = RemoveElemArray(teamPlayers, newPlayer)
		}
	}

	// cuando todos los jugadores me hayan enviado su jugada, calcular que pasa con la paridad
	team1Sum := 0
	team2Sum := 0
	number := 0
	for number < len(team1) {
		fmt.Println("indice: ", players_online[indexOfPlayers(team1[number])])
		fmt.Println("Jugada", players_online[indexOfPlayers(team1[number])].jugada)
		team1Sum += int(players_online[team1[number]].jugada)
		number ++
	}
	number = 0
	for number < len(team2) {
		fmt.Println("indice: ", players_online[indexOfPlayers(team2[number])])
		fmt.Println("Jugada", players_online[indexOfPlayers(team2[number])].jugada)
		team2Sum += int(players_online[team2[number]].jugada)
		number ++
	}

	

	team1Parity := team1Sum % 2
	team2Parity := team2Sum % 2

	fmt.Println("Paridad team 1: ", team1Parity)
	fmt.Println("Paridad team 2: ", team2Parity)

	// liberar el ping. (listo = true) y actualizar alive (decidiendo si murio o no, para cerrar procesos)

	if bossParity == team1Parity && bossParity == team2Parity {
		fmt.Println("Nadie muerio en otra ronda")
	} else if bossParity == team1Parity && bossParity != team2Parity {
		fmt.Println("Gano Team 1")
		for number := range team2 {
			RemovePlayer(number)
		}
	} else if bossParity != team1Parity && bossParity == team2Parity {
		fmt.Println("Gano Team 2")
		for number := range team1 {
			RemovePlayer(number)
		}
	} else {
		fmt.Println("Nadie gano, decidiendo a quien nos echamos...")
		if rand.Intn(2) == 0 {
			for number := range team1 {
				RemovePlayer(number)
			}
		} else {
			for number := range team2 {
				RemovePlayer(number)
			}
		}
	}
	

	// Hacer logica del juego 3
	started = true
	total_juego3 = 0

	for total_juego3 < len(players_online){
		fmt.Println("Aun no han jugado todos...")
		time.Sleep(2* time.Second)
	}
	// Esperar mensajes de vuelta 
	started = false
	rand.Seed(time.Now().UnixNano())
	fmt.Println("Bienvenido al juego número 3")
	if len(players_online)%2 == 1 {
		RemovePlayer(rand.Intn(len(players_online)))
	}

	fmt.Println("Elija un número del 1 al 10")
    fmt.Scanln(&bossNumber)

	gamePlayers := players_online
	var player1 int32
	var player2 int32
	var play1 int
	var play2 int
	for i := 0; i < len(players_online)/2; i++ {
		player1 = gamePlayers[rand.Intn(len(gamePlayers))].id
		gamePlayers = RemoveElemArray(gamePlayers, indexOf(gamePlayers, player1))
		player2 = gamePlayers[rand.Intn(len(gamePlayers))].id
		gamePlayers = RemoveElemArray(gamePlayers, indexOf(gamePlayers, player2))
		play1 = int(players_online[indexOfPlayers(player1)].jugada)
		play2 = int(players_online[indexOfPlayers(player2)].jugada)
		if math.Abs(float64(bossNumber - play1)) > math.Abs(float64(bossNumber - play2)) {
			RemovePlayer(indexOfPlayers(player1))
		} else if math.Abs(float64(bossNumber - play1)) < math.Abs(float64(bossNumber - play2)) {
			RemovePlayer(indexOfPlayers(player1))
		}
	}

	fmt.Println("Terminamos el for, y se removieron los malos!")

	fmt.Println("Escribe C para continuar: ")
	fmt.Scanln(&empezar)
	// Se liberan los procesos, para que puedan preguntar si siguen vivos...
	started = true

	fmt.Println("Escribe C para continuar: ")
	fmt.Scanln(&empezar)

	// Se termina el programa, luego de que todos hayan preguntado
	// anunciar el ganador
	// Y ver weas del pozo / mandar plata

}

