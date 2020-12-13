package kite_common

type (
	// Notifier is the instance being observed. Publisher is perhaps another decent
	// name, but naming things is hard.
	Notifier interface {
		// Register allows an instance to register itself to listen/observe
		// events.
		Register(Observer)
		// Deregister allows an instance to remove itself from the collection
		// of observers/listeners.
		Deregister(Observer)
		// Notify publishes new events to listeners. The method is not
		// absolutely necessary, as each implementation could define this itself
		// without losing functionality.
		Notify(Event, Observer)
		// Broadcast broadcast events to all connected endpoints.
		Close(Event)
	}

	EventNotifier struct {
		// Using a map with an empty struct allows us to keep the observers
		// unique while still keeping memory usage relatively low.
		Observers map[Observer]struct{}
	}
)

func (n *EventNotifier) Register(l Observer) {
	n.Observers[l] = struct{}{}
}

func (n *EventNotifier) Deregister(l Observer) {
	delete(n.Observers, l)
}

func (n *EventNotifier) Notify(e Event, l Observer, receiver Endpoint) {
	for o := range n.Observers {
		o.OnNotify(e, l, receiver)
	}
}

func (n *EventNotifier) Close(e Event) {
	for o := range n.Observers {
		o.OnClose(e)
	}
}
