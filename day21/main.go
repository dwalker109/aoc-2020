package main

import (
	"fmt"
	"sort"
	"strings"

	"github.com/dwalker109/aoc-2020/util"
)

func main() {
	pt1 := part1("./input.txt")
	pt2 := part2("./input.txt")

	fmt.Println("Part 1: ", pt1)
	fmt.Println("Part 2: ", pt2)
}

func part1(p string) int {
	i := make(chan string)
	go util.StreamInput(i, p)
	labelling, _, allIngredients := intersectLabelling(i)
	known := eliminate(labelling)

	sum := 0
	for i, c := range allIngredients {
		if _, exists := known[i]; !exists {
			sum += c
		}
	}

	return sum
}

func part2(p string) string {
	i := make(chan string)
	go util.StreamInput(i, p)
	labelling, _, _ := intersectLabelling(i)
	known := eliminate(labelling)

	cdilKeys := make([]string, 0, len(known))
	cdilVals := make([]string, 0, len(known))
	flippedKnown := make(map[string]string)
	for i, a := range known {
		flippedKnown[a] = i
		cdilKeys = append(cdilKeys, a)
	}
	sort.Strings(cdilKeys)
	for _, k := range cdilKeys {
		cdilVals = append(cdilVals, flippedKnown[k])
	}

	return strings.Join(cdilVals, ",")
}

func intersectLabelling(i <-chan string) (map[string]map[string]bool, map[string]int, map[string]int) {
	allAllergens := make(map[string]int)
	allIngredients := make(map[string]int)
	aToI := make(map[string]map[string]bool)

	for s := range i {
		r := strings.NewReplacer(",", "", "(", "", ")", "")
		ia := strings.Split(r.Replace(s), "contains ")
		ingredients, allergens := strings.Fields(ia[0]), strings.Fields(ia[1])

		for _, a := range allergens {
			allAllergens[a]++
		}
		for _, i := range ingredients {
			allIngredients[i]++
		}

		for _, ca := range allergens {
			if _, exists := aToI[ca]; !exists {
				aToI[ca] = map[string]bool{}
				for _, ci := range ingredients {
					aToI[ca][ci] = true
				}
			} else {
				retain := make(map[string]bool)
				for _, ci := range ingredients {
					if _, intersects := aToI[ca][ci]; intersects {
						retain[ci] = true
					}
				}
				aToI[ca] = retain
			}
		}
	}

	return aToI, allAllergens, allIngredients
}

func eliminate(l map[string]map[string]bool) map[string]string {
	known := make(map[string]string)
	for retry := true; retry; {
		retry = false
		for a, i := range l {
			for i := range i {
				if _, toClean := known[i]; toClean {
					delete(l[a], i)
					retry = true
				}
			}
			if len(i) == 1 {
				for i := range i {
					known[i] = a
					retry = true
				}
			}
		}
	}

	return known
}
