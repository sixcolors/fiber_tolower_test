package tolower

import (
	"testing"
)

func BenchmarkUtilsToLowerLowerCase(b *testing.B) {
	origins := []string{
		"https://example.com",
		"https://www.example.org",
		"https://sub.example.net",
		"http://chart.apis.google.com/chart?chs=500x500&chma=0,0,100,100&cht=p&chco=ff0000%2cffff00%7cff8000%2c00ff00%7c00ff00%2c0000ff&chd=t%3a122%2c42%2c17%2c10%2c8%2c7%2c7%2c7%2c7%2c6%2c6%2c6%2c6%2c5%2c5&chl=122%7c42%7c17%7c10%7c8%7c7%7c7%7c7%7c7%7c6%7c6%7c6%7c6%7c5%7c5&chdl=android%7cjava%7cstack-trace%7cbroadcastreceiver%7candroid-ndk%7cuser-agent%7candroid-webview%7cwebview%7cbackground%7cmultithreading%7candroid-source%7csms%7cadb%7csollections%7cactivity|chart", // not Origin header, but a more complex string, possiable Referer header
		// Add more realistic origins or referers as needed
		"lorem ipsum dolor sit amet, consectetur adipiscing elit. donec eu gravida purus, at interdum nulla. class aptent taciti sociosqu ad litora torquent per conubia nostra, per inceptos himenaeos. morbi tincidunt sapien ac elit convallis elementum.", // Lorem Ipsum scentence
	}

	b.ResetTimer()
	b.ReportAllocs()

	for i := 0; i < b.N; i++ {
		for _, origin := range origins {
			_ = ToLowerUtils(origin)
		}
	}
}

func BenchmarkUtilsToLowerMixedCase(b *testing.B) {
	origins := []string{
		"https://Example.com",
		"https://www.Example.org",
		"https://Sub.example.net",
		"http://chart.apis.google.com/chart?chs=500x500&chma=0,0,100,100&cht=p&chco=FF0000%2CFFFF00%7CFF8000%2C00FF00%7C00FF00%2C0000FF&chd=t%3A122%2C42%2C17%2C10%2C8%2C7%2C7%2C7%2C7%2C6%2C6%2C6%2C6%2C5%2C5&chl=122%7C42%7C17%7C10%7C8%7C7%7C7%7C7%7C7%7C6%7C6%7C6%7C6%7C5%7C5&chdl=android%7Cjava%7Cstack-trace%7Cbroadcastreceiver%7Candroid-ndk%7Cuser-agent%7Candroid-webview%7Cwebview%7Cbackground%7Cmultithreading%7Candroid-source%7Csms%7Cadb%7Csollections%7Cactivity|Chart",
		// Add more realistic origins or referers as needed
		"Lorem ipsum dolor sit amet, consectetur adipiscing elit. Donec eu gravida purus, at interdum nulla. Class aptent taciti sociosqu ad litora torquent per conubia nostra, per inceptos himenaeos. Morbi tincidunt sapien ac elit convallis elementum.",
	}

	b.ResetTimer()
	b.ReportAllocs()

	for i := 0; i < b.N; i++ {
		for _, origin := range origins {
			_ = ToLowerUtils(origin)
		}
	}
}

func BenchmarkStringsToLowerLowerCase(b *testing.B) {
	origins := []string{
		"https://example.com",
		"https://www.example.org",
		"https://sub.example.net",
		"http://chart.apis.google.com/chart?chs=500x500&chma=0,0,100,100&cht=p&chco=ff0000%2cffff00%7cff8000%2c00ff00%7c00ff00%2c0000ff&chd=t%3a122%2c42%2c17%2c10%2c8%2c7%2c7%2c7%2c7%2c6%2c6%2c6%2c6%2c5%2c5&chl=122%7c42%7c17%7c10%7c8%7c7%7c7%7c7%7c7%7c6%7c6%7c6%7c6%7c5%7c5&chdl=android%7cjava%7cstack-trace%7cbroadcastreceiver%7candroid-ndk%7cuser-agent%7candroid-webview%7cwebview%7cbackground%7cmultithreading%7candroid-source%7csms%7cadb%7csollections%7cactivity|chart",
		// Add more realistic origins or referers as needed
		"lorem ipsum dolor sit amet, consectetur adipiscing elit. donec eu gravida purus, at interdum nulla. class aptent taciti sociosqu ad litora torquent per conubia nostra, per inceptos himenaeos. morbi tincidunt sapien ac elit convallis elementum.",
	}

	b.ResetTimer()
	b.ReportAllocs()

	for i := 0; i < b.N; i++ {
		for _, origin := range origins {
			_ = ToLowerStrings(origin)
		}
	}
}

