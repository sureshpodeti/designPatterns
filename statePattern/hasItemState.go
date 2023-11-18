package main

import "fmt"

type HasItemState struct {
	vendingMachine *VendingMachine
}

func (n *HasItemState) addItem(count int) error {
	n.vendingMachine.IncrementItem(count)
	return nil
}

func (n *HasItemState) requestItem() error {
	if n.vendingMachine.itemCount == 0 {
		n.vendingMachine.SetState(n.vendingMachine.noItemState)
		return fmt.Errorf("no Item present")
	}
	fmt.Println("Item requested")
	n.vendingMachine.SetState(n.vendingMachine.itemRequestedState)
	return nil
}

func (n *HasItemState) insertMoney(amount int) error {
	return fmt.Errorf("please select item first")
}

func (n *HasItemState) dispenseItem() error {
	return fmt.Errorf("please select item first")
}
