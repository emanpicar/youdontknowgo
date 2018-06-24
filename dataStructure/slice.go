package dataStructure

import "fmt"

func sliceStruct() {
    // slice of `string`s of length `3` (initially zero-valued).
    s := make([]string, 3)
    fmt.Println("empty:", s)

    s[0] = "a"
    s[1] = "b"
    s[2] = "c"
    fmt.Println("set:", s)
    fmt.Println("get:", s[2])

    s = justAppend(s)
    justCopy(s)
    justSliceOperation(s)
    justInitializedSlice()
    justMultiDimSlice()
}

func justAppend(s []string) []string {
    fmt.Println("length:", len(s))

    s = append(s, "d")
    s = append(s, "e", "f")
    fmt.Println("append:", s)

    return s
}

func justCopy(s []string) {
    c := make([]string, len(s))
    copy(c, s)
    fmt.Println("copy:", c)
}

func justSliceOperation(s []string) {
    // this gets a slice of the elements `s[2]`, `s[3]`, and `s[4]`.
    l1 := s[2:5]
    fmt.Println("slice1:", l1)

    // This slices up to (but excluding) `s[5]`.
    l2 := s[:5]
    fmt.Println("slice2:", l2)

    // And this slices up from (and including) `s[2]`.
    l3 := s[2:]
    fmt.Println("slice3:", l3)
}

func justInitializedSlice() {
    // We can declare and initialize a variable for slice
    // in a single line as well.
    t := []string{"g", "h", "i"}
    fmt.Println("dcl:", t)
}

func justMultiDimSlice() {
    // Slices can be composed into multi-dimensional data
    // structures. The length of the inner slices can
    // vary, unlike with multi-dimensional arrays.
    twoD := make([][]int, 3)
    for i := 0; i < 3; i++ {
        innerLen := i + 1
        twoD[i] = make([]int, innerLen)
        for j := 0; j < innerLen; j++ {
            twoD[i][j] = i + j
        }
    }
    fmt.Println("2d: ", twoD)
}