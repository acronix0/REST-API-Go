package kafka

import "github.com/IBM/sarama"

type KafkaProducer struct {
	brokers []string
	producer sarama.SyncProducer
}

func NewProducer(brokers []string, retryMax int) (*KafkaProducer, error) {
	config := sarama.NewConfig()
  config.Producer.RequiredAcks = sarama.WaitForAll
  config.Producer.Retry.Max = retryMax
  config.Producer.Return.Successes = true

  producer, err := sarama.NewSyncProducer(brokers, config)
  if err!= nil {
    return nil, err
  }
	return &KafkaProducer{brokers: brokers, producer: producer}, nil
}

func (p *KafkaProducer) Close() error {
  return p.producer.Close()
}
func (p *KafkaProducer) SendMessage(topic string, key, value []byte) error {
	msg := &sarama.ProducerMessage{Topic: topic, Key: sarama.StringEncoder(key), Value: sarama.ByteEncoder(value)}
  _, _, err := p.producer.SendMessage(msg)
  return err
}