package family

import (
	"greatestworks/business/module"
)

var (
	Mod *Module
)

func init() {
	module.MManager.RegisterModule("", Mod)
}

type Module struct {
	families map[uint64]*Family
	IWorld
	ChIn  chan ManagerHandlerParam
	ChOut chan interface{}
}

func (m *Module) Loop() {
	for {
		select {
		case msg := <-m.ChOut:
			m.ForwardMsg(msg)
		}
	}
}

func (m *Module) Monitor() {
	for {
		select {
		case param := <-m.ChIn:
			m.Handler(param)
		}
	}
}