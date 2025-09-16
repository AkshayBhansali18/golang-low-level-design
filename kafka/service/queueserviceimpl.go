package service

import "golang-low-level-design/kafka/mediator"

type QueueServiceImpl struct {
}

func (q *QueueServiceImpl) CreatePublisher(name string) mediator.Mediator {
	publisher := mediator.NewMediator()
	return publisher
}

func (q *QueueServiceImpl) CreateSubscriber() mediator.Mediator {
	go func() {
		for {
			for topicName,
		}
	}()
}

