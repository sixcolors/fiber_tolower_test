package tolower

import (
	"bytes"
	"fmt"
	"strings"
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

func BenchmarkHybridToLowerLowerCase(b *testing.B) {
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
			_ = HybridToLower(origin)
		}
	}
}

func TestToLowerFunctions(t *testing.T) {
	cases := []struct {
		in, want string
	}{
		{"", ""},
		{"a", "a"},
		{"A", "a"},
		{"abc", "abc"},
		{"ABC", "abc"},
		{"AbC123", "abc123"},
		{"!@#ABCxyz", "!@#abcxyz"},
		{"https://Example.COM/Path", "https://example.com/path"},
		{"1234567890", "1234567890"},
		{"Lorem IPSUM", "lorem ipsum"},
	}

	funcs := []struct {
		name string
		fn   func(string) string
	}{
		{"ToLowerUtils", ToLowerUtils},
		{"ToLowerStrings", ToLowerStrings},
		{"HybridToLower", HybridToLower},
		{"ToLower", ToLower},
		{"ToLowerSWARString", ToLowerSWARString},
		{"ToLowerSWARv2", ToLowerSWARv2String},
		{"ToLowerUnsafeString", ToLowerUnsafeString},
		{"ToLowerGaby", ToLowerGabyString},
		{"OptimalToLower", OptimalToLower},
	}

	for _, f := range funcs {
		for _, c := range cases {
			got := f.fn(c.in)
			if got != c.want {
				t.Errorf("%s(%q) = %q; want %q", f.name, c.in, got, c.want)
			}
		}
	}
}
func BenchmarkToLowerSWARMixedCase(b *testing.B) {
	origins := []string{
		"https://Example.com",
		"https://www.Example.org",
		"https://Sub.example.net",
		"http://chart.apis.google.com/chart?chs=500x500&chma=0,0,100,100&cht=p&chco=FF0000%2CFFFF00%7CFF8000%2C00FF00%7C00FF00%2C0000FF&chd=t%3A122%2C42%2C17%2C10%2C8%2C7%2C7%2C7%2C7%2C6%2C6%2C6%2C6%2C5%2C5&chl=122%7C42%7C17%7C10%7C8%7C7%7C7%7C7%7C7%7C6%7C6%7C6%7C6%7C5%7C5&chdl=android%7Cjava%7Cstack-trace%7Cbroadcastreceiver%7Candroid-ndk%7Cuser-agent%7Candroid-webview%7Cwebview%7Cbackground%7Cmultithreading%7Candroid-source%7Csms%7Cadb%7Csollections%7Cactivity|Chart",
		"Lorem ipsum dolor sit amet, consectetur adipiscing elit. Donec eu gravida purus, at interdum nulla. Class aptent taciti sociosqu ad litora torquent per conubia nostra, per inceptos himenaeos. Morbi tincidunt sapien ac elit convallis elementum.",
	}

	b.ResetTimer()
	b.ReportAllocs()

	for i := 0; i < b.N; i++ {
		for _, origin := range origins {
			_ = ToLowerSWAR([]byte(origin))
		}
	}
}

func BenchmarkToLowerGabyMixedCase(b *testing.B) {
	origins := []string{
		"https://Example.com",
		"https://www.Example.org",
		"https://Sub.example.net",
		"http://chart.apis.google.com/chart?chs=500x500&chma=0,0,100,100&cht=p&chco=FF0000%2CFFFF00%7CFF8000%2C00FF00%7C00FF00%2C0000FF&chd=t%3A122%2C42%2C17%2C10%2C8%2C7%2C7%2C7%2C7%2C6%2C6%2C6%2C6%2C5%2C5&chl=122%7C42%7C17%7C10%7C8%7C7%7C7%7C7%7C7%7C6%7C6%7C6%7C6%7C5%7C5&chdl=android%7Cjava%7Cstack-trace%7Cbroadcastreceiver%7Candroid-ndk%7Cuser-agent%7Candroid-webview%7Cwebview%7Cbackground%7Cmultithreading%7Candroid-source%7Csms%7Cadb%7Csollections%7Cactivity|Chart",
		"Lorem ipsum dolor sit amet, consectetur adipiscing elit. Donec eu gravida purus, at interdum nulla. Class aptent taciti sociosqu ad litora torquent per conubia nostra, per inceptos himenaeos. Morbi tincidunt sapien ac elit convallis elementum.",
	}

	b.ResetTimer()
	b.ReportAllocs()

	for i := 0; i < b.N; i++ {
		for _, origin := range origins {
			_ = ToLowerGaby([]byte(origin))
		}
	}
}

func BenchmarkToLowerUnsafeMixedCase(b *testing.B) {
	origins := []string{
		"https://Example.com",
		"https://www.Example.org",
		"https://Sub.example.net",
		"http://chart.apis.google.com/chart?chs=500x500&chma=0,0,100,100&cht=p&chco=FF0000%2CFFFF00%7CFF8000%2C00FF00%7C00FF00%2C0000FF&chd=t%3A122%2C42%2C17%2C10%2C8%2C7%2C7%2C7%2C7%2C6%2C6%2C6%2C6%2C5%2C5&chl=122%7C42%7C17%7C10%7C8%7C7%7C7%7C7%7C7%7C6%7C6%7C6%7C6%7C5%7C5&chdl=android%7Cjava%7Cstack-trace%7Cbroadcastreceiver%7Candroid-ndk%7Cuser-agent%7Candroid-webview%7Cwebview%7Cbackground%7Cmultithreading%7Candroid-source%7Csms%7Cadb%7Csollections%7Cactivity|Chart",
		"Lorem ipsum dolor sit amet, consectetur adipiscing elit. Donec eu gravida purus, at interdum nulla. Class aptent taciti sociosqu ad litora torquent per conubia nostra, per inceptos himenaeos. Morbi tincidunt sapien ac elit convallis elementum.",
	}

	b.ResetTimer()
	b.ReportAllocs()

	for i := 0; i < b.N; i++ {
		for _, origin := range origins {
			_ = ToLowerUnsafe([]byte(origin))
		}
	}
}

