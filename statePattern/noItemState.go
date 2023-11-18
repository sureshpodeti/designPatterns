package main

import (
	"fmt"
)

type NoItemState struct {
	vendingMachine *VendingMachine
}

func (n *NoItemState) addItem(count int) error {
	//NoItem, and we got action to addItem
	n.vendingMachine.IncrementItem(count)
	return nil
}

func (n *NoItemState) requestItem() error {
	return fmt.Errorf("item out of stock")
}

func (n *NoItemState) insertMoney(amount int) error {
	return fmt.Errorf("item out of stock")
}

func (n *NoItemState) dispenseItem() error {
	return fmt.Errorf("item out of stock")
}
