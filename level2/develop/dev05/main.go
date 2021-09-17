package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"regexp"
	"strconv"
	"strings"
)

var (
	I, v, F, num, c bool
	A, B, C         int
)

func init() {
	flag.IntVar(&A, "A", 0, "after печатать +N строк после совпадения")          //+
	flag.IntVar(&B, "B", 0, "before печатать +N строк до совпадения")            //+
	flag.IntVar(&C, "C", 0, "context (A+B) печатать ±N строк вокруг совпадения") //+

	flag.BoolVar(&c, "c", false, "count (количество строк)")                        //+
	flag.BoolVar(&I, "i", false, "ignore-case (игнорировать регистр)")              //+
	flag.BoolVar(&v, "v", false, "invert (вместо совпадения, исключать)")           //+
	flag.BoolVar(&F, "F", false, "fixed, точное совпадение со строкой, не паттерн") //+
	flag.BoolVar(&num, "n", false, "line num, печатать номер строки")               //+
}

func main() {
	flag.Parse()

	pattern := flag.Arg(0)
	if pattern == "" {
		os.Exit(2)
	}
	filename := flag.Arg(1)
	content, err := ioutil.ReadFile(filename)
	if err != nil {
		fmt.Println("error opening file: err:", err)
		os.Exit(1)
	}

	rez := MyGrep(pattern, content)
	for _, line := range rez {
		if line == "" {
			continue
		}
		fmt.Println(line)
	}

}

func MyGrep(pattern string, content []byte) (rez []string) {
	shift := 0

	if I {
		rez = strings.Split(strings.ToLower(string(content)), "\n")
		pattern = strings.ToLower(pattern)
	} else {
		rez = strings.Split(string(content), "\n")
	}

	for i := 0; i < len(rez); i++ {
		if match := check(pattern, rez[i]); !match || (v && (B > 0 || C > 0)) {
			if B > 0 || C > 0 {
				flag := false
				for j := 0; (j <= B || j <= C) && i+j < len(rez); j++ {
					if match := check(pattern, rez[i+j]); match && !v {
						i = i + j - 1
						flag = true
						break
					}
				}
				if flag {
					continue
				}
			}
			if !v {
				if num {
					shift++
				}
				rez = append(rez[:i], rez[i+1:]...)
				i--
				continue
			}
		} else if A > 0 || C > 0 {

			for j := 1; (j <= A || j <= C) && i+j < len(rez); j++ {

				if match := check(pattern, rez[i+j]); match && !v {
					i = i + j - 1
					break
				} else if match && v {
					if num {
						shift++
					}
					rez = append(rez[:i], rez[i+1:]...)
					i--
				} else if i+j+1 == len(rez) {
					i = i + j
				}
			}
			continue
		} else if v && match {
			if num {
				shift++
			}
			rez = append(rez[:i], rez[i+1:]...)
			i--
		}
	}

	if num {
		var cont []string
		if I {
			cont = strings.Split(strings.ToLower(string(content)), "\n")
		} else {
			cont = strings.Split(string(content), "\n")
		}
		for i, j := 0, 0; i < len(cont) && j < len(rez); i++ {
			if cont[i] == rez[j] {
				rez[j] = fmt.Sprintf("%d:%s", i+1, cont[i])
				j++
			}
		}
	}
	if c {
		return []string{strconv.Itoa(len(rez))}

	}
	if len(rez) == 0 {
		return []string{""}
	}
	return rez
}

func check(pattern, line string) bool {
	if !F {
		re := regexp.MustCompile(pattern)
		return re.MatchString(line)
	} else {
		return pattern == line
	}
}
