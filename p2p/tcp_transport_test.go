package p2p

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTCPTransport(t *testing.T) {
	listAddr := ":4000"
	tr := NewTCPTransport(listAddr)

	assert.Equal(t, tr.listenAddress, listAddr)
	assert.Nil(t, tr.ListenAndAccept())

}
