package core

import (
	"log"
	"time"
)

func hasExpired(obj *Obj) bool {
	exp, ok := expires[obj]
	if !ok {
		return false
	}
	return exp <= uint64(time.Now().UnixMilli())
}

func getExpiry(obj *Obj) (uint64, bool) {
	exp, ok := expires[obj]
	return exp, ok
}

func expireSample() float32 {
	var limit int = 20
	var expiredCount int = 0

	for key, obj := range store {

		limit--
		// if the key is expired
		if hasExpired(obj) {
			Del(key)
			println(key)
			expiredCount++
		}

		// once we iterated to 20 keys that have some expiration set
		// we break the loop
		if limit == 0 {
			break
		}
	}
	return float32(expiredCount) / float32(20.0)
}

func DeleteExpiredKeys() {
	for {
		frac := expireSample()

		if frac < 25.0 {
			break
		}
	}
	log.Println("deleted the expired but undeleted keys. total keys", len(store))
}
