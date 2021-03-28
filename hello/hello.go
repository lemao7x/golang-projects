// go mod init -> initialize project
package main

import "fmt"
// go mod tidy -> download dependencies
import "rsc.io/quote"

func main() {
    fmt.Println(quote.Go())
}