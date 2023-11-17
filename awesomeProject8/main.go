package main

import (
	"awesomeProject8/context"
	"awesomeProject8/models"
	"awesomeProject8/strategies"
	"fmt"
)

func main() {

	//addStrategy := &strategies.AddStrategy{}
	//subtractStrategy := &strategies.SubtractStrategy{}
	//
	//arithmeticContext := &context2.ArithmeticContext{}
	//
	//arithmeticContext.SetStrategy(addStrategy)
	//
	//fmt.Println(arithmeticContext.Execute(10, 20))
	//
	//arithmeticContext.SetStrategy(subtractStrategy)
	//fmt.Println(arithmeticContext.Execute(20, 10))

	/*
			  PaymentStrategies: Credit Card, Debit Card, Net banking, Cash

			  Mode            Extra
			--------------------------
			  Cash               0
		   Net banking           10
			Debit Card           20
			Credit Card          35



	*/

	//cash := &strategies.Cash{Extra: 0}
	//netBank := &strategies.NetBanking{Extra: 10}
	//debitCard := &strategies.DebitCard{Extra: 20}
	//creditCard := &strategies.CreditCard{Extra: 30}
	//
	//paymentContext := &context2.PaymentContext{}
	//
	//var input int
	//
	//fmt.Println("Enter payment strategy: 1.Cash, 2. Net banking, 3. Debit Card, 4. Credit Card ")
	//fmt.Scanln(&input)
	//
	//switch input {
	//case 1:
	//	paymentContext.SetPaymentStrategy(cash)
	//	fmt.Println("Cash payment: ", paymentContext.Pay(10))
	//case 2:
	//	paymentContext.SetPaymentStrategy(netBank)
	//	fmt.Println("Netbanking: ", paymentContext.Pay(10))
	//
	//case 3:
	//	paymentContext.SetPaymentStrategy(debitCard)
	//	fmt.Println("Debit card: ", paymentContext.Pay(10))
	//
	//case 4:
	//	paymentContext.SetPaymentStrategy(creditCard)
	//	fmt.Println("Credit card: ", paymentContext.Pay(10))
	//
	//default:
	//	fmt.Println("Quit!. Wrong input!")
	//
	//}

	/*
	  Calculate time taken to reach from {s1, s2}, to {d1, d2}
	  assume, we have 3 modes of transportation, walk, bike, car

	  assume walk speed is: 10 km / 1hour
	         bike speed: 20 Km / hour
	         car speed: 50 km / hour
	*/

	walkStrategy := &strategies.Walk{Speed: 10}
	bikeStrategy := &strategies.Bike{Speed: 20}
	carStrategy := &strategies.Car{Speed: 50}

	source := models.Point{X: 1, Y: 1}
	destination := models.Point{X: 40, Y: 50}

	transportContext := &context.TransportContext{}

	transportContext.SetContext(walkStrategy)
	fmt.Printf("walking takes: %.2f hr\n", transportContext.CalculateTime(source, destination))

	transportContext.SetContext(bikeStrategy)
	fmt.Printf("bike takes: %.2f hr\n", transportContext.CalculateTime(source, destination))

	transportContext.SetContext(carStrategy)
	fmt.Printf("car takes: %.2f hr\n", transportContext.CalculateTime(source, destination))

}
