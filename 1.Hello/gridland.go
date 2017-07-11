package main

import (
	"fmt"
	"sort"
)

type pair struct {
	a, b uint64
}

type pairs []pair

func (p pairs) Len() int {
	return len(p)
}

func (p pairs) Swap(i, j int) {
	p[i], p[j] = p[j], p[i]
}

func (p pairs) Less(i, j int) bool {
	if p[i].a == p[j].a {
		return p[i].b < p[j].b
	}
	return p[i].a < p[j].a
}

func (p pairs) getActualTrackCount() uint64 {
	var connectdPairs pairs
	var current_c1 uint64
	var current_c2 uint64
	if len(p) > 1 {
		for _, each_pair := range p {
			if current_c1 == 0 && current_c2 == 0 {
				current_c1 = each_pair.a
				current_c2 = each_pair.b
			} else {
				if each_pair.a <= current_c2 && each_pair.b >= current_c2 {
					current_c2 = each_pair.b
				} else if each_pair.a <= current_c2 && each_pair.b < current_c2 {
					// Do Nothing
				} else {
					connectdPairs = append(connectdPairs, pair{a: current_c1, b: current_c2})
					current_c1 = each_pair.a
					current_c2 = each_pair.b
				}
			}
		}
		connectdPairs = append(connectdPairs, pair{a: current_c1, b: current_c2})
	} else {
		connectdPairs = p
	}
	var trackCount uint64
	for _, each_pair := range connectdPairs {
		trackCount += (each_pair.b - each_pair.a) + 1
	}
	return trackCount
}

func main() {
	var rows uint64
	var columns uint64
	var tracks uint64
	var r uint64
	var c1 uint64
	var c2 uint64
	fmt.Scan(&rows)
	fmt.Scan(&columns)
	fmt.Scan(&tracks)

	cityGrid := make(map[uint64]pairs, tracks)
	sum := rows * columns
	for i := uint64(0); i < tracks; i++ {
		fmt.Scan(&r)
		fmt.Scan(&c1)
		fmt.Scan(&c2)
		cityGrid[r] = append(cityGrid[r], pair{a: c1, b: c2})
	}
	for _, list_pairs := range cityGrid {
		sort.Sort(list_pairs)
	}

	var result uint64
	for _, list_pairs := range cityGrid {
		result += list_pairs.getActualTrackCount()
	}
	fmt.Println(sum - result)
}
