package main

type iDatabase interface {
	createToDoList(list ToDoList)
	addItemToList(item Item, listId string)
	removeItemFromList(itemId string, listId string)
}

type Database struct {
	database iDatabase
}

func (d *Database) createToDoList(list ToDoList){
	d.database.createToDoList(list)
}
func (d *Database)  addItemToList(item Item, listId string) {
	d.database.addItemToList(item, listId)
}
func (d *Database) removeItemFromList(itemId string, listId string) {
	d.database.removeItemFromList(itemId, listId)
}
