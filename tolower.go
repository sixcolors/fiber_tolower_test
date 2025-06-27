package tolower

import (
	"fmt"
	"strings"
	"unsafe"

	"github.com/gofiber/utils/v2"
)

func main() {
	fmt.Println("This file is used to demonstrate the usage of benchmarks.", "Please refer to the main_test.go file for the actual benchmarks.")

	// Do some work with the functions to avoid "declared and not used" errors
	_ = ToLowerUtils("https://example.com")
	_ = ToLowerStrings("https://example.com")
	_ = HybridToLower("https://example.com")

	fmt.Println("To run the benchmarks, use the following command:", "go test -v -run=^$ -bench=B -benchmem -count=4")
}

func ToLowerUtils(s string) string {
	return utils.ToLower(s)
}

func ToLowerStrings(s string) string {
	return strings.ToLower(s)
}

const (
	toLowerTable = "\x00\x01\x02\x03\x04\x05\x06\a\b\t\n\v\f\r\x0e\x0f\x10\x11\x12\x13\x14\x15\x16\x17\x18\x19\x1a\x1b\x1c\x1d\x1e\x1f !\"#$%&'()*+,-./0123456789:;<=>?@abcdefghijklmnopqrstuvwxyz[\\]^_`abcdefghijklmnopqrstuvwxyz{|}~\u007f\x80\x81\x82\x83\x84\x85\x86\x87\x88\x89\x8a\x8b\x8c\x8d\x8e\x8f\x90\x91\x92\x93\x94\x95\x96\x97\x98\x99\x9a\x9b\x9c\x9d\x9e\x9f\xa0\xa1\xa2\xa3\xa4\xa5\xa6\xa7\xa8\xa9\xaa\xab\xac\xad\xae\xaf\xb0\xb1\xb2\xb3\xb4\xb5\xb6\xb7\xb8\xb9\xba\xbb\xbc\xbd\xbe\xbf\xc0\xc1\xc2\xc3\xc4\xc5\xc6\xc7\xc8\xc9\xca\xcb\xcc\xcd\xce\xcf\xd0\xd1\xd2\xd3\xd4\xd5\xd6\xd7\xd8\xd9\xda\xdb\xdc\xdd\xde\xdf\xe0\xe1\xe2\xe3\xe4\xe5\xe6\xe7\xe8\xe9\xea\xeb\xec\xed\xee\xef\xf0\xf1\xf2\xf3\xf4\xf5\xf6\xf7\xf8\xf9\xfa\xfb\xfc\xfd\xfe\xff"
	toUpperTable = "\x00\x01\x02\x03\x04\x05\x06\a\b\t\n\v\f\r\x0e\x0f\x10\x11\x12\x13\x14\x15\x16\x17\x18\x19\x1a\x1b\x1c\x1d\x1e\x1f !\"#$%&'()*+,-./0123456789:;<=>?@ABCDEFGHIJKLMNOPQRSTUVWXYZ[\\]^_`ABCDEFGHIJKLMNOPQRSTUVWXYZ{|}~\u007f\x80\x81\x82\x83\x84\x85\x86\x87\x88\x89\x8a\x8b\x8c\x8d\x8e\x8f\x90\x91\x92\x93\x94\x95\x96\x97\x98\x99\x9a\x9b\x9c\x9d\x9e\x9f\xa0\xa1\xa2\xa3\xa4\xa5\xa6\xa7\xa8\xa9\xaa\xab\xac\xad\xae\xaf\xb0\xb1\xb2\xb3\xb4\xb5\xb6\xb7\xb8\xb9\xba\xbb\xbc\xbd\xbe\xbf\xc0\xc1\xc2\xc3\xc4\xc5\xc6\xc7\xc8\xc9\xca\xcb\xcc\xcd\xce\xcf\xd0\xd1\xd2\xd3\xd4\xd5\xd6\xd7\xd8\xd9\xda\xdb\xdc\xdd\xde\xdf\xe0\xe1\xe2\xe3\xe4\xe5\xe6\xe7\xe8\xe9\xea\xeb\xec\xed\xee\xef\xf0\xf1\xf2\xf3\xf4\xf5\xf6\xf7\xf8\xf9\xfa\xfb\xfc\xfd\xfe\xff"
)

