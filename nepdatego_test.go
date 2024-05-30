package nepdatego_test

import (
	nepdatego "nepDate"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestParse(t *testing.T) {
	input := "2024/5/29 - 14:04:49"
	pattern := "%Y/%M/%D - %h:%m:%s"
	date := nepdatego.GetDate(input, pattern, time.UTC)

	assert.EqualValues(t, date, time.Date(2082, 2, 14, 14, 4, 49, 0, time.UTC))
}
