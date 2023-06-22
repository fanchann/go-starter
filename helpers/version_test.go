package helpers_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/fanchann/go-starter/helpers"
)

func TestGetGoVersion(t *testing.T) {
	expectedVersion := "1.20"

	version := helpers.GetGoVersion()

	assert.Equal(t, expectedVersion, version, "Versi Go tidak sesuai")
}
