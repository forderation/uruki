[![godoc](https://godoc.org/github.com/golang/mock/gomock?status.svg)](https://pkg.go.dev/github.com/forderation/uruki)
# URUKI

<span style="color:pink">URUKI</span> (URi qUicK buIlder) uri / url parser wrapped with net/url. adjusted to safe mutate data with options of restricted scheme and automatically internal encode data. ref: https://en.wikipedia.org/wiki/Uniform_Resource_Identifier

Background using this wrapper:
- quick update part of uri, because there is multi step if we using net/url to add / update / delete part of uri
- keep consistent all uri parts (scheme, host, port, path, query, fragment)
- get result of uri as string with automatically encoding
- restrict space encode

Feature supports:
- automatically encode by default setting
- delete key query without force encode values
- restricted scheme url, and encoding method
- multi add key query param
- set / get fragment
- update base url with validation mechanism

## Uruki Option
Option while initiate builder

| Field | Type | Description |
|-|-|-|
| URL | string | URL existing to parse |
| RestrictScheme | []string | default no restrict scheme, for example if you want restrict scheme url only into http, https, and tokopedia. can use []string{"http", "https", "tokopedia"}|
| DefaultSpaceEncode | SpaceEncoding | space encoding method while escape query, refer to SpaceEncoding list below, default is keep space as is|
| UseEscapeAutomateURL | bool | automate escape existing query while init builder / SetURL(uri string), default false|

SpaceEncoding method build in
- WithoutEncoding = keep space as is
- PercentTwentyEncoding = change space into %20
- PlusEncoding = change space into +

## Example Initiate

```go
// without using options

ub, err := NewBuilder()
if err != nil {
    fmt.Println(err)
    return
}
got := ub.GetURLResult()

// got = ""
```

```go
// using options without UseEscapeAutomateURL

ub, err := NewBuilder(Option{
    URL:                   "https://www.tokopedia.com/search?condition=1&fcity=174,175,176,177,178,179&navsource=&rf=true&rt=4,5&srp_component_id=02.01.00.00&srp_page_id=&srp_page_title=&st=product&q=macbook air m2",
    UseEscapeAutomateURL: false,
})
if err != nil {
    fmt.Println(err)
    return
}
got = ub.GetURLResult()

// got = "https://www.tokopedia.com/search?condition=1&fcity=174,175,176,177,178,179&navsource=&rf=true&rt=4,5&srp_component_id=02.01.00.00&srp_page_id=&srp_page_title=&st=product&q=macbook air m2"
```

```go
// using options with UseEscapeAutomateURL and PercentTwentyEncoding

ub, err := NewBuilder(Option{
    URL:                   "https://www.tokopedia.com/search?condition=1&fcity=174,175,176,177,178,179&navsource=&rf=true&rt=4,5&srp_component_id=02.01.00.00&srp_page_id=&srp_page_title=&st=product&q=macbook air m2",
    UseEscapeAutomateURL: true,
    DefaultSpaceEncode:    PercentTwentyEncoding,
})
if err != nil {
    fmt.Println(err)
    return
}
got = ub.GetURLResult()

// got = "https://www.tokopedia.com/search?condition=1&fcity=174%2C175%2C176%2C177%2C178%2C179&navsource=&rf=true&rt=4%2C5&srp_component_id=02.01.00.00&srp_page_id=&srp_page_title=&st=product&q=macbook%20air%20m2"
```

```go
// using options with UseEscapeAutomateURL, PercentTwentyEncoding and RestrictScheme
_, err = NewBuilder(Option{
    URL:                   "http://www.tokopedia.com/search?condition=1&fcity=174,175,176,177,178,179&navsource=&rf=true&rt=4,5&srp_component_id=02.01.00.00&srp_page_id=&srp_page_title=&st=product&q=macbook air m2",
    UseEscapeAutomateURL: true,
    DefaultSpaceEncode:    PercentTwentyEncoding,
    RestrictScheme:        []string{"https"},
})
// this will be error because we restrict scheme only https
if err != nil {
    fmt.Println(err)
    return
}
```

## Featured List

### AddQueryParam
add query parameter values, internally will be encode the value of query, options param:
- Key string = key query
- Val string = value query
- SpaceEnc string = specify space encode if you use custom encoding
- UseDefaultEncode bool = if you want using DefaultSpaceEncoding as SpaceEnc
```go
// using ampersand value and some without encoding space
ub, err := NewBuilder(Option{
    URL:                "https://tokopedia.com/search",
    DefaultSpaceEncode: PercentTwentyEncoding,
})
if err != nil {
    fmt.Println(err)
    return
}
err := ub.AddQueryParam(AddQueryParamOpt{Key:"q",Val:"produk p&g",SpaceEnc:PlusEncoding})
if err != nil {
    fmt.Println(err)
    return
}
url := ub.GetURLResult()
// url = "https://tokopedia.com/search?q=produk+p%26g"
```

### SetBaseURL
change or update existing of base url only host and port
```go
ub, err := NewBuilder(Option{
    URL:                "https://www.tokopedia.com/acmic/acmic-usb-c-to-lightning-adapter-iphone-converter-connector-konektor?extParam=ivf%3Dfalse%26src%3Dsearch%26whid%3D13355454",
    DefaultSpaceEncode: PercentTwentyEncoding,
    RestrictScheme:     []string{"https", "tokopedia"},
})
err = ub.SetBaseURL("tokopedia://")
if err != nil {
    fmt.Println(err)
    return
}
gotURL := ub.GetURLResult()
// url = "tokopedia://acmic/acmic-usb-c-to-lightning-adapter-iphone-converter-connector-konektor?extParam=ivf%3Dfalse%26src%3Dsearch%26whid%3D13355454"
```

### SetPath
change or update path only of url
```go
ub, err := NewBuilder(Option{
    URL:                "https://www.tokopedia.com/acmic/acmic-usb-c-to-lightning-adapter-iphone-converter-connector-konektor?extParam=ivf%3Dfalse%26src%3Dsearch%26whid%3D13355454",
    DefaultSpaceEncode: PercentTwentyEncoding,
})
if err != nil {
    fmt.Println(err)
    return
}
ub.SetPath("/iphoneos/acmic-konektor")
gotURL := ub.GetURLResult()
// url = "https://www.tokopedia.com/iphoneos/acmic-konektor?extParam=ivf%3Dfalse%26src%3Dsearch%26whid%3D13355454"
```

### SetFragment
create / update existing fragment for references, without '#'. options:
- Fragment string = fragment value
- SpaceEnc string = specify space encode if you use custom encoding
- UseDefaultEncode bool = if you want using DefaultSpaceEncoding as SpaceEnc
```go
ub, err := NewBuilder(Option{
    URL:                "https://tokopedia.com/discovery",
    DefaultSpaceEncode: PercentTwentyEncoding,
})
if err != nil {
    fmt.Println(err)
    return
}
ub.SetFragment(SetFragmentOpt{Fragment: "top 10", SpaceEnc: PlusEncoding})
gotURL := ub.GetURLResult()
// url = "https://tokopedia.com/discovery#top+10"
```

### DeleteKeyQuery
delete key query parameter if exist
```go
ub, err := NewBuilder(Option{
    URL:                "https://www.tokopedia.com/search?st=product&q=produck%20p%26g&srp_component_id=01.07.00.00&srp_page_id=&srp_page_title=&navsource=&=exist_val_empty_key",
    DefaultSpaceEncode: PercentTwentyEncoding,
})
if err != nil {
    fmt.Println(err)
    return
}
ub.DeleteKeyQuery("st")
ub.DeleteKeyQuery("srp_page_id")
ub.DeleteKeyQuery("navsource")
url := ub.GetURLResult()
// url = "https://www.tokopedia.com/search?q=produck%20p%26g&srp_component_id=01.07.00.00&srp_page_title=&=exist_val_empty_key"
```

### SetURL
replace all url with new url based on parameter, if error keep old url
```go
ub, err := NewBuilder(Option{
    URL:                "https://www.tokopedia.com/acmic/acmic-usb-c-to-lightning-adapter-iphone-converter-connector-konektor?extParam=ivf%3Dfalse%26src%3Dsearch%26whid%3D13355454",
    DefaultSpaceEncode: PercentTwentyEncoding,
    RestrictScheme:     []string{"https", "tokopedia"},
})
if err != nil {
    fmt.Println(err)
    return
}
err = ub.SetURL("tokopedia://search?q=p+g&source=search#bottom")
if err != nil {
    fmt.Println(err)
    return
}
url := ub.GetURLResult()
// url = "tokopedia://search?q=p+g&source=search#bottom"

// expect error because scheme changes into http
err = ub.SetURL("http://www.tokopedia.com/now")
if err != nil {
    fmt.Println(err)
}
```

### DeleteFragment
remove existing fragment if any
```go
ub, err := NewBuilder(Option{
		URL:                "https://tokopedia.com/discovery#top",
		DefaultSpaceEncode: PercentTwentyEncoding,
	})
if err != nil {
    fmt.Println(err)
    return
}
ub.DeleteFragment()
url := ub.GetURLResult()
// url = "https://tokopedia.com/discovery"
```