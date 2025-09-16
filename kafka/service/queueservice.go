package service

type QueueService interface {
	CreateProducer()
	CreateConsumer()
	CreateTopic(topicName string)
}
