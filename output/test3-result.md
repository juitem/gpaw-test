```go
package main

import "testing"

func TestClamp(t *testing.T) {
    tests := []struct {
        v, min, max, want int
    }{
        {5, 1, 10, 5},   // 범위 내
        {0, 1, 10, 1},   // min보다 작음
        {15, 1, 10, 10}, // max보다 큼
        {1, 1, 10, 1},   // min 경계
        {10, 1, 10, 10}, // max 경계
        {-5, -10, -1, -5}, // 음수 범위 내
        {3, 3, 3, 3},    // min == max
    }

    for _, tc := range tests {
        got := clamp(tc.v, tc.min, tc.max)
        if got != tc.want {
            t.Errorf("clamp(%d, %d, %d) = %d, want %d", tc.v, tc.min, tc.max, got, tc.want)
        }
    }
}
```

테이블 드리븐 방식으로 작성했습니다. 범위 내, 하한/상한 초과, 경계값, 음수 범위, `min == max` 케이스를 커버합니다.
