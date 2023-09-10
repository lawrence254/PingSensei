package main

import "time"
type PingResults struct{
	MinLatency time.Duration
	MaxLatency time.Duration
	AvgLatency time.Duration
	Ping time.Duration
	PacketLoss float64
}