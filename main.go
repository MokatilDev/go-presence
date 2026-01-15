package main

import (
	"os"
	"os/signal"
	"time"

	"github.com/MokatilDev/go-presence/client"
	"github.com/MokatilDev/go-presence/utils"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	utils.CheckErr(err)

	c := make(chan os.Signal, 1)

	activity := client.Activity{
		Type:          client.Playing,
		StatusDisplay: client.Name,
		Timestamps: &client.ActivityTimestamps{
			Start: int(time.Now().Unix()),
		},
	}

	client_id := os.Getenv("CLIENT_ID")

	err = client.Login(client_id)
	utils.CheckErr(err)

	err = client.SetActivity(activity)
	utils.CheckErr(err)

	signal.Notify(c, os.Interrupt)
	<-c

	client.Logout()
}
