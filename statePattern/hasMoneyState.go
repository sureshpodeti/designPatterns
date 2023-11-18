package main

import "fmt"

type HasMoneyState struct {
	vendingMachine *VendingMachine
}

func (n *HasMoneyState) addItem(count int) error {
	return fmt.Errorf("item dispense in progress")
}

func (n *HasMoneyState) requestItem() error {
	return fmt.Errorf("item dispense in progress")
}

func (n *HasMoneyState) insertMoney(amount int) error {
	return fmt.Errorf("item dispense in progress")
}

func (n *HasMoneyState) dispenseItem() error {
	fmt.Println("dispensing item...")
	n.vendingMachine.DecrementItem()

	if n.vendingMachine.itemCount == 0 {
		n.vendingMachine.SetState(n.vendingMachine.noItemState)
	} else {
		n.vendingMachine.SetState(n.vendingMachine.hasItemState)
	}

	fmt.Println("item dispensed!")

	return nil
}
