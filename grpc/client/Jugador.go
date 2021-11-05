package main

import (
	"context"
	"fmt"
	pb "lab/game/proto"
	"log"
	"strconv"

	"google.golang.org/grpc"
)

type PlayerStruct struct {
	id    string
	alive bool
	round int32
	score int32
}

func main() {

	// Definicion de Variables
	//actualStage := "none"
	stage1 := "none"
	stage2 := "none"
	stage3 := "none"

	iniciado := false

	var lista_de_jugadores []PlayerStruct

	conn, err := grpc.Dial("localhost:8080", grpc.WithInsecure()) // Conectamos al IP de 10.6.43.109:8080, el lider.

	if err != nil {
		panic("cannot connect with server " + err.Error())
	}

	serviceClient := pb.NewCalamardoGameClient(conn)

	// Inscripcion del jugador

	player := "Jugador 1"
	message := "Hola, soy jugador"

	res, err := serviceClient.JoinGame(context.Background(), &pb.JoinRequest{Player: player, Message: message})

	if err != nil {
		panic("No se pudo anadir el jugador  " + err.Error())
	}

	lista_de_jugadores = append(lista_de_jugadores, PlayerStruct{player, true, 1, 0})

	stage1 = res.GetStage1()
	stage2 = res.GetStage2()
	stage3 = res.GetStage3()

	//actualStage = stage1
	iniciado = true

	log.Printf(stage1)
	log.Printf(stage2)
	log.Printf(stage3)

	if iniciado {
		fmt.Println("Bienvenidx al juego del Calamardo!")
	}


	// Inscribir 15 bots!

	for i := 1; i < 16; i++ {
		lista_de_jugadores = append(lista_de_jugadores, PlayerStruct{strconv.Itoa(i + 1), true, 1, 0})

		conn, err := grpc.Dial("localhost:8080", grpc.WithInsecure()) // Conectamos al IP de 10.6.43.109:8080, el lider.
		if err != nil {
			panic("cannot connect with server " + err.Error())
		}

		serviceClient := pb.NewCalamardoGameClient(conn)

		res, err := serviceClient.JoinGame(context.Background(), &pb.JoinRequest{Player: lista_de_jugadores[i].id, Message: "Soy Bot"})

		if err != nil {
			log.Fatalf("could not greet: %v", err)
		}

		log.Printf("Bot %d inscrito", i)

		lista_de_jugadores[i].id = strconv.Itoa(i + 1)
		lista_de_jugadores[i].alive = true

		stage2 = res.GetStage2()


	}

}
