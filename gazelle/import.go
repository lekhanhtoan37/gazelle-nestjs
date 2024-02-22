package nestjs

import "strings"

func PkgName(dirpath string) string {
	dirpath = strings.TrimRight(dirpath, "/")
	n := len(dirpath) - 1
	for i := n; i >= 0; i-- {
		if dirpath[i] == '/' {
			return dirpath[i+1 : n+1]
		}
	}
	return dirpath
}
