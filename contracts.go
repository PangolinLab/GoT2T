package got2t

type FromInt interface {
	ToBool(a int) (bool, error)
	ToString(a int) string
	ToBytes(a int) []byte
}

type FromInt8 interface {
	ToBool(a int8) (bool, error)
	ToString(a int8) string
	ToBytes(a int8) []byte
}

type FromInt16 interface {
	ToBool(a int16) (bool, error)
	ToString(a int16) string
	ToBytes(a int16) []byte
}

type FromInt32 interface {
	ToBool(a int32) (bool, error)
	ToString(a int32) string
	ToBytes(a int32) []byte
}

type FromInt64 interface {
	ToBool(a int64) (bool, error)
	ToString(a int64) string
	ToBytes(a int64) []byte
}

type FromUint interface {
	ToBool(a uint) (bool, error)
	ToString(a uint) string
	ToBytes(a uint) []byte
}

type FromUint8 interface {
	ToBool(a uint8) (bool, error)
	ToString(a uint8) string
	ToBytes(a uint8) []byte
}

type FromUint16 interface {
	ToBool(a uint16) (bool, error)
	ToString(a uint16) string
	ToBytes(a uint16) []byte
}

type FromUint32 interface {
	ToBool(a uint32) (bool, error)
	ToString(a uint32) string
	ToBytes(a uint32) []byte
}

type FromUint64 interface {
	ToBool(a uint64) (bool, error)
	ToString(a uint64) string
	ToBytes(a uint64) []byte
}

type FromFloat32 interface {
	ToBool(a float32) (bool, error)
	ToString(a float32) string
	ToBytes(a float32) []byte
}

type FromFloat64 interface {
	ToBool(a float64) (bool, error)
	ToString(a float64) string
	ToBytes(a float64) []byte
}

type FromBool interface {
	ToInt(a bool) int
	ToInt8(a bool) int8
	ToInt16(a bool) int16
	ToInt32(a bool) int32
	ToInt64(a bool) int64
	ToUint(a bool) (uint, error)
	ToUint8(a bool) (uint8, error)
	ToUint16(a bool) (uint16, error)
	ToUint32(a bool) (uint32, error)
	ToUint64(a bool) (uint64, error)
	ToFloat32(a bool) float32
	ToFloat64(a bool) float64
	ToString(a bool) string
	ToBytes(a bool) []byte
}

type FromString interface {
	ToInt(a string) (int, error)
	ToInt8(a string) (int8, error)
	ToInt16(a string) (int16, error)
	ToInt32(a string) (int32, error)
	ToInt64(a string) (int64, error)
	ToUint(a string) (uint, error)
	ToUint8(a string) (uint8, error)
	ToUint16(a string) (uint16, error)
	ToUint32(a string) (uint32, error)
	ToUint64(a string) (uint64, error)
	ToFloat32(a string) (float32, error)
	ToFloat64(a string) (float64, error)
	ToBool(a string) (bool, error)
}

type FromBytes interface {
	ToInt(a []byte) (int, error)
	ToInt8(a []byte) (int8, error)
	ToInt16(a []byte) (int16, error)
	ToInt32(a []byte) (int32, error)
	ToInt64(a []byte) (int64, error)
	ToUint(a []byte) (uint, error)
	ToUint8(a []byte) (uint8, error)
	ToUint16(a []byte) (uint16, error)
	ToUint32(a []byte) (uint32, error)
	ToUint64(a []byte) (uint64, error)
	ToFloat32(a []byte) (float32, error)
	ToFloat64(a []byte) (float64, error)
	ToBool(a []byte) (bool, error)
}