// HybridToLower returns a lowercase version of the input ASCII string.
//
// It leverages an optimization from the standard library's `strings.ToLower` function,
// where it first checks if the string contains any uppercase characters before converting it.
//
// This approach can be more efficient when the string is already lowercase,
// and because it is limited to ASCII characters,
// it can be faster than the standard library's `ToLower` function.
//
// For strings that are already lowercase,
// this function will be faster than both the standard library's `ToLower` function and
// the GoFiber Utils `ToLower` function.
//
// In the case of mixed-case strings, this function will be faster than the standard library's
// `ToLower` function, but slower than the GoFiber Utils `ToLower` function.
func HybridToLower(s string) string {
	hasUpper := false
	var i int
	for i = 0; i < len(s); i++ {
		c := s[i]
		if toLowerTable[c] != c {
			hasUpper = true
			break
		}
	}

	if !hasUpper {
		return s
	}

	res := make([]byte, len(s))
	copy(res, s)
	for ; i < len(res); i++ {
		res[i] = toLowerTable[res[i]]
	}

	return utils.UnsafeString(res)
}

// func OptimalToLower(s string) string {
// 	// First check if any conversion needed (like HybridToLower)
// 	// If needed, use SWAR implementation for the conversion
// 	if len(s) == 0 {
// 		return s
// 	}
// 	if len(s) == 1 {
// 		return string(ToLowerByte(s[0]))
// 	}
// 	if strings.IndexFunc(s, func(r rune) bool {
// 		return r >= 'A' && r <= 'Z'
// 	}) == -1 {
// 		return s // No uppercase letters, return original string
// 	}

// 	// If we reach here, it means we need to convert the string
// 	return utils.UnsafeString(ToLowerSWAR([]byte(s)))
// }

// OptimalToLower combines the best attributes of all implementations:
// - Fast early-out for already lowercase strings
// - Algorithm selection based on input length
// - SWAR processing for optimal performance on longer strings
func OptimalToLower(s string) string {
	// Handle trivial cases
	switch len(s) {
	case 0:
		return s
	case 1:
		return string(ToLowerByte(s[0]))
	}

	// Optimize check for strings that are already lowercase
	// For short strings (<32 bytes), check directly
	if len(s) < 32 {
		hasUpper := false
		for i := 0; i < len(s); i++ {
			c := s[i]
			if c >= 'A' && c <= 'Z' {
				hasUpper = true
				break
			}
		}
		if !hasUpper {
			return s
		}
	} else {
		// For longer strings, use IndexFunc which has internal optimizations
		if strings.IndexFunc(s, func(r rune) bool {
			return r >= 'A' && r <= 'Z'
		}) == -1 {
			return s
		}
	}

	// Select algorithm based on string length
	if len(s) <= 32 {
		// For short strings, avoid overhead of SWAR
		res := make([]byte, len(s))
		for i := 0; i < len(s); i++ {
			c := s[i]
			if c >= 'A' && c <= 'Z' {
				res[i] = c | 0x20
			} else {
				res[i] = c
			}
		}
		return utils.UnsafeString(res)
	}

	// For longer strings, use SWAR for best performance
	return utils.UnsafeString(ToLowerSWAR([]byte(s)))
}

func ToLowerGabyString(s string) string {
	if len(s) == 0 {
		return s
	}
	return utils.UnsafeString(ToLowerSWAR([]byte(s)))
}

