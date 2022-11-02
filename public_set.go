package uruki

import (
	"net/url"
	"strings"
)

// AddQueryParamOpt parameter for adding query params value in url
type AddQueryParamOpt struct {
	// key query
	Key string
	// value query
	Val string
	// if you want override DefaultSpaceEncoding
	SpaceEnc SpaceEncoding
}

// AddQueryParamOpt parameter for adding query params value in url
type SetFragmentOpt struct {
	// fragment value
	Fragment string
	// if you want override DefaultSpaceEncoding
	SpaceEnc SpaceEncoding
}

// SetBaseUrl change or update existing of base url only host and port
func (ub *UrukiBuilder) SetBaseUrl(baseUrl string) error {
	uri, err := ub.parseUrl(baseUrl)
	if err != nil {
		return err
	}
	ub.mutex.Lock()
	defer ub.mutex.Unlock()
	if len(uri.Host) < 1 && len(ub.url.Path) > 0 {
		firstChar := ub.url.Path[0]
		if firstChar == ("/"[0]) {
			ub.url.Path = strings.Replace(ub.url.Path, "/", "", 1)
		}
	}
	ub.url.Host = uri.Host
	ub.url.Scheme = uri.Scheme
	return nil
}

// SetPath change or update path only of url
func (ub *UrukiBuilder) SetPath(path string) {
	ub.mutex.Lock()
	defer ub.mutex.Unlock()
	ub.url.Path = path
}

// AddQueryParam add query parameter values, internally will be encode the value of query
func (ub *UrukiBuilder) AddQueryParam(opt AddQueryParamOpt) error {
	ub.mutex.Lock()
	defer ub.mutex.Unlock()
	key := strings.TrimSpace(opt.Key)
	if len(key) < 1 {
		return ErrorKeyEmpty
	}
	if strings.Contains(key, " ") {
		return ErrorKeyContainSpace
	}
	if opt.SpaceEnc < 1 {
		opt.SpaceEnc = ub.defaultSpaceEncode
	}
	value := url.QueryEscape(opt.Val)
	value = strings.ReplaceAll(value, PlusStr, opt.SpaceEnc.EncodeValue())
	rawQuery := ub.url.RawQuery
	if len(rawQuery) > 0 {
		rawQuery += AmpersandStr
	}
	ub.url.RawQuery = rawQuery + key + "=" + value
	return nil
}

// DeleteFragment remove existing fragment if any
func (ub *UrukiBuilder) DeleteFragment() {
	ub.mutex.Lock()
	defer ub.mutex.Unlock()
	ub.url.Fragment = ""
}

// SetFragment create / update existing fragment for references, without '#'
func (ub *UrukiBuilder) SetFragment(opt SetFragmentOpt) {
	ub.mutex.Lock()
	defer ub.mutex.Unlock()
	value := url.QueryEscape(opt.Fragment)
	if opt.SpaceEnc < 1 {
		opt.SpaceEnc = ub.defaultSpaceEncode
	}
	value = strings.ReplaceAll(value, PercentTwentyStr, opt.SpaceEnc.EncodeValue())
	ub.url.Fragment = value
}

// DeleteKeyQuery delete key query parameter if exist
func (ub *UrukiBuilder) DeleteKeyQuery(keyDelete string) {
	ub.mutex.Lock()
	defer ub.mutex.Unlock()
	buildRawResult := make([]string, 0)
	keyVal := strings.Split(ub.url.RawQuery, AmpersandStr)
	for _, queryParam := range keyVal {
		qv := strings.Split(queryParam, "=")
		if len(qv) > 0 {
			q := qv[0]
			v := ""
			if len(qv) > 1 {
				v = qv[1]
			}
			if q == keyDelete {
				continue
			}
			buildRawResult = append(buildRawResult, q+"="+v)
		}
	}
	if len(buildRawResult) > 0 {
		ub.url.RawQuery = strings.Join(buildRawResult, AmpersandStr)
	}
}
