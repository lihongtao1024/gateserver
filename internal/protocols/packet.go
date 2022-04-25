package protocols

type ClientHandShakeReq struct {
	Version uint32
	Zone    uint32
}

type ClientHandShakeRsp struct {
	Return uint32
	Load   uint32
	Delay  uint32
}

type ServerHandShakeReq struct {
	Type    uint16
	Index   uint16
	Zone    uint32
	Version uint16
	Lenght  uint16
	Extend  [32]byte
}

type ServerHandShakeRsp struct {
	Index   uint32
	Zone    uint32
	Version uint16
	Lenght  uint16
	Extend  [32]byte
}

type C2SProtoHeader struct {
	Sid uint32
	Uid uint32
	Aid uint32
}

type S2CProtoHeader struct {
	Client   uint8
	Reserved uint8
	Length   uint16
}

type S2CArray struct {
	Uid      uint32
	Sid      uint8
	Reserved uint8
	Aid      uint16
}
