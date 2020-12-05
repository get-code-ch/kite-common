package kite_common

type (
	// Observer defines a standard interface for instances that wish to list for
	// the occurrence of a specific event.
	Observer interface {
		// OnNotify allows an event to be "published" to interface implementations.
		// In the "real world", error handling would likely be implemented.
		OnNotify(Event, Observer, Endpoint)
		OnBroadcast(Event)
		Key() interface{}
	}

	// Event defines an indication of a point-in-time occurrence.
	Event struct {
		Data interface{}
	}
)
