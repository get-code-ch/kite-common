package kite_common

import "fmt"

type (
	GenericEventObserver struct {
		Observer
		Id int
	}
)

func (o *GenericEventObserver) OnNotify(e Event, r Observer, receiver Address) {
	fmt.Printf("*** Notify %s***\n", e.Data)
}

func (o *GenericEventObserver) OnBroadcast(e Event) {
	fmt.Printf("*** Broadcast %s***\n", e.Data)
}

func (o *GenericEventObserver) Key() interface{} {
	return o.Id
}
