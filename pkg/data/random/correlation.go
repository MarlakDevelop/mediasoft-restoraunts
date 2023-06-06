package random

import (
	"math/rand"
	"strconv"
)

func GetCorrelationId() string {
	return strconv.Itoa(rand.Intn(9999999999))
}
