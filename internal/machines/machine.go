package machines

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

type machineImpl struct {
	curState   State
	hostObject interface{}
}

func NewMachine[T any](host *T) Machine {
	return &machineImpl{
		hostObject: host,
	}
}

func (m *machineImpl) IsState(s int) bool {
	if m.curState == nil {
		return false
	}

	return m.curState.GetType() == s
}

func (m *machineImpl) SwitchState(s State) {
	if m.curState != nil {
		m.curState.OnLeave(m.hostObject)
	}

	m.curState = s
	m.curState.OnEnter(m.hostObject)
}

func (m *machineImpl) GetState() State {
	return m.curState
}
