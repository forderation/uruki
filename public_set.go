package uruki

import (
	"net/url"
	"strings"
	"sync"
)

// AddQueryParamOpt parameter for adding query params value in url
type AddQueryParamOpt struct {
	// key query
	Key string
	// value query
	Val string
	// if you want using DefaultSpaceEncoding as SpaceEnc
	UseDefaultEncode bool
	// specify space encode if you use custom encoding
	SpaceEnc string
}

// SetFragmentOpt parameter for adding query params value in url
type SetFragmentOpt struct {
	// fragment value
	Fragment string
	// if you want using DefaultSpaceEncoding as SpaceEnc
	UseDefaultEncode bool
	// specify space encode if you use custom encoding
	SpaceEnc string
}

// SetBaseURL change or update existing of base url only host and port
func (ub *Builder) SetBaseURL(baseURL string) error {
	uri, err := ub.parseURL(baseURL)
	if err != nil {
		return err
	}
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
func (ub *Builder) SetPath(path string) {
	ub.url.Path = path
}

// SetURL replace all url with new url based on parameter, if error keep old url
func (ub *Builder) SetURL(uri string) error {
	err := ub.setURL(uri)
	if err != nil {
		return err
	}
	if ub.useEscapeAutomateURL {
		ub.queryEscapeAutomate()
	}
	return nil
}

// AddQueryParam add query parameter values, internally will be encode the value of query
func (ub *Builder) AddQueryParam(opt AddQueryParamOpt) error {
	key := strings.TrimSpace(opt.Key)
	if len(key) < 1 {
		return ErrorKeyEmpty
	}
	if strings.Contains(key, " ") {
		return ErrorKeyContainSpace
	}
	if opt.UseDefaultEncode {
		opt.SpaceEnc = ub.defaultSpaceEncode
	}
	key = url.QueryEscape(key)
	key = strings.ReplaceAll(key, PlusEncoding, opt.SpaceEnc)
	value := url.QueryEscape(opt.Val)
	value = strings.ReplaceAll(value, PlusEncoding, opt.SpaceEnc)
	rawQuery := ub.url.RawQuery
	if len(rawQuery) > 0 {
		rawQuery += ampersandStr
	}
	ub.url.RawQuery = rawQuery + key + "=" + value
	return nil
}

// DeleteFragment remove existing fragment if any
func (ub *Builder) DeleteFragment() {
	ub.url.Fragment = ""
}

// SetFragment create / update existing fragment for references, without '#'
func (ub *Builder) SetFragment(opt SetFragmentOpt) {
	value := url.QueryEscape(opt.Fragment)
	if opt.UseDefaultEncode {
		opt.SpaceEnc = ub.defaultSpaceEncode
	}
	value = strings.ReplaceAll(value, PercentTwentyEncoding, opt.SpaceEnc)
	ub.url.Fragment = value
}

// DeleteKeyQuery delete key query parameter if exist
func (ub *Builder) DeleteKeyQuery(keyDelete string) {
	keyVal := strings.Split(ub.url.RawQuery, ampersandStr)
	buildRawResult := make([]string, len(keyVal))
	wg := sync.WaitGroup{}
	for i, queryParam := range keyVal {
		wg.Add(1)
		go func(qp string, idx int) {
			defer wg.Done()
			qv := strings.Split(qp, "=")
			if len(qv) > 0 {
				q := qv[0]
				v := ""
				if len(qv) > 1 {
					v = qv[1]
				}
				qDel, err := url.QueryUnescape(q)
				if err != nil {
					qDel = q
				}
				if qDel != keyDelete {
					buildRawResult[idx] = q + "=" + v
				}
			}
		}(queryParam, i)
	}
	wg.Wait()
	for idx, v := range buildRawResult {
		if v == "" {
			buildRawResult = removeIndex(buildRawResult, idx)
		}
	}
	if len(buildRawResult) > 0 {
		ub.url.RawQuery = strings.Join(buildRawResult, ampersandStr)
	}
}
