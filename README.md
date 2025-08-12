# letsgo

## LRU Cache

This project contains a thread-safe LRU (Least Recently Used) cache implementation in Go.

### Test-First Development

The development of the LRU cache followed a test-first approach:

1.  **Interface Definition**: Defined the `Cache` interface with `Put` and `Get` methods.
2.  **Test Case Design**: Wrote table-driven tests in `lru_test.go` to cover the main scenarios:
    *   **Eviction**: Ensured the least recently used item is removed when the cache is full.
    *   **Hit**: Verified that accessing an item makes it "recently used."
    *   **Concurrency**: Added tests using goroutines and the `-race` flag to ensure the implementation is thread-safe.
3.  **Implementation**: Wrote the `lruCache` struct and its methods to pass all the defined tests.
4.  **Verification**: Ran `go test -v -cover -race ./lru/...` to confirm that all tests passed, the race detector found no issues, and the test coverage met the requirement (≥70%).

Let's go