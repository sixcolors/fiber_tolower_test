# Expiremental Branch for Gaby:

# High-Performance ASCII ToLower for Go

This repository is a deep dive into creating and benchmarking highly optimized ASCII `ToLower` functions for Go, with a focus on performance for the GoFiber web framework. The goal was to explore various techniques, including Go-based SWAR (SIMD-Within-A-Register) and architecture-specific assembly, to find the absolute fastest implementations.

## Impressions

After extensive experimentation and benchmarking, I've written some optimal functions that serve different use cases:

0. **For `string` conversion: `UltimateToLower`**
    This function is the culmination of all optimizations, combining the best techniques from both Go and assembly implementations. It is designed to be the fastest possible `string` to `string` conversion function in Go, while still being allocation-aware.

1.  **For `string` conversion: `OptimalToLowerAsm`**
    This is the recommended function for general-purpose `string` to `string` conversion. It is both allocation-aware and extremely fast.
    - It performs a quick check to see if the string contains any uppercase characters. If not, it returns the original string immediately, incurring **zero allocations**.
    - If conversion is needed, it uses the highly optimized assembly core (`ToLowerAsm`) for maximum performance.

2.  **For in-place `[]byte` conversion: `ToLowerInPlaceOptimizedAsm`**
    This is the fastest possible function for converting a byte slice in-place, making it ideal for performance-critical internal operations where no allocations are allowed.
    - It performs a fast check for uppercase characters and returns immediately if none are found.
    - If conversion is needed, it calls the `ToLowerAsm` assembly routine to modify the slice in-place.

3. **For `string` conversion without assembly: `OptimalToLower`**
   This function takes a **SWAR (SIMD-Within-A-Register)** approach in pure Go, making it suitable for environments where assembly is not an option. It is still highly efficient and performs well for most use cases.

4. **For in-place `[]byte` conversion without assembly: `ToLowerInPlaceOptimized`**
   This function is the pure Go equivalent of `ToLowerInPlaceOptimizedAsm`, using a SWAR technique to convert a byte slice in-place without any assembly code. It is designed for environments where assembly cannot be used, while still providing good performance.

The other functions are still in the project and perhaps the techniques can be combined or serve for further optimizations and learning.

## Assembly Implementations

This project provides hand-tuned assembly for the two most common server architectures:

-   **ARM64 (Apple Silicon, etc.):** Implemented using a **SWAR (SIMD-Within-A-Register)** technique. This branchless approach uses 64-bit general-purpose registers to process 8 bytes at a time.
-   **AMD64 (Intel/AMD):** Implemented using **SSE2** instructions. This version processes 16-byte chunks and features loop unrolling and data prefetching for maximum throughput.

# Non-Updated README.md from main branch

# fiber_tolower_test

This repository contains a simple benchmark test for the `ToLower` function in the `strings` package and the `ToLower` function in the `utils` package of the `fiber` web framework.

It is used to compare the performance of the `ToLower` function in the `strings` package, the `ToLower` function in the `utils` package of the `fiber` web framework, and a hybrid approach that borrows the idea to check if the string contains any uppercase characters from the `ToLower` function in the standard `strings` package and if it does, then it uses the approach from the `utils` package of the `fiber` web framework.

## Running the benchmark

To run the benchmark, execute the following command:

```bash
go test -v -run=^$ -bench=B -benchmem -count=4
```

## Sample benchmark results

### Apple M2 Max

| Benchmark                              | Runs     | ns/op    | B/op   | allocs/op |
|----------------------------------------|----------|----------|--------|-----------|
| BenchmarkUtilsToLowerLowerCase-12      | 20728723.25 | 55.4275 | 72     | 3         |
| BenchmarkUtilsToLowerMixedCase-12      | 21040876.5  | 55.43   | 72     | 3         |
| BenchmarkStringsToLowerLowerCase-12    | 20979590.25 | 55.2825 | 0      | 0         |
| BenchmarkStringsToLowerMixedCase-12    | 8903027.75  | 132.175 | 72     | 3         |
| BenchmarkHybridToLowerLowerCase-12     | 48373012.75 | 23.72   | 0      | 0         |
| BenchmarkHybridToLowerMixedCase-12     | 17928221.5  | 65.115  | 72     | 3         |


