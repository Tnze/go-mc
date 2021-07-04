package screen

type EventsListener struct {
	Open    func(id int) error
	SetSlot func(id, index int) error
	Close   func(id int) error
}
