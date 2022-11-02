[![godoc](https://godoc.org/github.com/golang/mock/gomock?status.svg)](http://godoc.tkpd/pkg/github.com/tokopedia/tdk/go/uruki)
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
- safe concurrency set / get parts of uri
- update base url with validation mechanism

## Uruki Option
Option while initiate builder

| Field | Type | Description |
|-|-|-|
| URL | string | URL existing to parse |
| RestrictScheme | []string | default no restrict scheme, for example if you want restrict scheme url only into http, https, and tokopedia. can use []string{"http", "https", "tokopedia"}|
| DefaultSpaceEncode | SpaceEncoding | space encoding method while escape query, refer to SpaceEncoding list below, default is keep space as is|
| UseEscapeAutomateInit | bool | automate escape existing query while init builder, default false|

SpaceEncoding method
- WithoutEncoding = keep space as is
- PercentTwentyEncoding = change space into %20
- PlusEncoding = change space into +

## Example initiate

```go
// without using options

ub, err := NewUrukiBuilder()
if err != nil {
    fmt.Println(err)
    return
}
got := ub.GetUrlResult()

// got = ""
```

```go
// using options without UseEscapeAutomateInit

ub, err := NewUrukiBuilder(UrukiOption{
    URL:                   "https://www.tokopedia.com/search?condition=1&fcity=174,175,176,177,178,179&navsource=&rf=true&rt=4,5&srp_component_id=02.01.00.00&srp_page_id=&srp_page_title=&st=product&q=macbook air m2",
    UseEscapeAutomateInit: false,
})
if err != nil {
    fmt.Println(err)
    return
}
got = ub.GetUrlResult()

// got = "https://www.tokopedia.com/search?condition=1&fcity=174,175,176,177,178,179&navsource=&rf=true&rt=4,5&srp_component_id=02.01.00.00&srp_page_id=&srp_page_title=&st=product&q=macbook air m2"
```

```go
// using options with UseEscapeAutomateInit and PercentTwentyEncoding

ub, err := NewUrukiBuilder(UrukiOption{
    URL:                   "https://www.tokopedia.com/search?condition=1&fcity=174,175,176,177,178,179&navsource=&rf=true&rt=4,5&srp_component_id=02.01.00.00&srp_page_id=&srp_page_title=&st=product&q=macbook air m2",
    UseEscapeAutomateInit: true,
    DefaultSpaceEncode:    PercentTwentyEncoding,
})
if err != nil {
    fmt.Println(err)
    return
}
got = ub.GetUrlResult()

// got = "https://www.tokopedia.com/search?condition=1&fcity=174%2C175%2C176%2C177%2C178%2C179&navsource=&rf=true&rt=4%2C5&srp_component_id=02.01.00.00&srp_page_id=&srp_page_title=&st=product&q=macbook%20air%20m2"
```

```go
// using options with UseEscapeAutomateInit, PercentTwentyEncoding and RestrictScheme
_, err = NewUrukiBuilder(UrukiOption{
    URL:                   "http://www.tokopedia.com/search?condition=1&fcity=174,175,176,177,178,179&navsource=&rf=true&rt=4,5&srp_component_id=02.01.00.00&srp_page_id=&srp_page_title=&st=product&q=macbook air m2",
    UseEscapeAutomateInit: true,
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
```go
// using ampersand value and some without encoding space
ub, err := NewUrukiBuilder(UrukiOption{
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
url := ub.GetUrlResult()
// url = "https://tokopedia.com/search?q=produk+p%26g"
```

### SetBaseUrl
```go
ub, err := NewUrukiBuilder(UrukiOption{
    URL:                "https://www.tokopedia.com/acmic/acmic-usb-c-to-lightning-adapter-iphone-converter-connector-konektor?extParam=ivf%3Dfalse%26src%3Dsearch%26whid%3D13355454",
    DefaultSpaceEncode: PercentTwentyEncoding,
    RestrictScheme:     []string{"https", "tokopedia"},
})
err = ub.SetBaseUrl("tokopedia://")
if err != nil {
    fmt.Println(err)
    return
}
gotUrl := ub.GetUrlResult()
// url = "tokopedia://acmic/acmic-usb-c-to-lightning-adapter-iphone-converter-connector-konektor?extParam=ivf%3Dfalse%26src%3Dsearch%26whid%3D13355454"
```

### SetPath
```go
ub, err := NewUrukiBuilder(UrukiOption{
    URL:                "https://www.tokopedia.com/acmic/acmic-usb-c-to-lightning-adapter-iphone-converter-connector-konektor?extParam=ivf%3Dfalse%26src%3Dsearch%26whid%3D13355454",
    DefaultSpaceEncode: PercentTwentyEncoding,
})
if err != nil {
    fmt.Println(err)
    return
}
ub.SetPath("/iphoneos/acmic-konektor")
gotUrl := ub.GetUrlResult()
// url = "https://www.tokopedia.com/iphoneos/acmic-konektor?extParam=ivf%3Dfalse%26src%3Dsearch%26whid%3D13355454"
```

### SetFragment
```go
ub, err := NewUrukiBuilder(UrukiOption{
    URL:                "https://tokopedia.com/discovery",
    DefaultSpaceEncode: PercentTwentyEncoding,
})
if err != nil {
    fmt.Println(err)
    return
}
ub.SetFragment(SetFragmentOpt{Fragment: "top 10", SpaceEnc: PlusEncoding})
gotUrl := ub.GetUrlResult()
// url = "https://tokopedia.com/discovery#top+10"
```

### DeleteKeyQuery
```go
ub, err := NewUrukiBuilder(UrukiOption{
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
url := ub.GetUrlResult()
// url = "https://www.tokopedia.com/search?q=produck%20p%26g&srp_component_id=01.07.00.00&srp_page_title=&=exist_val_empty_key"
```

### DeleteFragment
```go
ub, err := NewUrukiBuilder(UrukiOption{
		URL:                "https://tokopedia.com/discovery#top",
		DefaultSpaceEncode: PercentTwentyEncoding,
	})
if err != nil {
    fmt.Println(err)
    return
}
ub.DeleteFragment()
url := ub.GetUrlResult()
// url = "https://tokopedia.com/discovery"
```
