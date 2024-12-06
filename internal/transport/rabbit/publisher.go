package rabbit

import (
	"fmt"
	"testing/rabbitmq/configs"

	"github.com/streadway/amqp"
)



func PrintEnv() {
    rabbitConfig := configs.NewRabbitConfig()
    fmt.Printf("os env print: %s", rabbitConfig.AmqpServerURL)
    
}

func PublishMessage (message string) {

    rabbitConfig := configs.NewRabbitConfig()
    fmt.Println(rabbitConfig.AmqpServerURL)
    // Create a new RabbitMQ connection.
    connectRabbitMQ, err := amqp.Dial(rabbitConfig.AmqpServerURL)
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

    rabbitMessage := amqp.Publishing{
        ContentType: "text/plain",
        Body:        []byte(message),
    }

    if err := channelRabbitMQ.Publish(
        "",              // exchange
        "QueueService1", // queue name
        false,           // mandatory
        false,           // immediate
        rabbitMessage,         // message to publish
    ); err != nil {
        
    }
    
}