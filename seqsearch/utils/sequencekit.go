package utils

import (
	"os"
	"bufio"
	"strings"
	"compress/gzip"
)

func GetfqScaner(f *os.File, path string, format string) (*bufio.Scanner, int, int) {
	var scanner *bufio.Scanner
	var k int
	var seq_line int

	if (strings.HasSuffix(path, ".gz")) {
		gz, err := gzip.NewReader(f)

		CheckError(err)

		scanner = bufio.NewScanner(gz)
	}else {
		scanner = bufio.NewScanner(f)
	}
    
	if format == "fastq" {
		k = 4
		seq_line = 2
	}
	if format == "fasta" {
		k = 2
		seq_line = 0
	}

	return scanner, k, seq_line
}

func Fasta2Map(fa_path string) (map[string]string, []string) {
	
    var record string
    var fa_seq, fa_name []string
	fa_map := make(map[string]string)
	f, err := os.Open(fa_path)
    
	defer f.Close()

	CheckError(err)
	input := bufio.NewScanner(f)
	
	for input.Scan() {
		record = input.Text()
	    if strings.HasPrefix(record, ">") {
		    fa_name = append(fa_name, strings.ReplaceAll(record, ">", ""))
	    }else{
		    fa_seq = append(fa_seq, record)
	    }
	}
	for inx, name := range fa_name {
		fa_map[name] = fa_seq[inx]
	}
	return fa_map, fa_name
}