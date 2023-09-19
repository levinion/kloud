package cache

import (
	"errors"
	"time"
)

type Nop struct{}

func UseNoCache() *Nop {
	nop := &Nop{}
	setDefaultCache(nop)
	return nop
}

func (n *Nop) Set(key []byte, value []byte, expiration time.Duration) {
	// Do nothing
}

func (n *Nop) Get(key []byte) ([]byte, error) {
	return nil, errors.New("nothing happens because you are using no cache")
}