func BenchmarkStringsToLowerMixedCase(b *testing.B) {
	origins := []string{
		"https://Example.com",
		"https://www.Example.org",
		"https://Sub.example.net",
		"http://chart.apis.google.com/chart?chs=500x500&chma=0,0,100,100&cht=p&chco=FF0000%2CFFFF00%7CFF8000%2C00FF00%7C00FF00%2C0000FF&chd=t%3A122%2C42%2C17%2C10%2C8%2C7%2C7%2C7%2C7%2C6%2C6%2C6%2C6%2C5%2C5&chl=122%7C42%7C17%7C10%7C8%7C7%7C7%7C7%7C7%7C6%7C6%7C6%7C6%7C5%7C5&chdl=android%7Cjava%7Cstack-trace%7Cbroadcastreceiver%7Candroid-ndk%7Cuser-agent%7Candroid-webview%7Cwebview%7Cbackground%7Cmultithreading%7Candroid-source%7Csms%7Cadb%7Csollections%7Cactivity|Chart",
		// Add more realistic origins or referers as needed
		"Lorem ipsum dolor sit amet, consectetur adipiscing elit. Donec eu gravida purus, at interdum nulla. Class aptent taciti sociosqu ad litora torquent per conubia nostra, per inceptos himenaeos. Morbi tincidunt sapien ac elit convallis elementum.",
	}

	b.ResetTimer()
	b.ReportAllocs()

	for i := 0; i < b.N; i++ {
		for _, origin := range origins {
			_ = ToLowerStrings(origin)
		}
	}
}

func BenchmarkHybribToLowerLowerCase(b *testing.B) {
	origins := []string{
		"https://example.com",
		"https://www.example.org",
		"https://sub.example.net",
		"http://chart.apis.google.com/chart?chs=500x500&chma=0,0,100,100&cht=p&chco=ff0000%2cffff00%7cff8000%2c00ff00%7c00ff00%2c0000ff&chd=t%3a122%2c42%2c17%2c10%2c8%2c7%2c7%2c7%2c7%2c6%2c6%2c6%2c6%2c5%2c5&chl=122%7c42%7c17%7c10%7c8%7c7%7c7%7c7%7c7%7c6%7c6%7c6%7c6%7c5%7c5&chdl=android%7cjava%7cstack-trace%7cbroadcastreceiver%7candroid-ndk%7cuser-agent%7candroid-webview%7cwebview%7cbackground%7cmultithreading%7candroid-source%7csms%7cadb%7csollections%7cactivity|chart",
		// Add more realistic origins or referers as needed
		"lorem ipsum dolor sit amet, consectetur adipiscing elit. donec eu gravida purus, at interdum nulla. class aptent taciti sociosqu ad litora torquent per conubia nostra, per inceptos himenaeos. morbi tincidunt sapien ac elit convallis elementum.",
	}

	b.ResetTimer()
	b.ReportAllocs()

	for i := 0; i < b.N; i++ {
		for _, origin := range origins {
			_ = HybribToLower(origin)
		}
	}
}

func BenchmarkHybribToLowerMixedCase(b *testing.B) {
	origins := []string{
		"https://Example.com",
		"https://www.Example.org",
		"https://Sub.example.net",
		"http://chart.apis.google.com/chart?chs=500x500&chma=0,0,100,100&cht=p&chco=FF0000%2CFFFF00%7CFF8000%2C00FF00%7C00FF00%2C0000FF&chd=t%3A122%2C42%2C17%2C10%2C8%2C7%2C7%2C7%2C7%2C6%2C6%2C6%2C6%2C5%2C5&chl=122%7C42%7C17%7C10%7C8%7C7%7C7%7C7%7C7%7C6%7C6%7C6%7C6%7C5%7C5&chdl=android%7Cjava%7Cstack-trace%7Cbroadcastreceiver%7Candroid-ndk%7Cuser-agent%7Candroid-webview%7Cwebview%7Cbackground%7Cmultithreading%7Candroid-source%7Csms%7Cadb%7Csollections%7Cactivity|Chart",
		// Add more realistic origins or referers as needed
		"Lorem ipsum dolor sit amet, consectetur adipiscing elit. Donec eu gravida purus, at interdum nulla. Class aptent taciti sociosqu ad litora torquent per conubia nostra, per inceptos himenaeos. Morbi tincidunt sapien ac elit convallis elementum.",
	}

	b.ResetTimer()
	b.ReportAllocs()

	for i := 0; i < b.N; i++ {
		for _, origin := range origins {
			_ = HybribToLower(origin)
		}
	}
}
