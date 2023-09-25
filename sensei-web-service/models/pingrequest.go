package models


type PingRequest struct{
	ServerIP string `json:"server"`
	GameID string `json:"game"`
	Provider string `json:"provider"`
}

type PingResponse struct {
	Req *PingRequest `json:"request"`
	Probe *PingRecord `json:"ping"`
}