package snappass_core

import (
	"errors"
	"fmt"
)

type MockDatabase struct {
	kv map[string][]byte
}

type MockKeyGenerator struct {
	pfx []byte
	i   int
}

func NewMockDatabase() (*MockDatabase, error) {
	mockdb := new(MockDatabase)
	return mockdb, nil
}

func (mockdb *MockDatabase) connect() error {
	mockdb.kv = make(map[string][]byte)
	return nil
}

func (mockdb *MockDatabase) destroy() error {
	mockdb.kv = make(map[string][]byte)
	return nil
}

func (mockdb *MockDatabase) setAndExpire(key, value []byte, ttl TTL) error {
	mockdb.kv[string(key)] = value
	return nil
}

func (mockdb *MockDatabase) getAndDelete(key []byte) ([]byte, error) {
	v, ok := mockdb.kv[string(key)]
	if !ok {
		return nil, errors.New("Key not found.")
	}
	delete(mockdb.kv, string(key))
	return v, nil
}

func (mockdb *MockDatabase) isKeySet(key []byte) (bool, error) {
	_, ok := mockdb.kv[string(key)]
	if ok {
		return true, nil
	} else {
		return false, nil
	}
}

func NewMockKeyGenerator(pfx []byte) (*MockKeyGenerator, error) {
	mockkg := new(MockKeyGenerator)
	mockkg.pfx = pfx
	return mockkg, nil
}

func (mockkg *MockKeyGenerator) gen() ([]byte, error) {
	key := fmt.Sprintf("%s%d", mockkg.pfx, mockkg.i)
	mockkg.i += 1
	return []byte(key), nil
}
