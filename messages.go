package messages

// Dispatcher dispatches messages to registered receivers.
type Dispatcher struct {
	receivers []IReceiver
}

// NewDispatch creates a new Dispatcher.
func NewDispatch() *Dispatcher {
	return &Dispatcher{}
}

// Register registers the provided receivers with the Dispatcher.
func (d *Dispatcher) Register(receivers ...IReceiver) {
	d.receivers = append(d.receivers, receivers...)
}

// Unregister unregisters the provided receivers with the Dispatcher.
func (d *Dispatcher) Unregister(receivers ...IReceiver) {
	for _, r := range receivers {
		for i, nr := range d.receivers {
			if nr == r {
				d.receivers[i] = nil
				d.receivers = append(d.receivers[:i], d.receivers[i+1:]...)
				break
			}
		}
	}
}

// Send sends the specified message to all Receivers in the Dispatcher.
func (d *Dispatcher) Send(msg IMessage) {

	for _, receiver := range d.receivers {
		d.handleMessage(msg, receiver.(IReceiver))
	}

}

// SendTo sends the specified message to all of the specified target IReceivers.
func (d *Dispatcher) SendTo(msg IMessage, targets ...IReceiver) {

	for _, receiver := range targets {
		d.handleMessage(msg, receiver)
	}

}

// handleMessage here handles messages that should be dispatched to target receivers, ensuring that the receiver
// does wish to subscribe to the specified message type.
func (d *Dispatcher) handleMessage(msg IMessage, target IReceiver) {
	if subscriber, ok := target.(ISubscriber); ok {
		subscriptions := subscriber.Subscribe()
		if msg.Type().Contains(subscriptions) {
			target.ReceiveMessage(msg)
		}
	} else {
		target.ReceiveMessage(msg)
	}
}

// An object implements IReceiver if it is capable of receiving a message.
type IReceiver interface {
	ReceiveMessage(msg IMessage)
}

// ISubscriber indicates an object subscribes to only a subset of all received Messages.
// The Subscribe() function returns the MessageType or MessageTypes (added together) that are desired.
// If no Subscribe() function is defined (so that the object does not fulfill ISubscriber), the object
// will receive all MessageTypes sent to it.
type ISubscriber interface {
	Subscribe() MessageType // Subscribe returns the MessageTypes (added together) that the IReceiver takes.
}

// MessageType is the type of message dispatched; it can be bitwise combined together using standard addition for subscriptions.
type MessageType uint64

// Contains returns if a MessageType contains another message type; this is done to allow for bitwise combination to determine
// what message types a Subscriber might be interested in.
func (msg MessageType) Contains(other MessageType) bool {
	return msg&other > 0
}

// IMessage indicates a contract for messages. IMessage-fulfilling objects can be anything internally, but they all must return a MessageType.
// MessageTypes should generally be constants that are bitwise add-able.
type IMessage interface {
	Type() MessageType
}

// Example message types:

// const (
// 	TypeMessageGameStart           MessageType = 1 << iota 	// 1
// 	TypeMessagePlayerHurt                             		// 2
// 	TypeMessageItemFound                         			// 4
// 	TypeMessageSceneChange                               	// 8
// )
