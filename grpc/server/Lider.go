package main

import (
	"context"
	"fmt"
	"log"
	"net"

	pb "github.com/ElKroko/CalamardoDistribuido/tree/main/grpc/proto"
	"google.golang.org/grpc"
)

type server struct {
	pb.UnimplementedWishListServiceServer
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
	// fmt.Printf("%v", players_online)
	return &pb.JoinReply{Codes1: "1rv", Codes2: "2tc", Codes3: "3tn"}, nil
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
