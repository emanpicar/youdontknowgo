package dataStructure

import "fmt"

func mapStruct() {
    // `make(map[key-type]val-type)`.
    m := make(map[string]int)

    m["k1"] = 7
    m["k2"] = 13

    fmt.Println("map:", m)

    v1 := m["k1"]
    fmt.Println("v1: ", v1)

    // returns the number of key/value pairs when called on a map.
    fmt.Println("len:", len(m))

    delete(m, "k2")
    fmt.Println("map:", m)

    // The optional second return value when getting a
    // value from a map indicates if the key was present
    // in the map. This can be used to disambiguate
    // between missing keys and keys with zero values
    // like `0` or `""`. Here we didn't need the value
    // itself, so we ignored it with the _blank identifier_
    // `_`.
    _, prs := m["k2"]
    fmt.Println("isKeyPresent:", prs)

    // declare and initialize a new map in the same line
    n := map[string]int{"foo": 1, "bar": 2}
    fmt.Println("map:", n)
}