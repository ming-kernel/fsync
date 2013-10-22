package main

import (
	"log"
	"net"
	"os"
)

func listenPeer() {
	gaddr := net.UDPAddr{IP: net.IPv4(224, 0, 0, 254), Port: 12345}
	conn, err := net.ListenMulticastUDP("udp", nil, &gaddr)
	if err != nil {
		log.Println("ListenMulticastUDP error")
		return
	}

	if !gaddr.IP.IsMulticast() {
		log.Println("gaddr not multicast")
		return
	}

	msg := make([]byte, 256)
	for {
		n, err := conn.Read(msg)
		if err != nil {
			log.Println("ReadFromUDP error")
		}
		log.Println("n:", n)
		if n != 0 {
			os.Exit(0)
		}
	}

}

func joinGroup() {
	gaddr := net.UDPAddr{IP: net.IPv4(224, 0, 0, 254), Port: 12345}
	if !gaddr.IP.IsMulticast() {
		log.Println("gaddr not multicast")
		return
	}

	conn, err := net.DialUDP(
		"udp",
		nil,
		&gaddr)

	if err != nil {
		log.Println("DialUDP error")
		return
	}

	msg := []byte("howdy\n")
	for i := 0; i < 5; i++ {
		n, err := conn.Write(msg)
		if err != nil {
			log.Println("Write error")
			return
		}

		log.Println("Wirted: ", n)

	}

}

func main() {

	log.Println("Listen Peer")

	go listenPeer()

	go joinGroup()

	log.Println("Process directory events")

	watchfs()
}
