package eventsingo

//adds a listener for an event.
func AddListener(listener EventListener, event *Event, maxFires int){
	infiniteListen := false
	if maxFires == 0{
		infiniteListen = true
		maxFires++
	}
	running := make(chan bool)
	for (maxFires > 0) {
		go func() {
			event.Lock()
			running <- true
			event.ListenEvent()
			listener()
			event.Unlock()
		}()

		if (!infiniteListen) {
			maxFires--
		}
		//block till your events are suspended
		<-running
	}
}

//remove a listener from an event
func removeListener(listener EventListener, event *Event, maxFires int){
	//todo: have a ds for storing map to listener list
}

