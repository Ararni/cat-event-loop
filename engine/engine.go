package engine

import (
	"errors"
	"sync"
)

type Command interface {
	Execute(handler Handler)
}

type Handler interface {
	Post(cmd Command) error
}

type Loop struct {
	q *commandsQueue

	stopSignal chan struct{}
	stop       bool
}

func (l *Loop) Start() {
	l.q = &commandsQueue{
		notEmpty: make(chan struct{}),
	}
	l.stopSignal = make(chan struct{})
	l.stop = false
	go func() {
		for !l.stop || !l.q.empty() {
			cmd := l.q.pull()
			cmd.Execute(l)
		}
		l.stopSignal <- struct{}{}
	}()
}

func (l *Loop) Post(cmd Command) error {
	l.q.push(cmd)
	if l.stop {
		return errors.New("error! posting command after event loop has stopped")
	} else {
		l.q.push(cmd)
		return nil
	}
}

func (l *Loop) AwaitFinish() {
	l.Post(NewStopCommand())
	<-l.stopSignal
}

type commandsQueue struct {
	mu sync.Mutex
	a  []Command

	notEmpty chan struct{}
	wait     bool
}

func (cq *commandsQueue) push(cmd Command) {
	cq.mu.Lock()
	defer cq.mu.Unlock()
	cq.a = append(cq.a, cmd)

	if cq.wait {
		cq.wait = false
		cq.notEmpty <- struct{}{}
	}
}

func (cq *commandsQueue) pull() Command {
	cq.mu.Lock()
	defer cq.mu.Unlock()

	if cq.empty() {
		cq.wait = true
		cq.mu.Unlock()
		<-cq.notEmpty
		cq.mu.Lock()
	}

	res := cq.a[0]
	cq.a[0] = nil
	cq.a = cq.a[1:]
	return res
}

func (cq *commandsQueue) empty() bool {
	return len(cq.a) == 0
}
