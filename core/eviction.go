package core

import "github.com/amrishkshah/dicedb/config"

func evictFirst() {
	for key := range store {
		println("evict")
		delete(store, key)
		return
	}
}

func evict() {
	switch config.EvictionStrategy {
	case "simple-first":
		evictFirst()
	}
}
