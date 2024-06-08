package main

import (
	"net/url"
	"sync"
)

type Server struct {
	URL               *url.URL
	ActiveConnections int
	Mutex             sync.Sync
	Healthy           bool
}