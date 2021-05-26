package landmap

import (
	"sync"
	"sync/atomic"
)

type LandMap struct {
	ops        uint32
	data       [][]int
	mutex      *sync.Mutex
	transforms [][]int
}

func New(landMapData [][]int, allowDiagonalIslands bool) *LandMap {
	// copy
	data := make([][]int, len(landMapData))
	for y, lat := range landMapData {
		data[y] = make([]int, len(lat))
		for x, point := range lat {
			data[y][x] = point
		}
	}
	lm := &LandMap{
		data:  data,
		mutex: &sync.Mutex{},
	}
	if allowDiagonalIslands {
		lm.transforms = [][]int{
			{1, 0}, {0, 1}, {-1, 0}, {0, -1},
			{1, 1}, {-1, -1}, {-1, 1}, {1, -1},
		}
	} else {
		lm.transforms = [][]int{
			{1, 0}, {0, 1}, {-1, 0}, {0, -1},
		}
	}
	return lm
}

func (lm *LandMap) paint(x, y int) {
	lm.mutex.Lock()
	lm.data[y][x] = 0
	atomic.AddUint32(&lm.ops, 1)
	lm.mutex.Unlock()
}

func (lm *LandMap) isOne(x, y int) bool {
	if y >= 0 && y < len(lm.data) {
		if x >= 0 && x < len(lm.data[y]) {
			return lm.data[y][x] == 1
		}
	}
	return false
}

func (lm *LandMap) paintSafe(x, y int) bool {
	if !lm.isOne(x, y) {
		return false
	}
	lm.paint(x, y)
	return true
}

func (lm *LandMap) markRec(x, y int) {
	lm.paintSafe(x, y)
	for _, transform := range lm.transforms {
		if lm.paintSafe(x+transform[0], y+transform[1]) {
			lm.markRec(x+transform[0], y+transform[1])
		}
	}
}

func (lm *LandMap) FindLands() (count int) {
	lm.ops = 0
	for y := range lm.data {
		for x := range lm.data[y] {
			if lm.data[y][x] == 1 {
				count++
				lm.markRec(x, y)
			}
		}
	}
	return
}

// GetOperationsCount returns the number of operations on the input of the last FindLands call
func (lm *LandMap) GetOperationsCount() uint32 {
	return lm.ops
}
