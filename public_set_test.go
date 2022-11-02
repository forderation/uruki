package uruki

import "testing"

func Test_AddQueryParam(t *testing.T) {
	type args struct {
		name         string
		opt          []AddQueryParamOpt
		wantRawQuery string
		wantErr      error
	}

	testCases := []args{
		{
			name: "empty key will be error",
			opt: []AddQueryParamOpt{
				{
					Key: "",
					Val: "",
				},
			},
			wantErr:      ErrorKeyEmpty,
			wantRawQuery: "https://tokopedia.com/search",
		},
		{
			name: "key contains space",
			opt: []AddQueryParamOpt{
				{
					Key: "k ey",
					Val: "",
				},
			},
			wantErr:      ErrorKeyContainSpace,
			wantRawQuery: "https://tokopedia.com/search",
		},
		{
			name: "space trailing key automatically",
			opt: []AddQueryParamOpt{
				{
					Key: " key ",
					Val: "value",
				},
			},
			wantRawQuery: "https://tokopedia.com/search?key=value",
		},
		{
			name: "space trailing key automatically",
			opt: []AddQueryParamOpt{
				{
					Key: " key ",
					Val: "value",
				},
			},
			wantRawQuery: "https://tokopedia.com/search?key=value",
		},
		{
			name: "using default space encode",
			opt: []AddQueryParamOpt{
				{
					Key: "key",
					Val: "space value",
				},
			},
			wantRawQuery: "https://tokopedia.com/search?key=space%20value",
		},
		{
			name: "using custom space encode and same key",
			opt: []AddQueryParamOpt{
				{
					Key:      "key",
					Val:      "space value",
					SpaceEnc: PercentTwentyEncoding,
				},
				{
					Key:      "plus_key",
					Val:      "space value",
					SpaceEnc: PlusEncoding,
				},
			},
			wantRawQuery: "https://tokopedia.com/search?key=space%20value&plus_key=space+value",
		},
		{
			name: "using ampersand value and some without encoding space",
			opt: []AddQueryParamOpt{
				{
					Key:      "q",
					Val:      "produk p&g",
					SpaceEnc: PlusEncoding,
				},
				{
					Key:      "navsource",
					Val:      "home campaign",
					SpaceEnc: WithoutEncoding,
				},
				{
					Key:      "srp_component_id",
					Val:      "02.01.00.00",
					SpaceEnc: PlusEncoding,
				},
			},
			wantRawQuery: "https://tokopedia.com/search?q=produk+p%26g&navsource=home campaign&srp_component_id=02.01.00.00",
		},
	}

	for _, tt := range testCases {
		t.Run(tt.name, func(t *testing.T) {
			ub, err := NewUrukiBuilder(UrukiOption{
				URL:                "https://tokopedia.com/search",
				DefaultSpaceEncode: PercentTwentyEncoding,
			})
			if err != nil {
				t.Error(err)
				return
			}
			for _, op := range tt.opt {
				err := ub.AddQueryParam(op)
				if err != tt.wantErr {
					t.Errorf("fail error test AddQueryParam() got %v want %v", err, tt.wantErr)
				}
			}
			url := ub.GetUrlResult()
			if url != tt.wantRawQuery {
				t.Errorf("fail value test AddQueryParam() got %v want %v", url, tt.wantRawQuery)
			}
		})
	}
}

