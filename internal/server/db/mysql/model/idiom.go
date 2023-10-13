package model

type Idiom struct {
	Id          string `gorm:"id;primarykey"`
	Word        string `gorm:"word"`
	Pinyin      string `gorm:"pinyin"`
	Abbr        string `gorm:"abbr"`
	Explanation string `gorm:"explanation"`
}

func (idiom Idiom) TableName() string {
	return "idiom"
}

type Quote struct {
	Id   string `gorm:"id;primarykey"`
	Text string `gorm:"text"`
	Book string `gorm:"book"`
}

func (quote Quote) TableName() string {
	return "quote"
}

type Source struct {
	Id   string `gorm:"id;primarykey"`
	Text string `gorm:"text"`
	Book string `gorm:"book"`
}

func (source Source) TableName() string {
	return "source"
}
