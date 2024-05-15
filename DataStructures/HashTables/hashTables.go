package main

import (
	"fmt"
)

type Keys struct {
	Key   string
	Value any
}

type HashTable struct {
	data   [][]Keys // store: (key,value) pair
	stored []bool
}

func (ht *HashTable) hash(key string) int {
	hash := 0
	for i, c := range key {
		hash = (hash + int(c)*i) % len(ht.data)
	}
	return hash
}

func (ht *HashTable) SetWithChaining(key string, data any) {
	address := ht.hash(key)
	kvPair := Keys{Key: key, Value: data}
	ht.data[address] = append(ht.data[address], kvPair)
}

func (ht *HashTable) GetForChaining(key string) any {
	hash := ht.hash(key)
	currentBucket := ht.data[hash]
	for _, item := range currentBucket {
		if item.Key == key {
			return item.Value
		}
	}
	return nil
}

func (ht *HashTable) keysForChaining() []string {
	var mapKeys []string
	for _, bucket := range ht.data {
		for _, keys := range bucket {
			mapKeys = append(mapKeys, keys.Key)
		}
	}

	return mapKeys
}

func (ht *HashTable) valuesForChaining() []any {
	var mapValues []any
	for _, bucket := range ht.data {
		for _, keys := range bucket {
			mapValues = append(mapValues, keys.Value)
		}
	}

	return mapValues
}

func New(size int) *HashTable {
	ht := HashTable{
		data:   make([][]Keys, size),
		stored: make([]bool, size),
	}
	return &ht
}
func main() {
	ht := New(200)
	ht.SetWithChaining("grape", 1)
	ht.SetWithChaining("bananas", 49)
	ht.SetWithChaining("oranges", 12) // we get a have collision here at index[1], if size of ht is 2.
	ht.SetWithChaining("apples", 8)   // we get a have collision here at index[1], if size of ht is 2.
	// fmt.Println(ht.data)
	fmt.Println(ht.GetForChaining("apples"))

	fmt.Println(ht.keysForChaining())
	fmt.Println(ht.valuesForChaining())
}
