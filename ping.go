package main

import (
	"fmt"
)

func main()  {
	valBahrain := getPingStats("75.2.105.73")

	formatedResponse := PingResults{
		MinLatency: valBahrain.MinRtt,
		MaxLatency: valBahrain.MaxRtt,
		AvgLatency: valBahrain.AvgRtt,
		Ping: valBahrain.StdDevRtt,
		PacketLoss: valBahrain.PacketLoss,
	}

	fmt.Println(formatedResponse)
}