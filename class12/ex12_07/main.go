package main

import (
	"fmt"
	"log"
	"sync"
)

var (
	waitGroup = sync.WaitGroup{}
)

// Order ...
type Order struct {
	Category string
	Amount   float64
}

// Actor ...
type Actor interface {
	Run()
}

// Producer ...
type Producer struct {
	MailBoxes []chan Order
}

// Run ...
func (p *Producer) Run() {
	defer waitGroup.Done()

	for i := 0; i < 100; i++ {
		category := fmt.Sprintf("cate-%d", i%7)
		amount := float64(i)

		order := Order{
			Category: category,
			Amount:   amount,
		}

		for _, m := range p.MailBoxes {
			m <- order
		}
	}

	for _, m := range p.MailBoxes {
		close(m)
	}
}

// CategorySum ...
type CategorySum struct {
	MailBox     chan Order
	CategorySum map[string]float64
}

// Run ...
func (c *CategorySum) Run() {
	defer waitGroup.Done()

	for order := range c.MailBox {
		c.CategorySum[order.Category] += order.Amount
	}
}

// SiteSum ...
type SiteSum struct {
	MailBox chan Order
	Total   float64
}

// Run ...
func (c *SiteSum) Run() {
	defer waitGroup.Done()

	for order := range c.MailBox {
		c.Total += order.Amount
	}
}

func main() {
	log.Println("start...")

	producer := &Producer{}
	waitGroup.Add(1)

	category := &CategorySum{
		MailBox:     make(chan Order),
		CategorySum: make(map[string]float64),
	}
	waitGroup.Add(1)
	producer.MailBoxes = append(producer.MailBoxes, category.MailBox)

	site := &SiteSum{
		MailBox: make(chan Order),
	}
	waitGroup.Add(1)
	producer.MailBoxes = append(producer.MailBoxes, site.MailBox)

	go producer.Run()
	go category.Run()
	go site.Run()

	waitGroup.Wait()

	total := 0.0

	for x, a := range category.CategorySum {
		log.Println(x, ":", a)
		total += a
	}

	log.Println("total: ", site.Total, total)

	log.Println("end")
}
