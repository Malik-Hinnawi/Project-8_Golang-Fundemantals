package product

import (
	"fmt"
	"time"
)

type Product struct {
	title     string
	id        int
	price     float64
	createdAt time.Time
}

func New(title string, id int, price float64) Product {
	return Product{
		title,
		id,
		price,
		time.Now(),
	}
}

func (p Product) Display() {
	fmt.Printf("title: %v, id: %v, price: %v, createdAt: %v", p.title, p.id, p.price, p.createdAt)
}
