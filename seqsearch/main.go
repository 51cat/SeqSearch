package main

import (
	//"fmt"
	"seqsearch/utils"
)

func main() {
	res:=utils.Search(
		"/SGRNJ06/randd/USER/liuzihao/randdScript/github/targetS/SeqSearch/test_data/test.fastq",
		"fastq",
		"/SGRNJ06/randd/USER/liuzihao/randdScript/github/targetS/SeqSearch/test_data/target.fa",
		"all",
		2,
		0,
		1000000,
	)
	//fmt.Println(res)
	//utils.WriteSearchRes(res,"./ANN.tsv")
	utils.CountSearchRes("CC", res,"./out.tsv")
}
