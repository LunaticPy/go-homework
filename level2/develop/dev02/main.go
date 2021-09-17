package unpack

import (
	"fmt"
	"strconv"
)

var (
	Separator = `\`
)

func Unpack(line string) string {
	rez := ""
	for i := 0; i < len(line); i++ {

		num, err := strconv.Atoi(line[i : i+1])
		if err != nil {
			if line[i:i+1] == Separator {
				i++
				if i == len(line) {
					break
				}
			}
			rez = rez + line[i:i+1]
			continue
		}
		if len(rez) == 0 {
			continue
		}
		if i > 1 && line[i-2:i-1] == Separator {
			num--
		}
		buf := rez[len(rez)-1:]
		for j := 0; j < num; j++ {
			rez = rez + buf
		}
	}
	fmt.Println(rez)
	return rez
}