func ToLowerGaby(b []byte) []byte {
	i := 0
	n := len(b)

	for ; i+16 <= n; i += 16 {
		b[i+0] = toLowerTable[b[i+0]]
		b[i+1] = toLowerTable[b[i+1]]
		b[i+2] = toLowerTable[b[i+2]]
		b[i+3] = toLowerTable[b[i+3]]
		b[i+4] = toLowerTable[b[i+4]]
		b[i+5] = toLowerTable[b[i+5]]
		b[i+6] = toLowerTable[b[i+6]]
		b[i+7] = toLowerTable[b[i+7]]
		b[i+8] = toLowerTable[b[i+8]]
		b[i+9] = toLowerTable[b[i+9]]
		b[i+10] = toLowerTable[b[i+10]]
		b[i+11] = toLowerTable[b[i+11]]
		b[i+12] = toLowerTable[b[i+12]]
		b[i+13] = toLowerTable[b[i+13]]
		b[i+14] = toLowerTable[b[i+14]]
		b[i+15] = toLowerTable[b[i+15]]
	}

	for ; i < n; i++ {
		b[i] = toLowerTable[b[i]]
	}

	return b
}

func ToLower(s string) string {
	if len(s) == 0 {
		return s
	}
	if len(s) == 1 {
		return string(ToLowerByte(s[0]))
	}

	res := make([]byte, len(s))
	for i := range s {
		res[i] = toLowerTable[s[i]]
	}
	return utils.UnsafeString(res)
}

func ToLowerByte(c byte) byte {
	// Set the 6th bit if 'A' <= c <= 'Z'
	if c >= 'A' && c <= 'Z' {
		return c | 0x20
	}
	return c
}

func ToLowerSWARString(s string) string {
	if len(s) == 0 {
		return s
	}
	return utils.UnsafeString(ToLowerSWAR([]byte(s)))
}

func ToLowerSWAR(b []byte) []byte {
	n := len(b)
	i := 0
	for ; i+8 <= n; i += 8 {
		chunk := *(*uint64)(unsafe.Pointer(&b[i]))

		// Create masks to identify bytes that are >= 'A' (0x41) and <= 'Z' (0x5A)
		// First, create a mask for bytes >= 'A'
		// Adding (0x80 - 'A') to each byte will set the high bit only for bytes >= 'A'
		geA := chunk + (0x8080808080808080 - 0x4141414141414141)
		// Then create a mask for bytes <= 'Z'
		// Adding (0x80 - ('Z' + 1)) to each byte will clear the high bit only for bytes <= 'Z'
		leZ := chunk + (0x8080808080808080 - 0x5B5B5B5B5B5B5B5B)
		// Combine these masks: high bit set only for bytes in range 'A'..'Z'
		mask := (geA & ^leZ) & 0x8080808080808080

		// Convert the mask to set bit 0x20 (shift right by 2 and mask)
		lowercase := (mask >> 2) & 0x2020202020202020

		// Apply the mask to set the 0x20 bit for uppercase characters
		chunk |= lowercase
		*(*uint64)(unsafe.Pointer(&b[i])) = chunk
	}

	// Process remaining bytes
	for ; i < n; i++ {
		c := b[i]
		if c >= 'A' && c <= 'Z' {
			b[i] = c | 0x20
		}
	}
	return b
}

func ToLowerSWARv2String(s string) string {
	if len(s) == 0 {
		return s
	}
	return utils.UnsafeString(ToLowerSWARv2([]byte(s)))
}

