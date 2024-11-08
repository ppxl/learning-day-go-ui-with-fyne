package main

import (
	"fyne.io/fyne/v2/test"
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_makeUI(t *testing.T) {
	t.Run("initial text should be equal", func(t *testing.T) {
		out, _ := makeUI()

		assert.Equal(t, "Hello world!", out.Text)
	})
	t.Run("should update the given text", func(t *testing.T) {
		out, in := makeUI()

		test.Type(in, "Test-äöüß")
		assert.Equal(t, "Hello Test-äöüß!", out.Text)
	})

}
