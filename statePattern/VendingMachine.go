package main

import (
	"fmt"
)

type VendingMachine struct {
	itemCount int
	itemPrice int

	currentState IState

	noItemState        IState
	hasItemState       IState
	itemRequestedState IState
	hasMoneyState      IState
}

func NewVendingMachine(count, price int) *VendingMachine {
	vendingMachine := &VendingMachine{
		itemCount: count,
		itemPrice: price,
	}

	vendingMachine.noItemState = &NoItemState{vendingMachine: vendingMachine}
	vendingMachine.hasItemState = &HasItemState{vendingMachine}
	vendingMachine.itemRequestedState = &ItemRequestedState{vendingMachine: vendingMachine}
	vendingMachine.hasMoneyState = &HasMoneyState{vendingMachine: vendingMachine}

	vendingMachine.SetState(vendingMachine.hasItemState)

	return vendingMachine

}

func (v *VendingMachine) IncrementItem(count int) {
	fmt.Printf("Adding %d items", count)
	v.itemCount = v.itemCount + count
}

func (v *VendingMachine) DecrementItem() error {
	if v.itemCount == 0 {
		return fmt.Errorf("can not decrement item")
	}
	v.itemCount--

	return nil
}
func (v *VendingMachine) AddItem() {

}

func (v *VendingMachine) RequestItem() error {
	return v.currentState.requestItem()
}

func (v *VendingMachine) SetState(s IState) {
	v.currentState = s
}

func (v *VendingMachine) InsertMoney(amount int) error {
	return v.currentState.insertMoney(amount)
}

func (v *VendingMachine) DispenseItem() error {
	return v.currentState.dispenseItem()
}
