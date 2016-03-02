package eventsingo

import(
	"errors"
	"sync"
	"crypto/rand"
	"encoding/base64"
)

//Event listeners take no input and give no output, they just listen and run
type EventListener func()

func NewEvent() (*Event, error){
	event := &Event{
		id_size:8,
	}
	err := event.New()
	if err != nil{
		return nil, err
	}
	return event, nil
}

type Event struct{
	m sync.Mutex
	trigger *sync.Cond //a condition variable is the root concept behind events
	id_size uint8
	id string
}

func (e *Event) New() error{
	e.trigger = sync.NewCond(&(e.m))
	err := e.setID()
	if err != nil{
		return err
	}
	return nil
}

func (e *Event) FireEvent(){
	e.trigger.Signal()
}

func (e *Event) ExplodeEvent(){
	e.trigger.Broadcast()
}

func (e *Event) ListenEvent(){
	e.trigger.Wait()
}

func (e * Event) Lock(){
	e.m.Lock()
}

func (e * Event) Unlock(){
	e.m.Unlock()
}

//get a base64 encoding for an array of random bytes
func (e * Event) setID() error{
	size := e.id_size // change the length of the generated random string here

	rb := make([]byte,size)
	_, err := rand.Read(rb)

	if err != nil {
		return errors.New("Error in setting the event ID")
	}

	rs := base64.URLEncoding.EncodeToString(rb)
	e.id = rs
	return nil
}

func (e * Event) GetID() string{
	return  e.id
}
