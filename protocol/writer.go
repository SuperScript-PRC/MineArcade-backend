package protocol

import (
	"bytes"
	"encoding/binary"
)

type Writer struct {
	buf []byte
	len int
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

func (w *Writer) Int8(x int8) {
	w.writeByte(byte(x))
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
	bs := make([]byte, 8)
	buf := bytes.NewBuffer(bs)
	binary.Write(buf, binary.BigEndian, x)
	w.writeBytes(bs)
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

func WriteSlice[T any](w *Writer, s []T, f func(T)) {
	w.Int32(int32(len(s)))
	for _, i := range s {
		f(i)
	}
}

func WriteSliceWithNewMarshaler[T any](w *Writer, s []T, f func(w *Writer, _ *T)) {
	w.Int32(int32(len(s)))
	for _, i := range s {
		f(w, &i)
	}
}
