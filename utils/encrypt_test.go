package utils

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestHashSHA256(t *testing.T) {
	stringToHash := "hello"
	hashExpected := "qvTGHdzF6KLavt4PO0gs2a6pQ00="
	hashGot := HashSHA256(stringToHash)
	assert.Equal(t, hashExpected, hashGot)
}
