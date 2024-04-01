package main

import (
	"fmt"
	"seqsearch/utils"
)


func main() {
	fa := "../test_data/test.fa"
	fa_map, fa_name := utils.Fasta2Map(fa)

	fmt.Println(fa_map)

	finder := utils.FinderFactory(
		fa_map,
		fa_name,
		1,
		0,
		100,
	)

	fmt.Println(finder.Find("mmmmCCCCCCCmmmm", "first"))
	fmt.Println(finder.Find("mmmmAAAAAAAAmmmm", "first"))
}
