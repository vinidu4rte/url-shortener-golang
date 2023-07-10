package request

import "fmt"

type CreateRedirectRequest struct {
	OriginalUrl string `json:"originalUrl"`
}

func (r *CreateRedirectRequest) Validate() error {
	if r.OriginalUrl == "" {
		return fmt.Errorf("it's necessary send originalUrl field")
	}

	return nil
}
