package core

import (
	"time"

	"github.com/amrishkshah/dicedb/config"
)

var store = make(map[string]*Obj)

// type Obj struct {
// 	Value     interface{}
// 	ExpiresAt int64
// }

func newObj(value interface{}, durationMs int64, oType uint8, oEnc uint8) *Obj {
	var expiresAt int64 = -1
	if durationMs > 0 {
		expiresAt = time.Now().UnixMilli() + durationMs
	}

	return &Obj{
		Value:        value,
		TypeEncoding: oType | oEnc,
		ExpiresAt:    expiresAt,
	}
}

func Put(k string, obj *Obj) {
	if len(store) >= config.MaxKeyLimit {
		evict()
	}
	store[k] = obj
}

func Get(k string) *Obj {
	v := store[k]
	if v != nil {
		if v.ExpiresAt != -1 && v.ExpiresAt <= time.Now().UnixMilli() {
			delete(store, k)
			return nil
		}
	}
	return v
}

func Del(k string) bool {

	if _, ok := store[k]; ok {
		delete(store, k)
		return true
	}
	return false
}
