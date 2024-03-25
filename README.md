# gofiber_tolower_test

This repository contains a simple benchmark test for the `ToLower` function in the `strings` package and the `ToLower` function in the `utils` package of the `fiber` web framework.

It is used to compare the performance of the `ToLower` function in the `strings` package, the `ToLower` function in the `utils` package of the `fiber` web framework, and a hybrid approach that borrows the idea to check if the string contains any uppercase characters from the `ToLower` function in the standard `strings` package and if it does, then it uses the approach from the `utils` package of the `fiber` web framework.

## Running the benchmark

To run the benchmark, execute the following command:

```bash
go test -bench=.
```

## Sample benchmark results

| Benchmark                              | Runs     | ns/op    | B/op   | allocs/op |
|----------------------------------------|----------|----------|--------|-----------|
| BenchmarkUtilsToLowerLowerCase-24      | 6193549  | 186.3    | 72     | 3         |
| BenchmarkUtilsToLowerMixedCase-24      | 6520929  | 183.1    | 72     | 3         |
| BenchmarkStringsToLowerLowerCase-24    | 9680548  | 121.3    | 0      | 0         |
| BenchmarkStringsToLowerMixedCase-24    | 3345157  | 357.8    | 72     | 3         |
| BenchmarkHybribToLowerLowerCase-24     | 13733870 | 86.15    | 0      | 0         |
| BenchmarkHybribToLowerMixedCase-24     | 4968349  | 237.4    | 72     | 3         |
