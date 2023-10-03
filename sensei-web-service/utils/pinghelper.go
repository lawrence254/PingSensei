package utils
import probing "github.com/prometheus-community/pro-bing"

func PingServer(ipAddress string) (*probing.Statistics, error){
	pinger, err := probing.NewPinger(ipAddress)

	if err != nil {
		return nil,err
	}

	pinger.Count =4

	err = pinger.Run()

	if err != nil {
		return nil, err
	}

	stats := pinger.Statistics()

	return stats, nil
}