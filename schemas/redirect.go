package schemas

import "gorm.io/gorm"

type Redirect struct {
	gorm.Model
	OriginalUrl string
	ShortUrl    string
	UrlHits     int64
}
