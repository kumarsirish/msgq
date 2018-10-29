package main

import (
        "fmt"
        "github.com/streadway/amqp"
        "log"
        "time"
)

func failOnError(err error, msgerr string, msgsuc string) {
        if err != nil {
                log.Fatalf("%s: %s",msgerr,err);
                
        } else {
                fmt.Printf("%s\n",msgsuc)
        }

}

func main() {
        
        fmt.Println("Connecting to RabbitMQ")
        conn, err := amqp.Dial(<connection string>) //Insert the  connection string
        failOnError(err,"Failed to connect to RabbitMQ server", "Successfully connected to RabbitMQ server")
        defer conn.Close()
        failOnError(err, "RabbitMQ connection failure", "RabbitMQ Connection Established")

        ch, err := conn.Channel()
        failOnError(err, "Failed to open a channel", "Opened the channel")
        defer ch.Close()

        q, err  := ch.QueueDeclare(
                "DemoQueue", //name
                //"ha.monitoring",
                true,
                false, //delete when unused
                false, //exclusive
                false, //no-wait
                nil, //arguements
        )
        failOnError(err, "Failed to declare the queue", "Declared the queue")


        msgs,err := ch.Consume(
                q.Name, //queue
                "", //consumer
                true, //auto-ack
                false, //exclusive
                false, //no-local
                false, //no-wait
                nil, //args
        )
        failOnError(err, "Failed to register a consumer ","Registered the consumer")


    msgCount :=0
        go func() {
                for d := range msgs {

                        msgCount++

                        fmt.Printf("\nMessage Count: %d, Message Body: %s\n", msgCount, d.Body)

                }
        }()

        select {
            case <-time.After(time.Second * 10):
                fmt.Printf("Total Messages Fetched: %d\n",msgCount)
                fmt.Println("No more messages in queue. Timing out...")
               
        }


}
