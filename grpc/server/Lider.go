package main

import (
	"context"
	"fmt"
	"log"
	"net"

	pb "lab/game/proto"
	"google.golang.org/grpc"
)

type server struct {
	pb.UnimplementedCalamardoGameServer
}

type PlayerStruct struct {
	id	string
	alive bool
	score int
}

// Variables
var totalPlayers int
var players_online []PlayerStruct
var started bool



func (s *server) JoinGame (ctx context.Context, in *pb.JoinRequest) (*pb.JoinReply, error) {
	log.Printf("Recibido: %s", in.GetMessage())
	totalPlayers +=1
	players_online = append(players_online, PlayerStruct{in.GetPlayer(), true, 0})

	log.Printf("Jugador %s agregado", in.GetPlayer())
	return &pb.JoinReply{Stage1: "1", Stage2: "2", Stage3: "3"}, nil
}


func (s *server) StartGame (ctx context.Context, in *pb.StartRequest) (*pb.StartReply, error) {
	return &pb.StartReply{Started: started}, nil
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
	Calamardo := ""

	var empezar string

	for totalPlayers != 16 {
		fmt.Println("escribe C para iniciar el Calamardo: ")
		fmt.Scanln(&Calamardo)

		if totalPlayers != 16{
			fmt.Println("Aun no hay suficientes calamares...")
		}
	}

	if totalPlayers == 16{
		fmt.Println("Escribe C para comenzar el primer juego: ")
		fmt.Scanln(&empezar)
		
		for i := 0; i < 16; i++ {
			fmt.Println(players_online[i].id)
		}

	}

}
