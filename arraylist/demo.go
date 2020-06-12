package arraylist

import (
	"errors"
	"fmt"
)

// ArrayList Apis
type IArray interface {
	Insert(val interface{}, index int) error
	Append(val interface{})
	Remove(index int) error
	SetValue(val interface{}, index int) error
	GetValue(index int) (interface{}, error)
	Size() int
}

// ArrayList struct that implements IArray
type ArrayList struct {
	data []interface{}
	size int
}

// NewArrayList to get an ArrayList
func NewArrayList() *ArrayList {
	return &ArrayList{
		data: make([]interface{}, 10), // default cap is 10
		size: 0,
	}
}

// Insert
func (arr *ArrayList) Insert(val interface{}, index int) error {
	// todo
	if err := arr.checkIndexValid(index); err != nil {
		return err
	}
	if arr.checkIfFull() {
		arr.increaseCap()
	}
	tmp := arr.data[index:]
	arr.data = append(arr.data[:index], val)
	arr.data = append(arr.data, tmp...)
	arr.size++
	return nil
}

// Append
func (arr *ArrayList) Append(val interface{}) {
	_ = arr.Insert(val, len(arr.data)) // Insert
}

// Remove
func (arr *ArrayList) Remove(index int) error {
	if err := arr.checkIndexValid(index); err != nil {
		return err
	}
	arr.data = append(arr.data[:index], arr.data[index+1:]...)
	arr.size--
	return nil
}

// SetValue
func (arr *ArrayList) SetValue(val interface{}, index int) error {
	if err := arr.checkIndexValid(index); err != nil {
		return err
	}
	arr.data[index] = val
	return nil
}

// GetValue
func (arr *ArrayList) GetValue(index int) (interface{}, error) {
	if err := arr.checkIndexValid(index); err != nil {
		return nil, err
	}
	return arr.data[index], nil
}

func (arr *ArrayList) String() string {
	return fmt.Sprintf("data:%v,size:%d", arr.data, arr.size)
}

// Size
func (arr *ArrayList) Size() int {
	return arr.size
}

// checkIndexValid check if index correct
func (arr *ArrayList) checkIndexValid(index int) error {
	if index < 0 || index >= arr.size {
		return errors.New("out of index")
	}
	return nil
}

// checkIfFull check if need to increase capacity
func (arr *ArrayList) checkIfFull() bool {
	return arr.size == cap(arr.data)
}

// increaseCap to increase capacity
func (arr *ArrayList) increaseCap() {
	var tmp []interface{}
	arr.data, tmp = make([]interface{}, cap(arr.data), cap(arr.data)*2), arr.data
	copy(arr.data, tmp) // keep origin data
}
