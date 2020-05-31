package main

import "fmt"

func main() {
        var b int = 100
        var a int

        for a < b {
                a++
                if (a % 3 == 0 && a % 5 == 0) {
                        fmt.Println("foobar")
                } else if (a % 5 == 0) {
                        fmt.Println("bar")
                } else if (a % 3 == 0) {
                        fmt.Println("foo")
                } else {
                        fmt.Println(a)
                }
        }
}
