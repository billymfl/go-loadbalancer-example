package util

import (
	"testing"
	"time"
)

func TestChatServer(t *testing.T) {
	url := "http://test.server.com:80"
	cpus := 8
	rooms := 0
	version := "1.0"

	cs := New(url, cpus, rooms, version)

	//Note: might not be deterministic enough
	tme := time.Now().UTC()
	cs.Updatetimestamp()

	if cs.URL() != url {
		t.Error("Url doesn't match")
	}
	if cs.Rooms() != rooms {
		t.Error("Rooms doesn't match")
	}

	if cs.Cpus() != cpus {
		t.Error("Cpus doesn't match")
	}
	if cs.Version() != version {
		t.Error("Version doesn't match")
	}
	if cs.Timestamp() != tme {
		t.Errorf("Timestamp doesn't match: %s = tme: %s", cs.Timestamp(), tme)
	}
	if cs.Capacity() != 8 {
		t.Errorf("Capacity doesn't match. Expected 8, actual %d", cs.Capacity())
	}

	cs.SetRooms(2)
	if cs.Capacity() != 6 {
		t.Errorf("Capacity doesn't match. Expected 6, actual %d", cs.Capacity())
	}

	time.Sleep(200 * time.Millisecond)
	cs.Updatetimestamp()
	if cs.Timestamp() == tme {
		t.Errorf("Updatetimestamp: %s = tme: %s", cs.Timestamp(), tme)
	}
}
