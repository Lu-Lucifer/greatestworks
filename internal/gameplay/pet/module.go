package pet

import (
	"greatestworks/internal"
	"sync"
)

const (
	ModuleName = "pet"
)

var (
	Mod         *Module
	onceInitMod sync.Once
)

func init() {
	internal.MManager.RegisterModule(ModuleName, GetMod())
}

type Module struct {
	*internal.BaseModule
}

func GetMod() *Module {
	Mod = &Module{internal.NewBaseModule()}

	return Mod
}