package dataStruct

import "fmt"

type Heap struct {
	list []int
}

func (h *Heap) Push(val int) {
	h.list = append(h.list, val)

	idx := len(h.list) - 1

	for idx > 0 {
		// 부모 인덱스 (N - 1) / 2
		parentIdx := (idx - 1) / 2
		if parentIdx < 0 {
			break
		}
		if h.list[parentIdx] < h.list[idx] {
			h.list[parentIdx], h.list[idx] = h.list[idx], h.list[parentIdx]
			idx = parentIdx
		} else {
			break
		}
	}
}

func (h *Heap) Pop() int {
	if len(h.list) == 0 {
		return 0
	}

	top := h.list[0]
	last := h.list[len(h.list)-1]
	h.list = h.list[:len(h.list)-1]

	h.list[0] = last

	idx := 0
	for idx < len(h.list) {
		swapIdx := -1
		// 왼쪽 자식 노드 N x 2 + 1
		leftIdx := idx*2 + 1
		if leftIdx >= len(h.list) {
			break
		}
		if h.list[leftIdx] > h.list[idx] {
			swapIdx = leftIdx
		}

		// 오른쪽 자식 노드 N x 2 + 2
		rightIdx := idx*2 + 2
		if rightIdx < len(h.list) {
			if h.list[rightIdx] > h.list[idx] {
				if swapIdx < 0 || h.list[swapIdx] < h.list[rightIdx] {
					swapIdx = rightIdx
				}
			}
		}

		if swapIdx < 0 {
			break
		}

		h.list[idx], h.list[swapIdx] = h.list[swapIdx], h.list[idx]
		idx = swapIdx
	}

	return top
}

func (h *Heap) PrintList() {
	fmt.Println(h.list)
}
