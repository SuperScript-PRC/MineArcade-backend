package protocol

import (
	"bytes"
	"encoding/binary"
)

type Writer struct {
	buf []byte
	len int
}

type SupportMarshal interface {
	Marshal(w *Writer)
}

func NewWriter() Writer {
	return Writer{buf: []byte{}, len: 0}
}

func (w *Writer) writeByte(b byte) {
	w.buf = append(w.buf, b)
	w.len++
}

func (w *Writer) writeBytes(bs []byte) {
	w.buf = append(w.buf, bs...)
	w.len += len(bs)
}

func (w *Writer) GetFullBytes() []byte {
	return w.buf
}

func (w *Writer) Bytes(bs []byte) {
	if len(bs) > 0x7FFFFFFF {
		panic("string length overflows a 32-bit integer")
	}
	w.Int32(int32(len(bs)))
	w.writeBytes(bs)
}

func (w *Writer) UInt8(x uint8) {
	w.writeByte(byte(x))
}

func (w *Writer) Int8(x int8) {
	w.writeByte(byte(uint8(x)))
}

func (w *Writer) Int16(x int16) {
	for i := range 2 {
		w.writeByte(byte(x >> uint((1-i)*8)))
	}
}

func (w *Writer) UInt32(x uint32) {
	for i := range 4 {
		w.writeByte(byte(x >> uint((3-i)*8)))
	}
}

func (w *Writer) Int32(x int32) {
	for i := range 4 {
		w.writeByte(byte(x >> uint((3-i)*8)))
	}
}

func (w *Writer) UInt64(x int64) {
	for i := range 8 {
		w.writeByte(byte(x >> uint((7-i)*8)))
	}
}

func (w *Writer) Double(x float64) {
	buf := bytes.NewBuffer([]byte{})
	binary.Write(buf, binary.BigEndian, x)
	w.writeBytes(buf.Bytes())
}

func (w *Writer) Bool(x bool) {
	if x {
		w.writeByte(1)
	} else {
		w.writeByte(0)
	}
}

func (w *Writer) StringUTF(x string) {
	if len(x) > 0xFFFF {
		panic("string length overflows a 16-bit integer")
	}
	w.Int16(int16(len(x)))
	w.writeBytes([]byte(x))
}

func WriteSlice[T SupportMarshal](w *Writer, s []T) {
	w.Int32(int32(len(s)))
	for _, i := range s {
		i.Marshal(w)
	}
}
