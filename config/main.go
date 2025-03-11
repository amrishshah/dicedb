package config

var Host string = "0.0.0.0"
var Port int = 7379
var MaxKeyLimit int = 20
var EvictionRatio float64 = 0.40
var EvictionStrategy string = "allkeys-lru"
var AOFFile string = "./dice-master.aof"
