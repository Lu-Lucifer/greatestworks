package pet

import (
	"greatestworks/aop/event"
	"greatestworks/business/module"
)

type EventHandle func(iEvent event.IEvent)

type EventWrap struct {
	IPlayer
	event.IEvent
}

func (m Module) OnEvent(c module.Character, event event.IEvent) {
	//TODO implement me
	panic("implement me")
}

func (m Module) SetEventCategoryActive(eventCategory int) {
	//TODO implement me
	panic("implement me")
}
