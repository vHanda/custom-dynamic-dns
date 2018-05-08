//usr/bin/go run $0 $@ ; exit
package main

import humanize "github.com/dustin/go-humanize"
import "fmt"

func main() {
	fmt.Printf("That file is %s.\n", humanize.Bytes(82854982)) // That file is 83 MB.
	fmt.Println("Hello World!")
}
