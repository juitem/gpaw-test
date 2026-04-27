아래 코드의 단위 테스트를 Go로 작성해줘:

```go
func clamp(v, min, max int) int {
    if v < min { return min }
    if v > max { return max }
    return v
}
```
