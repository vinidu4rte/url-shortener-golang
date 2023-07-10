package schemas

import "gorm.io/gorm"

type Redirect struct {
	gorm.Model
	OriginalUrl string `json:"originalUrl;unique_index"`
	ShortUrl    string `json:"shortUrl;unique_index"`
	UrlHits     int64  `json:"urlHits" gorm:"default: 0"`
}
