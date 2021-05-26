package landmap_test

import (
	"testing"

	"github.com/kiselev-nikolay/find-lands-example-go/landmap"
)

func TestLandmap(t *testing.T) {
	assertLandmap := func(expectedOutput int, input [][]int) {
		lm := landmap.New(input, false)
		output := lm.FindLands()
		t.Logf("ops=%d", lm.GetOperationsCount())
		if output != expectedOutput {
			t.Logf("expected %d, got %d\n", expectedOutput, output)
			t.Fail()
		}
	}
	t.Parallel()
	t.Run("Test with simple map", func(t *testing.T) {
		assertLandmap(1, [][]int{
			{1, 0, 0},
			{1, 1, 1},
			{0, 0, 0},
		})
	})
	t.Run("Test with 4x4 map", func(t *testing.T) {
		assertLandmap(3, [][]int{
			{0, 1, 0, 1},
			{1, 1, 0, 0},
			{0, 0, 0, 1},
			{1, 1, 1, 1},
		})
	})
	t.Run("Test with divided lands map", func(t *testing.T) {
		assertLandmap(3, [][]int{
			{0, 0, 0},
			{1, 0, 1},
			{0, 1, 0},
		})
	})
}
