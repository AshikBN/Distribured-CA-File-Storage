package p2p

// peer is an interface that represents a remote node
type Peer interface{}

// ransport can be anything that handles the commnucation
// between two nodes in the network.This can be of type
// tcp,udp,websockets..
type Transport interface{}
