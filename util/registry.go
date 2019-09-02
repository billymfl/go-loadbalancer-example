package util

import (
	"encoding/json"
	"fmt"
	"time"
)

var servers = make(map[string]*ChatServer)
var timeout = 30

type jsonServer struct {
	Version   string    `json:"version"`
	Rooms     int       `json:"rooms"`
	Cpus      int       `json:"cpus"`
	Capacity  int       `json:"capacity"`
	Timestamp time.Time `json:"timestamp"`
}

func newjsonServer(cs *ChatServer) *jsonServer {
	return &jsonServer{
		Rooms:     cs.Rooms(),
		Version:   cs.Version(),
		Cpus:      cs.Cpus(),
		Capacity:  cs.Capacity(),
		Timestamp: cs.Timestamp(),
	}
}

func init() {
	ticker := time.NewTicker(30 * time.Second)
	quit := make(chan bool)
	go func() {
		for {
			select {
			case <-ticker.C:
				cleanup()
			case <-quit:
				ticker.Stop()
				return
			}
		}
	}()
}

// Register adds or updates a chat server's info in our registry and returns if it was added or updated
func Register(url string, cpus int, rooms int, version string) string {
	cleanup()

	var cs *ChatServer
	msg := "updated"
	if _, ok := servers[url]; ok {
		cs = servers[url]
		cs.Updatetimestamp()
		cs.SetRooms(rooms)
	} else {
		cs = New(url, cpus, rooms, version)
		servers[url] = cs
		msg = "added"
	}
	return msg
}

// Servers returns a map of servers registered
func Servers() map[string]*ChatServer {
	return servers
}

// List returns string of servers in json format
func List() (string, error) {
	m := make(map[string]interface{})

	for url, cs := range servers {
		m[url] = newjsonServer(cs)
	}

	json, err := json.Marshal(m)
	if err != nil {
		//fmt.Printf("json error %s", err)
		return "", err
	}

	return string(json), nil
}

func cleanup() {
	fmt.Println("cleanup")
	now := time.Now().UTC()
	for url, cs := range servers {
		timestamp := cs.Timestamp().Add(time.Second * time.Duration(timeout))
		if timestamp.Before(now) {
			Unregister(url)
		}
	}
}

// Unregister removes a server
func Unregister(url string) {
	delete(servers, url)
}
