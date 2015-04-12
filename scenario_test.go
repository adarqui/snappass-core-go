package snappass_core

import (
	"bytes"
	"fmt"
	"testing"
)

func TestMock(t *testing.T) {
	mockdb, _ := NewMockDatabase()
	mockkg, _ := NewMockKeyGenerator([]byte("snap:"))
	testScenario("mock", mockdb, mockkg, t)
}

func TestRedis(t *testing.T) {
	redisdb, err := NewRedisDatabase("127.0.0.1:6379", "", 0)
	if err != nil {
		t.Fatal("TestRedis: Error: NewRedisDatabase", err)
	}
	uuidkg, _ := NewUUIDKeyGenerator([]byte("snap:"))
	testScenario("redis/uuid", redisdb, uuidkg, t)
}

func testScenario(name string, db Database, kg KeyGenerator, t *testing.T) {
	if name == "" {
		t.Fatal(`"TestScenario: Error: name == ""`)
	}
	t_sc := fmt.Sprintf("TestScenario (%s): Error: ", name)
	snap, err := New(db, kg)
	if err != nil {
		t.Fatal(t_sc+"New", err)
	}
	if snap == nil {
		t.Fatal(t_sc + "snap is nil")
	}

	key, err_setPassword := snap.SetPassword([]byte("password"), Day)
	if err_setPassword != nil {
		t.Fatal(t_sc+"SetPassword", err_setPassword)
	}
	if key == nil {
		t.Fatal(t_sc+`"SetPassword: key == ""`, key)
	}

	v, err_getPassword := snap.GetPassword(key)
	if err_getPassword != nil {
		t.Fatal(t_sc+"GetPassword", err_getPassword)
	}

	if bytes.Equal(v, []byte("password")) != true {
		t.Fatal(t_sc + "Password does not match")
	}

	b, err_isKeySet := snap.TestKey(key)
	if err_isKeySet != nil {
		t.Fatal(t_sc+"TestKey", err_isKeySet)
	}

	if b == true {
		t.Fatal(t_sc + "Key is still set when it should have been deleted.")
	}
}
