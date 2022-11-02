package uruki

import (
	"net/url"
	"strings"
)

// GetUrlResultUnescape get result url with unescape string
func (ub *UrukiBuilder) GetUrlResultUnescape() string {
	ub.mutex.RLock()
	defer ub.mutex.RUnlock()
	uri := ub.url.String()
	unescapeUrl, err := url.QueryUnescape(uri)
	if err != nil {
		unescapeUrl = uri
	}
	return unescapeUrl
}

// GetUrlResult get url result with escaped option
func (ub *UrukiBuilder) GetUrlResult() string {
	ub.mutex.RLock()
	defer ub.mutex.RUnlock()
	return ub.url.String()
}

// GetValueQuery get value of existing query parameter if any, return as decoded value
func (ub *UrukiBuilder) GetValueQuery(key string) string {
	ub.mutex.RLock()
	defer ub.mutex.RUnlock()
	query := ub.url.Query()
	return query.Get(key)
}

// GetAllQueryValue get all key-value of existing query parameter, return as map key and decoded value
func (ub *UrukiBuilder) GetAllQueryValue() map[string][]string {
	ub.mutex.RLock()
	defer ub.mutex.RUnlock()
	return ub.url.Query()
}

// GetInternalUrl get internal url that already parsed by uruki
func (ub *UrukiBuilder) GetInternalUrl() url.URL {
	ub.mutex.RLock()
	defer ub.mutex.RUnlock()
	return *ub.url
}

// GetPaths get each part of path in slice
func (ub *UrukiBuilder) GetPaths() []string {
	ub.mutex.RLock()
	paths := strings.Split(ub.url.Path, "/")
	ub.mutex.RUnlock()
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
func (ub *UrukiBuilder) GetFullPath() string {
	ub.mutex.RLock()
	defer ub.mutex.RUnlock()
	return ub.url.Path
}
