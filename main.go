package main

import "log"

func main() {
	var p Publisher
	p = newPublisher()

	p.broadcast("hello")

	s := newSubscriber("123")
	s2 := newSubscriber("456")
	p.addSubscriber(s)
	p.addSubscriber(s2)
	p.broadcast("hello again")

	p.removeSubscriber(s.id())
	p.broadcast("good morning")

}

type Publisher interface {
	addSubscriber(subscriber Subscriber)
	removeSubscriber(subId string)
	broadcast(msg string)
}
type Subscriber interface {
	id() string
	react(msg string)
}

// Implementation
type publisher struct {
	subscribers map[string]Subscriber
}

func newPublisher() publisher {
	return publisher{subscribers: make(map[string]Subscriber)}
}

func (p publisher) addSubscriber(subscriber Subscriber) {
	p.subscribers[subscriber.id()] = subscriber
}

func (p publisher) removeSubscriber(subId string) {
	delete(p.subscribers, subId)
}

func (p publisher) broadcast(msg string) {
	for _, subscriber := range p.subscribers {
		subscriber.react(msg)
	}
}

// Implementation >> Subscriber

type subscriber struct {
	subId string
}

func newSubscriber(subId string) subscriber {
	return subscriber{subId: subId}
}
func (s subscriber) id() string {
	return s.subId
}

func (s subscriber) react(msg string) {
	log.Printf("ID %s - recived: %s", s.subId, msg)
}
