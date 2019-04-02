package main

import (
	"container/list"
	"fmt"
)

func main() {
	cache := Constructor(2)
	cache.Put(1, 1)
	cache.Put(2, 2)
	fmt.Println(cache.Get(1))
	cache.Put(3, 3)
	fmt.Println(cache.Get(2))
	cache.Put(4, 4)
	fmt.Println(cache.Get(1))
	fmt.Println(cache.Get(3))
	fmt.Println(cache.Get(4))
}

type kv struct {
	key ,val int
} 

type LRUCache struct {
	capacity int
	list *list.List
	keyToPointer map[int]*list.Element
}


func Constructor(capacity int) LRUCache {
	return LRUCache{
		capacity:capacity,
		list:list.New(),
		keyToPointer:make(map[int]*list.Element),
	}
}


func (this *LRUCache) Get(key int) int {
	if _,ok := this.keyToPointer[key];!ok{
		return -1
	}
	this.list.MoveToFront(this.keyToPointer[key])
	return this.keyToPointer[key].Value.(kv).val
}


func (this *LRUCache) Put(key int, value int)  {
	if _,ok := this.keyToPointer[key];ok{
		this.keyToPointer[key].Value = kv{
			key:key,
			val:value,
		}
		this.list.MoveToFront(this.keyToPointer[key])
	}else{
		if this.capacity>this.list.Len(){
			this.list.PushFront(kv{
				key:key,
				val:value,
			})
			this.keyToPointer[key] = this.list.Front()
		}else{
			delete(this.keyToPointer,this.list.Back().Value.(kv).key)
			this.list.Remove(this.list.Back())
			this.list.PushFront(kv{
				key:key,
				val:value,
			})
			this.keyToPointer[key] = this.list.Front()
		}
	}
}


/**
 * Your LRUCache object will be instantiated and called as such:
 * obj := Constructor(capacity);
 * param_1 := obj.Get(key);
 * obj.Put(key,value);
 */