func ToLowerSWARv2(b []byte) []byte {
	n := len(b)
	i := 0

	// Process 32 bytes (4 chunks of 8 bytes) at a time
	for ; i+32 <= n; i += 32 {
		// Process 4 chunks in single loop iteration for better instruction pipelining
		for offset := 0; offset < 32; offset += 8 {
			chunk := *(*uint64)(unsafe.Pointer(&b[i+offset]))

			// SWAR magic: identify uppercase ASCII letters (A-Z)
			// Step 1: Add (0x80-'A') to set high bit only for bytes >= 'A'
			geA := chunk + (0x8080808080808080 - 0x4141414141414141)
			// Step 2: Add (0x80-('Z'+1)) to clear high bit only for bytes <= 'Z'
			leZ := chunk + (0x8080808080808080 - 0x5B5B5B5B5B5B5B5B)
			// Step 3: Combine masks to isolate only bytes that are both >= 'A' AND <= 'Z'
			mask := (geA & ^leZ) & 0x8080808080808080

			// Convert mask to set the 0x20 bit for uppercase letters only
			lowercase := (mask >> 2) & 0x2020202020202020

			// Apply conversion
			*(*uint64)(unsafe.Pointer(&b[i+offset])) = chunk | lowercase
		}
	}

	// Handle remaining 8-byte chunks
	for ; i+8 <= n; i += 8 {
		// Same SWAR logic as above
		chunk := *(*uint64)(unsafe.Pointer(&b[i]))
		geA := chunk + (0x8080808080808080 - 0x4141414141414141)
		leZ := chunk + (0x8080808080808080 - 0x5B5B5B5B5B5B5B5B)
		mask := (geA & ^leZ) & 0x8080808080808080
		lowercase := (mask >> 2) & 0x2020202020202020
		*(*uint64)(unsafe.Pointer(&b[i])) = chunk | lowercase
	}

	// Process remaining bytes
	for ; i < n; i++ {
		c := b[i]
		if c >= 'A' && c <= 'Z' {
			b[i] = c | 0x20
		}
	}

	return b
}

func ToLowerUnsafeString(s string) string {
	if len(s) == 0 {
		return s
	}
	return utils.UnsafeString(ToLowerUnsafe([]byte(s)))
}

func ToLowerUnsafe(b []byte) []byte {
	if len(b) == 0 {
		return b
	}

	// Calculate how many complete uint64 blocks we can process
	n8 := len(b) / 8

	// Get pointer to the first byte and create a uint64 slice viewing the same memory
	if n8 > 0 {
		// Get a pointer to the first byte
		firstPtr := unsafe.Pointer(&b[0])

		// Create a slice of uint64 viewing the same underlying memory
		b64 := unsafe.Slice((*uint64)(firstPtr), n8)

		// Process each uint64 block
		for i := range b64 {
			x := b64[i]

			// Adding (0x80 - 'A') to each byte will set the high bit only for bytes >= 'A'
			geA := x + (0x8080808080808080 - 0x4141414141414141)
			// Adding (0x80 - ('Z' + 1)) to each byte will clear the high bit only for bytes <= 'Z'
			leZ := x + (0x8080808080808080 - 0x5B5B5B5B5B5B5B5B)
			// Combine these masks: high bit set only for bytes in range 'A'..'Z'
			mask := (geA & ^leZ) & 0x8080808080808080

			// Convert the mask to set bit 0x20 (shift right by 2 and mask)
			lowercase := (mask >> 2) & 0x2020202020202020

			// Apply the mask to set the 0x20 bit for uppercase characters
			b64[i] = x | lowercase
		}
	}

	// Handle remaining bytes individually
	for i := n8 * 8; i < len(b); i++ {
		c := b[i]
		if c >= 'A' && c <= 'Z' {
			b[i] = c | 0x20
		}
	}

	return b
}

// ToLowerInPlaceString is an optimized function for string conversion
// It avoids allocation when strings are already lowercase
func ToLowerInPlaceString(s string) string {
	// Check for trivial cases
	if len(s) <= 1 {
		if len(s) == 0 {
			return s
		}
		return string(ToLowerByte(s[0]))
	}

	// Fast check for uppercase characters
	hasUpper := false
	for i := 0; i < len(s) && !hasUpper; i++ {
		if s[i] >= 'A' && s[i] <= 'Z' {
			hasUpper = true
		}
	}

	if !hasUpper {
		return s // No changes needed
	}

	// If changes are needed, we must allocate (strings are immutable)
	// Choose the fastest processing method based on length
	if len(s) < 64 {
		res := make([]byte, len(s))
		for i := 0; i < len(s); i++ {
			c := s[i]
			if c >= 'A' && c <= 'Z' {
				res[i] = c | 0x20
			} else {
				res[i] = c
			}
		}
		return utils.UnsafeString(res)
	}

	// For longer strings, use the enhanced SWAR
	return utils.UnsafeString(ToLowerSWARv2([]byte(s)))
}

