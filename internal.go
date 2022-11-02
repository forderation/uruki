package uruki

import (
	"net/url"
	"strings"
)

// setUrl setUrl parsing url and set internal url
func (ub *UrukiBuilder) setUrl(rawUrl string) error {
	uri, err := ub.parseUrl(rawUrl)
	if err != nil {
		return err
	}
	ub.mutex.Lock()
	ub.url = uri
	ub.mutex.Unlock()
	return nil
}

// parseUrl parsing given uri and check if scheme in restricted if using restricted scheme
func (ub *UrukiBuilder) parseUrl(rawUri string) (*url.URL, error) {
	uri, err := url.Parse(rawUri)
	if err != nil {
		return nil, err
	}
	// check it's an acceptable scheme
	ub.mutex.RLock()
	defer ub.mutex.RUnlock()
	if len(ub.restrictedScheme) > 0 && !ub.restrictedScheme[uri.Scheme] {
		return nil, ErrorInvalidSchemeUri
	}
	return uri, nil
}

// setRestrictedScheme lookup restricted scheme save into map
func (ub *UrukiBuilder) setRestrictedScheme(schemes []string) {
	mapScheme := make(map[string]bool, 0)
	for _, v := range schemes {
		v = strings.ToLower(strings.TrimSpace(v))
		mapScheme[v] = true
	}
	ub.mutex.Lock()
	ub.restrictedScheme = mapScheme
	ub.mutex.Unlock()
}

// queryEscapeAutomate escape all query parameter from existing url
func (ub *UrukiBuilder) queryEscapeAutomate() {
	ub.mutex.RLock()
	rawQuery := ub.url.RawQuery
	ub.mutex.RUnlock()
	if len(rawQuery) > 0 {
		keyVal := strings.Split(rawQuery, AmpersandStr)
		buildRawResult := make([]string, 0)
		for _, queryParam := range keyVal {
			qv := strings.Split(queryParam, "=")
			if len(qv) > 0 {
				q := qv[0]
				v := ""
				if len(qv) > 1 {
					v = qv[1]
				}
				vParsed, err := url.QueryUnescape(v)
				if err != nil {
					vParsed = v
				} else {
					vParsed = url.QueryEscape(vParsed)
					vParsed = strings.ReplaceAll(vParsed, PlusStr, ub.defaultSpaceEncode.EncodeValue())
				}
				buildRawResult = append(buildRawResult, q+"="+vParsed)
			}
		}
		if len(buildRawResult) > 0 {
			ub.mutex.Lock()
			ub.url.RawQuery = strings.Join(buildRawResult, AmpersandStr)
			ub.mutex.Unlock()
		}
	}
}
