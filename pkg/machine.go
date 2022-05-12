package pkg

type State interface {
	GetType() int
	OnEnter(o interface{})
	OnLeave(o interface{})
}

type Machine interface {
	IsState(s int) bool
	SwitchState(s State)
	GetState() State
}
