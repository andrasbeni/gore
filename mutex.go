package main

type response struct {
}

type request struct {
	response chan *response
}

func newRequest() *request {
	return &request{
		response: make(chan *response, 1),
	}
}

//****************************************************

type Mutex1 struct {
	lockRequests   chan *request
	unlockRequests chan *request
}

func NewMutex1() *Mutex1 {
	m := &Mutex1{
		lockRequests:   make(chan *request),
		unlockRequests: make(chan *request),
	}
	m.start()
	return m
}

func (m *Mutex1) Lock(who string) {
	l := newRequest()
	println(who + " acquiring lock")
	m.lockRequests <- l
	select {
	case _ = <-l.response:
		println(who + " acquired lock")
	}

}

func (m *Mutex1) Unlock(who string) {
	l := newRequest()
	println(who + " releasing lock")
	m.unlockRequests <- l
	select {
	case _ = <-l.response:
		println(who + " released lock")
	}
}

func (m *Mutex1) start() {
	go m.run()
}

func (m *Mutex1) run() {
	for {
		select {
		case l := <-m.lockRequests:
			l.response <- &response{}
		}
		select {
		case l := <-m.unlockRequests:
			l.response <- &response{}
		}
	}
}

// =========================================================

type Mutex2 struct {
	tokens chan *response
}

func NewMutex2() *Mutex2 {
	m := &Mutex2{
		tokens: make(chan *response, 1),
	}
	m.tokens <- &response{}
	return m
}

func (m *Mutex2) Lock(who string) {
	println(who + " acquiring lock")
	select {
	case _ = <-m.tokens:
		println(who + " acquired lock")
	}

}

func (m *Mutex2) Unlock(who string) {
	println(who + " releasing lock")
	m.tokens <- &response{}
	println(who + " released lock")
}
