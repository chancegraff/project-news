package collector

import (
	"context"
	"encoding/gob"
	"fmt"
	"log"
	"net"

	"github.com/chancegraff/project-news/internal/utils"
	"github.com/jinzhu/gorm"
)

type request struct {
	route string
	data  interface{}
}

var store *gorm.DB

// handleRequest will take a request from a connection and write a response
func handleRequest(ctx context.Context, connection net.Conn) {
	select {
	case <-ctx.Done():
		connection.Close()
		return
	default:
		// Defer closing connection
		defer connection.Close()

		// Decode request
		var request request
		decoder := gob.NewDecoder(connection)
		err := decoder.Decode(&request)
		if err != nil {
			connection.Write([]byte(err.Error()))
			return
		}

		// Switch into function
		var result []byte
		switch request.route {
		case "all":
			result, err = all(request.data)
			break
		case "get":
			result, err = get(request.data)
			break
		}

		// Write response
		if err != nil {
			connection.Write([]byte(err.Error()))
		} else {
			connection.Write([]byte(result))
		}
	}
}

// handleConnections waits for a connection and sends it
func handleConnections(ctx context.Context, listener net.Listener) {
	for {
		select {
		case <-ctx.Done():
			return
		default:
			// Wait for connections
			connection, err := listener.Accept()
			if err != nil {
				log.Panicln("Error", err)
				break
			}

			// Handle connection in separate thread
			go handleRequest(ctx, connection)
		}
	}
}

// Listen exposes a port and watches it for requests
func Listen(ctx context.Context, db *gorm.DB) {
	// Set database
	store = db

	// Build address
	port := utils.GetEnv("COLLECTOR_PORT", "6384")
	address := fmt.Sprintf(":%s", port)

	// Bind to address
	listener, err := net.Listen("tcp4", address)
	if err != nil {
		log.Panicln("Error", err)
		return
	}

	// Handle connections
	go handleConnections(ctx, listener)

	// Shutdown when parent does
	select {
	case <-ctx.Done():
		listener.Close()
		return
	}
}
