package utils

import (
	"fmt"
	"strings"
	//"strconv"
)

type Finder struct {
	Fasta_map  map[string]string
	Fasta_name []string
	Mismatch   int
	Start      int
	End        int
}

func FinderFactory(
	Fasta_map map[string]string,
	Fasta_name []string,
	Mismatch,
	Start,
	End int) Finder {
	r := Finder{Fasta_map, Fasta_name, Mismatch, Start, End}
	return r
}

func (f *Finder) Find(seq string, method string) string {
	
	var seq_name string

	if method == "all" {
		seq_name = f.findAll(seq)
	}

	if method == "first" {
		seq_name = f.findFirst(seq)
	}
	return seq_name
}

//func (f*Finder) intSlie2Str(s []int) string {/
//	var stringSlice []string
//	for _, num := range intSlice {
//		stringSlice = append(stringSlice, strconv.Itoa(num))
//	}
//	return strings.Join(stringSlice, ",")
//}


func (f *Finder) findAll(seq string) string {
	var seq_use string
	var target_name_name_out []string
	var locs []int

	for _, target_name := range f.Fasta_name {
		target_seq := f.Fasta_map[target_name]
		
		end := f.End + len(target_seq)
		
		if end >= len(seq){
			end = len(seq)
		}
		
		seq_use = seq[f.Start:end]
		
		switch {
			case f.Mismatch == 0:
				locs = KMPSearch(seq_use, target_seq, true)
			
			case f.Mismatch > 0:
				locs = HMSearch(seq_use, target_seq,f.Mismatch, true)
		}

		if len(locs) == 0 {
			continue
		}else{
			s := fmt.Sprintf("%s", target_name)
			target_name_name_out = append(target_name_name_out, s)
		}

	}

	switch {
		case len(target_name_name_out) == 0: return "NULL"
		case len(target_name_name_out) == 1: return target_name_name_out[0]
		case len(target_name_name_out) > 1: return strings.Join(target_name_name_out,"|") 
	}
	return "NULL"
}

func (f *Finder) findFirst(seq string) string {

	var seq_use string
	var locs []int

	for _, target_name := range f.Fasta_name {
		target_seq := f.Fasta_map[target_name]

		end := f.End + len(target_seq)
		
		if end >= len(seq){
			end = len(seq)
		}
		seq_use = seq[f.Start:end]
		
		switch {
		case f.Mismatch == 0:
			locs = KMPSearch(seq_use, target_seq, true)
		
		case f.Mismatch > 0:
			locs = HMSearch(seq_use, target_seq,f.Mismatch, true)
		}

		if len(locs) == 0 {
			continue
		}else {
			return fmt.Sprintf("%s", target_name)
		}
	}
	return "NULL"
}