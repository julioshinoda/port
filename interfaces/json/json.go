package json

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"sync"

	"github.com/julioshinoda/port/entity"
	"github.com/julioshinoda/port/service"
)

func ParseJson(ctx context.Context, filepath, dbURL string) {
	stream := NewJSONStream()
	portService, err := service.NewPortService(ctx, dbURL)
	if err != nil {
		log.Println(err)
		return
	}
	go func() {
		var wg sync.WaitGroup
		for data := range stream.Watch() {
			if data.Error != nil {
				log.Println(data.Error)
			}
			wg.Add(1)
			go func(wg *sync.WaitGroup) {
				if err := portService.PortDomainService(ctx, data.Port); err != nil {
					log.Println("error service:", err)
				}
				wg.Done()
			}(&wg)

			wg.Wait()
		}
	}()

	stream.Start(filepath)

}

// Entry represents stream itens
type Entry struct {
	Error error
	Port  entity.Port
}

// Stream helps transmit each streams withing a channel.
type Stream struct {
	stream chan Entry
}

// NewJSONStream returns a new `Stream` type.
func NewJSONStream() Stream {
	return Stream{
		stream: make(chan Entry),
	}
}

// Watch return a stream channel with Entry
func (s Stream) Watch() <-chan Entry {
	return s.stream
}

// Start starts streaming JSON file line by line and parse each json item
func (s Stream) Start(path string) {
	defer close(s.stream)

	file, err := os.Open(path)
	if err != nil {
		s.stream <- Entry{Error: fmt.Errorf("open file: %w", err)}
		return
	}
	defer file.Close()

	decoder := json.NewDecoder(file)

	_, err = decoder.Token()
	if err != nil {
		s.stream <- Entry{Error: fmt.Errorf("decode opening delimiter: %w", err)}
		return
	}

	for decoder.More() {

		token, err := decoder.Token()
		if err != nil {
			s.stream <- Entry{Error: fmt.Errorf("decode opening delimiter: %w", err)}
			return
		}

		var port entity.Port
		if err := decoder.Decode(&port); err != nil {
			s.stream <- Entry{Error: fmt.Errorf("decode line: %w", err)}
			return
		}
		port.ID = token.(string)

		s.stream <- Entry{Port: port}
	}

	if _, err := decoder.Token(); err != nil {
		s.stream <- Entry{Error: fmt.Errorf("decode closing delimiter: %w", err)}
		return
	}
}
