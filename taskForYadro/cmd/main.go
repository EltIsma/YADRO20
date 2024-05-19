package main

import (
	"bufio"
	"fmt"
	"os"
	"sortBalls/internal"
)

func main() {
	var in *bufio.Reader
	var out *bufio.Writer
	in = bufio.NewReader(os.Stdin)
	out = bufio.NewWriter(os.Stdout)
	defer out.Flush()
	scanner := bufio.NewScanner(in)
	scanner.Split(bufio.ScanWords)
	var n int8
	scanner.Scan()
    fmt.Sscan(scanner.Text(), &n)

	container := make([][]int32, n)
	var i int8
	for i = 0; i < n; i++ {
		container[i] = make([]int32, n)
		var j int8
		for j = 0; j < n; j++ {
			var k int32
            scanner.Scan()
            fmt.Sscan(scanner.Text(), &k)
			container[i][j] = k
		}
	}
	if internal.Checksortability(container) {
		fmt.Fprint(out, "yes\n")
	} else {
		fmt.Fprint(out, "no\n")
	}
}
