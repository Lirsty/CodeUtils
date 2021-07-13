package codereader

import (
	"strings"
)

type Block struct {
	Code []byte
}

func (b *Block) toString() []string {
	lines := make([]string, 0)
	s := 0
	for pos, v := range b.Code {
		if v == '\n' {
			lines = append(lines, string(b.Code[s:pos+1]))
			s = pos + 1
		}
	}
	return lines
}

func (b *Block) AddLine(s string) {
	new := append([]byte{'\n'}, []byte(s)...)
	b.Code = append(b.Code[0:len(b.Code)-2], new...)
	b.Code = append(b.Code, []byte("\n}\n")...)
}

func (b *Block) Remove(s string) {
	lines := b.toString()
	for i := 0; i < len(lines); i++ {
		if strings.Contains(lines[i], s) {
			lines = append(lines[:i], lines[i+1:]...)
			bytes := make([]byte, 0)
			for _, v := range lines {
				bytes = append(bytes, []byte(v)...)
			}
			b.Code = bytes
			break
		}
	}
}

func (b *Block) Replace(target, new string) {
	lines := b.toString()
	for i := 0; i < len(lines); i++ {
		if strings.Contains(lines[i], target) {
			lines[i] = new + "\r\n"
			bytes := make([]byte, 0)
			for _, v := range lines {
				bytes = append(bytes, []byte(v)...)
			}
			b.Code = bytes
			break
		}
	}
}
