/*
 * @lc app=leetcode id=460 lang=golang
 *
 * [460] LFU Cache
 */

// @lc code=start
type LFUCache struct {
	size     int
	keyToIdx map[int]int
	s        []valueTimes
}

type valueTimes struct {
	key   int
	value int
	times int
}

func Constructor(capacity int) LFUCache {
	return LFUCache{
		size:     0,
		keyToIdx: make(map[int]int, capacity),
		s:        make([]valueTimes, capacity),
	}
}

func (this *LFUCache) Get(key int) int {
	idx, ok := this.keyToIdx[key]
	if !ok {
		return -1
	}
	res := this.s[idx].value
	this.s[idx].times++
	this.moveToFront(idx)
	return res
}

func (this *LFUCache) Put(key int, value int) {
	if _, ok := this.keyToIdx[key]; !ok {
		if len(this.s) > this.size {
			this.s[this.size] = valueTimes{
				key:   key,
				value: value,
				times: 1,
			}
			this.keyToIdx[key] = this.size
			this.size++
			this.moveToFront(this.size - 1)
		} else {
			this.replaceNum(key, value, this.size-1)
		}
	} else {
		pos := this.keyToIdx[key]
		this.s[pos].times++
		this.s[pos].value = value
		this.moveToFront(pos)
	}
}

func (this *LFUCache) replaceNum(key, value, pos int) {
	oldKey := this.s[pos].key
	this.s[pos] = valueTimes{
		key:   key,
		value: value,
		times: 1,
	}
	this.keyToIdx[key] = pos
	delete(this.keyToIdx, oldKey)
	this.moveToFront(pos)
}

func (this *LFUCache) moveToFront(pos int) {
	for pos > 0 {
		if this.s[pos-1].times <= this.s[pos].times {
			this.s[pos-1], this.s[pos] = this.s[pos], this.s[pos-1]
			this.keyToIdx[this.s[pos].key] = pos
			this.keyToIdx[this.s[pos-1].key] = pos - 1
		} else {
			break
		}
		pos--
	}
}
