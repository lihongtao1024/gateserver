package networks

const (
	FragmentFatal    = -1
	FragmentContinue = 0
)

type Fragment interface {
	OnUnpack(data []byte) int
}
