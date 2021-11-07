package main

import (
  "log"
  "fmt"
  amqp "github.com/streadway/amqp"
)

func failOnError(err error, msg string) {
  if err != nil {
    log.Fatalf("%s: %s", msg, err)
  }
}

func main() {

	conn, err := amqp.Dial("amqp://test:test@10.6.43.110:5672/")
	failOnError(err, "Failed to connect to RabbitMQ")
	fmt.Println(ActualAmmount())("¿?")
	defer conn.Close()

	fmt.Println(ActualAmmount())("Conecto")

	ch, err := conn.Channel()
	failOnError(err, "Failed to open a channel")
	defer ch.Close()
	fmt.Println(ActualAmmount())("Conecto channel")

	q, err := ch.QueueDeclare(
		"hello", // name
		false,   // durable
		false,   // delete when unused
		false,   // exclusive
		false,   // no-wait
		nil,     // arguments
	)
	failOnError(err, "Failed to declare a queue")
	fmt.Println(ActualAmmount())("Declaro Qeue")

	i := "1"
	s := "2"

	body := "Jugador_" + i + " Ronda_" + s

	err = ch.Publish(
		"",     // exchange
		q.Name, // routing key
		false,  // mandatory
		false,  // immediate
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        []byte(body),
		})
	failOnError(err, "Failed to publish a message")
	fmt.Println(ActualAmmount())("Ha muerto: %s ", body)
}