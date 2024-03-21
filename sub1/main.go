package main

import (
	"fmt"
	"log"
	"time"

	beanstalk "github.com/beanstalkd/go-beanstalk"
)

const (
	tubeName1  = "testTube1"
	serverAddr = "localhost:11300"
)

func main() {
	conn, err := beanstalk.Dial("tcp", serverAddr)
	if err != nil {
		fmt.Println("Error connecting to Beanstalkd:", err)
		return
	}
	defer conn.Close()

	log.Println("Subscribe 1 connected to beanstalkd")

	tubeSet := beanstalk.NewTubeSet(conn, tubeName1)
	for {
		id, body, err := tubeSet.Reserve(10 * time.Second)
		if err != nil {
			if err == beanstalk.ErrTimeout {
				continue // No message available, retry
			}
			fmt.Println("Error reserving message:", err)
		} else {
			log.Printf("Subscriber 1 received msg : %s\n", body)
			conn.Delete(id)
		}
	}
}
