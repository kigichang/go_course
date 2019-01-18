package visible

import "fmt"

type supplier struct {
	ID   int
	Name string
}

// Product ...
type Product struct {
	supplier *supplier
	id       int
	name     string
}

// ResetProduct ...
func ResetProduct(p *Product) {
	p.id = 0
	p.name = ""
	p.supplier.ID = 0
	p.supplier.Name = ""
}

// GetSupplier ...
func GetSupplier(id int) *supplier {
	return &supplier{id, fmt.Sprintf("test-%d", id)}
}

// ProductSetSupplier ...
func ProductSetSupplier(p *Product, s *supplier) {
	p.supplier = s
}

func testPrivate() {
	fmt.Println("aa")
}
