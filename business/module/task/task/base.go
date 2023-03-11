package task

import (
	event2 "greatestworks/aop/event"
	"greatestworks/business/module/task"
)

type Base struct {
	Id       uint64
	ConfigId uint32
}

func (b *Base) SetStatus(status task.Status) {
}

func (b *Base) OnEvent(event event2.IEvent) {

}
