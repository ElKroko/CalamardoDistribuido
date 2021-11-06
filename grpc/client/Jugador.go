package main

import (
	"context"
	"fmt"
	pb "lab/game/proto"
	"log"
	"math/rand"
	"time"
	"google.golang.org/grpc"
	"strconv"
)

type PlayerStruct struct {
	id    int32
	alive bool
	round int32
	score int32
	etapa int32
}

func Juego1(jugador bool, player PlayerStruct) int {
	fmt.Println("Hola!")
	rand.Seed(time.Now().UnixNano())

	suma := player.score
	var numero int
	gano := 0
	ronda := player.round

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

		res, err := serviceClient.StartGame(context.Background(), &pb.StartRequest{Message: "Puedo jugar?"})

		if err != nil {
			panic("No pudimos chequear si esta inciado  " + err.Error())
		}

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

	// Decidir si sera Bot o Jugador
	var jugador bool

	fmt.Println("Es usted un jugador? [1] Si [0] no")
	fmt.Print("-> ")
	fmt.Scanln(&jugador)


	// Definicion de Variables

	iniciado := false


	conn, err := grpc.Dial("localhost:8080", grpc.WithInsecure()) // Conectamos al IP de 10.6.43.109:8080, el lider.

	if err != nil {
		panic("cannot connect with server " + err.Error())
	}

	serviceClient := pb.NewCalamardoGameClient(conn)

	// Inscripcion del jugador

	message := "Hola, soy jugador"

	res, err := serviceClient.JoinGame(context.Background(), &pb.JoinRequest{Message: message})

	if err != nil {
		panic("No se pudo anadir el jugador  " + err.Error())
	}

	jugador_struct := PlayerStruct{res.GetIdJugador(), res.GetAlive(), res.GetRound(), 0, 0}

	iniciado = true

	log.Printf("")
	log.Printf(strconv.Itoa(int(jugador_struct.id)))
	log.Printf(strconv.FormatBool(jugador_struct.alive))
	log.Printf(strconv.Itoa(int(jugador_struct.round)))
	log.Printf(strconv.Itoa(int(jugador_struct.score)))



	if iniciado {
		fmt.Println("Bienvenidx al juego del Calamardo!")
	}

	//actualStage = stage1

	flag1 := true
	var continuar string

	for flag1 {
		fmt.Println("Tamos aqui")

		res1, err := serviceClient.StartGame(context.Background(), &pb.StartRequest{Message: "Puedo jugar?"})
		if err != nil {
			panic("No pudimos chequear si esta inciado  " + err.Error())
		}

		if res1.GetStarted(){
			flag1 = false
			fmt.Println("Preparado? empezamos!")
		} else {
			fmt.Println("Espera, aun no hay suficientes jugadores...")
			fmt.Println("Desea reintentar? [1] Si [0] no")
			fmt.Print("-> ")
			fmt.Scanln(&continuar)
		}
	}
	
	// comienza juego 1

	// Se calcula el numero a mandar
	// Se manda el numero, y se espera respuesta

	


}
