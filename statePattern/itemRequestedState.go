package main

import "fmt"

type ItemRequestedState struct {
	vendingMachine *VendingMachine
}

func (n *ItemRequestedState) addItem(count int) error {
	return fmt.Errorf("item Dispense in progress")
}

func (n *ItemRequestedState) requestItem() error {
	return fmt.Errorf("item already requested")
}

func (n *ItemRequestedState) insertMoney(amount int) error {
	if amount < n.vendingMachine.itemPrice {
		return fmt.Errorf("inserted money is less. Please insert %d", n.vendingMachine.itemPrice)
	}
	fmt.Println("Money entered is ok")
	n.vendingMachine.SetState(n.vendingMachine.hasMoneyState)
	return nil
}

func (n *ItemRequestedState) dispenseItem() error {
	return fmt.Errorf("please insert money first")
}
