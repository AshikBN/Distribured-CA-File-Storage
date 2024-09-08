package p2p

import (
	"fmt"
	"net"
	"sync"
)

// TCPPeer represents a remote node over the tcp establised connection
type TCPPeer struct {
	conn net.Conn

	//if we dial and retrieved the connection outbound is true
	//if we accept and retrieved the connection then false
	outbound bool
}

func NewTCPPeer(conn net.Conn, outbound bool) *TCPPeer {
	return &TCPPeer{
		conn:     conn,
		outbound: outbound,
	}
}

type TCPTransportOpts struct {
	ListenAddr    string
	HandshakeFunc HandshakeFunc
	Decoder       Decoder
}

type TCPTransport struct {
	TCPTransportOpts
	listener net.Listener

	mu    sync.RWMutex
	peers map[net.Addr]Peer
}

func NewTCPTransport(opts TCPTransportOpts) *TCPTransport {
	return &TCPTransport{
		TCPTransportOpts: opts,
	}
}

func (t *TCPTransport) ListenAndAccept() error {
	var err error
	t.listener, err = net.Listen("tcp", t.ListenAddr)
	if err != nil {
		return err
	}
	go t.startAcceptLoop()
	return nil

}

func (t *TCPTransport) startAcceptLoop() {
	for {
		conn, err := t.listener.Accept()
		if err != nil {
			fmt.Printf("TCP accept error: %s\n", err)
		}
		go t.handleConnection(conn)

	}
}

func (t *TCPTransport) handleConnection(conn net.Conn) {
	var err error
	peer := NewTCPPeer(conn, true)
	if err = t.HandshakeFunc(peer); err != nil {
		conn.Close()
		fmt.Printf("TCP handshake error %s\n", err)
	}

	//read loop
	msg := &Message{}
	for {
		err := t.Decoder.Decode(conn, msg)
		if err != nil {
			fmt.Printf("TCP error %s\n", err)
			continue
		}
		fmt.Printf("Message:%+v\n", msg)

	}

}
