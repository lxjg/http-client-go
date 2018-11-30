package sdk

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewClient(t *testing.T) {
	_, err := NewClient()
	assert.Nil(t, err)
}

func TestNewClientWithAccessToken(t *testing.T) {
	_, err := NewClientWithAccessToken("1", "QfCAH04Cob7b71QCqy738vw5XGSnFZ9d")
	assert.Nil(t, err)
}
