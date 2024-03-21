package main

import (
	"fmt"
	"log"
	"time"

	beanstalk "github.com/beanstalkd/go-beanstalk"
)

const (
	tubeName1  = "testTube1"
	tubeName2  = "testTube2"
	serverAddr = "localhost:11300"
	message    = "Hello, Beanstalkd!"
)

func main() {
	conn, err := beanstalk.Dial("tcp", serverAddr)
	if err != nil {
		fmt.Println("Error connecting to Beanstalkd:", err)
		return
	}
	defer conn.Close()

	log.Println("Publisher is running")

	tube1 := beanstalk.Tube{Conn: conn, Name: tubeName1}
	tube2 := beanstalk.Tube{Conn: conn, Name: tubeName2}
	counter := 1

	for {
		msg := fmt.Sprintf("[%v] %v", counter, message)
		_, err := tube1.Put([]byte(msg), 1, 0, 0)
		if err != nil {
			fmt.Println("Error publishing message to tube1:", err)
		} else {
			log.Println("Published message to tube1:", msg)
		}

		_, err = tube2.Put([]byte(msg), 1, 0, 0)
		if err != nil {
			fmt.Println("Error publishing message to tube2:", err)
		} else {
			log.Println("Published message to tube2:", msg)
		}

		counter++
		time.Sleep(5 * time.Second)
	}
}
