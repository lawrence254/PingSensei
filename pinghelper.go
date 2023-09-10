package main
import probing "github.com/prometheus-community/pro-bing"

func getPingStats(ipAddress string) *probing.Statistics{
	pinger, err := probing.NewPinger(ipAddress)

	if err != nil {
		panic(err)
	}

	pinger.Count =4

	err = pinger.Run()

	if err != nil {
		panic(err)
	}

	stats := pinger.Statistics()

	return stats
}