package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"os"

	ckafka "github.com/confluentinc/confluent-kafka-go/kafka"
	// _ "github.com/go-sql-driver/mysql"
	"github.com/igor-gregori/imersao5-gateway/adapter/broker/kafka"
	"github.com/igor-gregori/imersao5-gateway/adapter/factory"
	"github.com/igor-gregori/imersao5-gateway/adapter/presenter/transaction"
	"github.com/igor-gregori/imersao5-gateway/usecase/process_transaction"
	_ "github.com/mattn/go-sqlite3"
)

func main() {
	fmt.Println("🚀 Entrou aqui na maain")
	os.Setenv("BOOTSTRAP_SERVERS", "host.docker.internal:9094")

	db, err := sql.Open("sqlite3", "test.db")
	if err != nil {
		log.Fatal(err)
	}

	repositoryFactory := factory.NewRepositoryDatabaseFactory(db)
	repository := repositoryFactory.CreateTransactionRepository()
	configMapProducer := &ckafka.ConfigMap{
		"bootstrap.servers": "host.docker.internal:9094",
	}
	kafkaPresenter := transaction.NewTransactionKafkaPresenter()
	producer := kafka.NewKafkaProducer(configMapProducer, kafkaPresenter)
	fmt.Println("🚀 Iniciou o producer")

	var msgChan = make(chan *ckafka.Message)
	configMapConsumer := &ckafka.ConfigMap{
		"bootstrap.servers": "host.docker.internal:9094",

		"client.id": "goapp",
		"group.id":  "goapp",
	}
	topics := []string{"transactions"}
	consumer := kafka.NewConsumer(configMapConsumer, topics)
	fmt.Println("🚀 Iniciou o consumer")
	go consumer.Consume(msgChan)

	usecase := process_transaction.NewProcessTransaction(repository, producer, "transactions_result")
	fmt.Println("🚀 Instanciou o caso de uso")

	for msg := range msgChan {
		var input process_transaction.TransactionDtoInput
		json.Unmarshal(msg.Value, &input)
		usecase.Execute(input)
	}
}
