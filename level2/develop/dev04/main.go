package main

import (
	"fmt"
	"sort"
	"strings"
)

var (
	dict = map[rune]int{
		'а': 0,
		'б': 1,
		'в': 2,
		'г': 3,
		'д': 4,
		'е': 5,
		'ё': 5,
		'ж': 6,
		'з': 7,
		'и': 8,
		'й': 9,
		'к': 10,
		'л': 11,
		'м': 12,
		'н': 13,
		'о': 14,
		'п': 15,
		'р': 16,
		'с': 17,
		'т': 18,
		'у': 19,
		'ф': 20,
		'х': 21,
		'ц': 22,
		'ч': 23,
		'ш': 24,
		'щ': 25,
		'ъ': 26,
		'ы': 27,
		'ь': 28,
		'э': 29,
		'ю': 30,
		'я': 31,
	}
)

func main() {
	data := []string{
		"листок",
		"тяпка",
		"пятка",
		"пятак",
		"Пятак",
		"пятак",
		"слиток",
		"столик",
	}

	fmt.Println(makeSet(data))
}

func makeSet(data []string) map[string][]string {
	set := make(map[string]*[]string)
	rez := make(map[string][]string)
	for _, w := range data {
		w = strings.ToLower(w)
		key := encode(w)

		if ans, ok := set[key]; !ok {
			l := make([]string, 0, 1)
			l = append(l, w)
			set[key] = &l
		} else {
			*ans = append(*ans, w)
		}
	}

	// filter
	for _, list := range set {
		un := map[string]struct{}{}
		for i := 0; i < len(*list); {
			if _, ok := un[(*list)[i]]; ok {
				if i < len(*list)-1 {
					*list = append((*list)[:i], (*list)[i+1:]...)
				} else {
					*list = (*list)[: i : len(*list)-1]
				}
			} else {
				un[(*list)[i]] = struct{}{}
				i++
			}
		}
	}
	// chek size
	for key := range set {
		if len(*set[key]) > 1 {
			rez[(*set[key])[0]] = *set[key]
			sort.Strings(*set[key])
		}
	}
	fmt.Println(rez)
	return rez
}

func encode(word string) (key string) {
	var bkey [32]byte
	for _, i := range word {
		bkey[dict[i]]++
	}
	key = string(bkey[:])
	return
}
