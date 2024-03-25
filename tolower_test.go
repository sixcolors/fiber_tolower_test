package tolower

import (
	"testing"
)

func BenchmarkUtilsToLowerLowerCase(b *testing.B) {
	origins := []string{
		"https://example.com",
		"https://www.example.org",
		"https://sub.example.net",
		// Add more realistic origins or referers as needed
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
		// Add more realistic origins or referers as needed
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
		// Add more realistic origins or referers as needed
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
		// Add more realistic origins or referers as needed
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
		// Add more realistic origins or referers as needed
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
		// Add more realistic origins or referers as needed
	}

	b.ResetTimer()
	b.ReportAllocs()

	for i := 0; i < b.N; i++ {
		for _, origin := range origins {
			_ = HybribToLower(origin)
		}
	}
}
