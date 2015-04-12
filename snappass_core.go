package snappass_core

import ()

type TTL int

const (
	Hour = 3600
	Day  = 86400
	Week = 604800
)

type Database interface {
	connect() error
	destroy() error
	setAndExpire([]byte, []byte, TTL) error
	getAndDelete([]byte) ([]byte, error)
	isKeySet([]byte) (bool, error)
}

type KeyGenerator interface {
	gen() ([]byte, error)
}

type SnapPass struct {
	db Database
	kg KeyGenerator
}

func New(db Database, kg KeyGenerator) (*SnapPass, error) {
	snap := new(SnapPass)
	snap.db = db
	snap.kg = kg
	err := snap.db.connect()
	if err != nil {
		return nil, err
	}
	return snap, nil
}

func (snap *SnapPass) Close() error {
	return snap.db.destroy()
}

func (snap *SnapPass) SetPassword(pass []byte, ttl TTL) ([]byte, error) {
	key, err := snap.kg.gen()
	if err != nil {
		return nil, err
	}
	return key, snap.db.setAndExpire(key, pass, ttl)
}

func (snap *SnapPass) SetPasswordStrTTL(pass []byte, ttl string) ([]byte, error) {
	_ttl, err_ttl := str2ttl(ttl)
	if err_ttl != nil {
		return nil, err_ttl
	}
	return snap.SetPassword(pass, _ttl)
}

func (snap *SnapPass) GetPassword(key []byte) ([]byte, error) {
	return snap.db.getAndDelete(key)
}

func (snap *SnapPass) TestKey(key []byte) (bool, error) {
	return snap.db.isKeySet(key)
}
