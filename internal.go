package uruki

import (
	"net/url"
	"strings"
)

// setURL setURL parsing url and set internal url
func (ub *Builder) setURL(rawURL string) error {
	uri, err := ub.parseURL(rawURL)
	if err != nil {
		return err
	}
	ub.url = uri
	return nil
}

// parseURL parsing given uri and check if scheme in restricted if using restricted scheme
func (ub *Builder) parseURL(rawURL string) (*url.URL, error) {
	uri, err := url.Parse(rawURL)
	if err != nil {
		return nil, err
	}
	// check it's an acceptable scheme
	if len(ub.restrictedScheme) > 0 && !ub.restrictedScheme[uri.Scheme] {
		return nil, ErrorInvalidSchemeURI
	}
	return uri, nil
}

// setRestrictedScheme lookup restricted scheme save into map
func (ub *Builder) setRestrictedScheme(schemes []string) {
	mapScheme := make(map[string]bool)
	for _, v := range schemes {
		v = strings.ToLower(strings.TrimSpace(v))
		mapScheme[v] = true
	}
	ub.restrictedScheme = mapScheme
}

// queryEscapeAutomate escape all query parameter from existing url
func (ub *Builder) queryEscapeAutomate() {
	rawQuery := ub.url.RawQuery
	if len(rawQuery) > 0 {
		keyVal := strings.Split(rawQuery, ampersandStr)
		buildRawResult := make([]string, 0)
		for _, queryParam := range keyVal {
			qv := strings.Split(queryParam, "=")
			if len(qv) > 0 {
				q := qv[0]
				v := ""
				if len(qv) > 1 {
					v = qv[1]
				}
				qParsed, err := url.QueryUnescape(q)
				if err != nil {
					qParsed = v
				} else {
					qParsed = url.QueryEscape(qParsed)
					qParsed = strings.ReplaceAll(qParsed, PlusEncoding, ub.defaultSpaceEncode)
				}
				vParsed, err := url.QueryUnescape(v)
				if err != nil {
					vParsed = v
				} else {
					vParsed = url.QueryEscape(vParsed)
					vParsed = strings.ReplaceAll(vParsed, PlusEncoding, ub.defaultSpaceEncode)
				}
				buildRawResult = append(buildRawResult, qParsed+"="+vParsed)
			}
		}
		if len(buildRawResult) > 0 {
			ub.url.RawQuery = strings.Join(buildRawResult, ampersandStr)
		}
	}
}

// removeIndex remove index from string slice
func removeIndex(s []string, index int) []string {
	return append(s[:index], s[index+1:]...)
}
