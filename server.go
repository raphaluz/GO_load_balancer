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
