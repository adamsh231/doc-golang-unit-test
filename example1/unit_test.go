package example1

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// Do not use `panic` for testing
// Fail -> Meneruskan code selanjutya
// FailNow -> Meneruskan test selanjutnya
// Error -> Error == Fail + Comment
// Fatal -> Fatal == FailNow + Comment
// Assert -> Assert == Error + detail from package
// Require -> Require == Fatal + detail from package

//! no return value for unit testing
func TestHalo(t *testing.T) {
	result := Halo("Nana")

	if result != "Hello Nana!!" {
		t.Error("Halo Error gan")
	}
}

func TestHaloAssert(t *testing.T) {
	result := Halo("Nana")

	if result != "Hello Nana!!" {
		assert.Equal(t, "Hello Nana!!", result, "Halo Error gan")
	}
}
