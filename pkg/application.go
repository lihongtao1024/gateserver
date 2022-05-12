package pkg

type Application interface {
	GetType() ServerType
	GetLogicName() string
	GetIndex() int
	GetEnvir() string
	GetRid() Guid
	SetRid(guid Guid)
	Start(typ1 ServerType, idx int, env string)
	Stop()
	IsClosed() bool
}

type ApplicationDispatcher interface {
	OnInit() bool
	OnUninit()
	OnClosing() bool
	OnWorking()
}
