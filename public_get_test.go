package uruki

import (
	"reflect"
	"testing"
)

func Test_GetURLResultUnescape(t *testing.T) {
	ub, err := NewBuilder(Option{
		URL:                "https://www.tokopedia.com/search?st=product&q=beras%20p%26g&srp_component_id=01.07.00.00&srp_page_id=1012&srp_page_title=beras+putih&navsource=tokonow",
		DefaultSpaceEncode: PercentTwentyEncoding,
	})
	if err != nil {
		t.Error(err)
		return
	}
	wantURL := "https://www.tokopedia.com/search?st=product&q=beras p&g&srp_component_id=01.07.00.00&srp_page_id=1012&srp_page_title=beras putih&navsource=tokonow"
	gotURL := ub.GetURLResultUnescape()
	if wantURL != gotURL {
		t.Errorf("fail test GetURLResultUnescape() got %v want %v", gotURL, wantURL)
	}
}

func Test_GetURLResult(t *testing.T) {
	ub, err := NewBuilder(Option{
		URL:                "https://www.tokopedia.com/search?st=product&q=beras+putih&srp_component_id=01.07.00.00&srp_page_id=1012&srp_page_title=beras+putih&navsource=tokonow",
		DefaultSpaceEncode: PlusEncoding,
	})
	if err != nil {
		t.Error(err)
		return
	}
	wantURL := "https://www.tokopedia.com/search?st=product&q=beras+putih&srp_component_id=01.07.00.00&srp_page_id=1012&srp_page_title=beras+putih&navsource=tokonow"
	gotURL := ub.GetURLResult()
	if wantURL != gotURL {
		t.Errorf("fail test GetURLResult() got %v want %v", gotURL, wantURL)
	}
}

func Test_GetValueQuery(t *testing.T) {
	ub, err := NewBuilder(Option{
		URL:                "https://www.tokopedia.com/search?st=product&q=beras%20p%26g&srp_component_id=01.07.00.00&srp_page_id=1012&srp_page_title=beras+putih&navsource=tokonow",
		DefaultSpaceEncode: PercentTwentyEncoding,
	})
	if err != nil {
		t.Error(err)
		return
	}
	type args struct {
		name      string
		queryKey  string
		wantValue string
	}
	testCases := []args{
		{
			name:      "exist key",
			queryKey:  "q",
			wantValue: "beras p&g",
		},
		{
			name:      "non exist key",
			queryKey:  "ob",
			wantValue: "",
		},
	}
	for _, tt := range testCases {
		t.Run(tt.name, func(t *testing.T) {
			got := ub.GetValueQuery(tt.queryKey)
			if tt.wantValue != got {
				t.Errorf("fail test GetValueQuery() got %v want %v", got, tt.wantValue)
			}
		})
	}
}

func Test_GetAllQueryValue(t *testing.T) {
	ub, err := NewBuilder(Option{
		URL:                "https://www.tokopedia.com/search?st=product&q=beras%20p%26g&srp_component_id=01.07.00.00&srp_page_id=1012&srp_page_title=beras+putih&navsource=tokonow&q=beras%20putih&source=",
		DefaultSpaceEncode: PercentTwentyEncoding,
	})
	if err != nil {
		t.Error(err)
		return
	}
	wantResult := map[string][]string{
		"navsource":        {"tokonow"},
		"q":                {"beras p&g", "beras putih"},
		"srp_component_id": {"01.07.00.00"},
		"srp_page_id":      {"1012"},
		"srp_page_title":   {"beras putih"},
		"st":               {"product"},
		"source":           {""},
	}
	got := ub.GetAllQueryValue()
	if !reflect.DeepEqual(got, wantResult) {
		t.Errorf("fail test GetValueQuery() got %v want %v", got, wantResult)
	}
}

func Test_GetInternalURL(t *testing.T) {
	rawURL := "https://www.tokopedia.com/search?st=product&q=beras%20p%26g&srp_component_id=01.07.00.00&srp_page_id=1012@dhome&srp_page_title=beras+putih&navsource=tokonow&q=beras putih"
	ub, err := NewBuilder(Option{
		URL:                  rawURL,
		DefaultSpaceEncode:   PlusEncoding,
		UseEscapeAutomateURL: true,
	})
	if err != nil {
		t.Error(err)
		return
	}
	uriUb := ub.GetInternalURL()
	wantURL := "https://www.tokopedia.com/search?st=product&q=beras+p%26g&srp_component_id=01.07.00.00&srp_page_id=1012%40dhome&srp_page_title=beras+putih&navsource=tokonow&q=beras+putih"
	gotURL := uriUb.String()
	if gotURL != wantURL {
		t.Errorf("fail test GetInternalURL() got %v want %v", gotURL, wantURL)
	}
}

func Test_GetPaths(t *testing.T) {
	ub, err := NewBuilder(Option{
		URL:                "https://www.tokopedia.com/acmic/acmic-usb-c-to//lightning-adapter/iphone-converter/connector-konektor?extParam=ivf%3Dfalse%26src%3Dsearch%26whid%3D13355454",
		DefaultSpaceEncode: PercentTwentyEncoding,
	})
	if err != nil {
		t.Error(err)
		return
	}
	want := []string{"acmic", "acmic-usb-c-to", "", "lightning-adapter", "iphone-converter", "connector-konektor"}
	got := ub.GetPaths()
	if !reflect.DeepEqual(got, want) {
		t.Errorf("fail test GetPaths() got %v want %v", got, want)
	}
}

func Test_GetFullPath(t *testing.T) {
	ub, err := NewBuilder(Option{
		URL:                "https://www.tokopedia.com/acmic/acmic-usb-c-to//lightning-adapter/iphone-converter/connector-konektor?extParam=ivf%3Dfalse%26src%3Dsearch%26whid%3D13355454",
		DefaultSpaceEncode: PercentTwentyEncoding,
	})
	if err != nil {
		t.Error(err)
		return
	}
	want := "/acmic/acmic-usb-c-to//lightning-adapter/iphone-converter/connector-konektor"
	got := ub.GetFullPath()
	if got != want {
		t.Errorf("fail test GetFullPath() got %v want %v", got, want)
	}
}
