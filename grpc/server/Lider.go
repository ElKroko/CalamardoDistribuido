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
)

type server struct {
	pb.UnimplementedCalamardoGameServer
}

type PlayerStruct struct {
	id    int32
	alive bool
	score int32
}

// Variables
var totalPlayers int								// Jugadores actuales
var players_online []PlayerStruct
var started bool
var total int = 3 									// Aqui se define el total de jugadores maximos!
var playsBoss []int

var etapaActual int




// Funciones de mensajeria
func (s *server) JoinGame(ctx context.Context, in *pb.JoinRequest) (*pb.JoinReply, error) {
	log.Printf("Recibido: %s", in.GetMessage())
	
	players_online = append(players_online, PlayerStruct{int32(totalPlayers), true, 0})
	totalPlayers += 1

	return &pb.JoinReply{IdJugador: int32(totalPlayers-1), Alive: true, Round: 0}, nil
}

func (s *server) StartGame(ctx context.Context, in *pb.StartRequest) (*pb.StartReply, error) {
	log.Printf("Recibido: %s", in.GetMessage())
	if len(players_online) != total {
		return &pb.StartReply{Started: false}, nil
	}
	return &pb.StartReply{Started: true}, nil
}

// Para el juego 1: 
// El lider debe esperar la jugada del jugador
// Al momento de recibirla, necesitamos calcular si este numero cumple las condiciones (>= y 21)
// 		Si es >=, eliminar el jugador
// 		Si es 21>, avisar que gano y cambiar su ronda



// Funcion para enviar y recibir jugadas desde el lider.

func (s *server) JuegoMsg(ctx context.Context, in *pb.JuegoRequest) (*pb.JuegoReply, error) {
	vivo := true
	etapa_jugador := in.GetEtapa()

	if etapaActual == 1 && etapa_jugador == 1{
		ronda := in.GetRound()
		id_jugador := in.GetId()
		log.Printf("El Jugador ", id_jugador, "en la ronda ", ronda )
		playPlayer := in.GetJugada()
		log.Printf("Ha jugado la carta",playPlayer)

		suma := in.GetScore()

		if playPlayer >= int32(playsBoss[ronda]) {
			RemovePlayer(indexOf(id_jugador))
			etapa_jugador = -1
			vivo = false
		} else {
			suma = suma + playPlayer
			if suma >= 21 {
				etapa_jugador += 1
			} else if ronda == 3 {
				RemovePlayer(indexOf(id_jugador))
			} else {
				ronda = ronda + 1
			}
		}
		return &pb.JuegoReply{Alive: vivo, Round: ronda, Score: suma, Etapa: etapa_jugador}, nil
	} else{
		return &pb.JuegoReply{Alive: vivo, Round: 0, Score: 0, Etapa: etapa_jugador}, nil
	}
	
}


// Funciones del Juego

// El lider escoge sus cartas e inicializa el array
func juego_1() {
	fmt.Println("Hola!, el lider eligira sus cartas")

	var numero int
	//gano := 0
	ronda := 0

	for ronda < 4 {
		fmt.Println("Elija un numero del 6 al 10")
		fmt.Print("-> ")
		fmt.Scanln(&numero)
		playsBoss = append(playsBoss, numero)
		ronda = ronda + 1
	}

	etapaActual = 1
	return 

}

// Funciones Auxiliares
func indexOf(element int32) (int) {
	for k, v := range players_online {
		if element == v.id {
			return k
		}
	}
	return -1    //not found.
 }

func RemovePlayer(i int){
	value := players_online[i].id 
    players_online[i] = players_online[len(players_online)-1]
	fmt.Printf("El jugador_%d fue eliminado\n",value)
	//avisar jugador
	//avisar a Namenode
}



func main() {

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
		time.Sleep(5 * time.Second)

		if totalPlayers != total {
			fmt.Println("Aun no hay suficientes calamares... Total:", totalPlayers, "\t esperamos:", total)
		}
	}


	fmt.Println("Escribe C para comenzar el primer juego: ")
	fmt.Scanln(&empezar)

	for i := 0; i < total; i++ {
		fmt.Println(players_online[i].id)
	}




}
