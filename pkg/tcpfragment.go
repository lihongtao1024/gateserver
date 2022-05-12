package pkg

type TcpFragmentType int

const (
	TcpFragmentFatal = TcpFragmentType(iota - 1)
	TcpFragmentContinue
)

type TcpFragment interface {
	OnUnpack(data []byte) int
}
