package kafka

import (
	"codegen-service/internal/config"
	"github.com/confluentinc/confluent-kafka-go/kafka"
	"log/slog"
)

type KafkaProducer struct {
	producer      *kafka.Producer
	log           *slog.Logger
	schemaManager *SchemaManager
}

func NewKafkaProducer(
	cfg *config.Config,
	log *slog.Logger,
	schemaManager *SchemaManager,
) *KafkaProducer {
	producer, err := kafka.NewProducer(&kafka.ConfigMap{"bootstrap.servers": cfg.KafkaHost})
	if err != nil {
		panic("Error creating kafka producer")
	}

	return &KafkaProducer{
		producer:      producer,
		schemaManager: schemaManager,
		log:           log,
	}
}

func (kp *KafkaProducer) ProduceNewZip(key string, native map[string]interface{}) error {
	kp.log.Info("Sending NewZip message: ", key, native)

	err := kp.produce("NewZip", key, native)
	if err != nil {
		kp.log.Error("Error sending NewZip message: ", key, err)
		return err
	}

	kp.log.Error("Successfully sent NewZip message: ", key)
	return nil
}

func (kp *KafkaProducer) produce(topic string, key string, native map[string]interface{}) error {
	codec, err := kp.schemaManager.GetCodec(topic)
	if err != nil {
		return err
	}

	binValue, err := codec.TextualFromNative(nil, native)
	if err != nil {
		return err
	}

	msg := &kafka.Message{
		TopicPartition: kafka.TopicPartition{Topic: &topic, Partition: kafka.PartitionAny},
		Key:            []byte(key),
		Value:          binValue,
	}

	err = kp.producer.Produce(msg, nil)
	return err
}
