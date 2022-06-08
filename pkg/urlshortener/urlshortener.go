package urlshortener

import (
	"github.com/segmentio/ksuid"
)

var CodeLength = 8

func Generate() string {
	randomId := ksuid.New().String()
	code := randomId[0:CodeLength]
	return code
}
