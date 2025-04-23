package py

import (
	"fmt"
	"github.com/markadc/pgg"
	"testing"
)

func TestJson(t *testing.T) {
	m1 := pgg.A{"name": "CLOS", "age": 22}
	s1, err := Dumps(m1)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(s1)

	s2 := "{\"flag\": \"pgg\", \"ids\": [1, 2, 3]}"
	m2, err := Loads(s2)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(m2)
}
