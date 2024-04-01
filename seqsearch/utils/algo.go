package utils

import "fmt"

func hm(seq1 string, seq2 string, max_dist int) int {
	
	if len(seq1) != len(seq2) {
		err_info := fmt.Sprintf("%s must equal to %s", seq1, seq2)
		panic(err_info)
	}

		dis := 0
		for i := 0; i < len(seq1); i++ {
			if seq1[i] != seq2[i] {
				dis++
				if dis > max_dist {
					return dis
				}
			}	
		}
		return dis
	}

func HMSearch(seq, target_seq string, max_dist int, isfirst bool) []int {
	var locs []int
	loop_n := len(seq) - len(target_seq) + 1
		
	for i := 0; i != loop_n; i++ {
		hm_dis := hm(seq[i : i + len(target_seq)], target_seq, max_dist)
		
		if hm_dis <= max_dist {
			locs = append(locs, i)
			if isfirst {
				return locs
			}
		}
	}
	return locs
}

func computeLPSArray(target_seq string) []int {
	length := len(target_seq)
	lps := make([]int, length)
	lps[0] = 0
	i := 1
	j := 0

	for i < length {
		if target_seq[i] == target_seq[j] {
			j++
			lps[i] = j
			i++
		} else {
			if j != 0 {
				j = lps[j-1]
			} else {
				lps[i] = 0
				i++
			}
		}
	}
	return lps
}

func KMPSearch(seq, target_seq string, isfirst bool) []int {
	m := len(target_seq)
	n := len(seq)
	var result []int

	lps := computeLPSArray(target_seq)
	i, j := 0, 0
	for i < n {
		if target_seq[j] == seq[i] {
			i++
			j++
		}
		if j == m {
			result = append(result, i-j)
			
			if isfirst {
				return result
			}
			
			j = lps[j-1]
		} else if i < n && target_seq[j] != seq[i] {
			if j != 0 {
				j = lps[j-1]
			} else {
				i++
			}
		}
	}
	return result
}
