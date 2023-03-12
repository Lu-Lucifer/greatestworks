package email

import (
	"greatestworks/business/module"
	"sync"
)

const (
	ModuleName = "email"
)

var (
	Mod         *Module
	onceInitMod sync.Once
)

func init() {
	module.MManager.RegisterModule(ModuleName, GetMod())
}

type Module struct {
	*module.BaseModule
}

func GetMod() *Module {
	Mod = &Module{module.NewBaseModule()}

	return Mod
}
