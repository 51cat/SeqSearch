package main

import (
	//"fmt"
	"seqsearch/utils"
)

func main() {
	res:=utils.Search(
		"../test_data/test.fastq",
		"fastq",
		"../test_data/target.fa",
		"all",
		2,
		0,
		1000000,
	)
	//fmt.Println(res)
	//utils.WriteSearchRes(res,"./ANN.tsv")
	utils.CountSearchRes("CC", res,"./out.tsv")
}
