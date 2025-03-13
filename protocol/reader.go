package protocol

import (
	"bytes"
	"encoding/binary"
	"fmt"
)

type SupportUnmarshal interface {
	Unmarshal(r *Reader)
}

type Reader struct {
	buf    []byte
	offset int
	len    int
}

func NewReader(bs []byte) Reader {
	r := Reader{}
	r.buf = bs
	r.len = len(bs)
	return r
}

func (r *Reader) AddFullBytes(bs []byte, len int) {
	r.buf = bs
	r.len = len
}

func (r *Reader) SetFullBytes(bs []byte, len int) {
	r.offset = 0
	r.buf = bs
	r.len = len
}

func (r *Reader) readByte() (byte, error) {
	if r.offset >= r.len {
		return 0, fmt.Errorf("not enough bytes")
	}
	b := r.buf[r.offset]
	r.offset++
	return b, nil
}

func (r *Reader) readBytes(n int) ([]byte, error) {
	if r.offset+n > r.len {
		return []byte{}, fmt.Errorf("not enough bytes")
	}
	bs := r.buf[r.offset : r.offset+n]
	r.offset += n
	return bs, nil
}

func (r *Reader) Bytes(bs *[]byte) error {
	var length int32
	err := r.Int32(&length)
	if err != nil {
		return err
	}
	*bs, err = r.readBytes(int(length))
	if err != nil {
		return err
	}
	return nil
}

func (r *Reader) UInt8(x *uint8) error {
	b, err := r.readByte()
	if err != nil {
		return err
	}
	*x = b
	return nil
}

func (r *Reader) Int8(x *int8) error {
	b, err := r.readByte()
	if err != nil {
		return err
	}
	*x = int8(b)
	return nil
}

func (r *Reader) Int16(x *int16) error {
	var res int16
	for i := range 2 {
		b, err := r.readByte()
		if err != nil {
			return err
		}
		res |= int16(b) << uint((1-i)*8)
	}
	*x = res
	return nil
}

func (r *Reader) UInt32(x *uint32) error {
	var res uint32
	for i := range 4 {
		b, err := r.readByte()
		if err != nil {
			return err
		}
		res |= uint32(b) << uint((3-i)*8)
	}
	*x = res
	return nil
}

func (r *Reader) Int32(x *int32) error {
	var res int32
	for i := range 4 {
		b, err := r.readByte()
		if err != nil {
			return err
		}
		res |= int32(b) << uint((3-i)*8)
	}
	*x = res
	return nil
}

func (r *Reader) Double(x *float64) error {
	var f float64
	bs, err := r.readBytes(8)
	if err != nil {
		return err
	}
	buf := bytes.NewBuffer(bs)
	err = binary.Read(buf, binary.BigEndian, &f)
	if err != nil {
		return err
	}
	*x = f
	return nil
}

func (r *Reader) Bool(x *bool) error {
	b, err := r.readByte()
	if err != nil {
		return err
	}
	*x = b != 0
	return nil
}

func (r *Reader) StringUTF(x *string) error {
	var length int16
	err := r.Int16(&length)
	if err != nil {
		return err
	}
	if length <= 0 {
		*x = ""
		return nil
	}
	strbytes, err := r.readBytes(int(length))
	if err != nil {
		return err
	}
	*x = string(strbytes)
	return nil
}

func (r *Reader) End() bool {
	return r.offset >= r.len
}

func ReadSlice[T any, TPtr interface {
	*T
	SupportUnmarshal
}](r *Reader, s *[]T) (e error) {
	defer func() {
		if err := recover(); err != nil {
			*s = nil
			e = err.(error)
		}
	}()
	var length int32
	arr := make([]T, length)
	r.Int32(&length)
	for i := range length {
		var t T
		var tp TPtr = &t
		tp.Unmarshal(r)
		arr[i] = *tp
	}
	*s = arr
	return
}
