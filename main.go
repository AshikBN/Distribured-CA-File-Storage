package main

import (
	"log"

	"github.com/AshikBN/Distribured-CA-File-Storage/p2p"
)

func main() {
	tcpOpts := p2p.TCPTransportOpts{
		ListenAddr:    ":3000",
		HandshakeFunc: p2p.NOPHandshakeFunc,
		Decoder:       p2p.DefaultDecoder{},
	}
	tr := p2p.NewTCPTransport(tcpOpts)
	err := tr.ListenAndAccept()
	if err != nil {
		log.Fatal(err)
	}
	select {}

}
