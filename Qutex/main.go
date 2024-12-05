package main

// DO NOT USE ANY IMPORT

type Qutex struct {
	locked bool
}

func NewQutex() *Qutex {
	return &Qutex{locked: false}
}

func (q *Qutex) Lock() {
	for q.locked {
	}
	q.locked = true
}

func (q *Qutex) Unlock() {
	if q.locked {
		q.locked = false
	} else {
		panic("Q is not locked")
	}
}
