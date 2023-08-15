package snow_flake

import (
	"gin-pro/app/global/consts"
	"sync"
	"time"
)

func NewSnowFlake() *Snowflake {
	return &Snowflake{
		timestamp: 0,
		machineId: 1,
		sequence:  0,
	}
}

type Snowflake struct {
	sync.Mutex
	timestamp int64
	machineId int64
	sequence  int64
}

func (s *Snowflake) GetId() int64 {
	s.Lock()
	defer func() {
		s.Unlock()
	}()
	now := time.Now().UnixNano() / 1e6
	if s.timestamp == now {
		s.sequence = (s.sequence + 1) & consts.SequenceMask
		if s.sequence == 0 {
			for now <= s.timestamp {
				now = time.Now().UnixNano() / 1e6
			}
		}
	} else {
		s.sequence = 0
	}
	s.timestamp = now
	r := (now-consts.StartTimeStamp)<<consts.TimestampShift | (s.machineId << consts.MachineIdShift) | (s.sequence)
	return r
}