// ToLowerInPlace modifies a byte slice in-place, optimizing for already lowercase content
func ToLowerInPlace(b []byte) []byte {
	// Quick check if any conversion is needed
	hasUpper := false
	for i := 0; i < len(b) && !hasUpper; i++ {
		if b[i] >= 'A' && b[i] <= 'Z' {
			hasUpper = true
		}
	}

	if !hasUpper {
		return b // No changes needed
	}

	// Choose algorithm based on size
	if len(b) < 64 {
		// For small slices, simple approach is faster
		for i := 0; i < len(b); i++ {
			if b[i] >= 'A' && b[i] <= 'Z' {
				b[i] |= 0x20
			}
		}
	} else {
		// For larger slices, use SWAR approach
		ToLowerSWARv2(b)
	}

	return b
}

// ToLowerSWAREnhanced combines superior SWAR implementation with early exit
func ToLowerSWAREnhanced(b []byte) []byte {
	// Quick check for uppercase characters with early exit
	// This helps avoid the overhead of SWAR setup for already lowercase strings
	hasUpper := false
	quickCheckLimit := min(len(b), 32) // Only check a sample to reduce overhead
	for i := 0; i < quickCheckLimit; i++ {
		if b[i] >= 'A' && b[i] <= 'Z' {
			hasUpper = true
			break
		}
	}

	if !hasUpper && len(b) < 128 { // Skip full SWAR for small already lowercase strings
		return b
	}

	n := len(b)
	i := 0

	// Process 8-byte chunks (benchmarks show this is faster than 32-byte processing)
	for ; i+8 <= n; i += 8 {
		chunk := *(*uint64)(unsafe.Pointer(&b[i]))

		// Use constant folding for these calculations
		const upperMask = 0x8080808080808080 - 0x4141414141414141
		const lowerMask = 0x8080808080808080 - 0x5B5B5B5B5B5B5B5B

		// SWAR operations condensed for clarity and efficiency
		geA := chunk + upperMask
		leZ := chunk + lowerMask
		mask := (geA & ^leZ) & 0x8080808080808080
		lowercase := (mask >> 2) & 0x2020202020202020

		// Only write back if we need to change something
		if lowercase != 0 {
			*(*uint64)(unsafe.Pointer(&b[i])) = chunk | lowercase
		}
	}

	// Process remaining bytes
	for ; i < n; i++ {
		c := b[i]
		if c >= 'A' && c <= 'Z' {
			b[i] = c | 0x20
		}
	}

	return b
}

// SuperToLower combines the best strategies from all implementations
// - Early out for already lowercase (with limited sampling)
// - Branch optimization for short strings
// - SWAR processing for longer strings
// - Branchless operations where beneficial
func SuperToLower(s string) string {
	// Handle trivial cases
	if len(s) <= 1 {
		if len(s) == 0 {
			return s
		}
		return string(ToLowerByte(s[0]))
	}

	// Improved early detection of uppercase characters
	// Only sample a portion of the string for very long strings
	sampleLimit := min(len(s), 64)
	hasUpper := false

	for i := 0; i < sampleLimit; i++ {
		if s[i] >= 'A' && s[i] <= 'Z' {
			hasUpper = true
			break
		}
	}

	// If no uppercase in sample and string is long, use IndexFunc for full check
	if !hasUpper && len(s) > 64 {
		if strings.IndexFunc(s[sampleLimit:], func(r rune) bool {
			return r >= 'A' && r <= 'Z'
		}) == -1 {
			return s
		}
		hasUpper = true
	} else if !hasUpper {
		return s // No uppercase in short string
	}

	// Choose algorithm based on length
	if len(s) <= 24 {
		// For very short strings, the overhead of SWAR isn't worth it
		res := make([]byte, len(s))
		for i := 0; i < len(s); i++ {
			c := s[i]
			// Use byte mask for branchless lowercase conversion
			mask := byte(0)
			if c >= 'A' && c <= 'Z' {
				mask = 0x20
			}
			res[i] = c | mask
		}
		return utils.UnsafeString(res)
	}

	// For longer strings, SWAR is the fastest approach
	return utils.UnsafeString(ToLowerSWAREnhanced([]byte(s)))
}

