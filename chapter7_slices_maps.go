// Chapter 7: Slices, Maps & Custom Data Structures
package main

import "fmt"

func main() {
    // Slices
    arr := [5]int{1, 2, 3, 4, 5}
    s := arr[1:4]
    fmt.Println("slice:", s, "len:", len(s), "cap:", cap(s))
    s = append(s, 6, 7)
    fmt.Println("appended slice:", s)

    // Maps
    m := map[string]int{"apple": 5}
    m["banana"] = 7
    for k, v := range m {
        fmt.Println(k, v)
    }

    // Structs
    type User struct {
        ID   int
        Name string
    }
    u := User{1, "Alice"}
    fmt.Println("User:", u)

    // Linked List
    type Node struct {
        Value int
        Next  *Node
    }
    head := &Node{Value: 1, Next: &Node{Value: 2}}
    for n := head; n != nil; n = n.Next {
        fmt.Println("Node:", n.Value)
    }
}
