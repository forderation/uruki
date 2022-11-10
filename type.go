package uruki

import "errors"

// constants of SpaceEncoding
const (
	WithoutEncoding       = " "   // Keep space as is
	PercentTwentyEncoding = "%20" // Change space into %20
	PlusEncoding          = "+"   // Change space into +

	ampersandStr = "&"
)

var (
	// ErrorInvalidSchemeURI error invalid scheme url not in restricted schemes
	ErrorInvalidSchemeURI = errors.New("invalid scheme url not in restricted schemes")
	// ErrorKeyEmpty key query parameter cannot be empty
	ErrorKeyEmpty = errors.New("key query parameter cannot be empty")
	// ErrorKeyContainSpace key query parameter cannot contains space
	ErrorKeyContainSpace = errors.New("key query parameter cannot contains space")
)
