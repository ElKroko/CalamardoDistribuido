package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"time"

	pb "lab/game/proto"

	"google.golang.org/grpc"
)

type server struct {
	pb.UnimplementedCalamardoGameServer
}

type PlayerStruct struct {
	id    string
	alive bool
	score int
}

// Variables
var totalPlayers int								// Jugadores actuales
var players_online []PlayerStruct
var started bool
var total int = 3 									// Aqui se define el total de jugadores maximos!


func (s *server) JoinGame(ctx context.Context, in *pb.JoinRequest) (*pb.JoinReply, error) {
	log.Printf("Recibido: %s", in.GetMessage())
	totalPlayers += 1
	players_online = append(players_online, PlayerStruct{in.GetPlayer(), true, 0})

	log.Printf("Jugador %s agregado", in.GetPlayer())
	return &pb.JoinReply{Stage1: "1", Stage2: "2", Stage3: "3"}, nil
}

func (s *server) StartGame(ctx context.Context, in *pb.StartRequest) (*pb.StartReply, error) {
	log.Printf("Recibido: %s", in.GetMessage())
	if len(players_online) != total {
		return &pb.StartReply{Started: false}, nil
	}
	return &pb.StartReply{Started: true}, nil
}

func (s *server) Juego1msg(ctx context.Context, in *pb.Juego1Request, muere bool) (*pb.Juego1Reply, error) {
	log.Printf("Iniciamos el primer juego")

	log.Printf("La jugada recibida fue: %s", in.GetJugada())

	return &pb.Juego1Reply{Muere: muere}, nil
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
	//Calamardo := ""
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

	for i := 0; i < 16; i++ {
		fmt.Println(players_online[i].id)
	}




}
