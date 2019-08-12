package kata

import "strings"

type FParam func(string) string

func Oper(f FParam, x string) string {
	return f(x)
}

func VertMirror(s string) (mirrored string) {
	slice := strings.Split(s, "\n")
	for i, row := range slice {
		for sub := len(row); sub > 0; sub-- {
			mirrored += string(row[sub-1])
		}
		if i != len(slice)-1 {
			mirrored += "\n"
		}
	}
	return
}

func HorMirror(s string) (mirrored string) {
	slice := strings.Split(s, "\n")
	for i := len(slice); i > 0; i-- {
		for _, v := range slice[i-1] {
			mirrored += string(v)
		}
		if i != 1 {
			mirrored += "\n"
		}
	}
	return
}
