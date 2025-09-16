package producer

import "golang-low-level-design/kafka/mediator"

type ProducerImpl struct {
	mediator mediator.Mediator
}

func (p *ProducerImpl) Publish(topic string, message string) error {
	p.mediator.PublishMessage()
}
