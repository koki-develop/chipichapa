package main

import (
	"fmt"
	"math/rand"
	"strings"
)

func main() {
	items := []string{"ﾁﾋﾟ", "ﾁｬﾊﾟ", "ﾄﾞｩﾋﾞ", "ﾀﾞﾊﾞ", "ﾏﾋ", "ｺﾐ"}
	want := "ﾁﾋﾟﾁﾋﾟﾁｬﾊﾟﾁｬﾊﾟﾄﾞｩﾋﾞﾄﾞｩﾋﾞﾀﾞﾊﾞﾀﾞﾊﾞﾏﾋｺﾐﾄﾞｩﾋﾞﾄﾞｩﾋﾞ"

	s := new(strings.Builder)
	for {
		item := items[rand.Intn(len(items))]
		fmt.Print(item)
		s.WriteString(item)

		got := s.String()
		if !strings.HasPrefix(want, got) {
			s.Reset()
			continue
		}
		if got == want {
			break
		}
	}

	fmt.Println("ﾌﾞｰﾝﾌﾞｰﾝﾌﾞｰﾝﾌﾞｰﾝ")
}
