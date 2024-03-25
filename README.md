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
| BenchmarkHybribToLowerLowerCase-12     | 48373012.75 | 23.72   | 0      | 0         |
| BenchmarkHybribToLowerMixedCase-12     | 17928221.5  | 65.115  | 72     | 3         |


### Intel Xeon X5675

| Benchmark                              | Runs     | ns/op    | B/op   | allocs/op |
|----------------------------------------|----------|----------|--------|-----------|
| BenchmarkUtilsToLowerLowerCase-24      | 6282977.75 | 187.925 | 72     | 3         |
| BenchmarkUtilsToLowerMixedCase-24      | 6369836.25 | 187.525 | 72     | 3         |
| BenchmarkStringsToLowerLowerCase-24    | 9776561.5  | 121.625 | 0      | 0         |
| BenchmarkStringsToLowerMixedCase-24    | 3359413.75 | 353.325 | 72     | 3         |
| BenchmarkHybribToLowerLowerCase-24     | 15111381.25 | 76.72   | 0      | 0         |
| BenchmarkHybribToLowerMixedCase-24     | 5217725   | 227.8    | 72     | 3         |

## Conclusion

Hybrid approach ported to gofiber utils package as `IfToLower` and `IfToUpper` functions.

PR [#76](https://github.com/gofiber/utils/pull/76)