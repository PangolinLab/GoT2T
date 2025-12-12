package got2t

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

type FromIntImpl struct{}
type FromInt8Impl struct{}
type FromInt16Impl struct{}
type FromInt32Impl struct{}
type FromInt64Impl struct{}
type FromUintImpl struct{}
type FromUint8Impl struct{}
type FromUint16Impl struct{}
type FromUint32Impl struct{}
type FromUint64Impl struct{}
type FromFloat32Impl struct{}
type FromFloat64Impl struct{}
type FromBoolImpl struct{}
type FromStringImpl struct{}
type FromBytesImpl struct{}

var _ FromInt = (*FromIntImpl)(nil)
var _ FromInt8 = (*FromInt8Impl)(nil)
var _ FromInt16 = (*FromInt16Impl)(nil)
var _ FromInt32 = (*FromInt32Impl)(nil)
var _ FromInt64 = (*FromInt64Impl)(nil)
var _ FromUint = (*FromUintImpl)(nil)
var _ FromUint8 = (*FromUint8Impl)(nil)
var _ FromUint16 = (*FromUint16Impl)(nil)
var _ FromUint32 = (*FromUint32Impl)(nil)
var _ FromUint64 = (*FromUint64Impl)(nil)
var _ FromFloat32 = (*FromFloat32Impl)(nil)
var _ FromFloat64 = (*FromFloat64Impl)(nil)
var _ FromBool = (*FromBoolImpl)(nil)
var _ FromString = (*FromStringImpl)(nil)
var _ FromBytes = (*FromBytesImpl)(nil)

func (int *FromIntImpl) ToBool(a int) (bool, error) {
	if a == 0 {
		return false, nil
	} else if a == 1 {
		return true, nil
	} else {
		return false, errors.New("FromIntToBool: Not 0 or 1")
	}
}

func (int *FromIntImpl) ToString(a int) string {
	return fmt.Sprintf("%d", a)
}

func (int *FromIntImpl) ToBytes(a int) []byte {
	return strconv.AppendInt(nil, int64(a), 10)
}

func (int8 *FromInt8Impl) ToBool(a int8) (bool, error) {
	instance := FromIntImpl{}
	return instance.ToBool(int(a))
}

func (int8 *FromInt8Impl) ToString(a int8) string {
	instance := FromIntImpl{}
	return instance.ToString(int(a))
}

func (int8 *FromInt8Impl) ToBytes(a int8) []byte {
	instance := FromIntImpl{}
	return instance.ToBytes(int(a))
}

func (int16 *FromInt16Impl) ToBool(a int16) (bool, error) {
	instance := FromIntImpl{}
	return instance.ToBool(int(a))
}

func (int16 *FromInt16Impl) ToString(a int16) string {
	instance := FromIntImpl{}
	return instance.ToString(int(a))
}

func (int16 *FromInt16Impl) ToBytes(a int16) []byte {
	instance := FromIntImpl{}
	return instance.ToBytes(int(a))
}

func (int32 *FromInt32Impl) ToBool(a int32) (bool, error) {
	instance := FromIntImpl{}
	return instance.ToBool(int(a))
}

func (int32 *FromInt32Impl) ToString(a int32) string {
	instance := FromIntImpl{}
	return instance.ToString(int(a))
}

func (int32 *FromInt32Impl) ToBytes(a int32) []byte {
	instance := FromIntImpl{}
	return instance.ToBytes(int(a))
}

func (int64 *FromInt64Impl) ToBool(a int64) (bool, error) {
	instance := FromIntImpl{}
	return instance.ToBool(int(a))
}

func (int64 *FromInt64Impl) ToString(a int64) string {
	instance := FromIntImpl{}
	return instance.ToString(int(a))
}

func (int64 *FromInt64Impl) ToBytes(a int64) []byte {
	instance := FromIntImpl{}
	return instance.ToBytes(int(a))
}

func (uint *FromUintImpl) ToBool(a uint) (bool, error) {
	instance := FromIntImpl{}
	return instance.ToBool(int(a))
}

