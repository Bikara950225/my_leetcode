package consts

const (
	MinInt64 = int64(-1) << (64 - 1)
	MaxInt64 = int64(^uint64(0) >> 1)
	MinUint64 = uint64(0)
	MaxUint64 = ^uint64(0)

	MinInt32 = int32(-1) << (32 - 1)
	MaxInt32 = int32(^uint32(0) >> 1)
	MinUint32 = uint32(0)
	MaxUint32 = ^uint32(0)
)
