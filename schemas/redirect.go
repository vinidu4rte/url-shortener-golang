package schemas

type Redirect struct {
	BasicModel
	OriginalUrl string `json:"originalUrl" gorm:"unique"`
	ShortUrl    string `json:"shortUrl" gorm:"unique"`
	UrlHits     int64  `json:"urlHits" gorm:"default: 0"`
}
