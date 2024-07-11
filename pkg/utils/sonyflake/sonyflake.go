package sonyflake

import (
	"time"

	"github.com/sony/sonyflake"
)

var (
	sf *sonyflake.Sonyflake
)

func startTime() time.Time {
	return time.Now()
}

func getMachineID() (uint16, error) {
	return 0, nil
}

func checkMachineID(uint16) bool {
	return false
}

// InitSonyflake int sonyflake
func InitSonyflake() {
	settings := sonyflake.Settings{
		/*StartTime: startTime(),
		MachineID:      getMachineID,
		CheckMachineID: checkMachineID,*/
	}
	sf = sonyflake.NewSonyflake(settings)
	if sf == nil {
		panic("sonyflake not created")
	}
}

// GenerateSonyflake  generate Sonyflake ID
func GenerateSonyflake() uint64 {
	if sf == nil {
		InitSonyflake()
	}
	if sf == nil {
		return 0
	}
	id, err := sf.NextID()
	if err != nil {
		return 0
	}
	return id
}
