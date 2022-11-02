package uruki

import (
	"errors"
	"net/url"
	"sync"
)

var (
	ErrInvalidUrl = errors.New("invalid scheme url, it should https or http prefix")
)

// UrukiBuilder base struct uruki
type UrukiBuilder struct {
	url                *url.URL
	defaultSpaceEncode SpaceEncoding
	restrictedScheme   map[string]bool
	mutex              sync.RWMutex
}

// UrukiOption options to create new UrukiBuilder
type UrukiOption struct {
	// URL: url to proceed
	URL string
	// RestrictScheme: if you want your url scheme keep as is, example []string{"https"} if any url set with "http" options this will be return error
	RestrictScheme []string
	// DefaultSpaceEncode: default space encoding for internally encode. see SpaceEncoding const for more the details
	DefaultSpaceEncode SpaceEncoding
	// UseEscapeAutomateInit to automate query escape on existing url query parameter while initiating builder
	UseEscapeAutomateInit bool
}

// NewUrukiBuilder create uruki (URi qUicK buIlder) parser & wrapper of net/url
func NewUrukiBuilder(options ...UrukiOption) (*UrukiBuilder, error) {
	ub := &UrukiBuilder{mutex: sync.RWMutex{}, url: &url.URL{}}
	if len(options) > 0 {
		opt := options[0]
		ub.defaultSpaceEncode = opt.DefaultSpaceEncode
		ub.setRestrictedScheme(opt.RestrictScheme)
		err := ub.setUrl(opt.URL)
		if err != nil {
			return nil, err
		}
		if opt.UseEscapeAutomateInit {
			ub.queryEscapeAutomate()
		}
	}
	return ub, nil
}
