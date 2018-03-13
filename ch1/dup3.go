package main

import (
    "fmt"
    "io/ioutil"
    "os"
    "strings"
)

func main() {
    counts := make(map[string]int)
    files := os.Args[1:]
    for _, filename := range files {
        data, err := ioutil.ReadFile(filename)
        if err != nil {
            fmt.Fprintf(os.Stderr, "dup2: %v\n", err)
            continue
        }

        for _, line := range strings.Split(string(data), "\n") {
            counts[line]++
        }

    }

    for line, n := range counts {
        if n > 1 {
            fmt.Printf("%d\t%s\n", n, line)
        }
    }
}
