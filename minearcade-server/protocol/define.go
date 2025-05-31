package protocol

type SupportUnmarshal interface {
	Unmarshal(r *Reader)
}
