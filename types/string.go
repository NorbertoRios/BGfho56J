package types

import (
	"strconv"
	"strings"
)

//NewString ...
func NewString(_data string) *String {
	return &String{
		Data: _data,
	}
}

//String str for string
type String struct {
	Data string
}

//BitmaskStringToInt32 string value to byte
func (str *String) BitmaskStringToInt32() int32 {
	result := int32(0)
	bitsArr := strings.Split(str.Data, "")
	lngth := len(bitsArr) - 1
	for i := lngth; i > -1; i-- {
		if bitsArr[i] == "1" {
			result = 1<<(lngth-i) | result
		}
	}
	return result
}

//BitmaskStringToByte string value to byte
func (str *String) BitmaskStringToByte() byte {
	result := byte(0)
	bitsArr := strings.Split(str.Data, "")
	lngth := len(bitsArr) - 1
	for i := lngth; i > -1; i-- {
		if bitsArr[i] == "1" {
			result = 1<<(lngth-i) | result
		}
	}
	return result
}

//Byte returns byte from string
func (str *String) Byte() byte {
	value, err := strconv.ParseUint(str.Data, 10, 8)
	if err != nil {
		return 0
	}
	return byte(value)
}

//Float ...
func (str *String) Float(bitSize int) interface{} {
	value, _ := strconv.ParseFloat(str.Data, bitSize)
	return float32(value)
}

//IntValue ...
func (str *String) IntValue() interface{} {
	value, _ := strconv.ParseInt(str.Data, 10, 32)
	return int(value)
}

//Int ...
func (str *String) Int(bitSize int) interface{} {
	value, _ := strconv.ParseInt(str.Data, 10, bitSize)
	return int32(value)
}

//UInt ...
func (str *String) UInt(bitSize int) interface{} {
	value, _ := strconv.ParseInt(str.Data, 10, bitSize)
	return value
}
