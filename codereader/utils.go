package codereader

import (
	"strings"
)

func getFuncName(s string) string {
	b := []byte(s)
	pos := 0
	isStruct := false
	for i, v := range b {
		if v == '(' {
			pos = i
			if b[i-1] == ' ' {
				isStruct = true
			}
		}
	}
	if isStruct {
		count := 0
		p := 0
		for i, v := range b {
			if v == ' ' {
				count++
			}
			if count == 2 {
				p = i
				if b[i+1] == '*' {
					p = i + 1
				}
				break
			}
		}
		i := strings.Index(s, ") ")
		return string(append(append(b[p+1:i], '.'), b[i+2:pos]...))
	}
	return string(b[5:pos])
}