func TestToLowerSWAR(t *testing.T) {
	cases := []struct {
		in, want string
	}{
		{"", ""},
		{"a", "a"},
		{"A", "a"},
		{"abc", "abc"},
		{"ABC", "abc"},
		{"AbC123", "abc123"},
		{"!@#ABCxyz", "!@#abcxyz"},
		{"https://Example.COM/Path", "https://example.com/path"},
		{"1234567890", "1234567890"},
		{"Lorem IPSUM", "lorem ipsum"},
	}

	for _, c := range cases {
		got := string(ToLowerSWAR([]byte(c.in)))
		if got != c.want {
			t.Errorf("ToLowerSWAR(%q) = %q; want %q", c.in, got, c.want)
		}
	}
}

func TestToLowerUnsafe(t *testing.T) {
	cases := []struct {
		in, want string
	}{
		{"", ""},
		{"a", "a"},
		{"A", "a"},
		{"abc", "abc"},
		{"ABC", "abc"},
		{"AbC123", "abc123"},
		{"!@#ABCxyz", "!@#abcxyz"},
		{"https://Example.COM/Path", "https://example.com/path"},
		{"1234567890", "1234567890"},
		{"Lorem IPSUM", "lorem ipsum"},
	}

	for _, c := range cases {
		got := string(ToLowerUnsafe([]byte(c.in)))
		if got != c.want {
			t.Errorf("ToLowerUnsafe(%q) = %q; want %q", c.in, got, c.want)
		}
	}
}

// Test SWAR functions with various string lengths
func TestToLowerSWARLengths(t *testing.T) {
	// Test with strings of various lengths to check boundary conditions
	for i := 0; i < 20; i++ {
		s := strings.Repeat("A", i)
		want := strings.Repeat("a", i)
		got := string(ToLowerSWAR([]byte(s)))
		if got != want {
			t.Errorf("ToLowerSWAR(repeat 'A' %d times) = %q; want %q", i, got, want)
		}
	}
}

// Test Unsafe functions with various string lengths
func TestToLowerUnsafeLengths(t *testing.T) {
	// Test with strings of various lengths to check boundary conditions
	for i := 0; i < 20; i++ {
		s := strings.Repeat("A", i)
		want := strings.Repeat("a", i)
		got := string(ToLowerUnsafe([]byte(s)))
		if got != want {
			t.Errorf("ToLowerUnsafe(repeat 'A' %d times) = %q; want %q", i, got, want)
		}
	}
}

// Test with mixed ASCII and non-ASCII characters
func TestToLowerMixedChars(t *testing.T) {
	cases := []struct {
		in, want string
	}{
		{"HELLO world", "hello world"},
		{"123ABC", "123abc"},
		{"!@#$%^&*()", "!@#$%^&*()"},
		{"A-Z a-z", "a-z a-z"},
		{"AbCdEfGhIjKlMnOpQrStUvWxYz", "abcdefghijklmnopqrstuvwxyz"},
	}

	funcs := []struct {
		name string
		fn   func([]byte) []byte
	}{
		{"ToLowerSWAR", ToLowerSWAR},
		{"ToLowerUnsafe", ToLowerUnsafe},
	}

	for _, f := range funcs {
		for _, c := range cases {
			got := string(f.fn([]byte(c.in)))
			if got != c.want {
				t.Errorf("%s(%q) = %q; want %q", f.name, c.in, got, c.want)
			}
		}
	}
}

// Test for specific edge cases that might cause issues
func TestToLowerEdgeCases(t *testing.T) {
	cases := []struct {
		in, want string
	}{
		{"\x00\x01\x02\x03\x04\x05", "\x00\x01\x02\x03\x04\x05"},                         // Control characters
		{string([]byte{0x41, 0x42, 0x00, 0x43}), string([]byte{0x61, 0x62, 0x00, 0x63})}, // With null bytes
		{string([]byte{0xFF, 0xFE, 0xFD}), string([]byte{0xFF, 0xFE, 0xFD})},             // Non-ASCII bytes
	}

	funcs := []struct {
		name string
		fn   func([]byte) []byte
	}{
		{"ToLowerSWAR", ToLowerSWAR},
		{"ToLowerUnsafe", ToLowerUnsafe},
	}

	for _, f := range funcs {
		for _, c := range cases {
			got := string(f.fn([]byte(c.in)))
			if got != c.want {
				t.Errorf("%s(%q) = %q; want %q", f.name, c.in, got, c.want)
			}
		}
	}
}
func BenchmarkHybridToLowerMixedCase(b *testing.B) {
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
			_ = HybridToLower(origin)
		}
	}
}

func BenchmarkToLowerMixedCase(b *testing.B) {
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
			_ = ToLower(origin)
		}
	}
}

func BenchmarkToLowerSWARStringMixedCase(b *testing.B) {
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
			_ = ToLowerSWARString(origin)
		}
	}
}

func BenchmarkToLowerUnsafeStringMixedCase(b *testing.B) {
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
			_ = ToLowerUnsafeString(origin)
		}
	}
}

func BenchmarkToLowerLowerCase(b *testing.B) {
	origins := []string{
		"https://example.com",
		"https://www.example.org",
		"https://sub.example.net",
		"http://chart.apis.google.com/chart?chs=500x500&chma=0,0,100,100&cht=p&chco=ff0000%2cffff00%7cff8000%2c00ff00%7c00ff00%2c0000ff&chd=t%3a122%2c42%2c17%2c10%2c8%2c7%2c7%2c7%2c7%2c6%2c6%2c6%2c6%2c5%2c5&chl=122%7c42%7c17%7c10%7c8%7c7%7c7%7c7%7c7%7c6%7c6%7c6%7c6%7c5%7c5&chdl=android%7cjava%7cstack-trace%7cbroadcastreceiver%7candroid-ndk%7cuser-agent%7candroid-webview%7cwebview%7cbackground%7cmultithreading%7candroid-source%7csms%7cadb%7csollections%7cactivity|chart",
		"lorem ipsum dolor sit amet, consectetur adipiscing elit. donec eu gravida purus, at interdum nulla. class aptent taciti sociosqu ad litora torquent per conubia nostra, per inceptos himenaeos. morbi tincidunt sapien ac elit convallis elementum.",
	}
	b.ResetTimer()
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		for _, origin := range origins {
			_ = ToLower(origin)
		}
	}
}

