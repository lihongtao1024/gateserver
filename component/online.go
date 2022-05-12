package component

type Online interface {
	AddOnline(client Client)
	GetOnline(sid uint8, uid uint32) Client
	DeleteOnline(client Client)
	Close()
}
