package snappass_core

import (
	"code.google.com/p/go-uuid/uuid"
	"fmt"
)

type UUIDKeyGenerator struct {
	pfx []byte
}

func NewUUIDKeyGenerator(pfx []byte) (*UUIDKeyGenerator, error) {
	uuidkg := new(UUIDKeyGenerator)
	uuidkg.pfx = pfx
	return uuidkg, nil
}

func (uuidkg *UUIDKeyGenerator) gen() ([]byte, error) {
	_uuid := uuid.NewRandom()
	key := fmt.Sprintf("%s%s", uuidkg.pfx, _uuid)
	return []byte(key), nil
}
