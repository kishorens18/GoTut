package main

import (
	"fmt"
	"strings"
)

var pr = fmt.Println

func main() {
	str := "Iam kishore"
	pr(str)

	modStr := strings.NewReplacer("kishore", "bala")
	newStr := modStr.Replace(str)

	pr(newStr)

	pr("length:", len(newStr))

	pr("contains:", strings.Contains(newStr, "bala"))

	pr("index:", strings.Index(newStr, "a"))

	pr("replace:", strings.Replace(newStr, "bala", "kishore", -1))

	newStr2 := "iam katuvasiii"
	pr(len(newStr2))

	newStr3 := strings.TrimSpace(newStr2)

	pr(newStr3)
	pr(len(newStr3))

	pr("split:", strings.Split(newStr2, "a"))

	pr("toLower", strings.ToLower("KISHORE"))

	pr("toUpper", strings.ToUpper("kishore"))

	pr("hasPrefix", strings.HasPrefix("kishore", "ki"))

	pr("hasSuffix", strings.HasSuffix("kishore", "re"))

}
