package utils


import (
	"os"
	"strings"
	"log"
	"sync"
	"time"
	"runtime"
	"bufio"
	"fmt"
)	

var counter []map[string]string
var wg sync.WaitGroup
var mutex sync.Mutex

func Search(fa_path, format, target_fa, method string, mismatch, start, end int) map[string]string {

	var seqs []string
	var seq_name string
	var max = 500000
	var line int
	var group int
	var total_reads int
	seq_task := make(map[int][]string)
	
	fa_map, fa_name := Fasta2Map(target_fa)

	finder := FinderFactory(
		fa_map,
		fa_name,
		mismatch,
		start,
		end,
	)

	log.SetPrefix("[INFO] ")
	log.SetFlags(log.Ldate | log.Lmicroseconds)

	f, err := os.Open(fa_path)
	CheckError(err)
	f_scanner, k, seq_line := GetfqScaner(f, fa_path, format)

	for f_scanner.Scan() {
		line++

		if strings.HasPrefix(f_scanner.Text(),"@") {
			seq_name = f_scanner.Text() 
		}
		
		if line % k == seq_line {
			total_reads++
			seq := fmt.Sprintf("%s||%s", seq_name, f_scanner.Text())
			seqs = append(seqs, seq)
		}
		
		if len(seqs) == max {
			seq_task[group] = seqs
			group++
			seqs = []string{}
		}

		if len(seqs) != 0 {
			seq_task[group + 1] = seqs
		}

	}
	

	log.Println("Total Reads: ", total_reads)
    log.Println("Total Group: ", len(seq_task))

	// multi run
	wg.Add(len(seq_task))
	runtime.GOMAXPROCS(12)

	for k, _ := range seq_task {
		go startSearch(seq_task[k], finder, method)
	}
	wg.Wait()
	
	// merge
	merge_map := make(map[string]string)

	for _, r := range counter {
		for k, v := range r {
			merge_map[k] = v
		}
	}
	return merge_map
}

func CountSearchRes(prfx string, search_res map[string]string, outfile string) {
	count_map := make(map[string]int)
	file, err := os.OpenFile(outfile, os.O_WRONLY|os.O_CREATE, 0666)
	CheckError(err)
	defer file.Close()
	write := bufio.NewWriter(file)

	for _,v := range search_res {
		count_map[v]++
	}

	fmt.Fprintf(write,"%s\t%s\t%s\n", "sample_name", "target_name", "nReads")
	for k, v := range count_map {
		s := fmt.Sprintf("%s\t%s\t%d\n",prfx, k, v)
		fmt.Fprintf(write,s)
		write.Flush()
	}
	write.Flush()

}

func WriteSearchRes(search_res map[string]string, outfile string) {
	file, err := os.OpenFile(outfile, os.O_WRONLY|os.O_CREATE, 0666)
	CheckError(err)
	defer file.Close()
	write := bufio.NewWriter(file)

	for k, v := range search_res {
		s := fmt.Sprintf("%s\t%s\n",k, v)
		fmt.Fprintf(write,s)
		write.Flush()
	}
	write.Flush()
}

func startSearch(seq_q []string, finder Finder, method string) {
	defer wg.Done()
	res := make(map[string]string)

	for _, s := range seq_q {
		arr := strings.Split(s,"||")
		arr2 := strings.Split(arr[0]," ")
		target_name := finder.Find(arr[1], method)
		res[arr2[0]] = target_name
	}
	

	mutex.Lock()
    counter = append(counter, res)
	mutex.Unlock()
	time.Sleep(time.Millisecond)
	log.Println("Finish Searching reads number: ",len(seq_q) )
}