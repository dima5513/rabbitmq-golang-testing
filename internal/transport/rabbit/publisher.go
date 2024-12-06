package main

import (
	"log"
	"os"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/streadway/amqp"
)


func main () {
    amqpServerURL := os.Getenv("AMQP_SERVER_URL")


    // Create a new RabbitMQ connection.
    connectRabbitMQ, err := amqp.Dial(amqpServerURL)
    if err != nil {
        panic(err)
    }
    defer connectRabbitMQ.Close()

    // Let's start by opening a channel to our RabbitMQ
    // instance over the connection we have already
    // established.
    channelRabbitMQ, err := connectRabbitMQ.Channel()
    if err != nil {
        panic(err)
    }
    defer channelRabbitMQ.Close()



     // With the instance and declare Queues that we can
    // publish and subscribe to.
    _, err = channelRabbitMQ.QueueDeclare(
        "QueueService1", // queue name
        true,            // durable
        false,           // auto delete
        false,           // exclusive
        false,           // no wait
        nil,             // arguments
    )
    if err != nil {
        panic(err)
    }


    app := echo.New()

    app.Use(middleware.Logger())


     // Add route.
     app.GET("/send", func(c echo.Context) error {
      
        //  Create a message to publish.
         message := amqp.Publishing{
            ContentType: "text/plain",
            Body:        []byte(c.QueryParam("msg")),
        }

        // Attempt to publish a message to the queue.
        if err := channelRabbitMQ.Publish(
            "",              // exchange
            "QueueService1", // queue name
            false,           // mandatory
            false,           // immediate
            message,         // message to publish
        ); err != nil {
            return err
        }
        return nil
     })

    // Start Fiber API server.
    log.Fatal(app.Start(":3000"))
}