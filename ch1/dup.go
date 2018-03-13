package main

import (
    "bufio"
    "fmt"
    "os"
    "strings"
)

func main() {
    counts := make(map[string]int)
    input := bufio.NewScanner(os.Stdin)
    for input.Scan() {
        if strings.EqualFold(input.Text(), "quit") {
            break
        }
        counts[input.Text()]++
    }

    for line, n := range counts {
        if n > 1 {
            fmt.Printf("%d\t%s\n", n, line)
        }
    }
}
