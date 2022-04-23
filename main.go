package main

import goloc "aboxofsox/goloc/pkg"

func main() {
	// files := goloc.Load()

	// for k, v := range files {
	// 	fmt.Printf("%s: %d\n", k, v)
	// }

	goloc.Printer(goloc.Load())
}
