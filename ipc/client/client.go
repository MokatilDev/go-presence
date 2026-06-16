package client

import (
	"errors"

	"github.com/MokatilDev/go-presence/rpc"
)

type Connection interface {
	Write(opcode int32, data []byte) error
	Read() ([]byte, error)
	Close() error
	Handshake(clientID int64) error
	Update(packet *rpc.Packet) error
	Clear() error
}

type RichClient struct {
	ClientID     int64         `json:"client_id"`
	Conn         Connection    `json:"-"`
	LastActivity *rpc.Activity `json:"last_activity"`
}

func NewClientInstance(clientID int64, conn Connection) (*RichClient, error) {
	if conn == nil {
		return nil, errors.New("Invalid connection")
	}
	return &RichClient{
		ClientID:     clientID,
		Conn:         conn,
		LastActivity: nil,
	}, nil
}

func (c *RichClient) Update(packet *rpc.Packet) error {
	c.LastActivity = packet.Args.Activity
	return c.Conn.Update(packet)
}

func (c *RichClient) Close() error {
	return c.Conn.Close()
}

func (c *RichClient) Handshake() error {
	return c.Conn.Handshake(c.ClientID)
}
