package main

import (
  "log"
  "fmt"
  "github.com/streadway/amqp"
)


func failOnError(err error, msgerr string, msgsuc string) {
        if err != nil {
                log.Fatalf("%s: %s",msgerr,err);
                
        } else {
                fmt.Printf("%s\n",msgsuc)
        }

}

func main() {

	// Connect to RabbitMQ server
    fmt.Println("Connecting to RabbitMQ ...")
    conn, err := amqp.Dial(<connection string>) //Insert the  connection string
 	failOnError(err, "RabbitMQ connection failure", "RabbitMQ Connection Established")
    defer conn.Close()

    //Connect to the channel

    ch,err := conn.Channel()
    failOnError(err, "Failed to open a channel", "Opened the channel")
    defer ch.Close()

    //Declare the queue where messages need to be sent. Queue will be created if not already there
    q, err := ch.QueueDeclare(
    	"DemoQueue", //name
    	true, //durable
    	false, //delete when unused
    	false, //exclusive
    	false, //no-wait
    	nil, //arguements
    )

    failOnError(err, "Failed to declare the queue", "Declared the queue")


    body := "Hello Tony!"

    //Publish to the queue

    err = ch.Publish(
    	"", //exchange
    	q.Name, //routing key
    	false, //mandatory
    	false, //immediate
    	amqp.Publishing {
    		ContentType: "text/plain",
    		Body: []byte(body),
    })
        
    failOnError(err, "Failed to publish a message ","Published the message")
      
    	
}
