package common

import (
	"sync"
)

type JSON = map[string]interface{}

type Once struct {
	m    sync.Mutex
	done uint32
}