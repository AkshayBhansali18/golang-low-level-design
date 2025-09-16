package consumer

import "golang-low-level-design/kafka/mediator"

var subscriberTopics []string

type ConsumerImpl struct {
	Mediator mediator.Mediator
}

func NewConsumer(mediator mediator.Mediator) Consumer {
	subscriberTopics = []string{}
	return &ConsumerImpl{
		Mediator: mediator,
	}
}

func (c *ConsumerImpl) readMessage() {
	subscriberTopics = append(subscriberTopics, topic)
}
