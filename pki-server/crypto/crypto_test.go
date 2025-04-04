package crypto

import (
	"fmt"
	"testing"
)

func TestGetInstance(t *testing.T) {
	sub := getCryptoInstance()
	fmt.Println(sub != nil)
}
