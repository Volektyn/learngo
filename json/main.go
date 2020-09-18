package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type Address struct {
	Street  string
	City    string
	State   string
	Zipcode int
}

type Item struct {
	Name        string `json:"itemname"`
	Description string `json:"desc,omitempty"`
	Quantity    int    `json:"qty"`
	Price       int    `json:"price"`
}

type Order struct {
	TotalPrice  int
	IsPaid      bool
	Fragile     bool `json:",omitempty"`
	OrderDetail []Item
}

type Customer struct {
	UserName      string
	Password      string `json:"-"`
	Token         string `json:"-"`
	ShipTo        Address
	PurchaseOrder Order `json:"order"`
}

func (customer *Customer) countTotalPrice() {
	total := 0
	for _, item := range customer.PurchaseOrder.OrderDetail {
		total += item.Quantity * item.Price
	}
	customer.PurchaseOrder.TotalPrice = total
}

func main() {
	jsonData := []byte(`
	{
		"username": "blackhat",
		"shipto": {
		  "street": "Sulphur Springs Rd",
		  "city": "Park City",
		  "state": "VA",
		  "zipcode": 12345
		},
		"order": {
		  "paid": false,
		  "orderdetail": [
			{
			  "itemname": "A Guide to the World of zeros and ones",
			  "desc": "book",
			  "qty": 3,
			  "price": 50
			}
		  ]
		}
	  }
		`)

	if !json.Valid(jsonData) {
		fmt.Printf("JSON is not valid: %s", jsonData)
		os.Exit(1)
	}

	var customer Customer
	err := json.Unmarshal(jsonData, &customer)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Println(customer)

	game := Item{}
	game.Name = "Final Fantasy The Zodiac Age"
	game.Description = "Nintendo Switch Game"
	game.Quantity = 1
	game.Price = 50

	glass := Item{}
	glass.Name = "Crystal Drinking Glass"
	glass.Quantity = 11
	glass.Price = 25

	customer.PurchaseOrder.OrderDetail = append(customer.PurchaseOrder.OrderDetail, game)
	customer.PurchaseOrder.OrderDetail = append(customer.PurchaseOrder.OrderDetail, glass)

	customer.countTotalPrice()

	customer.PurchaseOrder.IsPaid = true
	customer.PurchaseOrder.Fragile = true

	customerOrder, err := json.MarshalIndent(customer, "", "    ")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	fmt.Println(string(customerOrder))
}
