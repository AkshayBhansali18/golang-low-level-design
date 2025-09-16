package mediator

import (
	"fmt"
	"golang-low-level-design/kafka/entities"
)

type MediatorImpl struct {
	topicList map[string][]*entities.Message
}

func NewMediator() *MediatorImpl {
	return &MediatorImpl{
		topicList: make(map[string][]*entities.Message),
	}
}

func (t *MediatorImpl) CreateTopic(name string) error {
	t.topicList[name] = make([]*entities.Message, 0)
	return nil
}

func (s *MediatorImpl) Subscribe(subscribedTopics map[string]uint32, topic string) error {
	subscribedTopics[topic] = 0
	return nil
}

func (p *MediatorImpl) Publish(topic string, message string) error {
	_, ok := p.topicList[topic]
	if ok {
		p.topicList[topic] = append(p.topicList[topic], &entities.Message{
			Topic:   topic,
			Message: message,
		})
		return nil
	}
	return fmt.Errorf("topic %s not found", topic)
}
