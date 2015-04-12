package snappass_core

import (
	"errors"
	"strings"
)

func str2ttl(ttl string) (TTL, error) {
	switch strings.ToLower(ttl) {
	case "hour":
		return Hour, nil
	case "day":
		return Day, nil
	case "week":
		return Week, nil
	}
	return 0, errors.New("Unable to convert TTL")
}
