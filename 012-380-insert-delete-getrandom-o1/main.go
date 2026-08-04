package main

import "math/rand"

type RandomizedSet struct {
	pointer map[int]int //value to index slice
	arr     []int
}

func Constructor() RandomizedSet {
	return RandomizedSet{
		pointer: make(map[int]int),
		arr:     []int{},
	}
}

func (this *RandomizedSet) Insert(val int) bool {
	_, ok := this.pointer[val]
	if ok {
		return false
	}
	this.pointer[val] = len(this.arr)
	this.arr = append(this.arr, val)
	return true
}

func (this *RandomizedSet) Remove(val int) bool {
	valIndex, ok := this.pointer[val]
	if !ok {
		return false
	}

	lastVal := this.arr[len(this.arr)-1]
	this.arr[valIndex] = lastVal
	this.arr = this.arr[0 : len(this.arr)-1]
	this.pointer[lastVal] = valIndex
	delete(this.pointer, val)

	return true
}

func (this *RandomizedSet) GetRandom() int {
	return this.arr[rand.Intn(len(this.arr))]
}
