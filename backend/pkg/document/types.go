package document

type Page struct {
	Number int
	Text   string
}

type Chunk struct {
	Index      int
	Content    string
	PageNumber int
}
