package main

import (
	"context"
	"fmt"
	pb "lab/game/proto"
	"google.golang.org/grpc"
	"math/rand"
	"strconv"
	"time"
)

func generateID() string {
	rand.Seed(time.Now().Unix())
	return "ID: " + strconv.Itoa(rand.Int())
}

func main() {
	conn, err := grpc.Dial("10.6.43.109:8080", grpc.WithInsecure())

	if err != nil {
		panic("cannot connect with server " + err.Error())
	}

	serviceClient := pb.NewCalamardoGameClient(conn)

	player := "Jugador 1"
	message := "Hola, quiero entrar al juego"

	res, err := serviceClient.JoinGame(context.Background(), &pb.JoinRequest{Player: player, Message: message})

	if err != nil {
		panic("No se pudo anadir el jugador  " + err.Error())
	}

	fmt.Println(res.Unido)

}