func BenchmarkToLowerSWARLowerCase(b *testing.B) {
	origins := []string{
		"https://example.com",
		"https://www.example.org",
		"https://sub.example.net",
		"http://chart.apis.google.com/chart?chs=500x500&chma=0,0,100,100&cht=p&chco=ff0000%2cffff00%7cff8000%2c00ff00%7c00ff00%2c0000ff&chd=t%3a122%2c42%2c17%2c10%2c8%2c7%2c7%2c7%2c7%2c6%2c6%2c6%2c6%2c5%2c5&chl=122%7c42%7c17%7c10%7c8%7c7%7c7%7c7%7c7%7c6%7c6%7c6%7c6%7c5%7c5&chdl=android%7cjava%7cstack-trace%7cbroadcastreceiver%7candroid-ndk%7cuser-agent%7candroid-webview%7cwebview%7cbackground%7cmultithreading%7candroid-source%7csms%7cadb%7csollections%7cactivity|chart",
		"lorem ipsum dolor sit amet, consectetur adipiscing elit. donec eu gravida purus, at interdum nulla. class aptent taciti sociosqu ad litora torquent per conubia nostra, per inceptos himenaeos. morbi tincidunt sapien ac elit convallis elementum.",
	}
	b.ResetTimer()
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		for _, origin := range origins {
			_ = ToLowerSWAR([]byte(origin))
		}
	}
}

func BenchmarkToLowerGabyLowerCase(b *testing.B) {
	origins := []string{
		"https://example.com",
		"https://www.example.org",
		"https://sub.example.net",
		"http://chart.apis.google.com/chart?chs=500x500&chma=0,0,100,100&cht=p&chco=ff0000%2cffff00%7cff8000%2c00ff00%7c00ff00%2c0000ff&chd=t%3a122%2c42%2c17%2c10%2c8%2c7%2c7%2c7%2c7%2c6%2c6%2c6%2c6%2c5%2c5&chl=122%7c42%7c17%7c10%7c8%7c7%7c7%7c7%7c7%7c6%7c6%7c6%7c6%7c5%7c5&chdl=android%7cjava%7cstack-trace%7cbroadcastreceiver%7candroid-ndk%7cuser-agent%7candroid-webview%7cwebview%7cbackground%7cmultithreading%7candroid-source%7csms%7cadb%7csollections%7cactivity|chart",
		"lorem ipsum dolor sit amet, consectetur adipiscing elit. donec eu gravida purus, at interdum nulla. class aptent taciti sociosqu ad litora torquent per conubia nostra, per inceptos himenaeos. morbi tincidunt sapien ac elit convallis elementum.",
	}
	b.ResetTimer()
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		for _, origin := range origins {
			_ = ToLowerGaby([]byte(origin))
		}
	}
}

func BenchmarkToLowerUnsafeLowerCase(b *testing.B) {
	origins := []string{
		"https://example.com",
		"https://www.example.org",
		"https://sub.example.net",
		"http://chart.apis.google.com/chart?chs=500x500&chma=0,0,100,100&cht=p&chco=ff0000%2cffff00%7cff8000%2c00ff00%7c00ff00%2c0000ff&chd=t%3a122%2c42%2c17%2c10%2c8%2c7%2c7%2c7%2c7%2c6%2c6%2c6%2c6%2c5%2c5&chl=122%7c42%7c17%7c10%7c8%7c7%7c7%7c7%7c7%7c6%7c6%7c6%7c6%7c5%7c5&chdl=android%7cjava%7cstack-trace%7cbroadcastreceiver%7candroid-ndk%7cuser-agent%7candroid-webview%7cwebview%7cbackground%7cmultithreading%7candroid-source%7csms%7cadb%7csollections%7cactivity|chart",
		"lorem ipsum dolor sit amet, consectetur adipiscing elit. donec eu gravida purus, at interdum nulla. class aptent taciti sociosqu ad litora torquent per conubia nostra, per inceptos himenaeos. morbi tincidunt sapien ac elit convallis elementum.",
	}
	b.ResetTimer()
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		for _, origin := range origins {
			_ = ToLowerUnsafe([]byte(origin))
		}
	}
}

func BenchmarkToLowerSWARStringLowerCase(b *testing.B) {
	origins := []string{
		"https://example.com",
		"https://www.example.org",
		"https://sub.example.net",
		"http://chart.apis.google.com/chart?chs=500x500&chma=0,0,100,100&cht=p&chco=ff0000%2cffff00%7cff8000%2c00ff00%7c00ff00%2c0000ff&chd=t%3a122%2c42%2c17%2c10%2c8%2c7%2c7%2c7%2c7%2c6%2c6%2c6%2c6%2c5%2c5&chl=122%7c42%7c17%7c10%7c8%7c7%7c7%7c7%7c7%7c6%7c6%7c6%7c6%7c5%7c5&chdl=android%7cjava%7cstack-trace%7cbroadcastreceiver%7candroid-ndk%7cuser-agent%7candroid-webview%7cwebview%7cbackground%7cmultithreading%7candroid-source%7csms%7cadb%7csollections%7cactivity|chart",
		"lorem ipsum dolor sit amet, consectetur adipiscing elit. donec eu gravida purus, at interdum nulla. class aptent taciti sociosqu ad litora torquent per conubia nostra, per inceptos himenaeos. morbi tincidunt sapien ac elit convallis elementum.",
	}
	b.ResetTimer()
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		for _, origin := range origins {
			_ = ToLowerSWARString(origin)
		}
	}
}

