package httpconf

import (
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
	"time"
)

func TestGetConfig(t *testing.T) {
	expected := &ServerConfig{}
	expected.Server.Host = "localhost"
	expected.Server.Port = "8080"
	expected.Server.Timeout.Write = time.Minute
	expected.Server.Timeout.Read = time.Minute
	expected.Server.Timeout.Idle = time.Minute * 5

	got, err := GetConfig("example.yaml")
	require.NoError(t, err)
	assert.Equal(t, expected, got)
}
