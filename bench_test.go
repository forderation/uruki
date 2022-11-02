package uruki

import (
	"testing"
)

func BenchmarkPublicGet(b *testing.B) {
	var ub *UrukiBuilder
	var err error
	b.Run("NewUrukiBuilder()", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			ub, err = NewUrukiBuilder(UrukiOption{
				URL:            "https://www.tokopedia.com/search?navsource=campaign&q=sepatu&srp_component_id=02.04.01.02&srp_disco_url=sportacular-shopathon&srp_page_id=44691&srp_page_title=Sportacular+Shopathon&st=product&srp_ext_ref=campaigncode.thco0009684#gid=123",
				RestrictScheme: []string{"http", "tokopedia", "https"},
			})
			if err != nil {
				b.Error(err)
				return
			}
		}
	})
	b.Run("GetUrlResultUnescape()", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			ub.GetUrlResultUnescape()
		}
	})
	b.Run("GetUrlResult()", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			ub.GetUrlResult()
		}
	})
	b.Run("GetAllQueryValue()", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			ub.GetAllQueryValue()
		}
	})
	b.Run("GetInternalUrl()", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			ub.GetInternalUrl()
		}
	})
	b.Run("GetFullPath()", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			ub.GetFullPath()
		}
	})
	b.Run("GetPaths()", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			ub.GetPaths()
		}
	})
	b.Run("GetValueQuery()", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			queries := []string{"q", "st", "srp_ext_ref", "navsource", "srp_disco_url"}
			for _, v := range queries {
				ub.GetValueQuery(v)
			}
		}
	})
}

func BenchmarkPublicSet(b *testing.B) {
	var ub *UrukiBuilder
	var err error
	b.Run("NewUrukiBuilder()", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			ub, err = NewUrukiBuilder(UrukiOption{
				URL:            "https://www.tokopedia.com/search?navsource=campaign&q=sepatu&srp_component_id=02.04.01.02&srp_disco_url=sportacular-shopathon&srp_page_id=44691&srp_page_title=Sportacular+Shopathon&st=product&srp_ext_ref=campaigncode.thco0009684#gid=123",
				RestrictScheme: []string{"http", "tokopedia", "https"},
			})
			if err != nil {
				b.Error(err)
				return
			}
		}
	})
	b.Run("SetBaseUrl()", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			ub.SetBaseUrl("tokopedia://")
		}
	})
	b.Run("SetPath()", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			ub.SetPath("/search/v4/handle")
		}
	})
	b.Run("DeleteFragment()", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			ub.DeleteFragment()
		}
	})
	b.Run("AddQueryParam()", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			ub.AddQueryParam(AddQueryParamOpt{Key: "source", Val: "search feed ads", SpaceEnc: PlusEncoding})
			ub.AddQueryParam(AddQueryParamOpt{Key: "source_line", Val: "search feed ads", SpaceEnc: PercentTwentyEncoding})
			ub.AddQueryParam(AddQueryParamOpt{Key: "src", Val: "search feed ads"})
		}
	})
	b.Run("SetFragment()", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			ub.SetFragment(SetFragmentOpt{Fragment: "gid=123 456 789", SpaceEnc: PercentTwentyEncoding})
			ub.SetFragment(SetFragmentOpt{Fragment: "gid=123 456 789", SpaceEnc: PlusEncoding})
			ub.SetFragment(SetFragmentOpt{Fragment: "gid=123 456 789"})
		}
	})
	b.Run("DeleteKeyQuery()", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			ub.DeleteKeyQuery("navsource")
			ub.DeleteKeyQuery("")
			ub.DeleteKeyQuery("srp_ext_reff")
			ub.DeleteKeyQuery("q")
		}
	})
}