func BenchmarkToLowerUnsafeStringLowerCase(b *testing.B) {
	origins := []string{
		"https://example.com",
		"https://www.example.org",
		"https://sub.example.net",
		"http://chart.apis.google.com/chart?chs=500x500&chma=0,0,100,100&cht=p&chco=ff0000%2cffff00%7cff8000%2c00ff00%7c00ff00%2c0000ff&chd=t%3a122%2c42%2c17%2c10%2c8%2c7%2c7%2c7%2c7%2c6%2c6%2c6%2c6%2c5%2c5&chl=122%7c42%7c17%7c10%7c8%7c7%7c7%7c7%7c7%7c6%7c6%7c6%7c6%7c5%7c5&chdl=android%7cjava%7cstack-trace%7cbroadcastreceiver%7candroid-ndk%7cuser-agent%7candroid-webview%7cwebview%7cbackground%7cmultithreading%7candroid-source%7csms%7cadb%7csollections%7cactivity|chart",
		"lorem ipsum dolor sit amet, consectetur adipiscing elit. donec eu gravida purus, at interdum nulla. class aptent taciti sociosqu ad litora torquent per conubia nostra, per inceptos himenaeos. morbi tincidunt sapien ac elit convallis elementum.",
	}
	b.ResetTimer()
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		for _, origin := range origins {
			_ = ToLowerUnsafeString(origin)
		}
	}
}

func BenchmarkOptimalToLowerLowerCase(b *testing.B) {
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
			_ = OptimalToLower(origin)
		}
	}
}
func BenchmarkOptimalToLowerMixedCase(b *testing.B) {
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
			_ = OptimalToLower(origin)
		}
	}
}

func BenchmarkToLowerSWARv2LowerCase(b *testing.B) {
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
			_ = ToLowerSWARv2([]byte(origin))
		}
	}
}
func BenchmarkToLowerSWARv2MixedCase(b *testing.B) {
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
			_ = ToLowerSWARv2([]byte(origin))
		}
	}
}

func BenchmarkToLowerInPlaceStringLowerCase(b *testing.B) {
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
			_ = ToLowerInPlaceString(origin)
		}
	}
}
func BenchmarkToLowerInPlaceStringMixedCase(b *testing.B) {
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
			_ = ToLowerInPlaceString(origin)
		}
	}
}

func BenchmarkToLowerSWAREnhancedLowerCase(b *testing.B) {
	origins := []string{
		"https://example.com",
		"https://www.example.org",
		"https://sub.example.net",
		"http://chart.apis.google.com/chart?chs=500x500&chma=0,0,100,100&cht=p&chco=ff0000%2cffff00%7cff8000%2c00ff00%7c00ff00%2c0000ff&chd=t%3a122%2c42%2c17%2c10%2c8%2c7%2c7%2c7%2c7%2c6%2c6%2c6%2c6%2c5%2c5&chl=122%7c42%7c17%7c10%7c8%7c7%7c7%7c7%7c7%7c6%7c6%7c6%7c6%7c5%7c5&chdl=android%7cjava%7cstack-trace%7cbroadcastreceiver%7candroid-ndk%7cuser-agent%7candroid-webview%7cwebview%7cbackground%7cmultithreading%7candroid-source%7csms%7cadb%7csollections%7cactivity|chart",
		"lorem ipsum dolor sit amet, consectetur adipiscing elit. donec eu gravida purus, at interdum nulla. class aptent taciti sociosqu ad litora torquent per conubia nostra, per inceptos himenaeos. morbi tincidunt sapien ac elit convallis elementum.",
	}
	b.ResetTimer()
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		for _, origin := range origins {
			_ = ToLowerSWAREnhanced([]byte(origin))
		}
	}
}

func BenchmarkToLowerSWAREnhancedMixedCase(b *testing.B) {
	origins := []string{
		"https://Example.com",
		"https://www.Example.org",
		"https://Sub.example.net",
		"http://chart.apis.google.com/chart?chs=500x500&chma=0,0,100,100&cht=p&chco=FF0000%2CFFFF00%7CFF8000%2C00FF00%7C00FF00%2C0000FF&chd=t%3A122%2C42%2C17%2C10%2C8%2C7%2C7%2C7%2C7%2C6%2C6%2C6%2C6%2C5%2C5&chl=122%7C42%7C17%7C10%7C8%7C7%7C7%7C7%7C7%7C6%7C6%7C6%7C6%7C5%7C5&chdl=android%7Cjava%7Cstack-trace%7Cbroadcastreceiver%7Candroid-ndk%7Cuser-agent%7Candroid-webview%7Cwebview%7Cbackground%7Cmultithreading%7Candroid-source%7Csms%7Cadb%7Csollections%7Cactivity|Chart",
		"Lorem ipsum dolor sit amet, consectetur adipiscing elit. Donec eu gravida purus, at interdum nulla. Class aptent taciti sociosqu ad litora torquent per conubia nostra, per inceptos himenaeos. Morbi tincidunt sapien ac elit convallis elementum.",
	}
	b.ResetTimer()
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		for _, origin := range origins {
			_ = ToLowerSWAREnhanced([]byte(origin))
		}
	}
}

func BenchmarkSuperToLowerLowerCase(b *testing.B) {
	origins := []string{
		"https://example.com",
		"https://www.example.org",
		"https://sub.example.net",
		"http://chart.apis.google.com/chart?chs=500x500&chma=0,0,100,100&cht=p&chco=ff0000%2cffff00%7cff8000%2c00ff00%7c00ff00%2c0000ff&chd=t%3a122%2c42%2c17%2c10%2c8%2c7%2c7%2c7%2c7%2c6%2c6%2c6%2c6%2c5%2c5&chl=122%7c42%7c17%7c10%7c8%7c7%7c7%7c7%7c7%7c6%7c6%7c6%7c6%7c5%7c5&chdl=android%7cjava%7cstack-trace%7cbroadcastreceiver%7candroid-ndk%7cuser-agent%7candroid-webview%7cwebview%7cbackground%7cmultithreading%7candroid-source%7csms%7cadb%7csollections%7cactivity|chart",
		"lorem ipsum dolor sit amet, consectetur adipiscing elit. donec eu gravida purus, at interdum nulla. class aptent taciti sociosqu ad litora torquent per conubia nostra, per inceptos himenaeos. morbi tincidunt sapien ac elit convallis elementum.",
	}
	b.ResetTimer()
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		for _, origin := range origins {
			_ = SuperToLower(origin)
		}
	}
}

