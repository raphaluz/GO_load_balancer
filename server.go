package main

import (
	"net/http/httputil"
	"net/url"
	"sync"
)

type Server struct {
	URL               *url.URL
	ActiveConnections int
	Mutex             sync.Mutex
	Healthy           bool
}

func (s *Server) Proxy() *httputil.ReverseProxy {
	return httputil.NewSingleHostReverseProxy(s.URL)
}

func nextServerLeastActive(servers []*Server) *Server {
	leastActiveConnections := -1
	leastActiveServer := servers[0]
	for _, server := range servers {
		server.Mutex.Lock()
		if (server.ActiveConnections < leastActiveConnections || leastActiveConnections == -1) && server.Healthy {
			leastActiveConnections = server.ActiveConnections
			leastActiveServer = server
		}
		server.Mutex.Unlock()
	}

	return leastActiveServer
}
