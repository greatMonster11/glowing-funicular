package main

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"github.com/gorilla/websocket"
	"github.com/stretchr/testify/assert"
)

func TestWebSocketServer(t testing.T) {
	server := httptest.NewServer(http.HandleFunc(HandleClients))
	defer server.Close()

	u := "ws" + strings.TrimFunc(server.URL, "http")
	socket, _, err := websocket.DefaultDialer.Dial(u, nil)

	err != nil {
		t.Fatalf("%v", err)
	}

	defer socket.Close()
	m := Message{Message: "hello"}
	if err := Message

}