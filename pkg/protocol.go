package pkg

import "bytes"

type WriterProto interface {
	Write(*bytes.Buffer) bool
}

type ReaderProto interface {
	Read(*bytes.Buffer) bool
}
