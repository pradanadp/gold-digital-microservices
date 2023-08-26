package broker

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/pradanadp/gold-digital-microservices/topup-storage/repository"
	"github.com/pradanadp/gold-digital-microservices/topup-storage/service"
	"github.com/segmentio/kafka-go"
)

func SubscribeMessage(db *sql.DB) {
	r := kafka.NewReader(kafka.ReaderConfig{
		Brokers:   []string{"localhost:9092"},
		Topic:     "topup",
		Partition: 0,
		MaxBytes:  10e6,
	})

	offsetFilePath := "./kafka_last_read_offset.txt"
	lastOffsetStr, err := readLastOffsetFromFile(offsetFilePath)
	if err != nil {
		log.Println("Failed to read last offset from file:", err)
	}

	lastOffset, _ := strconv.ParseInt(lastOffsetStr, 10, 64)
	r.SetOffset(lastOffset)

	for {
		m, err := r.ReadMessage(context.Background())
		if err != nil {
			break
		}
		fmt.Printf("message at offset %d: %s = %s\n", m.Offset, string(m.Key), string(m.Value))

		lastOffset = m.Offset + 1 // Update the last read offset
		writeLastOffsetToFile(offsetFilePath, strconv.FormatInt(lastOffset, 10))

		transactionRepository := repository.NewTransactionRepository()
		transactionService := service.NewTransactionService(transactionRepository, db)
		transactionService.Create(context.Background(), m.Value)

		// No need to implement Account interface

		// accountRepository := repository.NewAccountRepository()
		// accountService := service.NewAccountService(accountRepository, db)
		// accountService.Update(context.Background(), m.Value)
	}

	if err := r.Close(); err != nil {
		log.Fatal("failed to close reader:", err)
	}
}

func readLastOffsetFromFile(filePath string) (string, error) {
	data, err := os.ReadFile(filePath)
	if err != nil {
		return "", err
	}
	return string(data), nil
}

func writeLastOffsetToFile(filePath, offset string) error {
	// return os.WriteFile(filePath, []byte(offset), 0644)
	file, err := os.Create(filePath)
	if err != nil {
		return err
	}
	defer file.Close()

	_, err = file.WriteString(offset)
	if err != nil {
		return err
	}

	if err := file.Sync(); err != nil {
		return err
	}

	return nil
}
