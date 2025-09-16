package mediator

type Mediator interface {
	CreateTopic()
	PublishMessage()
	ReadMessage()
}
