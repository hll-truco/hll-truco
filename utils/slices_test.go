package utils_test

import (
	"testing"

	"github.com/hll-truco/hll-truco/utils"
)

func TestCopyWithoutThese(t *testing.T) {
	xs := []int{1, 2, 3, 4, 5}
	ys := []int{3, 5}
	zs := utils.CopyWithoutThese(xs, ys...)
	if ok := zs[0] == 1 && zs[1] == 2 && zs[2] == 4 && len(zs) == 3; !ok {
		t.Error("el resultado no es el esperado")
	}
}
