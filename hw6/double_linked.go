package main

import (
	"errors"
	"fmt"
	"unsafe"
)

/*
 https://en.wikipedia.org/wiki/Doubly_linked_list
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

	if dl.length > 0 {
		firstItem := (*Item)(ptrItem(dl.first))
		firstItem.prevItem = ptrItem(newItem)
		newItem.nextItem = dl.first
	} else {
		dl.last = ptrItem(newItem)
	}
	dl.first = ptrItem(newItem)
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
	dl.length++
}

//Remove - удаление элемента списка
func Remove(dl *DblLnkList) error {

	if dl.length > 0 {

		currItem := (*Item)(ptrItem(dl.current))

		if currItem.nextItem != nil {
			nextItem := (*Item)(ptrItem(currItem.nextItem))
			nextItem.prevItem = currItem.prevItem
		} else {
			// удаляем последний элемент, свигаем указатель конца списка
			dl.last = currItem.prevItem
		}

		if currItem.prevItem != nil {
			prevItem := (*Item)(ptrItem(currItem.prevItem))
			prevItem.nextItem = currItem.nextItem
			dl.current = ptrItem(prevItem) // сдвигаем текущую позицию на предыдущий элемент
		} else {
			// удаляем первый элемент, сдвигаем указатель начала списка
			dl.first = currItem.nextItem
			dl.current = dl.first // сдвигаем текущую позицию на первый элемент
		}

		//dl.current = dl.first
		currItem.nextItem = nil
		currItem.prevItem = nil
		clearValue(currItem)
		dl.length--
	} else {
		return errors.New("double list is empty")
	}
	return nil
}

// Next - итерация вперед по списку
func Next(dl *DblLnkList) error {

	currItem := *(*Item)(ptrItem(dl.current))

	if currItem.nextItem == nil {
		return errors.New("last item reached")
	}
	dl.current = currItem.nextItem
	return nil
}

// Prev - итерация назад по списку
func Prev(dl *DblLnkList) error {

	currItem := *(*Item)(ptrItem(dl.current))

	if currItem.prevItem == nil {
		return errors.New("first item reached")
	}
	dl.current = currItem.prevItem
	return nil
}

// Len возвращает длину списка
func Len(dl *DblLnkList) int { return dl.length }

// Value возвращает текущий элемент
func Value(dl *DblLnkList) interface{} { return getValue((*Item)(ptrItem(dl.current))) }

// First Установка указателя на первый элемент
func First(dl *DblLnkList) { dl.current = dl.first }

// Last Установка указателя на последний элемент
func Last(dl *DblLnkList) { dl.current = dl.last }

//Test is for testing
func Test(dl *DblLnkList) {

	var currItem Item
	First(dl)
	for i := 0; i < Len(dl); i++ {
		currItem = *(*Item)(ptrItem(dl.current))
		fmt.Printf("[%d] current &%#x prev %#x next %#x data %v (%T)\n", i, ptrItem(dl.current), currItem.prevItem, currItem.nextItem, getValue(&currItem), getValue(&currItem))
		Next(dl)
	}
	fmt.Println()
}

func main() {
	var dl = new(DblLnkList)

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
	Remove(dl)
	Remove(dl)
	Remove(dl)
	Remove(dl)
	Test(dl)

	fmt.Printf("DoubleLink (Elements: %d, Current: %#x First: %#x Last: %#x)", dl.length, dl.current, dl.first, dl.last)
}