func (uint *FromUintImpl) ToString(a uint) string {
	instance := FromIntImpl{}
	return instance.ToString(int(a))
}

func (uint *FromUintImpl) ToBytes(a uint) []byte {
	return strconv.AppendUint(nil, uint64(a), 10)
}

func (uint8 *FromUint8Impl) ToBool(a uint8) (bool, error) {
	instance := FromIntImpl{}
	return instance.ToBool(int(a))
}

func (uint8 *FromUint8Impl) ToString(a uint8) string {
	instance := FromIntImpl{}
	return instance.ToString(int(a))
}

func (uint8 *FromUint8Impl) ToBytes(a uint8) []byte {
	instance := FromUintImpl{}
	return instance.ToBytes(uint(a))
}

func (uint16 *FromUint16Impl) ToBool(a uint16) (bool, error) {
	instance := FromIntImpl{}
	return instance.ToBool(int(a))
}

func (uint16 *FromUint16Impl) ToString(a uint16) string {
	instance := FromIntImpl{}
	return instance.ToString(int(a))
}

func (uint16 *FromUint16Impl) ToBytes(a uint16) []byte {
	instance := FromInt16Impl{}
	return instance.ToBytes(int16(a))
}

func (uint32 *FromUint32Impl) ToBool(a uint32) (bool, error) {
	instance := FromIntImpl{}
	return instance.ToBool(int(a))
}

func (uint32 *FromUint32Impl) ToString(a uint32) string {
	instance := FromIntImpl{}
	return instance.ToString(int(a))
}

func (uint32 *FromUint32Impl) ToBytes(a uint32) []byte {
	instance := FromInt32Impl{}
	return instance.ToBytes(int32(a))
}

func (uint64 *FromUint64Impl) ToBool(a uint64) (bool, error) {
	instance := FromIntImpl{}
	return instance.ToBool(int(a))
}

func (uint64 *FromUint64Impl) ToString(a uint64) string {
	instance := FromIntImpl{}
	return instance.ToString(int(a))
}

func (uint64 *FromUint64Impl) ToBytes(a uint64) []byte {
	instance := FromInt64Impl{}
	return instance.ToBytes(int64(a))
}

func (float32 *FromFloat32Impl) ToBool(a float32) (bool, error) {
	instance := FromIntImpl{}
	return instance.ToBool(int(a))
}

func (float32 *FromFloat32Impl) ToString(a float32) string {
	return fmt.Sprintf("%f", a)
}

func (float32 *FromFloat32Impl) ToBytes(a float32) []byte {
	return strconv.AppendFloat(nil, float64(a), 'f', -1, 32)
}

func (float64 *FromFloat64Impl) ToBool(a float64) (bool, error) {
	instance := FromIntImpl{}
	return instance.ToBool(int(a))
}

func (float64 *FromFloat64Impl) ToString(a float64) string {
	return fmt.Sprintf("%f", a)
}

func (float64 *FromFloat64Impl) ToBytes(a float64) []byte {
	return strconv.AppendFloat(nil, a, 'f', -1, 64)
}

func (bool *FromBoolImpl) ToInt(a bool) int {
	if !a {
		return 0
	} else {
		return 1
	}
}

func (bool *FromBoolImpl) ToInt8(a bool) int8 {
	return int8(bool.ToInt(a))
}

func (bool *FromBoolImpl) ToInt16(a bool) int16 {
	return int16(bool.ToInt(a))
}

func (bool *FromBoolImpl) ToInt32(a bool) int32 {
	return int32(bool.ToInt(a))
}

func (bool *FromBoolImpl) ToInt64(a bool) int64 {
	return int64(bool.ToInt(a))
}

func (bool *FromBoolImpl) ToUint(a bool) (uint, error) {
	intData := bool.ToInt(a)
	if intData < 0 {
		return 0, errors.New("FromBoolToUint: A negative number")
	}
	return uint(intData), nil
}

