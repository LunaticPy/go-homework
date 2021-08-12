package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"sort"
	"strconv"
	"strings"

	"github.com/facette/natsort" // natural sorting
)

var (
	reverse, numeric, unic, mnth, b, c, h bool
	col                                   int
	colSeparator                          string = " "
)

func init() {
	flag.IntVar(&col, "k", 0, "указание колонки для сортировки")            //+
	flag.BoolVar(&numeric, "n", false, "сортировать по числовому значению") //+
	flag.BoolVar(&reverse, "r", false, "сортировать в обратном порядке")    //+
	flag.BoolVar(&unic, "u", false, "не выводить повторяющиеся строки")     //+

	flag.BoolVar(&mnth, "M", false, "сортировать по названию месяца")                    //+
	flag.BoolVar(&b, "b", false, "игнорировать хвостовые пробелы")                       //+
	flag.BoolVar(&c, "c", false, "проверять отсортированы ли данные")                    //+
	flag.BoolVar(&h, "h", false, "сортировать по числовому значению с учётом суффиксов") //+
	flag.StringVar(&colSeparator, "t", " ", "таб сепаратор")                             //+
}

func main() {
	flag.Parse()

	filename := flag.Arg(0)
	content, err := ioutil.ReadFile(filename)
	if err != nil {
		fmt.Println("error opening file: err:", err)
		os.Exit(1)
	}

	rez := MySort(content)

	for i := 0; i < len(rez); i++ {

		fmt.Println(rez[i])
	}

}

func MySort(content []byte) []string {
	if b {
		content = []byte(strings.TrimSpace(string(content)))
	}
	rez := strings.Split(string(content), "\n")

	if col > 0 {
		buf := make([]string, 0, len(rez))
		for _, r := range rez {
			buf = append(buf, strings.Split(r, colSeparator)[col-1])
		}
		rez = buf
	}

	if unic {
		un := map[string]struct{}{}
		for _, r := range rez {
			un[r] = struct{}{}
		}
		rez = make([]string, 0, len(un))
		for k := range un {
			rez = append(rez, k)
		}
	}
	if numeric {
		nums := make([]int, 0, len(rez))
		if c {
			if !sort.IntsAreSorted(nums) {
				return []string{"Need to sort"}
			} else {
				return []string{"No need to sort"}
			}
		}
		for k := range rez {
			n, err := strconv.Atoi(rez[k])
			if err != nil {
				fmt.Println("error sort numeric:", err)
				os.Exit(1)
			}
			nums = append(nums, n)
		}
		if reverse {
			sort.Sort(sort.Reverse(sort.IntSlice(nums)))
		} else {
			sort.Ints(nums)
		}

		rez = make([]string, 0, len(nums))
		for _, n := range nums {

			rez = append(rez, strconv.Itoa(n))
		}
	}
	if mnth {
		sF := func(i, j int) bool {
			return getMonth(rez[i]) < getMonth(rez[j])
		}
		rsF := func(i, j int) bool {
			return getMonth(rez[i]) > getMonth(rez[j])
		}

		if reverse {
			sort.SliceStable(rez, rsF)
		} else {
			sort.SliceStable(rez, sF)
		}
	}

	if !numeric && !mnth && !h {
		if c {
			if !sort.StringsAreSorted(rez) {
				return []string{"Need to sort"}
			} else {
				return []string{"No need to sort"}
			}
		}
		if reverse {
			sort.Sort(sort.Reverse(sort.StringSlice(rez)))
		} else {
			sort.Strings(rez)
		}

	}
	if h {
		natsort.Sort(rez)
	}
	if col > 0 {
		buf := make(map[string]string, len(rez))
		for _, c := range strings.Split(string(content), "\n") {

			line := strings.Split(c, colSeparator)
			buf[line[col-1]] = c
		}

		for i := 0; i < len(rez); i++ {
			rez[i] = buf[rez[i]]
		}

	}

	return rez
}
