package agent

import (
	"bytes"
	"io"
	"log"
	"net"
)

var (
	server = "127.0.0.1:9876"
)

type Sender struct {
	//	buffer chan bytes.Buffer
	client TcpClient
}

func NewSender() *Sender {
	TcpClient := NewTcpClient(server)

	return &Sender{
		client: TcpClient,
		// buffer: make(chan bytes.Buffer 200)
	}

}

func (s *Sender) SendData(data string) {
	s.client.SendReportPacket(data)
	log.Print("send data success!")
}