func (bool *FromBoolImpl) ToUint8(a bool) (uint8, error) {
	uintData, err := bool.ToUint(a)
	if err != nil {
		return 0, errors.New("FromBoolToUint8: A negative number")
	}
	return uint8(uintData), nil
}

func (bool *FromBoolImpl) ToUint16(a bool) (uint16, error) {
	uintData, err := bool.ToUint(a)
	if err != nil {
		return 0, errors.New("FromBoolToUint16: A negative number")
	}
	return uint16(uintData), nil
}

func (bool *FromBoolImpl) ToUint32(a bool) (uint32, error) {
	uintData, err := bool.ToUint(a)
	if err != nil {
		return 0, errors.New("FromBoolToUint32: A negative number")
	}
	return uint32(uintData), nil
}

func (bool *FromBoolImpl) ToUint64(a bool) (uint64, error) {
	uintData, err := bool.ToUint(a)
	if err != nil {
		return 0, errors.New("FromBoolToUint64: A negative number")
	}
	return uint64(uintData), nil
}

func (bool *FromBoolImpl) ToFloat32(a bool) float32 {
	return float32(bool.ToInt(a))
}

func (bool *FromBoolImpl) ToFloat64(a bool) float64 {
	return float64(bool.ToInt(a))
}

func (bool *FromBoolImpl) ToString(a bool) string {
	if !a {
		return "false"
	} else {
		return "true"
	}
}

func (bool *FromBoolImpl) ToBytes(a bool) []byte {
	if !a {
		return []byte("false")
	} else {
		return []byte("true")
	}
}

func (string *FromStringImpl) ToInt(a string) (int, error) {
	return strconv.Atoi(a)
}

func (string *FromStringImpl) ToInt8(a string) (int8, error) {
	intData, err := string.ToInt(a)
	if err != nil {
		return 0, err
	}
	return int8(intData), nil
}

func (string *FromStringImpl) ToInt16(a string) (int16, error) {
	intData, err := string.ToInt(a)
	if err != nil {
		return 0, err
	}
	return int16(intData), nil
}

func (string *FromStringImpl) ToInt32(a string) (int32, error) {
	intData, err := string.ToInt(a)
	if err != nil {
		return 0, err
	}
	return int32(intData), nil
}

func (string *FromStringImpl) ToInt64(a string) (int64, error) {
	intData, err := string.ToInt(a)
	if err != nil {
		return 0, err
	}
	return int64(intData), nil
}

func (string *FromStringImpl) ToUint(a string) (uint, error) {
	intData, err := string.ToInt(a)
	if err != nil {
		return 0, err
	}
	if intData < 0 {
		return 0, errors.New("FromStringToUint: A negative number")
	}
	return uint(intData), nil
}

func (string *FromStringImpl) ToUint8(a string) (uint8, error) {
	intData, err := string.ToInt(a)
	if err != nil {
		return 0, err
	}
	if intData < 0 {
		return 0, errors.New("FromStringToUint8: A negative number")
	}
	return uint8(intData), nil
}

func (string *FromStringImpl) ToUint16(a string) (uint16, error) {
	intData, err := string.ToInt(a)
	if err != nil {
		return 0, err
	}
	if intData < 0 {
		return 0, errors.New("FromStringToUint16: A negative number")
	}
	return uint16(intData), nil
}

func (string *FromStringImpl) ToUint32(a string) (uint32, error) {
	intData, err := string.ToInt(a)
	if err != nil {
		return 0, err
	}
	if intData < 0 {
		return 0, errors.New("FromStringToUint32: A negative number")
	}
	return uint32(intData), nil
}

func (string *FromStringImpl) ToUint64(a string) (uint64, error) {
	intData, err := string.ToInt(a)
	if err != nil {
		return 0, err
	}
	if intData < 0 {
		return 0, errors.New("FromStringToUint64: A negative number")
	}
	return uint64(intData), nil
}

func (string *FromStringImpl) ToFloat32(a string) (float32, error) {
	floatData, err := strconv.ParseFloat(a, 32)
	if err != nil {
		return 0, err
	}
	return float32(floatData), nil
}

