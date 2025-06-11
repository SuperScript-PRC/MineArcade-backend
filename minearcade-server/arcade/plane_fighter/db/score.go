package db

import (
	"bytes"
	"encoding/binary"
)

type PlayerScore struct {
	Score int32
}

func (ps *PlayerScore) Marshal() []byte {
	buf := bytes.NewBuffer([]byte{})
	binary.Write(buf, binary.BigEndian, ps.Score)
	return buf.Bytes()
}

func (ps *PlayerScore) Unmarshal(bs []byte) error {
	buf := bytes.NewBuffer(bs)
	return binary.Read(buf, binary.BigEndian, &ps.Score)
}
