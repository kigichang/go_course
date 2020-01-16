package visible_test

import (
	"fmt"
	"os"
	"testing"
	"visible"
)

func TestMain(m *testing.M) {
	p := &visible.Product{}
	s := visible.GetSupplier(1)
	fmt.Println(s.ID)

	visible.ProductSetSupplier(p, s)
	os.Exit(m.Run())
}
