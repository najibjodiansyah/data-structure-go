package ds

import "fmt"

type DynamicArray struct {
	Length int64
	Data   map[int64]any
}

type Array interface {
	Get(index int64) any
	Push(item any)
	Pop() any
	Delete(index int64) any
	ShiftItems(index int64)
	Print()
	Reverse()
}

func NewDynamycArray() *DynamicArray {
	return &DynamicArray{
		Length: 0,
		Data:   make(map[int64]any),
	}
}

func (a *DynamicArray) Get(index int64) any {
	return a.Data[index]
}

func (a *DynamicArray) Push(item any) {
	a.Data[a.Length] = item
	a.Length++
}

func (a *DynamicArray) Pop() any {
	lastItem := a.Data[a.Length-1]
	delete(a.Data, a.Length-1)
	a.Length--
	return lastItem
}

func (a *DynamicArray) Delete(index int64) any {
	item := a.Data[index]
	a.ShiftItems(index)
	return item
}

func (a *DynamicArray) ShiftItems(index int64) {
	for i := index; i < a.Length-1; i++ {
		a.Data[i] = a.Data[i+1]
	}
	delete(a.Data, a.Length-1)
	a.Length--
}

func (a *DynamicArray) Print() {
	for i := int64(0); i < a.Length; i++ {
		fmt.Println(a.Data[i])
	}
}

func (a *DynamicArray) Reverse() {
	for i, j := int64(0), a.Length-1; i < j; i, j = i+1, j-1 {
		a.Data[i], a.Data[j] = a.Data[j], a.Data[i]
	}
}
