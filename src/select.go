package main

import (
	"fmt"
	"time"
)

/*
Imagine we have two pizza shop, an old one and a new one,
OldPizzaShop is the first branch with no specialist chefs and also it doesn't adopt latest technology
NewPizzaShop is the newest branch, adopting latest technology and having specialist chefs.

Based on the above scenoria, the work efficiency of NewPizzaShop is faster than OldPizzaShop
*/

// PizzaShop contains neccessary fields
type PizzaShop struct {
	IsExpertChefAvailable    bool
	IsNewTechnologyAvailable bool
	OrderQuantity            int
}

// OldShopDelivery delivers the pizza with a delay to the customers
func (o PizzaShop) OldShopDelivery(ch chan PizzaShop) {

	time.Sleep(10 * time.Second)

	// Send PizzaShop to the channel ch
	ch <- PizzaShop{
		IsExpertChefAvailable:    false,
		IsNewTechnologyAvailable: false,
		OrderQuantity:            20,
	}

}

// NewShopDelivery delivers the pizza without a delay to the customers
func (n PizzaShop) NewShopDelivery(ch chan PizzaShop) {

	time.Sleep(2 * time.Second)

	// Send PizzaShop to the channel ch
	ch <- PizzaShop{
		IsExpertChefAvailable:    true,
		IsNewTechnologyAvailable: true,
		OrderQuantity:            50,
	}

}

func main() {

	// Create reference of PizzaShop
	var pizzashop PizzaShop

	//Create OldShop and NewShop Channels
	oldShop, newShop := make(chan PizzaShop), make(chan PizzaShop)

	// Call two routines
	go pizzashop.OldShopDelivery(oldShop)
	go pizzashop.NewShopDelivery(newShop)

	//The select statement lets a goroutine wait on multiple communication operations
	select {
	// Receive from the oldshop channel and assignes to oldshop
	case oldshop := <-oldShop:
		{
			fmt.Printf("%d pizzas are ready to deliver from OldPizzaShop\nHence, The old pizza shop delivers the pizzas slower than the new pizza shop to the customers", oldshop.OrderQuantity)
		}
	// Receive from the newShop channel and assignes to newShop
	case newshop := <-newShop:
		fmt.Printf("%d pizzas are ready to deliver from NewPizzaShop\nHence, The new pizza shop delivers the pizzas faster than the old pizza shop to the customers", newshop.OrderQuantity)

	}

	// close the channels
	close(oldShop)
	close(newShop)

}
