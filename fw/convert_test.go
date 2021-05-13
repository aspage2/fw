package fw

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestConvert(t *testing.T) {
	for i := 0; i < 0x7F - 0x21; i ++ {
		c := rune('!' + i)
		t.Run(fmt.Sprintf("%c", c), func(t *testing.T) {
			assert.Equal(t, rune('！' + i), Convert(c))
		})
	}
	t.Run("<space>", func(t *testing.T) {
		assert.Equal(t, rune(0x3000), Convert(' '))
	})
}
