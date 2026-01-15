package ipc

import (
	"bytes"
	"encoding/binary"
	"errors"
	"fmt"
	"net"
	"os"
	"time"

	"github.com/MokatilDev/go-presence/utils"
	"github.com/elastic/go-sysinfo"
)

var socket net.Conn

func GetIpcPath() (path string, err error) {
	host, err := sysinfo.Host()
	utils.CheckErr(err)

	OSType := host.Info().OS.Type

	for i := range 10 {
		for _, path := range PathFormats[OSType] {
			fullPath := os.ExpandEnv(fmt.Sprintf(path+"%d", i))
			if _, err := os.Stat(fullPath); err == nil {
				return fullPath, nil
			}
		}
	}

	return "", errors.New("Discord IPC not found")
}

func OpenSocket() error {
	path, err := GetIpcPath()
	if err != nil {
		return err
	}

	s, err := net.DialTimeout("unix", path, 5*time.Second)
	if err != nil {
		return nil
	}

	socket = s
	return nil
}

func CloseSocket() error {
	if socket != nil {
		socket.Close()
		socket = nil
	}
	return nil
}

func Read() string {
	buf := make([]byte, 1024)

	payloadLength, err := socket.Read(buf)
	utils.PrintErr(err)

	buffer := new(bytes.Buffer)

	for i := 8; i < payloadLength; i++ {
		buffer.WriteByte(buf[i])
	}

	return buffer.String()
}

func Send(opcode int, payload string) string {
	buf := new(bytes.Buffer)

	err := binary.Write(buf, binary.LittleEndian, int32(opcode))
	utils.PrintErr(err)

	err = binary.Write(buf, binary.LittleEndian, int32(len(payload)))
	utils.PrintErr(err)

	buf.Write([]byte(payload))
	_, err = socket.Write(buf.Bytes())
	utils.PrintErr(err)

	return Read()
}
