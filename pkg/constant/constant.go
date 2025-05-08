package constant

const (
	// username password addr database
	DsnFormat = "%s:%s@tcp(%s)/%s?charset=utf8mb4&parseTime=True&loc=Local"

	Salt = "save0814."

	// snowflake
	Epoch          = 903024000000
	TimestampBits  = 41
	MachineIDBits  = 10
	SequenceBits   = 12
	MaxTimestamp   = -1 ^ (-1 << TimestampBits)
	MaxMachineID   = -1 ^ (-1 << MachineIDBits)
	MaxSequence    = -1 ^ (-1 << SequenceBits)
	TimestampShift = MachineIDBits + SequenceBits
	MachineIDShift = SequenceBits
)
