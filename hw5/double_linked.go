package main

import (
	"errors"
	"fmt"
	"unsafe"
)

/*
 https://en.wikipedia.org/wiki/Doubly_linked_list

 Len()  длина списка
 First() первый Item
 Last() последний Item
 PushFront(v interface{})  добавить значение в начало
 PushBack(v interface{})  добавить значение в конец
 Remove(i Item) удалить элемент
 Value() interface{} возвращает значение
 Next() *Item следующий Item
 Prev() *Item предыдущий

*/

// Объявление и реализация интерфейса
type ifValue = interface {
	setValue(interface{})
	getValue() interface{}
	clearValue()
}

type multiValue = struct {
	sData string
	iData int
	fData float64
}

// DblLnkList Сервисная структура двусвязного списка
type DblLnkList = struct {
	length               int
	current, first, last ptrItem
}

// Указатель на элемент Item (невозможно непосредственно хранить указатель на Item внутри Item)
type ptrItem unsafe.Pointer

// Item - Базовая структура (элемент) списка
type Item = struct {
	//id       int
	nextItem ptrItem
	prevItem ptrItem
	value    multiValue
}

func setValue(m *Item, i interface{}) {
	switch i.(type) {
	case int:
		m.value.iData = i.(int)
	case float64:
		m.value.fData = i.(float64)
	case string:
		m.value.sData = i.(string)
	case nil:
		m.value.sData = "nil"
	default:
		m.value.sData = fmt.Sprintf("%v", i)
	}
}

func getValue(i *Item) interface{} {

	switch {
	case i.value.sData != "":
		return i.value.sData
	case i.value.iData != 0:
		return i.value.iData
	case i.value.fData != 0:
		return i.value.fData
	default:
		return nil
	}
}

func clearValue(i *Item) {
	i.value.sData = ""
	i.value.iData = 0
	i.value.fData = 0
}

// PushFront Добавить элемент в начало списка
func PushFront(dl *DblLnkList, i interface{}) {

	newItem := new(Item)
	setValue(newItem, i)

	/*if dl.first != nil {
		currItem.prevItem = ptrItem(newItem)
		newItem.nextItem = dl.first
	} else if dl.last == nil {
		dl.last = dl.first
	}*/

	if dl.length > 0 {
		firstItem := (*Item)(ptrItem(dl.first))
		firstItem.prevItem = ptrItem(newItem)
		newItem.nextItem = dl.first
	} else {
		dl.last = ptrItem(newItem)
	}
	dl.first = ptrItem(newItem)

	//dl.current = dl.first
	dl.length++
}

// PushBack Добавить элемент в конец списка
func PushBack(dl *DblLnkList, i interface{}) {

	newItem := new(Item)
	setValue(newItem, i)

	if dl.length > 0 {
		lastItem := (*Item)(ptrItem(dl.last))
		lastItem.nextItem = ptrItem(newItem)
		newItem.prevItem = dl.last
	} else {
		dl.first = ptrItem(newItem)
	}
	dl.last = ptrItem(newItem)

	//dl.current = dl.last
	dl.length++
}

// Next - итерация вперед по списку
func Next(dl *DblLnkList) error /*(Item, error)*/ {

	var currItem Item
	//var nextItem Item

	//currItem = *(*Item)(ptrItem(dl.current))
	currItem = *(*Item)(ptrItem(dl.current))

	if currItem.nextItem == nil {
		return /*Item{},*/ errors.New("last item reached")
	}

	//nextItem = *(*Item)(unsafe.Pointer(currItem.nextItem))
	dl.current = currItem.nextItem
	//return nextItem, nil
	return nil
}

// Prev - итерация назад по списку
func Prev(dl *DblLnkList) error /*(Item, error)*/ {

	var currItem Item
	//var prevItem Item

	currItem = *(*Item)(ptrItem(dl.current))

	if currItem.prevItem == nil {
		return /*Item{},*/ errors.New("first item reached")
	}

	//prevItem = *(*Item)(unsafe.Pointer(currItem.prevItem))
	dl.current = currItem.prevItem
	//return prevItem, nil
	return nil
}

// Len возвращает длину списка
func Len(dl *DblLnkList) int { return dl.length }

// Value возвращает текущий элемент
func Value(dl *DblLnkList) interface{} {

	var currItem Item

	currItem = *(*Item)(ptrItem(dl.current))
	return getValue(&currItem)
}

// First Установка указателя на первый элемент
func First(dl *DblLnkList) { dl.current = dl.first }

// Last Установка указателя на последний элемент
func Last(dl *DblLnkList) { dl.current = dl.last }

/*
func (dl dblList) First() itemValue {}
func (dl dblList) Last() itemValue {}
func (dl dblList) Prev() itemValue, error {}
func (dl dblList) Next() itemValue, error {}
func (dl dblList) Value(i Item) itemValue, error {}
func (dl dblList) Remove(i Item) error {}
*/

//Test is for testing
func Test(dl *DblLnkList) {
	//var val interface{}
	var currItem Item
	//TestData := []interface{}{5,"text",0.5,nil,7}
	First(dl)
	for i := 0; i < Len(dl); i++ {
		currItem = *(*Item)(ptrItem(dl.current))
		fmt.Printf("[%d] current &%#x prev %#x next %#x data %v (%T)\n", i, ptrItem(dl.current), currItem.prevItem, currItem.nextItem, getValue(&currItem), getValue(&currItem))
		//val = Value(dl)
		//fmt.Printf("[%d] %v is type %t\n", i, val, val)
		Next(dl)
	}
	fmt.Println()
}

func main() {
	var dl *DblLnkList
	//var val interface{}
	dl = new(DblLnkList)

	fmt.Printf("%v\n", dl)
	PushFront(dl, 5)
	Test(dl)
	PushBack(dl, "last!")
	Test(dl)
	PushFront(dl, 0.5)
	Test(dl)
	PushBack(dl, nil)
	Test(dl)
	PushFront(dl, 7)
	Test(dl)
	fmt.Printf("%v\n", dl)
}
