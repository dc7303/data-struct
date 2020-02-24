package dataStruct

type pair struct {
	key string
	val string
}

type HashMap struct {
	list [3999][]*pair
}

func (h *HashMap) Put(key string, val string) {
	hash := Hash(key)
	h.list[hash] = append(h.list[hash], &pair{key, val})
}

func (h *HashMap) Get(key string) string {
	hash := Hash(key)
	pairList := h.list[hash]

	for i := 0; i < len(pairList); i++ {
		if pairList[i].key == key {
			return pairList[i].val
		}
	}
	return ""
}

func Hash(v string) int {
	// Rolling Hash Algorithm
	h := 0
	A := 256
	B := 3999

	for i := 0; i < len(v); i++ {
		// (H x A + S[i]) % B
		h = (h*A + int(v[i])) % B
	}
	return h
}
