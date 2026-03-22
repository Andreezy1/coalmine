package miners

import "context"

type Coal int

type MinerInfo struct {
	Id         int
	MinerClass string
	Energy     int
}

type Miner interface {
	Run(ctx context.Context) <-chan Coal
	Info() MinerInfo
}
