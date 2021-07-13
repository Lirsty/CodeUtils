package codereader

import "strings"

func (c *CodeReader) read(code []byte) {
	var (
		inFunc, inStruct bool
		f_deep, s_deep   int
		f_name, s_name   string
		func_p, struct_p int //startPoss
		s1               int //string.startPos
		blockPos         int
	)
	for pos, b := range code {
		//onNewLine
		if b == '\n' {
			newLine := string(code[s1:pos])
			//func
			if strings.Index(newLine, "func ") == 0 {
				inFunc = true
				func_p = s1
				f_deep = 1
				f_name = getFuncName(newLine)
				//addBlock
				c.code = append(c.code, &Block{Code: code[blockPos : s1-1]})
			}
			//struct
			if strings.Index(newLine, "type ") == 0 {
				s := strings.Split(newLine, " ")
				if s[2] == "struct" {
					inStruct = true
					struct_p = s1
					s_deep = 1
					s_name = s[1]
					//addBlock
					c.code = append(c.code, &Block{Code: code[blockPos : s1-1]})
				}
			}
			s1 = pos + 1
		}
		//onReadByte
		if inFunc {
			if b == '{' {
				f_deep++
			} else if b == '}' {
				if f_deep--; f_deep == 0 {
					inFunc = false
					part := code[func_p : pos+1]
					cc := make([]byte, len(part), (cap(part)+1)*2)
					copy(cc, part)
					newBlock := &Block{Code: cc}
					c.Func[f_name] = newBlock
					c.code = append(c.code, newBlock)
					blockPos = pos + 2
				}
			}
		}
		if inStruct {
			if b == '{' {
				s_deep++
			} else if b == '}' {
				if s_deep--; s_deep == 0 {
					inStruct = false
					part := code[struct_p : pos+1]
					cc := make([]byte, len(part), (cap(part)+1)*2)
					copy(cc, part)
					newBlock := &Block{Code: cc}
					c.Struct[s_name] = newBlock
					c.code = append(c.code, newBlock)
					blockPos = pos + 2
				}
			}
		}
	}
}
