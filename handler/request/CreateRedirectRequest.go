package request

import (
	"fmt"
	"github.com/vinidu4rte/url-shortener-golang/helpers"
)

type CreateRedirectRequest struct {
	OriginalUrl string `json:"originalUrl"`
}

func (r *CreateRedirectRequest) Validate() error {
	if r.OriginalUrl == "" || !helpers.IsUrl(r.OriginalUrl) {
		return fmt.Errorf("it's necessary send originalUrl field with a valid URL")
	}

	return nil
}
