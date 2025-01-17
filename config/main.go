package config

var Host string = "0.0.0.0"
var Port int = 7379
var MaxKeyLimit int = 5
var EvictionStrategy string = "simple-first"
var AOFFile string = "./dice-master.aof"
