package machines

import "gateserver/pkg"

type machineImpl struct {
	curState   pkg.State
	hostObject interface{}
}

func NewMachine[T any](host *T) pkg.Machine {
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

func (m *machineImpl) SwitchState(s pkg.State) {
	if m.curState != nil {
		m.curState.OnLeave(m.hostObject)
	}

	m.curState = s
	m.curState.OnEnter(m.hostObject)
}

func (m *machineImpl) GetState() pkg.State {
	return m.curState
}
