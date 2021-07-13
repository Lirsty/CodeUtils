package codereader

type CodeReader struct {
	code   []*Block
	Func   map[string]*Block
	Struct map[string]*Block
}

func New(code []byte) *CodeReader {
	reader := &CodeReader{Func: make(map[string]*Block), Struct: make(map[string]*Block)}
	reader.read(code)
	return reader
}
