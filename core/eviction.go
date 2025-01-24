package core

import "github.com/amrishkshah/dicedb/config"

func evictFirst() {
	for key := range store {
		println("evict")
		Del(key)
		return
	}
}

func evictAllkeysRandom() {
	evictCount := int64(config.EvictionRatio * float64(config.MaxKeyLimit))
	for k := range store {
		Del(k)
		evictCount--
		if evictCount <= 0 {
			break
		}
	}
}

func evict() {
	switch config.EvictionStrategy {
	case "simple-first":
		evictFirst()
	case "allkeys-random":
		evictAllkeysRandom()
	}

}
