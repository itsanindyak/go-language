package main

import "fmt"

const (
	Recived   = "Recived"
	Confirmed = "Confirmed"
)

type order struct {
	item        string
	price       int
	orderstatus string
}

func (o order) currentStatus() {
	fmt.Println("Your order is", o.orderstatus)
}

func (o *order) changeStatus(status string) {
	o.orderstatus = status

}

func main() {
	newOrder := order{item: "BlackMan", price: 199, orderstatus: "Recived"}

	newOrder.currentStatus()
	newOrder.changeStatus(Recived)
	newOrder.currentStatus()

}
