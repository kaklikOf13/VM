package vm

import (
	"encoding/binary"
	"fmt"
)

const (
	T_NullType  uint8 = 0
	T_Int       uint8 = 1
	T_Int8      uint8 = 2
	T_Int16     uint8 = 3
	T_Int32     uint8 = 4
	T_Int64     uint8 = 5
	T_UInt      uint8 = 6
	T_UInt8     uint8 = 7
	T_UInt16    uint8 = 8
	T_UInt32    uint8 = 9
	T_UInt64    uint8 = 10
	T_String    uint8 = 11
	T_Interface uint8 = 12
	T_Struct    uint8 = 13
	T_Type      uint8 = 14
	T_Instance  uint8 = 15
	T_Array     uint8 = 16
	T_Map       uint8 = 17
	T_Function  uint8 = 18
	T_Pointer   uint8 = 19
	T_Slice     uint8 = 20
	T_Node      uint8 = 21
)

type Value interface {
	String() string
	Bytes() []byte
	Type() uint8
}

type NullType struct {
}

func (null *NullType) String() string {
	return "null"
}
func (null *NullType) Bytes() []byte {
	return []byte{}
}
func (null *NullType) Type() uint8 {
	return T_NullType
}

type Int struct {
	NullType
	value int64
}

func (t *Int) String() string {
	return fmt.Sprint(t.value)
}
func (t *Int) Bytes() []byte {
	v := make([]byte, 8)
	binary.LittleEndian.PutUint64(v, uint64(t.value))
	return v
}
func (t *Int) Type() uint8 {
	return T_Int
}

type Int64 struct {
	Int
}

func (t *Int64) Type() uint8 {
	return T_Int64
}

type Int32 struct {
	NullType
	value int32
}

func (t *Int32) String() string {
	return fmt.Sprint(t.value)
}
func (t *Int32) Bytes() []byte {
	v := make([]byte, 4)
	binary.LittleEndian.PutUint32(v, uint32(t.value))
	return v
}
func (t *Int32) Type() uint8 {
	return T_Int32
}

type Int16 struct {
	NullType
	value int16
}

func (t *Int16) String() string {
	return fmt.Sprint(t.value)
}
func (t *Int16) Bytes() []byte {
	v := make([]byte, 2)
	binary.LittleEndian.PutUint16(v, uint16(t.value))
	return v
}
func (t *Int16) Type() uint8 {
	return T_Int16
}

type Int8 struct {
	NullType
	value int8
}

func (t *Int8) String() string {
	return fmt.Sprint(t.value)
}
func (t *Int8) Bytes() []byte {
	return []byte{byte(t.value)}
}
func (t *Int8) Type() uint8 {
	return T_Int8
}

type UInt struct {
	Int64
}

func (t *UInt) Type() uint8 {
	return T_UInt
}

type UInt64 struct {
	UInt
}

func (t *UInt64) Type() uint8 {
	return T_UInt64
}

type UInt32 struct {
	Int32
}

func (t *UInt32) Type() uint8 {
	return T_UInt32
}

type UInt16 struct {
	Int16
}

func (t *UInt16) Type() uint8 {
	return T_UInt16
}

type UInt8 struct {
	Int8
}

func (t *UInt8) Type() uint8 {
	return T_UInt8
}

type String struct {
	NullType
	value string
}

func (t *String) String() string {
	return t.value
}
func (t *String) Bytes() []byte {
	return []byte(t.value)
}
func (t *String) Type() uint8 {
	return T_String
}

type Node struct {
	NullType
	Value    []Value
	NodeType string
}

func (n *Node) String() string {
	txt := ""
	for _, v := range n.Value {
		if txt != "" {
			txt += ","
		}
		txt += v.String()
	}
	return "<" + n.NodeType + ":[" + txt + "]>"
}
func (n *Node) Type() uint8 {
	return T_Node
}
