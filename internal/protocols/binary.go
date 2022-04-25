package protocols

import (
	"bytes"
	"encoding/binary"
)

type Writer interface {
	Write(*bytes.Buffer) bool
}

type Reader interface {
	Read(*bytes.Buffer) bool
}

type BuiltinIdx interface {
	uint8 | uint16 | uint32 | uint64
}

type BuiltinInt interface {
	int8 | uint8 | int16 | uint16 | int32 | uint32 | int64 | uint64
}

var builtinIntegers = map[string]struct{}{
	"int8":   {},
	"uint8":  {},
	"int16":  {},
	"uint16": {},
	"int32":  {},
	"uint32": {},
	"int64":  {},
	"uint64": {},
}

var builtinIntegerArrays = map[string]struct{}{
	"[]int8":   {},
	"[]uint8":  {},
	"[]int16":  {},
	"[]uint16": {},
	"[]int32":  {},
	"[]uint32": {},
	"[]int64":  {},
	"[]uint64": {},
}

var builtinIntegerMaxs = map[string]string{
	"int8":   "",
	"uint8":  "255",
	"int16":  "",
	"uint16": "65535",
	"int32":  "",
	"uint32": "4294967295",
	"int64":  "",
	"uint64": "18446744073709551615",
}

func IsBuiltinInteger(s string) bool {
	_, result := builtinIntegers[s]
	return result
}

func IsBuiltinIntegerArray(s string) bool {
	_, result := builtinIntegerArrays[s]
	return result
}

func GetBuiltinIntegerMax(s string) string {
	ss, result := builtinIntegerMaxs[s]
	if !result {
		return ""
	}

	return ss
}

func ProtoWriteInteger[T BuiltinInt](buff *bytes.Buffer, i T) bool {
	return binary.Write(buff, binary.LittleEndian, i) == nil
}

func ProtoReadInteger[T BuiltinInt](buff *bytes.Buffer, i *T) bool {
	return binary.Read(buff, binary.LittleEndian, i) == nil
}

func ProtoWriteIntegerArray[T BuiltinInt, I BuiltinIdx](buff *bytes.Buffer, array []T, limit I) bool {
	l := len(array)
	if uint64(l) > uint64(limit) {
		return false
	}

	if err := binary.Write(buff, binary.LittleEndian, I(l)); err != nil {
		return false
	}

	return binary.Write(buff, binary.LittleEndian, array) == nil
}

func ProtoReadIntegerArray[T BuiltinInt, I BuiltinIdx](buff *bytes.Buffer, array *[]T, limit I) bool {
	l := I(0)
	if err := binary.Read(buff, binary.LittleEndian, &l); err != nil {
		return false
	}

	if l > limit {
		return false
	}

	*array = make([]T, l)
	return binary.Read(buff, binary.LittleEndian, *array) == nil
}

func ProtoWriteString(buff *bytes.Buffer, s string, limit uint16) bool {
	l := len(s)
	if l > int(limit) {
		return false
	}

	if err := binary.Write(buff, binary.LittleEndian, uint16(l)); err != nil {
		return false
	}

	_, err := buff.WriteString(s)
	return err == nil
}

func ProtoReadString(buff *bytes.Buffer, s *string, limit uint16) bool {
	l := uint16(0)
	if err := binary.Read(buff, binary.LittleEndian, &l); err != nil {
		return false
	}

	if l > limit {
		return false
	}

	d := make([]byte, l)
	if _, err := buff.Read(d); err != nil {
		return false
	}

	*s = string(d)
	return true
}

func ProtoWriteCustom(buff *bytes.Buffer, t Writer) bool {
	return t.Write(buff)
}

func ProtoReadCustom(buff *bytes.Buffer, t Reader) bool {
	return t.Read(buff)
}

func ProtoWriteCustomArray[T any, I BuiltinIdx](buff *bytes.Buffer, tt []T, limit I) bool {
	l := len(tt)
	if uint64(l) > uint64(limit) {
		return false
	}

	if err := binary.Write(buff, binary.LittleEndian, I(l)); err != nil {
		return false
	}

	var t interface{}
	for i := 0; i < l; i++ {
		t = &tt[i]
		writer, ok := t.(Writer)
		if !ok {
			return false
		}

		if !writer.Write(buff) {
			return false
		}
	}

	return true
}

func ProtoReadCustomArray[T any, I BuiltinIdx](buff *bytes.Buffer, tt *[]T, limit I) bool {
	l := I(0)
	if err := binary.Read(buff, binary.LittleEndian, &l); err != nil {
		return false
	}

	if l > limit {
		return false
	}

	var t interface{}
	*tt = make([]T, l)
	for i := I(0); i < l; i++ {
		t = &(*tt)[i]
		reader, ok := t.(Reader)
		if !ok {
			return false
		}

		if !reader.Read(buff) {
			return false
		}
	}

	return true
}
