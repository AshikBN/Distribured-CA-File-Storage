package p2p

// message holds any arbitrary data that is being sent over each transport
// between two node in the network
type Message struct {
	Payload []byte
}
