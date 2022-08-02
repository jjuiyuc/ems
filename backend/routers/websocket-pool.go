package routers

import (
	log "github.com/sirupsen/logrus"
)

// Pool godoc
type Pool struct {
	Register   chan *Client
	Unregister chan *Client
	Clients    map[*Client]bool
}

func newPool() *Pool {
	return &Pool{
		Register:   make(chan *Client),
		Unregister: make(chan *Client),
		Clients:    make(map[*Client]bool),
	}
}

func (pool *Pool) start() {
	for {
		select {
		case client := <-pool.Register:
			pool.Clients[client] = true
			log.Debug("size of connection pool: ", len(pool.Clients))
		case client := <-pool.Unregister:
			delete(pool.Clients, client)
			log.Debug("size of connection pool: ", len(pool.Clients))
		}
	}
}
