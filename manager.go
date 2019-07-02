package daikin

type Manager struct {
	SerialChan chan *ACConfig
	State      *ACConfig
}

func NewManager() *Manager {
	return &Manager{
		SerialChan: make(chan *ACConfig),
		State:      &ACConfig{},
	}
}
