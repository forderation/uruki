package uruki

import "testing"

func Test_NewBuilder(t *testing.T) {
	// without using options
	ub, err := NewBuilder()
	if err != nil {
		t.Error(err)
		return
	}
	got := ub.GetURLResult()
	want := ""
	if got != want {
		t.Errorf("failed NewBuilder() without using options got: %v, want: %v", got, want)
	}

	// using options without UseEscapeAutomateURL
	ub, err = NewBuilder(Option{
		URL:                  "https://www.tokopedia.com/search?condition=1&fcity=174,175,176,177,178,179&navsource=&rf=true&rt=4,5&srp_component_id=02.01.00.00&srp_page_id=&srp_page_title=&st=product&q=macbook air m2",
		UseEscapeAutomateURL: false,
	})
	if err != nil {
		t.Error(err)
		return
	}
	got = ub.GetURLResult()
	want = "https://www.tokopedia.com/search?condition=1&fcity=174,175,176,177,178,179&navsource=&rf=true&rt=4,5&srp_component_id=02.01.00.00&srp_page_id=&srp_page_title=&st=product&q=macbook air m2"
	if got != want {
		t.Errorf("failed NewBuilder() using options without UseEscapeAutomateURL got: %v, want: %v", got, want)
	}

	// using options with UseEscapeAutomateURL and PercentTwentyEncoding
	ub, err = NewBuilder(Option{
		URL:                  "https://www.tokopedia.com/search?condition=1&fcity=174,175,176,177,178,179&navsource=&rf=true&rt=4,5&srp_component_id=02.01.00.00&srp_page_id=&srp_page_title=&st=product&q=macbook air m2",
		UseEscapeAutomateURL: true,
		DefaultSpaceEncode:   PercentTwentyEncoding,
	})
	if err != nil {
		t.Error(err)
		return
	}
	got = ub.GetURLResult()
	want = "https://www.tokopedia.com/search?condition=1&fcity=174%2C175%2C176%2C177%2C178%2C179&navsource=&rf=true&rt=4%2C5&srp_component_id=02.01.00.00&srp_page_id=&srp_page_title=&st=product&q=macbook%20air%20m2"
	if got != want {
		t.Errorf("failed NewBuilder() using options with UseEscapeAutomateURL and PercentTwentyEncoding got: %v, want: %v", got, want)
	}

	// using options with UseEscapeAutomateURL, PercentTwentyEncoding and RestrictScheme
	_, err = NewBuilder(Option{
		URL:                  "http://www.tokopedia.com/search?condition=1&fcity=174,175,176,177,178,179&navsource=&rf=true&rt=4,5&srp_component_id=02.01.00.00&srp_page_id=&srp_page_title=&st=product&q=macbook air m2",
		UseEscapeAutomateURL: true,
		DefaultSpaceEncode:   PercentTwentyEncoding,
		RestrictScheme:       []string{"https"},
	})
	// this will be error because we restrict scheme only https
	if err == nil {
		t.Error(err)
		return
	}
}
