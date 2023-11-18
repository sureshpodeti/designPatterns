package main

type IState interface {
	addItem(count int) error
	requestItem() error
	insertMoney(amount int) error
	dispenseItem() error
}