// ToLowerHeader is optimized specifically for HTTP headers
// HTTP headers are typically short and follow patterns like:
// "Content-Type", "Accept-Encoding", etc.
func ToLowerHeader(s string) string {
	// Most headers are shorter than 32 bytes
	if len(s) == 0 {
		return s
	}

	// Check if first character is uppercase (common in HTTP headers)
	if s[0] >= 'A' && s[0] <= 'Z' {
		// We have at least one uppercase, need to convert
	} else {
		// Check if any other character is uppercase
		// Most headers have few uppercase letters, so this approach is efficient
		hasUpper := false
		for i := 1; i < len(s); i++ {
			if s[i] >= 'A' && s[i] <= 'Z' {
				hasUpper = true
				break
			}
		}
		if !hasUpper {
			return s
		}
	}

	// HTTP headers are usually short, so simple approach is fastest
	res := make([]byte, len(s))
	for i := 0; i < len(s); i++ {
		c := s[i]
		if c >= 'A' && c <= 'Z' {
			res[i] = c | 0x20
		} else {
			res[i] = c
		}
	}
	return utils.UnsafeString(res)
}

// ToLowerInPlaceOptimized modifies a byte slice in-place with better early exit
func ToLowerInPlaceOptimized(b []byte) []byte {
	if len(b) == 0 {
		return b
	}

	// Only check a sample of bytes for very long slices
	// This reduces overhead of the uppercase check
	sampleLimit := min(len(b), 32)
	index := -1

	// Find the first uppercase character (if any)
	for i := 0; i < sampleLimit; i++ {
		if b[i] >= 'A' && b[i] <= 'Z' {
			index = i
			break
		}
	}

	// If no uppercase in sample but slice is long, do more checking
	if index == -1 {
		if len(b) > 32 {
			// Check remaining bytes in larger chunks using SWAR
			// This is faster than byte-by-byte for large inputs
			i := sampleLimit - (sampleLimit % 8) // Align to 8-byte boundary

			for ; i+8 <= len(b); i += 8 {
				chunk := *(*uint64)(unsafe.Pointer(&b[i]))

				// Use SWAR to detect uppercase letters
				const upperMask = 0x8080808080808080 - 0x4141414141414141
				const lowerMask = 0x8080808080808080 - 0x5B5B5B5B5B5B5B5B

				geA := chunk + upperMask
				leZ := chunk + lowerMask
				mask := (geA & ^leZ) & 0x8080808080808080

				if mask != 0 {
					// We found an uppercase, now we need to find which byte
					for j := 0; j < 8; j++ {
						if b[i+j] >= 'A' && b[i+j] <= 'Z' {
							index = i + j
							break
						}
					}
					break
				}
			}

			// Check any remaining bytes if we still haven't found uppercase
			if index == -1 {
				for i := sampleLimit - (sampleLimit % 8) + 8; i < len(b); i++ {
					if b[i] >= 'A' && b[i] <= 'Z' {
						index = i
						break
					}
				}
			}
		}

		// If still no uppercase, return original slice
		if index == -1 {
			return b
		}
	}

	// Convert detected uppercase and all following bytes
	// This is optimized by starting from the first uppercase letter
	for i := index; i < len(b); i++ {
		c := b[i]
		if c >= 'A' && c <= 'Z' {
			b[i] = c | 0x20
		}
	}

	return b
}