func BenchmarkSuperToLowerMixedCase(b *testing.B) {
	origins := []string{
		"https://Example.com",
		"https://www.Example.org",
		"https://Sub.example.net",
		"http://chart.apis.google.com/chart?chs=500x500&chma=0,0,100,100&cht=p&chco=FF0000%2CFFFF00%7CFF8000%2C00FF00%7C00FF00%2C0000FF&chd=t%3A122%2C42%2C17%2C10%2C8%2C7%2C7%2C7%2C7%2C6%2C6%2C6%2C6%2C5%2C5&chl=122%7C42%7C17%7C10%7C8%7C7%7C7%7C7%7C7%7C6%7C6%7C6%7C6%7C5%7C5&chdl=android%7Cjava%7Cstack-trace%7Cbroadcastreceiver%7Candroid-ndk%7Cuser-agent%7Candroid-webview%7Cwebview%7Cbackground%7Cmultithreading%7Candroid-source%7Csms%7Cadb%7Csollections%7Cactivity|Chart",
		"Lorem ipsum dolor sit amet, consectetur adipiscing elit. Donec eu gravida purus, at interdum nulla. Class aptent taciti sociosqu ad litora torquent per conubia nostra, per inceptos himenaeos. Morbi tincidunt sapien ac elit convallis elementum.",
	}
	b.ResetTimer()
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		for _, origin := range origins {
			_ = SuperToLower(origin)
		}
	}
}

func BenchmarkToLowerHeaderLowerCase(b *testing.B) {
	headers := []string{
		"accept-encoding",
		"content-type",
		"user-agent",
		"x-forwarded-for",
		"authorization",
		"referer",
		"cookie",
		"origin",
		"host",
		"connection",
	}
	b.ResetTimer()
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		for _, header := range headers {
			_ = ToLowerHeader(header)
		}
	}
}

func BenchmarkToLowerHeaderMixedCase(b *testing.B) {
	headers := []string{
		"Accept-Encoding",
		"Content-Type",
		"User-Agent",
		"X-Forwarded-For",
		"Authorization",
		"Referer",
		"Cookie",
		"Origin",
		"Host",
		"Connection",
	}
	b.ResetTimer()
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		for _, header := range headers {
			_ = ToLowerHeader(header)
		}
	}
}

func BenchmarkToLowerInPlaceOptimizedLowerCase(b *testing.B) {
	origins := []string{
		"https://example.com",
		"https://www.example.org",
		"https://sub.example.net",
		"http://chart.apis.google.com/chart?chs=500x500&chma=0,0,100,100&cht=p&chco=ff0000%2cffff00%7cff8000%2c00ff00%7c00ff00%2c0000ff&chd=t%3a122%2c42%2c17%2c10%2c8%2c7%2c7%2c7%2c7%2c6%2c6%2c6%2c6%2c5%2c5&chl=122%7c42%7c17%7c10%7c8%7c7%7c7%7c7%7c7%7c6%7c6%7c6%7c6%7c5%7c5&chdl=android%7cjava%7cstack-trace%7cbroadcastreceiver%7candroid-ndk%7cuser-agent%7candroid-webview%7cwebview%7cbackground%7cmultithreading%7candroid-source%7csms%7cadb%7csollections%7cactivity|chart",
		"lorem ipsum dolor sit amet, consectetur adipiscing elit. donec eu gravida purus, at interdum nulla. class aptent taciti sociosqu ad litora torquent per conubia nostra, per inceptos himenaeos. morbi tincidunt sapien ac elit convallis elementum.",
	}
	b.ResetTimer()
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		for _, origin := range origins {
			// Need to create a new slice for each test since ToLowerInPlaceOptimized modifies the input
			buf := []byte(origin)
			_ = ToLowerInPlaceOptimized(buf)
		}
	}
}

func BenchmarkToLowerInPlaceOptimizedMixedCase(b *testing.B) {
	origins := []string{
		"https://Example.com",
		"https://www.Example.org",
		"https://Sub.example.net",
		"http://chart.apis.google.com/chart?chs=500x500&chma=0,0,100,100&cht=p&chco=FF0000%2CFFFF00%7CFF8000%2C00FF00%7C00FF00%2C0000FF&chd=t%3A122%2C42%2C17%2C10%2C8%2C7%2C7%2C7%2C7%2C6%2C6%2C6%2C6%2C5%2C5&chl=122%7C42%7C17%7C10%7C8%7C7%7C7%7C7%7C7%7C6%7C6%7C6%7C6%7C5%7C5&chdl=android%7Cjava%7Cstack-trace%7Cbroadcastreceiver%7Candroid-ndk%7Cuser-agent%7Candroid-webview%7Cwebview%7Cbackground%7Cmultithreading%7Candroid-source%7Csms%7Cadb%7Csollections%7Cactivity|Chart",
		"Lorem ipsum dolor sit amet, consectetur adipiscing elit. Donec eu gravida purus, at interdum nulla. Class aptent taciti sociosqu ad litora torquent per conubia nostra, per inceptos himenaeos. Morbi tincidunt sapien ac elit convallis elementum.",
	}
	b.ResetTimer()
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		for _, origin := range origins {
			// Need to create a new slice for each test since ToLowerInPlaceOptimized modifies the input
			buf := []byte(origin)
			_ = ToLowerInPlaceOptimized(buf)
		}
	}
}

// Add tests for our new functions
func TestToLowerSWAREnhanced(t *testing.T) {
	cases := []struct {
		in, want string
	}{
		{"", ""},
		{"a", "a"},
		{"A", "a"},
		{"abc", "abc"},
		{"ABC", "abc"},
		{"AbC123", "abc123"},
		{"!@#ABCxyz", "!@#abcxyz"},
		{"https://Example.COM/Path", "https://example.com/path"},
		{"1234567890", "1234567890"},
		{"Lorem IPSUM", "lorem ipsum"},
	}

	for _, c := range cases {
		got := string(ToLowerSWAREnhanced([]byte(c.in)))
		if got != c.want {
			t.Errorf("ToLowerSWAREnhanced(%q) = %q; want %q", c.in, got, c.want)
		}
	}
}

