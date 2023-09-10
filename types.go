package main

import (
	"time"
)
type PingResults struct{
	MinLatency time.Duration
	MaxLatency time.Duration
	AvgLatency time.Duration
	Ping time.Duration
	PacketLoss float64
}

type PingRecord struct{
	ID string `json:"id"`
	AvgLatency string `json:"avglatency"`
	MinLatency string `json:"minlatency"`
	MaxLatency string `json:"maxlatency"`
	Ping time.Duration `json:"ping"`
	PacketLoss float64 `json:"packetloss"`
	GameID string `json:"gameid"`
	ProviderID string `json:"providerid"`
	ServerIP string `json:"serverip"`
	Date time.Time `json:"date"` 
}