package node

import "sync"

type Storage interface {

}

type MemoryStorage struct {
	sync.Mutex

}