func TestSuperToLower(t *testing.T) {
	cases := []struct {
		in, want string
	}{
		{"", ""},
		{"a", "a"},
		{"A", "a"},
		{"abc", "abc"},
		{"ABC", "abc"},
		{"AbC123", "abc123"},
		{"!@#ABCxyz", "!@#abcxyz"},
		{"https://Example.COM/Path", "https://example.com/path"},
		{"1234567890", "1234567890"},
		{"Lorem IPSUM", "lorem ipsum"},
	}

	for _, c := range cases {
		got := SuperToLower(c.in)
		if got != c.want {
			t.Errorf("SuperToLower(%q) = %q; want %q", c.in, got, c.want)
		}
	}
}

func TestToLowerHeader(t *testing.T) {
	cases := []struct {
		in, want string
	}{
		{"", ""},
		{"Content-Type", "content-type"},
		{"X-Forwarded-For", "x-forwarded-for"},
		{"User-Agent", "user-agent"},
		{"Host", "host"},
		{"Accept-Encoding", "accept-encoding"},
	}

	for _, c := range cases {
		got := ToLowerHeader(c.in)
		if got != c.want {
			t.Errorf("ToLowerHeader(%q) = %q; want %q", c.in, got, c.want)
		}
	}
}

func TestToLowerInPlaceOptimized(t *testing.T) {
	cases := []struct {
		in, want string
	}{
		{"", ""},
		{"a", "a"},
		{"A", "a"},
		{"abc", "abc"},
		{"ABC", "abc"},
		{"AbC123", "abc123"},
		{"!@#ABCxyz", "!@#abcxyz"},
		{"https://Example.COM/Path", "https://example.com/path"},
		{"1234567890", "1234567890"},
		{"Lorem IPSUM", "lorem ipsum"},
	}

	for _, c := range cases {
		input := []byte(c.in)
		got := string(ToLowerInPlaceOptimized(input))
		if got != c.want {
			t.Errorf("ToLowerInPlaceOptimized(%q) = %q; want %q", c.in, got, c.want)
		}
	}
}

// TestToLowerByte tests the byte-level conversion function
func TestToLowerByte(t *testing.T) {
	testCases := []struct {
		in   byte
		want byte
	}{
		{'A', 'a'},
		{'Z', 'z'},
		{'a', 'a'},
		{'z', 'z'},
		{'0', '0'},
		{'!', '!'},
		{0, 0},
		{255, 255},
	}

	for _, tc := range testCases {
		got := ToLowerByte(tc.in)
		if got != tc.want {
			t.Errorf("ToLowerByte(%q) = %q; want %q", tc.in, got, tc.want)
		}
	}
}

// TestToLowerSWARv2 tests the enhanced SIMD within a register implementation
func TestToLowerSWARv2(t *testing.T) {
	testCases := []struct {
		in   string
		want string
	}{
		{"", ""},
		{"A", "a"},
		{"ABC", "abc"},
		{"ABCDEFGHIJKLMNOPQRSTUVWXYZ", "abcdefghijklmnopqrstuvwxyz"},
		{"ABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890", "abcdefghijklmnopqrstuvwxyz1234567890"},
		{"ABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890!@#$%^&*()", "abcdefghijklmnopqrstuvwxyz1234567890!@#$%^&*()"},
		{"mixed CASE string WITH numbers 123", "mixed case string with numbers 123"},
		// Test with lengths crossing SWAR processing boundaries
		{strings.Repeat("A", 7), strings.Repeat("a", 7)},
		{strings.Repeat("A", 8), strings.Repeat("a", 8)},
		{strings.Repeat("A", 9), strings.Repeat("a", 9)},
		{strings.Repeat("A", 31), strings.Repeat("a", 31)},
		{strings.Repeat("A", 32), strings.Repeat("a", 32)},
		{strings.Repeat("A", 33), strings.Repeat("a", 33)},
		// Test with null bytes and non-ASCII characters
		{string([]byte{0x41, 0x00, 0x42}), string([]byte{0x61, 0x00, 0x62})},
	}

	for _, tc := range testCases {
		input := []byte(tc.in)
		got := string(ToLowerSWARv2(input))
		if got != tc.want {
			t.Errorf("ToLowerSWARv2(%q) = %q; want %q", tc.in, got, tc.want)
		}
	}
}

// TestToLowerInPlace tests the in-place byte slice conversion function
func TestToLowerInPlace(t *testing.T) {
	testCases := []struct {
		in   string
		want string
	}{
		{"", ""},
		{"a", "a"},
		{"A", "a"},
		{"ABC123", "abc123"},
		{"Mixed CASE", "mixed case"},
		// Test with strings that are already lowercase
		{"already lowercase", "already lowercase"},
		// Test with lengths that cross the algorithm selection threshold
		{strings.Repeat("A", 63), strings.Repeat("a", 63)},
		{strings.Repeat("A", 64), strings.Repeat("a", 64)},
		{strings.Repeat("A", 65), strings.Repeat("a", 65)},
	}

	for _, tc := range testCases {
		input := []byte(tc.in)
		got := string(ToLowerInPlace(input))
		if got != tc.want {
			t.Errorf("ToLowerInPlace(%q) = %q; want %q", tc.in, got, tc.want)
		}

		// Also verify that the input slice was actually modified in-place
		if !bytes.Equal(input, []byte(tc.want)) {
			t.Errorf("ToLowerInPlace did not modify the input slice in-place for %q", tc.in)
		}
	}
}

// TestToLowerInPlaceString tests string conversion that avoids allocation
// for already lowercase strings
func TestToLowerInPlaceString(t *testing.T) {
	testCases := []struct {
		in   string
		want string
	}{
		{"", ""},
		{"a", "a"},
		{"A", "a"},
		{"abc", "abc"},
		{"ABC", "abc"},
		{"Mixed CASE 123", "mixed case 123"},
		// Test with lengths that cross the algorithm selection threshold
		{strings.Repeat("A", 63), strings.Repeat("a", 63)},
		{strings.Repeat("A", 64), strings.Repeat("a", 64)},
		{strings.Repeat("A", 65), strings.Repeat("a", 65)},
		// Test with already lowercase strings to verify early return
		{"already lowercase", "already lowercase"},
	}

	for _, tc := range testCases {
		got := ToLowerInPlaceString(tc.in)
		if got != tc.want {
			t.Errorf("ToLowerInPlaceString(%q) = %q; want %q", tc.in, got, tc.want)
		}

		// For already lowercase strings, verify if it's the same instance (no allocation)
		if tc.in == tc.want && tc.in != "" {
			// This test is a bit tricky since we can't directly compare pointers for strings
			// But we can verify the implementation logic separately
			hasUpper := false
			for i := 0; i < len(tc.in); i++ {
				if tc.in[i] >= 'A' && tc.in[i] <= 'Z' {
					hasUpper = true
					break
				}
			}
			if !hasUpper && got != tc.in {
				t.Errorf("ToLowerInPlaceString(%q) returned a different instance for already lowercase string", tc.in)
			}
		}
	}
}

