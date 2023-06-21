# stringx

The `stringx` package provides optimized string operations in Go, including an efficient prefix matching function. It leverages bitwise operations to achieve improved performance.

## Usage

1. Utilize the `HasPrefixFast` function to perform prefix matching:
```go
s := "example string"
prefix := "example"

if stringx.HasPrefixFast(s, prefix) {
...
}
```

## Functions

### HasPrefixFast

The `HasPrefixFast` function efficiently checks if a string `s` has the specified prefix. It employs bitwise operations and avoids unnecessary calls to `runtime.memequal`, resulting in faster execution.

```go
func HasPrefixFast(s, prefix string) bool
```

- `s`: The target string to check for the prefix.
- `prefix`: The prefix to match against `s`.
- Returns: `true` if `s` has the specified prefix, `false` otherwise.

## Benchmarks

Benchmarks were conducted to compare the performance of the `stringx` package against the standard Go string operations.

```
BenchmarkHasPrefixFastLongStringLongPrefix-12      	297485900	         3.942 ns/op
BenchmarkHasPrefixStdLongStringLongPrefix-12       	315866076	         3.682 ns/op
BenchmarkHasPrefixFastLongStringShortPrefix-12     	815244889	         1.453 ns/op
BenchmarkHasPrefixStdLongStringShortPrefix-12      	456260694	         2.514 ns/op
BenchmarkHasPrefixFastShortStringLongPrefix-12     	1000000000	         0.9196 ns/op
BenchmarkHasPrefixStdShortStringLongPrefix-12      	1000000000	         0.4695 ns/op
BenchmarkHasPrefixFastShortStringShortPrefix-12    	866156751	         1.400 ns/op
BenchmarkHasPrefixStdShortStringShortPrefix-12     	452262331	         2.539 ns/op
BenchmarkHasPrefixFastInvalidPrefix-12             	896777925	         1.339 ns/op
BenchmarkHasPrefixStdInvalidPrefix-12              	451083560	         2.557 ns/op
```

These benchmarks compare the performance of the `HasPrefixFast` function from the `stringx` package with the standard Go `HasPrefix` function. The results indicate the time taken per operation (`ns/op`) for different scenarios involving long and short strings, as well as valid and invalid prefixes.

Based on the benchmark results, the following observations can be made:

- For long strings with long prefixes, both `HasPrefixFast` and `HasPrefixStd` perform similarly, with `HasPrefixStd` being slightly faster in this specific case.
- For long strings with short prefixes, `HasPrefixFast` demonstrates significantly better performance compared to `HasPrefixStd`, being approximately twice as fast.
- For short strings with long prefixes, `HasPrefixFast` outperforms `HasPrefixStd`, showing noticeable performance improvements.
- For short strings with short prefixes, `HasPrefixFast` and `HasPrefixStd` exhibit similar performance.

These results highlight the optimized implementation of `HasPrefixFast` in the `stringx` package, particularly in scenarios involving long strings and short prefixes, which [minimizes](https://godbolt.org/z/1ce5dvq63) the use of `runtime.memequal` and reduces overhead, resulting in faster execution..


## License

This project is licensed under the [MIT License](LICENSE).
