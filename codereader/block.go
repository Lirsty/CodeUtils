package codereader

type Block struct {
	Code []byte
}

func (b *Block) AddLine(s string) {
	new := []byte(s)
	b.Code = append(b.Code[0:len(b.Code)-2], new...)
	b.Code = append(b.Code, []byte("\n}\n")...)
}