// TestToLowerComparison runs comprehensive comparison tests for all ToLower functions
func TestToLowerComparison(t *testing.T) {
	testCases := []struct {
		name string
		in   string
		want string
	}{
		{"empty", "", ""},
		{"single lowercase", "a", "a"},
		{"single uppercase", "A", "a"},
		{"all lowercase", "abcdefghijklmnopqrstuvwxyz", "abcdefghijklmnopqrstuvwxyz"},
		{"all uppercase", "ABCDEFGHIJKLMNOPQRSTUVWXYZ", "abcdefghijklmnopqrstuvwxyz"},
		{"mixed case", "AbCdEfGhIjKlMnOpQrStUvWxYz", "abcdefghijklmnopqrstuvwxyz"},
		{"mixed case with non-letters", "Hello, World! 123", "hello, world! 123"},
		{"URL", "https://Example.COM/Path?Query=Value", "https://example.com/path?query=value"},
		{"Edge cases", string([]byte{0x00, 0x41, 0x7F, 0xFF}), string([]byte{0x00, 0x61, 0x7F, 0xFF})},
	}

	// Define all our ToLower functions
	stringFuncs := []struct {
		name string
		fn   func(string) string
	}{
		{"ToLowerUtils", ToLowerUtils},
		{"ToLowerStrings", ToLowerStrings},
		{"HybridToLower", HybridToLower},
		{"ToLower", ToLower},
		{"OptimalToLower", OptimalToLower},
		{"ToLowerSWARString", ToLowerSWARString},
		{"ToLowerSWARv2String", ToLowerSWARv2String},
		{"ToLowerUnsafeString", ToLowerUnsafeString},
		{"ToLowerGabyString", ToLowerGabyString},
		{"ToLowerInPlaceString", ToLowerInPlaceString},
		{"SuperToLower", SuperToLower},
		{"ToLowerHeader", ToLowerHeader},
	}

	byteFuncs := []struct {
		name string
		fn   func([]byte) []byte
	}{
		{"ToLowerSWAR", ToLowerSWAR},
		{"ToLowerSWARv2", ToLowerSWARv2},
		{"ToLowerUnsafe", ToLowerUnsafe},
		{"ToLowerGaby", ToLowerGaby},
		{"ToLowerInPlace", ToLowerInPlace},
		{"ToLowerSWAREnhanced", ToLowerSWAREnhanced},
		{"ToLowerInPlaceOptimized", ToLowerInPlaceOptimized},
	}

	// Test all string functions
	for _, tc := range testCases {
		for _, fn := range stringFuncs {
			t.Run(fn.name+"_"+tc.name, func(t *testing.T) {
				got := fn.fn(tc.in)
				if got != tc.want {
					t.Errorf("%s(%q) = %q; want %q", fn.name, tc.in, got, tc.want)
				}
			})
		}
	}

	// Test all byte slice functions
	for _, tc := range testCases {
		for _, fn := range byteFuncs {
			t.Run(fn.name+"_"+tc.name, func(t *testing.T) {
				input := []byte(tc.in)
				got := string(fn.fn(input))
				if got != tc.want {
					t.Errorf("%s(%q) = %q; want %q", fn.name, tc.in, got, tc.want)
				}
			})
		}
	}
}

// TestToLowerBoundaryConditions tests various functions with boundary conditions
func TestToLowerBoundaryConditions(t *testing.T) {
	// Generate test cases with specific lengths to test boundary conditions
	testLengths := []int{0, 1, 7, 8, 9, 15, 16, 17, 31, 32, 33, 63, 64, 65, 127, 128, 129}
	testCases := make([]struct {
		length int
		upper  string
		lower  string
	}, len(testLengths))

	for i, length := range testLengths {
		testCases[i].length = length
		testCases[i].upper = strings.Repeat("A", length)
		testCases[i].lower = strings.Repeat("a", length)
	}

	byteFuncs := []struct {
		name string
		fn   func([]byte) []byte
	}{
		{"ToLowerSWAR", ToLowerSWAR},
		{"ToLowerSWARv2", ToLowerSWARv2},
		{"ToLowerUnsafe", ToLowerUnsafe},
		{"ToLowerInPlace", ToLowerInPlace},
		{"ToLowerSWAREnhanced", ToLowerSWAREnhanced},
		{"ToLowerInPlaceOptimized", ToLowerInPlaceOptimized},
	}

	stringFuncs := []struct {
		name string
		fn   func(string) string
	}{
		{"OptimalToLower", OptimalToLower},
		{"ToLowerInPlaceString", ToLowerInPlaceString},
		{"SuperToLower", SuperToLower},
	}

	// Test byte slice functions
	for _, tc := range testCases {
		for _, fn := range byteFuncs {
			t.Run(fn.name+fmt.Sprintf("_length_%d", tc.length), func(t *testing.T) {
				input := []byte(tc.upper)
				got := string(fn.fn(input))
				if got != tc.lower {
					t.Errorf("%s(%s) = %q; want %q", fn.name, tc.upper, got, tc.lower)
				}
			})
		}
	}

	// Test string functions with boundary conditions
	for _, tc := range testCases {
		for _, fn := range stringFuncs {
			t.Run(fn.name+fmt.Sprintf("_length_%d", tc.length), func(t *testing.T) {
				got := fn.fn(tc.upper)
				if got != tc.lower {
					t.Errorf("%s(%s) = %q; want %q", fn.name, tc.upper, got, tc.lower)
				}
			})
		}
	}
}

