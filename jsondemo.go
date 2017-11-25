package main

import (
    "github.com/json-iterator/go"
    "fmt"
)
type Animal struct {
    Name  string `json:"name"`
    Order string `json:"order"`
}
func main() {
    var animals []Animal
    animals = append(animals, Animal{Name: "Platypus", Order: "Monotremata"})
    animals = append(animals, Animal{Name: "Quoll", Order: "Dasyuromorphia"})

    jsonStr, err := jsoniter.Marshal(animals)
    if err != nil {
        fmt.Println("error:", err)
    }

    fmt.Println(string(jsonStr))
}