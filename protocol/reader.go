package protocol

import (
	"bytes"
	"encoding/binary"
	"fmt"
)

type Reader struct {
	buf    []byte
	offset int
	len    int
}

func (r *Reader) SetFullBytes(bs []byte) {
	r.buf = bs
	r.len = len(bs)
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

func ReadSlice[T any](r *Reader, f func(*T) error) ([]T, error) {
	var length int32
	var s []T
	var t T
	r.Int32(&length)
	for range length {
		err := f(&t)
		if err != nil {
			return nil, err
		}
		s = append(s, t)
	}
	return s, nil
}
