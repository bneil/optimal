package config

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGetConfig(t *testing.T) {
	cfg := GetConfig()
	assert.Equal(t, cfg.App.Environment, "local", "they should match")
}
