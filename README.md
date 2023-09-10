# PINGSENSEI

**PINGSENSEI** is a Go backend for an application that is intended to get latency metrics from gamers for specific games on specific servers. This tool is meant to give gamers an idea of which ISPs are best placed to offer services that meet their latency requirements on the games that they play.

ISPs in the long run can also use it as a benchmark against other providers to see where they need to improve if they need to.

## Data Captured

Once the frontend is provided to the users, they will provide the following metrics:

```go
type PingRecord struct {
	ID         string        `json:"id"`
	MinLatency string        `json:"minlatency"`
	MaxLatency string        `json:"maxlatency"`
	AvgLatency string        `json:"avglatency"`
	Ping       time.Duration `json:"ping"`
	PacketLoss float64       `json:"packetloss"`
	GameID     string        `json:"gameid"`
	ProviderID string        `json:"providerid"`
	ServerIP   string        `json:"serverip"`
	Date       string        `json:"date"`
}
```

Metrics will be captured with little to no user identifiable data. The following is a quick breakdown of some of the fields.
The `ID` field is a general UUID-V4 field that will act as the id for the given data point.
`GameID`,is the ID for the game that the user wishes to test, the ID is provided from a preexisting database of games or once that are user added.

The `ProviderID` on the other hand is going to be an Internet Service Provider identifier also from a pre-seeded database. It can be expanded to include more providers over time.

`ServerIP` is the publicly listed IP Address for a given game or one that the use has found for the game that they intend to test for.

## Usage

To run the application, ensure that you have [Go](https://go.dev/) installed.

After cloning the repository, you can run the following command and you should get a response with the pre-seeded server for testing.

```go
go run .
```

You can change the default test server by editing the following line in the `ping.go` file:

```go
valBahrain := getPingStats("75.2.105.73") // Provide your IP here
```

## Status

In Development

## Contributing

Pull requests are welcome. For major changes, please open an issue first
to discuss what you would like to change.

Please make sure to update tests as appropriate.

## License

[MIT](https://choosealicense.com/licenses/mit/)
