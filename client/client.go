package client

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/MokatilDev/go-presence/ipc"
	"github.com/MokatilDev/go-presence/utils"
	"github.com/google/uuid"
)

var (
	logged bool
	nonce  string
)

func Login(clientId string) error {
	if !logged {
		payload, err := json.Marshal(&HandShake{"1", clientId})
		if err != nil {
			return err
		}

		err = ipc.OpenSocket()
		if err != nil {
			return err
		}

		ipc.Send(0, string(payload))
	}

	logged = true
	SetNonce()

	return nil
}

func Logout() {
	logged = false
	err := ipc.CloseSocket()
	utils.PrintErr(err)
}

func SetActivity(activity Activity) error {
	if !logged {
		return nil
	}

	payload := Packet{
		Cmd: "SET_ACTIVITY",
		Args: Args{
			Pid:      os.Getpid(),
			Activity: activity,
		},
		Nonce: nonce,
	}

	p, err := json.Marshal(payload)
	if err != nil {
		return err
	}

	res := ipc.Send(1, string(p))
	fmt.Println("response form discord client:\n", res)

	return nil
}

func SetNonce() {
	if nonce == "" {
		nonce = uuid.New().String()
	}
}
