package venv

import (
	"testing"

	"github.com/georgettica/venv/mock"
	"github.com/georgettica/venv/os"
	"github.com/stretchr/testify/assert"
)

func TestOS(t *testing.T) {
	assert.IsType(t, &os.OsEnv{}, OS())
}

func TestMock(t *testing.T) {
	assert.IsType(t, &mock.MockEnv{}, Mock())
}
