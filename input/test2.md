다음 함수의 시간 복잡도를 분석하고 개선안을 제시해줘 (마크다운):

```go
func findDuplicates(nums []int) []int {
    result := []int{}
    for i := 0; i < len(nums); i++ {
        for j := i + 1; j < len(nums); j++ {
            if nums[i] == nums[j] {
                result = append(result, nums[i])
            }
        }
    }
    return result
}
```