func (string *FromStringImpl) ToFloat64(a string) (float64, error) {
	floatData, err := strconv.ParseFloat(a, 64)
	if err != nil {
		return 0, err
	}
	return floatData, nil
}

func (string *FromStringImpl) ToBool(a string) (bool, error) {
	if strings.ToLower(a) == "true" {
		return true, nil
	} else if strings.ToLower(a) == "false" {
		return false, nil
	} else {
		return false, errors.New("FromStringToBool: Not 'true' or 'false'")
	}
}

func (bytes *FromBytesImpl) ToInt(a []byte) (int, error) {
	i, err := strconv.ParseInt(string(a), 10, 64)
	if err != nil {
		return 0, err
	}
	return int(i), nil
}

func (bytes *FromBytesImpl) ToInt8(a []byte) (int8, error) {
	intData, err := bytes.ToInt(a)
	if err != nil {
		return 0, err
	}
	return int8(intData), nil
}

func (bytes *FromBytesImpl) ToInt16(a []byte) (int16, error) {
	intData, err := bytes.ToInt(a)
	if err != nil {
		return 0, err
	}
	return int16(intData), nil
}

func (bytes *FromBytesImpl) ToInt32(a []byte) (int32, error) {
	intData, err := bytes.ToInt(a)
	if err != nil {
		return 0, err
	}
	return int32(intData), nil
}

func (bytes *FromBytesImpl) ToInt64(a []byte) (int64, error) {
	intData, err := bytes.ToInt(a)
	if err != nil {
		return 0, err
	}
	return int64(intData), nil
}

func (bytes *FromBytesImpl) ToUint(a []byte) (uint, error) {
	intData, err := bytes.ToInt(a)
	if err != nil {
		return 0, err
	}
	if intData < 0 {
		return 0, errors.New("FromBytesToUint: A negative number")
	}
	return uint(intData), nil
}

func (bytes *FromBytesImpl) ToUint8(a []byte) (uint8, error) {
	intData, err := bytes.ToInt(a)
	if err != nil {
		return 0, err
	}
	if intData < 0 {
		return 0, errors.New("FromBytesToUint8: A negative number")
	}
	return uint8(intData), nil
}

func (bytes *FromBytesImpl) ToUint16(a []byte) (uint16, error) {
	intData, err := bytes.ToInt(a)
	if err != nil {
		return 0, err
	}
	if intData < 0 {
		return 0, errors.New("FromBytesToUint16: A negative number")
	}
	return uint16(intData), nil
}

func (bytes *FromBytesImpl) ToUint32(a []byte) (uint32, error) {
	intData, err := bytes.ToInt(a)
	if err != nil {
		return 0, err
	}
	if intData < 0 {
		return 0, errors.New("FromBytesToUint32: A negative number")
	}
	return uint32(intData), nil
}

func (bytes *FromBytesImpl) ToUint64(a []byte) (uint64, error) {
	intData, err := bytes.ToInt(a)
	if err != nil {
		return 0, err
	}
	if intData < 0 {
		return 0, errors.New("FromBytesToUint64: A negative number")
	}
	return uint64(intData), nil
}

func (bytes *FromBytesImpl) ToFloat32(a []byte) (float32, error) {
	floatData, err := strconv.ParseFloat(string(a), 32)
	if err != nil {
		return 0, err
	}
	return float32(floatData), nil
}

func (bytes *FromBytesImpl) ToFloat64(a []byte) (float64, error) {
	floatData, err := strconv.ParseFloat(string(a), 64)
	if err != nil {
		return 0, err
	}
	return floatData, nil
}

func (bytes *FromBytesImpl) ToBool(a []byte) (bool, error) {
	if strings.ToLower(string(a)) == "true" {
		return true, nil
	} else if strings.ToLower(string(a)) == "false" {
		return false, nil
	} else {
		return false, errors.New("FromBytesToBool: Not 'true' or 'false'")
	}
}