// BenchmarkToLowerByte benchmarks the low-level byte conversion function
func BenchmarkToLowerByte(b *testing.B) {
	chars := []byte{'A', 'N', 'Z', 'a', 'n', 'z', '0', '!', '@'}
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		for _, c := range chars {
			_ = ToLowerByte(c)
		}
	}
}

// BenchmarkToLowerInPlaceVariousLengths measures performance across different string lengths
func BenchmarkToLowerInPlaceVariousLengths(b *testing.B) {
	lengths := []int{8, 16, 32, 64, 128, 256, 512, 1024}

	for _, length := range lengths {
		b.Run(fmt.Sprintf("Length_%d_Mixed", length), func(b *testing.B) {
			s := strings.Repeat("A", length/2) + strings.Repeat("a", length/2)
			b.ResetTimer()
			b.ReportAllocs()

			for i := 0; i < b.N; i++ {
				buf := []byte(s)
				_ = ToLowerInPlace(buf)
			}
		})

		b.Run(fmt.Sprintf("Length_%d_Upper", length), func(b *testing.B) {
			s := strings.Repeat("A", length)
			b.ResetTimer()
			b.ReportAllocs()

			for i := 0; i < b.N; i++ {
				buf := []byte(s)
				_ = ToLowerInPlace(buf)
			}
		})

		b.Run(fmt.Sprintf("Length_%d_Lower", length), func(b *testing.B) {
			s := strings.Repeat("a", length)
			b.ResetTimer()
			b.ReportAllocs()

			for i := 0; i < b.N; i++ {
				buf := []byte(s)
				_ = ToLowerInPlace(buf)
			}
		})
	}
}

// TestToLowerMalformedUTF8 tests handling of malformed UTF-8 sequences
func TestToLowerMalformedUTF8(t *testing.T) {
	// Create some malformed UTF-8 sequences
	malformed := []string{
		// Incomplete UTF-8 sequences
		string([]byte{0xC0}),             // Incomplete 2-byte sequence
		string([]byte{0xE0, 0x80}),       // Incomplete 3-byte sequence
		string([]byte{0xF0, 0x80, 0x80}), // Incomplete 4-byte sequence

		// Mixed ASCII and malformed UTF-8
		string([]byte{'A', 'B', 0xC0, 'C', 'D'}),
		string([]byte{'A', 0xE0, 0x80, 'B', 'C'}),
	}

	// All our functions should handle malformed UTF-8 without crashing
	funcs := []struct {
		name string
		fn   func(string) string
	}{
		{"ToLowerSWARString", ToLowerSWARString},
		{"OptimalToLower", OptimalToLower},
		{"ToLowerUnsafeString", ToLowerUnsafeString},
		{"SuperToLower", SuperToLower},
	}

	for _, s := range malformed {
		for _, fn := range funcs {
			t.Run(fn.name, func(t *testing.T) {
				// Test that the function doesn't panic
				defer func() {
					if r := recover(); r != nil {
						t.Errorf("%s panicked on malformed UTF-8: %v", fn.name, r)
					}
				}()

				result := fn.fn(s)

				// Verify it at least processed ASCII characters correctly
				for i := 0; i < len(s); i++ {
					if i < len(result) && s[i] >= 'A' && s[i] <= 'Z' {
						if result[i] != s[i]|0x20 {
							t.Errorf("%s didn't convert ASCII uppercase to lowercase in malformed UTF-8", fn.name)
						}
					}
				}
			})
		}
	}
}

// Comprehensive benchmark for all ToLower functions
func BenchmarkToLowerComparison(b *testing.B) {
	testCases := []struct {
		name string
		in   string
	}{
		{"empty", ""},
		{"single lowercase", "a"},
		{"single uppercase", "A"},
		{"all lowercase", "abcdefghijklmnopqrstuvwxyz"},
		{"all uppercase", "ABCDEFGHIJKLMNOPQRSTUVWXYZ"},
		{"mixed case", "AbCdEfGhIjKlMnOpQrStUvWxYz"},
		{"mixed case with non-letters", "Hello, World! 123"},
		{"HTTP Headers", "Content-Type: application/json"},
		{"URL with mixed case", "https://Example.COM/Path?Query=Value"},
		{"URL", "https://Example.COM/Path?Query=Value"},
		{"Edge cases", string([]byte{0x00, 0x41, 0x7F, 0xFF})},
	}

	stringFuncs := []struct {
		name string
		fn   func(string) string
	}{
		{"ToLowerUtils", ToLowerUtils},
		{"ToLowerStrings", ToLowerStrings},
		{"HybridToLower", HybridToLower},
		{"ToLower", ToLower},
		{"OptimalToLower", OptimalToLower},
		{"ToLowerSWARString", ToLowerSWARString},
		{"ToLowerSWARv2String", ToLowerSWARv2String},
		{"ToLowerUnsafeString", ToLowerUnsafeString},
		{"ToLowerGabyString", ToLowerGabyString},
		{"ToLowerInPlaceString", ToLowerInPlaceString},
		{"SuperToLower", SuperToLower},
		{"ToLowerHeader", ToLowerHeader},
	}

	byteFuncs := []struct {
		name string
		fn   func([]byte) []byte
	}{
		{"ToLowerSWAR", ToLowerSWAR},
		{"ToLowerSWARv2", ToLowerSWARv2},
		{"ToLowerUnsafe", ToLowerUnsafe},
		{"ToLowerGaby", ToLowerGaby},
		{"ToLowerInPlace", ToLowerInPlace},
		{"ToLowerSWAREnhanced", ToLowerSWAREnhanced},
		{"ToLowerInPlaceOptimized", ToLowerInPlaceOptimized},
	}

	for _, tc := range testCases {
		for _, fn := range stringFuncs {
			b.Run(fn.name+"_"+tc.name, func(b *testing.B) {
				for i := 0; i < b.N; i++ {
					_ = fn.fn(tc.in)
				}
			})
		}
		for _, fn := range byteFuncs {
			b.Run(fn.name+"_"+tc.name, func(b *testing.B) {
				input := []byte(tc.in)
				for i := 0; i < b.N; i++ {
					_ = fn.fn(input)
				}
			})
		}
	}
}
