package visible_test

import (
	"fmt"
	"go_course/class09/visible"
	"os"
	"testing"
)

func TestMain(m *testing.M) {
	p := &visible.Product{}
	s := visible.GetSupplier(1)
	fmt.Println(s.ID)

	visible.ProductSetSupplier(p, s)
	os.Exit(m.Run())
}
