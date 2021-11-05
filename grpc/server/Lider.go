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



func (s *server) JoinGame (ctx context.Context, in *pb.JoinRequest) (*pb.JoinReply, error) {
	log.Printf("Recibido: %s", in.GetMessage())
	totalPlayers +=1
	players_online = append(players_online, PlayerStruct{in.GetPlayer(), true, 0})

	log.Printf("Jugador agregado")
	fmt.Println("Hola")
	// fmt.Printf("%v", players_online)
	return &pb.JoinReply{Unido: "si"}, nil
}

func main() {
	listner, err := net.Listen("tcp", ":8080")

	if err != nil {
		panic("cannot create tcp connection " + err.Error())
	}
	

	serv := grpc.NewServer()
	pb.RegisterCalamardoGameServer(serv, &server{})
	if err = serv.Serve(listner); err != nil {
		panic("cannot initialize the server" + err.Error())
	}

}