### Intel Xeon X5675

| Benchmark                              | Runs     | ns/op    | B/op   | allocs/op |
|----------------------------------------|----------|----------|--------|-----------|
| BenchmarkUtilsToLowerLowerCase-24      | 6282977.75 | 187.925 | 72     | 3         |
| BenchmarkUtilsToLowerMixedCase-24      | 6369836.25 | 187.525 | 72     | 3         |
| BenchmarkStringsToLowerLowerCase-24    | 9776561.5  | 121.625 | 0      | 0         |
| BenchmarkStringsToLowerMixedCase-24    | 3359413.75 | 353.325 | 72     | 3         |
| BenchmarkHybridToLowerLowerCase-24     | 15111381.25 | 76.72   | 0      | 0         |
| BenchmarkHybridToLowerMixedCase-24     | 5217725   | 227.8    | 72     | 3         |

## Conclusion

Hybrid approach ported to gofiber utils package as `IfToLower` and `IfToUpper` functions.

PR [#76](https://github.com/gofiber/utils/pull/76)

# Update

## Enhanced Performance of `HybridToLower` Function

The `HybridToLower` function for has been optimized to improve performance, particularly when handling HTTP header values such as the "Referer" header, which can contain lengthy URLs, where the initial part of the string is often in lowercase with uppercase characters appearing later in the path component.

The optimization strategy involves utilizing the same index variable across both loops within the function. The first loop scans the string for any uppercase characters. If it finds any, it breaks, and the index variable retains its current position. The second loop then commences from this index, eliminating the need to reiterate over the initial part of the string already confirmed to be in lowercase. This approach significantly enhances performance for long strings that commence with an extended sequence of lowercase characters, such as URLs.

## Updated Benchmarks

The benchmarks have been enhanced to include a "Referer" header and a Lorem Ipsum text string. The "Referer" header string starts with a lowercase string, followed by a lengthy path component with uppercase characters. The Lorem Ipsum text string begins with an uppercase character.

## Sample Benchmark Results

The following results, generated by the `benchstat` tool, compare the performance of the `HybridToLower` function before and after the optimization. 

## Apple M2 Max

| Benchmark                   | Old (sec/op) | New (sec/op) | Change |
|-----------------------------|--------------|--------------|--------|
| UtilsToLowerLowerCase-12    | 355.7n ± 2%  | 354.3n ± 0%  | -0.39% |
| UtilsToLowerMixedCase-12    | 356.1n ± 0%  | 354.9n ± 0%  | -0.31% |
| StringsToLowerLowerCase-12  | 615.0n ± 0%  | 615.1n ± 0%  | ~      |
| StringsToLowerMixedCase-12  | 1.054µ ± 0%  | 1.054µ ± 0%  | ~      |
| HybridToLowerLowerCase-12   | 282.8n ± 2%  | 282.4n ± 0%  | ~      |
| HybridToLowerMixedCase-12   | 382.2n ± 1%  | 353.8n ± 0%  | -7.46% |
| Geomean                     | 455.0n       | 448.5n       | -1.42% |

| Benchmark                   | Old (B/op)   | New (B/op)   | Change |
|-----------------------------|--------------|--------------|--------|
| UtilsToLowerLowerCase-12    | 808.0 ± 0%   | 808.0 ± 0%   | ~      |
| UtilsToLowerMixedCase-12    | 808.0 ± 0%   | 808.0 ± 0%   | ~      |
| StringsToLowerLowerCase-12  | 0.000 ± 0%   | 0.000 ± 0%   | ~      |
| StringsToLowerMixedCase-12  | 808.0 ± 0%   | 808.0 ± 0%   | ~      |
| HybridToLowerLowerCase-12   | 0.000 ± 0%   | 0.000 ± 0%   | ~      |
| HybridToLowerMixedCase-12   | 808.0 ± 0%   | 808.0 ± 0%   | ~      |
| Geomean                     | ~            | ~            | +0.00% |

| Benchmark                   | Old (allocs/op) | New (allocs/op) | Change |
|-----------------------------|-----------------|-----------------|--------|
| UtilsToLowerLowerCase-12    | 5.000 ± 0%      | 5.000 ± 0%      | ~      |
| UtilsToLowerMixedCase-12    | 5.000 ± 0%      | 5.000 ± 0%      | ~      |
| StringsToLowerLowerCase-12  | 0.000 ± 0%      | 0.000 ± 0%      | ~      |
| StringsToLowerMixedCase-12  | 5.000 ± 0%      | 5.000 ± 0%      | ~      |
| HybridToLowerLowerCase-12   | 0.000 ± 0%      | 0.000 ± 0%      | ~      |
| HybridToLowerMixedCase-12   | 5.000 ± 0%      | 5.000 ± 0%      | ~      |
| Geomean                     | ~               | ~               | +0.00% |

## Intel Xeon X5675

| Benchmark                  | old.txt (sec/op) | new.txt (sec/op) | vs base (sec/op) |
|----------------------------|------------------|------------------|------------------|
| UtilsToLowerLowerCase-24   | 1.133µ ± 1%      | 1.137µ ± 1%      | ~ (p=0.083 n=10) |
| UtilsToLowerMixedCase-24   | 1.131µ ± 2%      | 1.129µ ± 0%      | ~ (p=0.269 n=10) |
| StringsToLowerLowerCase-24 | 2.377µ ± 0%      | 2.392µ ± 1%      | ~ (p=0.084 n=10) |
| StringsToLowerMixedCase-24 | 3.388µ ± 0%      | 3.363µ ± 0%      | -0.75% (p=0.000 n=10) |
| HybridToLowerLowerCase-24  | 747.4n ± 0%      | 750.1n ± 1%      | ~ (p=0.289 n=10) |
| HybridToLowerMixedCase-24  | 1.281µ ± 0%      | 1.146µ ± 0%      | -10.54% (p=0.000 n=10) |
| geomean                    | 1.465µ           | 1.439µ           | -1.76%           |

| Benchmark                  | old.txt (B/op) | new.txt (B/op) | vs base (B/op) |
|----------------------------|----------------|----------------|----------------|
| UtilsToLowerLowerCase-24   | 808.0 ± 0%     | 808.0 ± 0%     | ~ (p=1.000 n=10) ¹ |
| UtilsToLowerMixedCase-24   | 808.0 ± 0%     | 808.0 ± 0%     | ~ (p=1.000 n=10) ¹ |
| StringsToLowerLowerCase-24 | 0.000 ± 0%     | 0.000 ± 0%     | ~ (p=1.000 n=10) ¹ |
| StringsToLowerMixedCase-24 | 808.0 ± 0%     | 808.0 ± 0%     | ~ (p=1.000 n=10) ¹ |
| HybridToLowerLowerCase-24  | 0.000 ± 0%     | 0.000 ± 0%     | ~ (p=1.000 n=10) ¹ |
| HybridToLowerMixedCase-24  | 808.0 ± 0%     | 808.0 ± 0%     | ~ (p=1.000 n=10) ¹ |
| geomean                    |                |                | +0.00%          |

| Benchmark                  | old.txt (allocs/op) | new.txt (allocs/op) | vs base (allocs/op) |
|----------------------------|---------------------|---------------------|---------------------|
| UtilsToLowerLowerCase-24   | 5.000 ± 0%          | 5.000 ± 0%          | ~ (p=1.000 n=10) ¹  |
| UtilsToLowerMixedCase-24   | 5.000 ± 0%          | 5.000 ± 0%          | ~ (p=1.000 n=10) ¹  |
| StringsToLowerLowerCase-24 | 0.000 ± 0%          | 0.000 ± 0%          | ~ (p=1.000 n=10) ¹  |
| StringsToLowerMixedCase-24 | 5.000 ± 0%          | 5.000 ± 0%          | ~ (p=1.000 n=10) ¹  |
| HybridToLowerLowerCase-24  | 0.000 ± 0%          | 0.000 ± 0%          | ~ (p=1.000 n=10) ¹  |
| HybridToLowerMixedCase-24  | 5.000 ± 0%          | 5.000 ± 0%          | ~ (p=1.000 n=10) ¹  |
| geomean                    |                     |                     | +0.00%              |