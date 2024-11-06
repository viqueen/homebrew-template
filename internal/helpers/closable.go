package helpers

import (
	"log"
)

type Closable interface {
	Close() error
}

func GracefulClose(c Closable, handleErr func(err error)) {
	if err := c.Close(); err != nil {
		handleErr(err)
	}
}

func LogError(err error) {
	if err != nil {
		log.Printf("failed with %v", err)
	}
}
