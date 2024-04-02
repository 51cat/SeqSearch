package main

import (
	"flag"
	"seqsearch/utils"
)

func main() {

	var input string
	var format string
	var target_fa string
	var mode string
	var out string
	var sample_name string

	var mistmatch int
	var start int
	var end int

	flag.StringVar(&input, "input", "", "Input Your fasta/fastq path")
	flag.StringVar(&format, "format", "", "fastq/fasta")
	flag.StringVar(&mode, "mode", "first", "all/first")
	flag.StringVar(&target_fa, "target_fa", "", "Target fasta path")
	flag.StringVar(&out, "out", "./count_out.tsv", "path of out file")
	flag.StringVar(&sample_name, "sample_name", "sample", "sample name")

	flag.IntVar(&mistmatch, "mistmatch", 0, "mismatch")
	flag.IntVar(&start, "start",0 , "search region: start")
	flag.IntVar(&end, "end", 99999, "search region: end")
	flag.Parse()

	res:=utils.Search(
		input,
		format,
		target_fa,
		mode,
		mistmatch,
		start,
		end,
	)
	//fmt.Println(res)
	//utils.WriteSearchRes(res,"./ANN.tsv")
	utils.CountSearchRes(sample_name, res, out)
}
