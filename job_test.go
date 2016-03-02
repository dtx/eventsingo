package eventsingo

import (
	"testing"
	"fmt"
	"github.com/stretchr/testify/assert"
)

//Output:
//Hello World
//Just heard an event
//Just heard an event
//Bye World
func TestAddListener(t *testing.T){
	fmt.Println("Hello World")
	n:=2
	done := make(chan bool, n)
	//create an event
	//event := &Event{}
	event, err := NewEvent()
	if err != nil{
		t.Fatalf("Could not create new event")
	}
	//add a Listener to it to get invoked at most 2 times.
	AddListener(func() {fmt.Println("Just heard an event"); done <- true}, event, 2)
	for n > 0{
		event.Lock()
		event.FireEvent()
		event.Unlock()
		//groutine comes out of suspension writes to channel
		//or the code goes into deadlock
		<-done
		select {
		//make sure only 1 goroutine has awoken
		case <-done:
			fmt.Println("Too many goroutines awake")
		default:
		}
		n--
	}
	fmt.Println("Bye World")
}

func TestGetID(t *testing.T){
	event, err := NewEvent()
	if err != nil{
		t.Fatalf("Could not create new event")
	}
	id := event.GetID()
	assert.IsType(t, "string", id)
}