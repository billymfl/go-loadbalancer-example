package chatserver

import "time"

// ChatServer represents a chat server
type ChatServer struct {
	cpus      int       // number of cpu cores
	capacity  int       // how many rooms can still be created
	rooms     int       // number of active rooms
	name      string    // room name
	version   string    // version of app
	timestamp time.Time // heartbeat

}

// New creates a new ChatServer instance
func New(name string, rooms int, cpus int, version string) *ChatServer {
	cs := ChatServer{
		name:    name,
		rooms:   rooms,
		cpus:    cpus,
		version: version,
		//timestamp: time.Now().UTC(),
		//capacity: cpus - rooms,
	}
	cs.Updatecapacity()
	cs.Updatetimestamp()
	return &cs
}

// Name returns room's name
func (cs *ChatServer) Name() string {
	return cs.name
}

// Cpus returns number of cpus
func (cs *ChatServer) Cpus() int {
	return cs.cpus
}

// SetCpus sets cpu cores
func (cs *ChatServer) SetCpus(cpus int) {
	cs.cpus = cpus
}

// Rooms returns number of rooms
func (cs *ChatServer) Rooms() int {
	return cs.rooms
}

// SetRooms sets cpu cores
func (cs *ChatServer) SetRooms(rooms int) {
	cs.rooms = rooms
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
