package utils

import (
	"github.com/alg-grindel/saveSocial/pkg/constant"
	"github.com/alg-grindel/saveSocial/pkg/errno"
	"log"
	"sync"
	"time"
)

type Snowflake struct {
	mu          sync.Mutex
	timestamp   int64
	machineId   int64
	sequenceNum int64
}

func NewSnowflake(machineId int64) (*Snowflake, error) {
	if machineId < 0 || machineId > constant.MaxMachineID {
		log.Printf("[ERROR] utils.snowflake: machineID must between 0 and %s", constant.MaxMachineID-1)
		return nil, errno.ParmaErr
	}
	return &Snowflake{
		timestamp:   0,
		machineId:   machineId,
		sequenceNum: 0,
	}, nil
}

func (s *Snowflake) GenerateID() (int64, error) {
	s.mu.Lock()
	now := time.Now().UnixNano() / 1e6
	if s.timestamp == now {
		s.sequenceNum = (s.sequenceNum + 1) & constant.MaxSequence
		if s.sequenceNum == 0 {
			for now <= s.timestamp {
				now = time.Now().UnixNano() / 1e6
			}
		}
	} else {
		s.sequenceNum = 0
	}
	t := now - constant.Epoch
	if t > constant.MaxTimestamp {
		s.mu.Unlock()
		log.Printf("[ERROR] utils.snowflake: epoch must between 0 and %d", constant.MaxTimestamp-1)
		return -1, errno.SystemErr
	}
	s.timestamp = now
	id := ((t) << constant.TimestampShift) | (s.machineId << constant.MachineIDShift) | (s.sequenceNum)
	s.mu.Unlock()
	return id, nil
}
