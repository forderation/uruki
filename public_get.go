package uruki

import (
	"net/url"
	"strings"
)

// GetURLResultUnescape get result url with unescape string
func (ub *Builder) GetURLResultUnescape() string {
	uri := ub.url.String()
	unescapeURL, err := url.QueryUnescape(uri)
	if err != nil {
		unescapeURL = uri
	}
	return unescapeURL
}

// GetURLResult get url result with escaped option
func (ub *Builder) GetURLResult() string {
	return ub.url.String()
}

// GetValueQuery get value of existing query parameter if any, return as decoded value
func (ub *Builder) GetValueQuery(key string) string {
	query := ub.url.Query()
	return query.Get(key)
}

// GetAllQueryValue get all key-value of existing query parameter, return as map key and decoded value
func (ub *Builder) GetAllQueryValue() map[string][]string {
	return ub.url.Query()
}

// GetInternalURL get internal url that already parsed by uruki
func (ub *Builder) GetInternalURL() url.URL {
	return *ub.url
}

// GetPaths get each part of path in slice
func (ub *Builder) GetPaths() []string {
	paths := strings.Split(ub.url.Path, "/")
	pathsStrip := make([]string, 0)
	for i, v := range paths {
		if i == 0 {
			continue
		}
		pathsStrip = append(pathsStrip, v)
	}
	return pathsStrip
}

// GetFullPath get full path. exclude base url, query parameter and fragment
func (ub *Builder) GetFullPath() string {
	return ub.url.Path
}
