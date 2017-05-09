package agent

import (
	"bufio"
	"crypto/rand"
	"encoding/json"
	"fmt"
	"net"
	"os"
	"sync"
	"time"
)

//数据包类型
const (
	HEART_BEAT_PACKET = 0x00
	REPORT_PACKET     = 0x01
)

//数据包
type Packet struct {
	PacketType    byte
	PacketContent []byte
}

//心跳包
type HeartPacket struct {
	Version   string `json:"version"`
	Timestamp int64  `json:"timestamp"`
}

//数据包
type ReportPacket struct {
	Content   string `json:"content"`
	Rand      int    `json:"rand"`
	Timestamp int64  `json:"timestamp"`
}

//客户端对象
type TcpClient struct {
	connection *net.TCPConn
	hawkServer *net.TCPAddr
}

func NewTcpClient(server string) *TcpClient {

	hawkServer, err := net.ResolveTCPAddr("tcp", server)
	if err != nil {
		fmt.Printf("hawk server [%s] resolve error: [%s]", server, err.Error())

	}

	connection, err := net.DialTCP("tcp", nil, hawkServer)
	if err != nil {
		fmt.Printf("connect to hawk server error: [%s]", err.Error())

	}
	return &TcpClient{
		connection: connection,
		hawkServer: hawkServer,
	}

}

//发送数据包
func (client *TcpClient) SendReportPacket(data string) {
	reportPacket := ReportPacket{
		Content:   data,
		Timestamp: time.Now().Unix(),
		Rand:      rand.Int(),
	}
	packetBytes, err := json.Marshal(reportPacket)
	if err != nil {
		fmt.Println(err.Error())
	}

	packet := Packet{
		PacketType:    REPORT_PACKET,
		PacketContent: packetBytes,
	}
	sendBytes, err := json.Marshal(packet)
	if err != nil {
		fmt.Println(err.Error())
	}
	//发送
	client.connection.Write(EnPackSendData(sendBytes))
	fmt.Println("Send metric data success!")
}

//使用的协议与服务器端保持一致
func EnPackSendData(sendBytes []byte) []byte {
	packetLength := len(sendBytes) + 8
	result := make([]byte, packetLength)
	result[0] = 0xFF
	result[1] = 0xFF
	result[2] = byte(uint16(len(sendBytes)) >> 8)
	result[3] = byte(uint16(len(sendBytes)) & 0xFF)
	copy(result[4:], sendBytes)
	sendCrc := crc32.ChecksumIEEE(sendBytes)
	result[packetLength-4] = byte(sendCrc >> 24)
	result[packetLength-3] = byte(sendCrc >> 16 & 0xFF)
	result[packetLength-2] = 0xFF
	result[packetLength-1] = 0xFE
	fmt.Println(result)
	return result
}

//发送心跳包
func (client *TcpClient) sendHeartPacket() {
	heartPacket := HeartPacket{
		Version:   "1.0",
		Timestamp: time.Now().Unix(),
	}
	packetBytes, err := json.Marshal(heartPacket)
	if err != nil {
		fmt.Println(err.Error())
	}
	packet := Packet{
		PacketType:    HEART_BEAT_PACKET,
		PacketContent: packetBytes,
	}
	sendBytes, err := json.Marshal(packet)
	if err != nil {
		fmt.Println(err.Error())
	}
	client.connection.Write(EnPackSendData(sendBytes))
	fmt.Println("Send heartbeat data success!")
}
