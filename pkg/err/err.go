package err

import (
	"errors"
	"log"
)

var (
	NotSet = errors.New("environment has not been set")
)

func Throw(msg error, key string) {
	log.Panicf("%v: %s", msg, key)
}
