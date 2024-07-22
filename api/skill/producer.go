package skill

import (
	"log"
	"os"

	_ "net/http/pprof"

	_ "github.com/rcrowley/go-metrics"

	"github.com/IBM/sarama"
)

// Sarama configuration options
var (
	topic    = os.Getenv("TOPIC")
	producer sarama.SyncProducer
)

func init() {
	if len(topic) == 0 {
		panic("no topic given to be consumed, please set the -topic flag")
	}
}

func Producer(message []byte, key string) error {
	config := sarama.NewConfig()
	config.Producer.Return.Successes = true
	config.Producer.Return.Errors = true
	config.Producer.Partitioner = sarama.NewRoundRobinPartitioner
	config.Producer.RequiredAcks = sarama.WaitForAll

	var err error
	producer, err = sarama.NewSyncProducer([]string{"localhost:9092", "localhost:9093", "localhost:9094"}, config)
	if err != nil {
		log.Fatalln(err)
	}
	defer func() {
		if err := producer.Close(); err != nil {
			log.Fatalln(err)
		}
	}()

	// for range 10 {
	msg := &sarama.ProducerMessage{Topic: "my-topic", Value: sarama.ByteEncoder(message), Key: sarama.StringEncoder(key)}
	partition, offset, error := producer.SendMessage(msg)
	if error != nil {
		log.Printf("FAILED to send message: %s\n", error)
		return error
	} else {
		log.Printf("> message sent to partition %d at offset %d\n", partition, offset)
		return nil
	}
	// }

	// return nil
	// for range 10 {
	// msg := &sarama.ProducerMessage{Topic: "my-topic", Value: sarama.StringEncoder("testing 123")}
	// partition, offset, err := producer.SendMessage(msg)
	// if err != nil {
	// 	log.Printf("FAILED to send message: %s\n", err)
	// } else {
	// 	log.Printf("> message sent to partition %d at offset %d\n", partition, offset)
	// }
	// }
}

// func SendMessage(message []byte, key string) error {
// 	msg := &sarama.ProducerMessage{Topic: "my-topic", Value: sarama.ByteEncoder(message), Key: sarama.StringEncoder(key)}
// 	partition, offset, err := producer.SendMessage(msg)
// 	if err != nil {
// 		log.Printf("FAILED to send message: %s\n", err)
// 		return err
// 	} else {
// 		log.Printf("> message sent to partition %d at offset %d\n", partition, offset)
// 		return nil
// 	}
// }