func Test_DeleteKeyQuery(t *testing.T) {
	type args struct {
		name         string
		keyDelete    []string
		wantRawQuery string
		url          string
	}

	testCases := []args{
		{
			name:         "delete multi key",
			keyDelete:    []string{"st", "srp_page_id", "navsource"},
			url:          "https://www.tokopedia.com/search?st=product&q=produck%20p%26g&srp_component_id=01.07.00.00&srp_page_id=&srp_page_title=&navsource=&=exist_val_empty_key",
			wantRawQuery: "https://www.tokopedia.com/search?q=produck%20p%26g&srp_component_id=01.07.00.00&srp_page_title=&=exist_val_empty_key",
		},
		{
			name:         "delete non exist key",
			keyDelete:    []string{"refer"},
			url:          "https://www.tokopedia.com/search?st=product&q=produck%20p%26g&srp_component_id=01.07.00.00&srp_page_id=&srp_page_title=&navsource=",
			wantRawQuery: "https://www.tokopedia.com/search?st=product&q=produck%20p%26g&srp_component_id=01.07.00.00&srp_page_id=&srp_page_title=&navsource=",
		},
	}

	for _, tt := range testCases {
		t.Run(tt.name, func(t *testing.T) {
			ub, err := NewUrukiBuilder(UrukiOption{
				URL:                tt.url,
				DefaultSpaceEncode: PercentTwentyEncoding,
			})
			if err != nil {
				t.Error(err)
				return
			}
			for _, key := range tt.keyDelete {
				ub.DeleteKeyQuery(key)
			}
			url := ub.GetUrlResult()
			if url != tt.wantRawQuery {
				t.Errorf("fail value test DeleteKeyQuery() got %v want %v", url, tt.wantRawQuery)
			}
		})
	}
}

func Test_DeleteFragment(t *testing.T) {
	wantUrl := "https://tokopedia.com/discovery"
	ub, err := NewUrukiBuilder(UrukiOption{
		URL:                "https://tokopedia.com/discovery#top",
		DefaultSpaceEncode: PercentTwentyEncoding,
	})
	if err != nil {
		t.Error(err)
		return
	}
	ub.DeleteFragment()
	gotUrl := ub.GetUrlResult()
	if wantUrl != gotUrl {
		t.Errorf("fail test DeleteFragment() got %v want %v", gotUrl, wantUrl)
	}
}

func Test_SetFragment(t *testing.T) {
	wantUrl := "https://tokopedia.com/discovery#top+10"
	ub, err := NewUrukiBuilder(UrukiOption{
		URL:                "https://tokopedia.com/discovery",
		DefaultSpaceEncode: PercentTwentyEncoding,
	})
	if err != nil {
		t.Error(err)
		return
	}
	ub.SetFragment(SetFragmentOpt{Fragment: "top 10", SpaceEnc: PlusEncoding})
	gotUrl := ub.GetUrlResult()
	if wantUrl != gotUrl {
		t.Errorf("fail test SetFragment() got %v want %v", gotUrl, wantUrl)
	}
}

func Test_SetPath(t *testing.T) {
	wantUrl := "https://www.tokopedia.com/iphoneos/acmic-konektor?extParam=ivf%3Dfalse%26src%3Dsearch%26whid%3D13355454"
	ub, err := NewUrukiBuilder(UrukiOption{
		URL:                "https://www.tokopedia.com/acmic/acmic-usb-c-to-lightning-adapter-iphone-converter-connector-konektor?extParam=ivf%3Dfalse%26src%3Dsearch%26whid%3D13355454",
		DefaultSpaceEncode: PercentTwentyEncoding,
	})
	if err != nil {
		t.Error(err)
		return
	}
	ub.SetPath("/iphoneos/acmic-konektor")
	gotUrl := ub.GetUrlResult()
	if wantUrl != gotUrl {
		t.Errorf("fail test SetPath() got %v want %v", gotUrl, wantUrl)
	}
}

func Test_SetBaseUrl(t *testing.T) {
	wantUrl := "tokopedia://acmic/acmic-usb-c-to-lightning-adapter-iphone-converter-connector-konektor?extParam=ivf%3Dfalse%26src%3Dsearch%26whid%3D13355454"
	ub, err := NewUrukiBuilder(UrukiOption{
		URL:                "https://www.tokopedia.com/acmic/acmic-usb-c-to-lightning-adapter-iphone-converter-connector-konektor?extParam=ivf%3Dfalse%26src%3Dsearch%26whid%3D13355454",
		DefaultSpaceEncode: PercentTwentyEncoding,
		RestrictScheme:     []string{"https", "tokopedia"},
	})
	if err != nil {
		t.Error(err)
		return
	}
	err = ub.SetBaseUrl("tokopedia://")
	if err != nil {
		t.Error(err)
		return
	}
	gotUrl := ub.GetUrlResult()
	if wantUrl != gotUrl {
		t.Errorf("fail test SetBaseUrl() got %v want %v", gotUrl, wantUrl)
	}
}
