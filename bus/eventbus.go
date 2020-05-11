package bus

import (
	"fmt"
	"sync"
)

type DataEvent struct {
	Data interface{}
	Topic string
}

type DataChannel chan DataEvent

type DataSliceChannel []DataChannel

type EventBus struct {
	subscribers map[string] DataSliceChannel
	rm sync.Mutex
}

func (eb *EventBus) Publish(topic string, data interface{})  {
	eb.rm.Lock()
	if chans, found := eb.subscribers[topic]; found {
		channels := append(DataSliceChannel{}, chans...)
		go func(data DataEvent, dataSliceChannel DataSliceChannel) {
			for _,ch := range dataSliceChannel  {
				ch <- data
			}
		}(DataEvent{Data: data, Topic:topic}, channels)
	}
	eb.rm.Unlock()
}

func (eb *EventBus)Subscribe(topic string, ch DataChannel)  {
	eb.rm.Lock()
	if prev, found := eb.subscribers[topic]; found{
		eb.subscribers[topic] = append(prev, ch)
	} else {
		eb.subscribers[topic] = append([]DataChannel{}, ch)
	}
	eb.rm.Unlock()
}

var eb = &EventBus{
	subscribers: map[string]DataSliceChannel{},
}
func printDataEvent(ch string, data DataEvent)  {
	fmt.Println("Channle: %s; Topic %s; DataEvent: %v\n", ch, data.Topic, data.Data)
}

func publishTo(topic string, data string)  {
	for {
		eb.Publish(topic, data)
	}
}

func example()  {
	ch1 := make(chan DataEvent)
	ch2 := make(chan DataEvent)
	ch3 := make(chan DataEvent)

	eb.Subscribe("topic1", ch1)
	eb.Subscribe("topic2", ch2)
	eb.Subscribe("topic2", ch3)

	go publishTo("topic1", "Hi topic 1")
	go publishTo("topic2", "welcome to topic 2")

	for {
		select {
		case d := <- ch1:
			go printDataEvent("ch1", d)
		}
	}
}