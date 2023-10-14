package main

import (
	"fmt"
	"strings"
)

type house struct {
	size    float64
	filling filling
}

type filling struct {
	furniture []string
	technics  []string
	members   []string
}

func main() {
	house1 := &house{
		size: 100,
		filling: filling{
			furniture: []string{"sofa", "wardrobe", "bed", "armchair", "billiard"},
			technics:  []string{"kettle", "microwave", "laptop", "telly", "stove"},
			members:   []string{"mom", "dad", "dog", "me"},
		},
	}
	house1.seeHouse()
}

func (h house) seeHouse() {
	fmt.Println("Square is", h.size, "square meters")
	fmt.Println("Members are:", strings.Trim(fmt.Sprint(h.filling.members), "[]"))
	fmt.Println("Technics are:", strings.Trim(fmt.Sprint(h.filling.technics), "[]"))
	fmt.Println("Furniture is:", strings.Trim(fmt.Sprint(h.filling.furniture), "[]"))
}
