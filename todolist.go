package main

type ToDoList struct {
	Id string
	ToDos          []Item
	CompletedItems []Item
}

func (list *ToDoList) countAll() int {
	return len(list.ToDos) + len(list.CompletedItems)
}

func (list *ToDoList) countToDos() int {
	return len(list.ToDos)
}

func (list *ToDoList) countCompletedItems() int {
	return len(list.CompletedItems)
}

func (list *ToDoList) addItem(item Item) {
	if item.Completed {
		list.CompletedItems = append(list.CompletedItems, item)
	} else {
		list.ToDos = append(list.ToDos, item)
	}
}

func (list *ToDoList) removeItem(item Item) {
	l := &list.ToDos
	if item.Completed {
		l = &list.CompletedItems
	}
	for i := range *l {
		if item == (*l)[i] {
			*l = append((*l)[:i], (*l)[i+1:]...)
		}
	}
}

func (list *ToDoList) getCompletedItems() []Item {
	return list.CompletedItems
}

func (list *ToDoList) getToDos() []Item {
	return list.ToDos
}

func (list *ToDoList) getItem(id string) Item {
	for i := range list.ToDos {
		if list.ToDos[i].Id == id {
			return list.ToDos[i]
		}
	}
	for i := range list.CompletedItems {
		if list.CompletedItems[i].Id == id {
			return list.CompletedItems[i]
		}
	}
	return Item{}
}
