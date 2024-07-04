package natcmp

import (
	"cmp"
	"regexp"
	"slices"
	"strconv"
)

var (
	slowExpr = regexp.MustCompile(`(?<num>\d+)|(?<rest>\D+)`)
	slowNum  = slowExpr.SubexpIndex("num")
	slowRest = slowExpr.SubexpIndex("rest")
)

func slowCompare(a, b string) int {
	am := slowExpr.FindAllStringSubmatch(a, -1)
	bm := slowExpr.FindAllStringSubmatch(b, -1)
	return slices.CompareFunc(am, bm, func(a []string, b []string) int {
		if a[slowNum] != "" && b[slowNum] != "" {
			aNum, _ := strconv.Atoi(a[slowNum])
			bNum, _ := strconv.Atoi(b[slowNum])
			return cmp.Compare(aNum, bNum)
		}
		return cmp.Compare(a[slowRest], b[slowRest])
	})
}
