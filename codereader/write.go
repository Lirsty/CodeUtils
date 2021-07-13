package codereader

func (c *CodeReader) Write() []byte {
	result := make([]byte, 0)
	for _, b := range c.code {
		result = append(result, b.Code...)
	}
	return result
}
