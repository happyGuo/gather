package agent

import (
	"log"
	"net"
)

type SingleConnection struct {
	addr        string
	bakAddr     string
	currentAddr string
	conn        *net.TCPConn

	failedTimes int
	maxTryTimes int
}

type Connection interface {
	reconnect(conn *net.TCPConn)
	getConn() *net.TCPConn
	close()
}

func NewSingleConn(address string, bakAddress string) (sc *SingleConnection) {

	sc = new(SingleConnection)
	sc.addr = address
	sc.currentAddr = sc.addr
	sc.bakAddr = bakAddress

	sc.failedTimes = 0
	sc.maxTryTimes = 30

	sc.initConn()

	return sc
}

func createSingleConnection(address string) (conn *net.TCPConn, err error) {

	tcpAddr, err := net.ResolveTCPAddr("tcp4", address)
	if err != nil {
		return nil, err
	}
	conn, err = net.DialTCP("tcp", nil, tcpAddr)
	if err != nil {
		log.Fatal("get connection from " + address + " failed! Error:" + err.Error())
		return nil, err
	} else {
		log.Println("get connection from " + address + " success! remote addr " + conn.RemoteAddr().String())
	}

	return conn, nil
}

//init conn list from addrMap
func (sc *SingleConnection) initConn() {
	newConn, err := createSingleConnection(sc.currentAddr)
	if err != nil {
		log.Fatal("init err:" + err.Error())
	} else {
		sc.conn = newConn
	}
}
