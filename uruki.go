package uruki

import (
	"errors"
	"net/url"
)

var (
	// ErrInvalidURL invalid url
	ErrInvalidURL = errors.New("invalid scheme url, it should https or http prefix")
)

// Builder base struct uruki
type Builder struct {
	url                  *url.URL
	defaultSpaceEncode   string
	restrictedScheme     map[string]bool
	useEscapeAutomateURL bool
}

// Option options to create new Builder
type Option struct {
	// URL: url to proceed
	URL string
	// RestrictScheme: if you want your url scheme keep as is, example []string{"https"} if any url set with "http" options this will be return error
	RestrictScheme []string
	// DefaultSpaceEncode: default space encoding for internally encode. see SpaceEncoding const for more the details
	DefaultSpaceEncode string
	// UseEscapeAutomateURL to automate query escape on existing url query parameter while initiating builder / SetURL(uri string)
	UseEscapeAutomateURL bool
}

// NewBuilder create uruki (URi qUicK buIlder) parser & wrapper of net/url
func NewBuilder(options ...Option) (*Builder, error) {
	ub := &Builder{url: &url.URL{}}
	if len(options) > 0 {
		opt := options[0]
		ub.defaultSpaceEncode = opt.DefaultSpaceEncode
		ub.useEscapeAutomateURL = opt.UseEscapeAutomateURL
		ub.setRestrictedScheme(opt.RestrictScheme)
		err := ub.setURL(opt.URL)
		if err != nil {
			return nil, err
		}
		if ub.useEscapeAutomateURL {
			ub.queryEscapeAutomate()
		}
	}
	return ub, nil
}
