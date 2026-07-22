//go:build !amd64 && !arm64

package tolower

// toLowerAsmWithOffset is the portable fallback when assembly is unavailable.
func toLowerAsmWithOffset(b []byte, offset, length int) {
	ToLowerSWARWithOffset(b, offset, length)
}
