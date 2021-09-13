package mytest_test

import (
	"github.com/stretchr/testify/assert"
	"mytest"
	"testing"
)

func TestAdd(t *testing.T) {
	assert.Equal(t, 2, mytest.Add(1, 1))

	// 故意出錯，看看錯誤訊息
	assert.NotEqual(t, 4, mytest.Add(1, 3))
}

func TestPanic(t *testing.T) {
	assert.Panics(t, func() { mytest.Panic() })

	// 故意出錯，看看錯誤訊息
	assert.NotPanics(t, func() { mytest.Panic() })
}

func TestLarge(t *testing.T) {
	assert.True(t, mytest.Large(1, -1))

	// 故意出錯，看看錯誤訊息
	assert.False(t, mytest.Large(1, -1))
}
