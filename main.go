package main

import (
	"log"
	"os"
	"os/signal"
	"strconv"
	"time"

	client "github.com/MokatilDev/go-presence/ipc/client"
	"github.com/MokatilDev/go-presence/ipc/platform"
	"github.com/MokatilDev/go-presence/rpc"
	"github.com/google/uuid"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	c := make(chan os.Signal, 1)
	clientID := os.Getenv("CLIENT_ID")

	id, err := strconv.Atoi(clientID)
	if err != nil {
		log.Fatal("Invalid client id")
	}

	connection, err := platform.NewUnixConnection(int64(id))
	client, err := client.NewClientInstance(int64(id), connection)
	if err != nil {
		log.Fatal(err.Error())
	}

	err = client.Handshake()
	if err != nil {
		log.Println(err)
	}

	err = client.Update(&rpc.Packet{
		Cmd: "SET_ACTIVITY",
		Args: rpc.Args{
			Pid: os.Getpid(),
			Activity: &rpc.Activity{
				Type:          rpc.Playing,
				StatusDisplay: rpc.Name,
				Timestamps: &rpc.ActivityTimestamps{
					Start: int(time.Now().Unix()),
				},
			},
		},
		Nonce: uuid.New().String(),
	})

	signal.Notify(c, os.Interrupt)
	<-c

	client.Close()
}
