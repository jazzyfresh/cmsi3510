package dinner

import "sync"

type ChopStick struct {
	sync.Mutex
}
