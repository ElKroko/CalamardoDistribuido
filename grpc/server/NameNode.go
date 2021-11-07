package main

import ("math/rand"
		"time"
		"log"
		"strconv"
		"os"
)

func (s *server) SearchPlayerPlays(ctx context.Context, in *pb.JoinRequest) (*pb.JoinReply, error) {
	log.Printf("Recibido: %s", in.GetMessage())

    f, err := os.Open("lugarJugador.txt")
    var splitText string
    if err != nil {
        log.Fatal(err)
    }
    defer f.Close()
	scanner := bufio.NewScanner(f)
    for scanner.Scan() {
		splitText = strings.Split(scanner.Text(), " ")
		if splitText[0] == "Jugador_" + player {
        	llamarDataNode(splitText[0][-1], splitText[1][-1], splitText[2])
		}
    }

	return &pb.JoinReply{IdJugador: int32(totalPlayers-1), Alive: true, Round: 0}, nil
}

func main() {
	rand.Seed(time.Now().UnixNano())
	f, err := os.Create("lugarJugadores.txt")
    if err != nil {
        log.Fatal(err)
    }
    defer f.Close()
	next := false
	var index int64 = 0
	var ip string
	for next == false {
		//esperar datos
		jugador, ronda := 1, 2
		selectDataNode := rand.Intn(3)
		if selectDataNode == 0 {
			ip = "10.6.43.109"
		} else if selectDataNode == 1 {
			ip = "10.6.43.110"
		} else {
			ip = "10.6.43.111"
		}
		val := "Jugador_" + strconv.Itoa(jugador) + " Ronda_" + strconv.Itoa(ronda) + " " + ip + "\n"
		data := []byte(val)
		_, err := f.WriteAt(data, index)
		index = int64(len(data))+index
		if err != nil {
			log.Fatal(err)
		}
		//mandar jugador y ronda a ip
	}
	
}