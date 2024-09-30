package Mock

import (
	"github.com/stretchr/testify/mock"
)

type Log struct {
	mock.Mock
}

func (m *Log) LogError(value string) {
	m.Called(value)
}

func (m *Log) LogDebug(value string) {
	m.Called(value)
}

func (m *Log) LogInfo(value string) {
	m.Called(value)
}
