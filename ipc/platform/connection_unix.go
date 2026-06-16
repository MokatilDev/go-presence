package platform

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net"
	"os"
	"path/filepath"
	"strconv"

	"github.com/MokatilDev/go-presence/rpc"
	"github.com/MokatilDev/go-presence/utils"
	"github.com/google/uuid"
)

type UnixConnection struct {
	socket net.Conn
}

func NewUnixConnection(clientID int64) (*UnixConnection, error) {
	dirs := []string{"/tmp"}
	envVars := []string{"XDG_RUNTIME_DIR", "TMPDIR", "TMP", "TEMP"}

	for _, env := range envVars {
		if v := os.Getenv(env); v != "" {
			dirs = append(dirs, v)
		}
	}

	for _, dir := range dirs {
		for i := 0; i < 10; i++ {
			socketPath := filepath.Join(dir, fmt.Sprintf("discord-ipc-%d", i))
			socket, err := net.Dial("unix", socketPath)

			if err == nil {
				return &UnixConnection{
					socket: socket,
				}, nil
			}
		}
	}

	return nil, errors.New("Unable to connect with discord")
}

func (conn *UnixConnection) Write(opcode int32, data []byte) error {
	if conn.socket == nil {
		return errors.New("socket not found")
	}

	payload := utils.Encode(opcode, int32(len(data)))
	payload = append(payload, data...)

	_, err := conn.socket.Write(payload)

	return err
}

func (conn *UnixConnection) Read() ([]byte, error) {
	if conn.socket == nil {
		return nil, errors.New("socket not found")
	}

	header := make([]byte, 8)
	_, err := io.ReadFull(conn.socket, header)
	if err != nil {
		return nil, err
	}

	size := utils.Decode(header)
	buffer := make([]byte, size)

	_, err = io.ReadFull(conn.socket, buffer)
	if err != nil {
		return nil, err
	}

	return buffer, nil
}

func (conn *UnixConnection) Close() error {
	if conn.socket == nil {
		return errors.New("socket not found")
	}

	err := conn.socket.Close()
	if err != nil {
		return err
	}

	return nil
}

func (conn *UnixConnection) Handshake(clientID int64) error {
	if conn.socket == nil {
		return errors.New("socket not found")
	}

	handshake := rpc.HandShake{
		V:        "1",
		ClientID: strconv.Itoa(int(clientID)),
	}

	payload, err := json.Marshal(&handshake)
	if err != nil {
		return err
	}

	return conn.Write(0, payload)
}

func (conn *UnixConnection) Update(packet *rpc.Packet) error {
	if conn.socket == nil {
		return errors.New("socket not found")
	}

	data, err := json.Marshal(packet)
	if err != nil {
		return err
	}

	err = conn.Write(1, data)
	if err != nil {
		return err
	}

	return nil
}

func (conn *UnixConnection) Clear() error {
	if conn.socket == nil {
		return errors.New("socket not found")
	}

	packet := rpc.Packet{
		Cmd: "SET_ACTIVITY",
		Args: rpc.Args{
			Pid:      os.Getpid(),
			Activity: nil,
		},
		Nonce: uuid.New().String(),
	}

	data, err := json.Marshal(packet)
	if err != nil {
		return err
	}

	err = conn.Write(1, data)
	if err != nil {
		return err
	}

	return nil
}
