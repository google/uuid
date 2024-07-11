package uuid

import (
	"sync"
)

// CustomNodeId is a node id with a few tweaks:
// It allows the creation of UUIDs using distinct node ids simultaneously from a single host, without
// calling SetNodeId() repeatedly.
// It avoids collisions with version 1 UUIDs generated based on hardware IEEE 802 addresses.
type CustomNodeId struct {
	nodeid            // inherit the nodeid
	mu     sync.Mutex // hold an own mutex
}

// CustomNodeIdError enables clean error handling during CustomNodeId operation.
type CustomNodeIdError string

func (e CustomNodeIdError) Error() string {
	return string(e)
}

const (
	ErrIncompleteId CustomNodeIdError = "given node id is missing data"
	ErrInvalidFlag  CustomNodeIdError = "multicast flag not set"
)

// NewCustomNodeId returns a new custom node id or an error, in case invalid or insufficient data is given.
func NewCustomNodeId(id []byte) (*CustomNodeId, error) {
	ret := &CustomNodeId{}
	err := ret.SetNodeId(id)
	if err != nil {
		return nil, err
	}
	return ret, nil
}

// NewUUID generates a new version 1 UUID based on this CustomNodeId.
// This implementation mimics the global NewUUID function.
func (c *CustomNodeId) NewUUID() (UUID, error) {
	// Use the plain version 1 uuid generation …
	uuid, err := newVersion1UUID()
	if err != nil {
		return uuid, err
	}

	// _ And fill in our own node id:
	c.mu.Lock()
	defer c.mu.Unlock()
	copy(uuid[10:], c.nodeid[:])

	return uuid, nil
}

// SetNodeId sets the custom node id to the given value.
// Note: RFC-4122 suggest to use a "47-bit cryptographic quality random number" with the least significant bit of the
// first byte set to 1. It also states, the data is "system specific".
func (c *CustomNodeId) SetNodeId(id []byte) error {
	if len(id) < 6 {
		return ErrIncompleteId
	}

	// According to RFC-4122 section 4.5 the unicast/multicast flag of the IEEE 802 MAC has to be set
	// for custom node ids in order to avoid conflicts.
	// Thereby we should reject node ids with this flag being unset:
	if id[0]&0x01 == 0 {
		return ErrInvalidFlag
	}

	c.mu.Lock()
	defer c.mu.Unlock()
	copy(c.nodeid[:], id)

	return nil
}

// NodeID returns the 6 byte node id.
func (c *CustomNodeId) NodeID() []byte {
	c.mu.Lock()
	defer c.mu.Unlock()
	var node nodeid
	copy(node[:], c.nodeid[:])
	return node[:]
}
