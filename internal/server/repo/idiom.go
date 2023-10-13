package repo

type Idiom struct {
	Word        string `json:"word"`
	Pinyin      string `json:"pinyin"`
	Abbr        string `json:"abbr"`
	Explanation string `json:"explanation"`
	Quote       Quote  `json:"quote"`
	Source      Source `json:"source"`
}

type Quote struct {
	Text string `json:"text"`
	Book string `json:"book"`
}
type Source struct {
	Text string `json:"text"`
	Book string `json:"book"`
}
