package util

import "time"

// ChatServer represents a chat server
type ChatServer struct {
	cpus      int       // number of cpu cores
	capacity  int       // how many rooms can still be created
	rooms     int       // number of active rooms
	url       string    // server url. Acts as unique key
	version   string    // version of app
	timestamp time.Time // heartbeat
}

// New creates a new ChatServer instance
func New(url string, cpus int, rooms int, version string) *ChatServer {
	cs := ChatServer{
		url:     url,
		rooms:   rooms,
		cpus:    cpus,
		version: version,
	}
	cs.Updatecapacity()
	cs.Updatetimestamp()
	return &cs
}

// URL returns url
func (cs *ChatServer) URL() string {
	return cs.url
}

// Cpus returns number of cpus
func (cs *ChatServer) Cpus() int {
	return cs.cpus
}

// Rooms returns number of rooms
func (cs *ChatServer) Rooms() int {
	return cs.rooms
}

// SetRooms sets cpu cores
func (cs *ChatServer) SetRooms(rooms int) {
	cs.rooms = rooms
	cs.Updatecapacity()
}

// Capacity returns capacity left on server
func (cs *ChatServer) Capacity() int {
	return cs.capacity
}

// Updatecapacity updates capaicty left on server
func (cs *ChatServer) Updatecapacity() {
	cs.capacity = cs.cpus - cs.rooms
}

// Updatetimestamp updates to current time
func (cs *ChatServer) Updatetimestamp() {
	cs.timestamp = time.Now().UTC()
}

// Timestamp returns timestamp
func (cs *ChatServer) Timestamp() time.Time {
	return cs.timestamp
}

// Version returns version
func (cs *ChatServer) Version() string {
	return cs.version
}
