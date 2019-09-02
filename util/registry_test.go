package util

import (
	"encoding/json"
	"testing"
)

func TestRegistry(t *testing.T) {
	url := "http://test.server.com:80"
	cpus := 8
	rooms := 0
	version := "1.0"
	msg := Register(url, cpus, rooms, version)
	if msg != "added" {
		t.Error("Failed to add server")
	}

	list := Servers()
	if list[url] == nil {
		t.Errorf("%s was not Registered", url)
	}

	//adding same one should return updated msg
	msg = Register(url, cpus, rooms, version)
	if msg != "updated" {
		t.Error("Failed to update server")
	}

	//Need to redo this
	// timeout = 0
	// cleanup()
	// time.Sleep(time.Second * 5)
	// //delete(servers, url)
	// list = Servers()

	// if _, ok := list[url]; ok {
	// 	t.Errorf("Failed to clean up %s, %v, %v", url, ok, Registry)
	// }

	Register("http://server2.com", 16, 2, "1.1")

	data := List()

	result := make(map[string]interface{})
	err := json.Unmarshal([]byte(data), &result)
	if err != nil {
		t.Errorf("failed to unmarshal server data. %s", err)
	}

	if _, ok := result["http://server2.com"]; !ok {
		t.Errorf("server %s not in list data", "http://server2.com")
	}

	if _, ok := result[url]; !ok {
		t.Errorf("server %s not in list data", url)
	}

}
