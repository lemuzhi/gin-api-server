package test

import (
	"fmt"
	"github.com/google/uuid"
	"testing"
)

func TestUUID(t *testing.T) {
	for i := 0; i < 2; i++ {
		id := uuid.New()

		fmt.Println(id.ID())
		fmt.Printf("%T,%v\n", id, id)
		fmt.Printf("%v,%v,%v\n", id.String(), id.Version(), id.Version().String())
	}

	fmt.Println("-------------------------------")

	for i := 0; i < 2; i++ {
		id2, err := uuid.NewRandom()
		if err != nil {
			fmt.Println(err)
		}
		fmt.Printf("%T,%v\n", id2, id2)
		fmt.Printf("%v,%v,%v\n", id2.String(), id2.Version(), id2.Version().String())
	}
}
