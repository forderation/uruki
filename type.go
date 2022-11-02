package uruki

import "errors"

type SpaceEncoding int

// constants of SpaceEncoding
const (
	WithoutEncoding       SpaceEncoding = iota + 1 // Keep space as is
	PercentTwentyEncoding                          // Change space into %20
	PlusEncoding                                   // Change space into +
)

// EncodeValue get encode translated value
func (se SpaceEncoding) EncodeValue() string {
	switch se {
	case PercentTwentyEncoding:
		return PercentTwentyStr
	case PlusEncoding:
		return PlusStr
	default:
		return SpaceStr
	}
}

const (
	SpaceStr         = " "
	PercentTwentyStr = "%20"
	PlusStr          = "+"
	AmpersandStr     = "&"
)

var (
	ErrorInvalidSchemeUri = errors.New("invalid scheme url not in restricted schemes")
	ErrorKeyEmpty         = errors.New("key query parameter cannot be empty")
	ErrorKeyContainSpace  = errors.New("key query parameter cannot contains space")
)
