## 시간 복잡도 분석

### 현재 코드: O(n²)

이중 루프로 모든 쌍을 비교하므로 `n*(n-1)/2`번 비교 수행.

**추가 문제점:**
- 같은 숫자가 3개 이상이면 중복 결과 포함 (e.g. `[1,1,1]` → `[1,1,1]`)
- 공간 복잡도: O(n) (result slice)

---

## 개선안

### 방법 1 — HashMap 사용: O(n) 시간, O(n) 공간

```go
func findDuplicates(nums []int) []int {
    seen := make(map[int]bool)
    result := []int{}
    for _, n := range nums {
        if seen[n] {
            result = append(result, n)
            delete(seen, n) // 중복 결과 방지
        } else {
            seen[n] = true
        }
    }
    return result
}
```

한 번의 순회로 완료. `delete`로 3회 이상 등장하는 숫자의 중복 추가를 방지.

---

### 방법 2 — 정렬 후 인접 비교: O(n log n) 시간, O(1) 추가 공간

```go
func findDuplicates(nums []int) []int {
    sort.Ints(nums)
    result := []int{}
    for i := 1; i < len(nums); i++ {
        if nums[i] == nums[i-1] && (i < 2 || nums[i] != nums[i-2]) {
            result = append(result, nums[i])
        }
    }
    return result
}
```

입력 배열 수정이 허용될 때 적합. 추가 메모리 불필요.

---

## 비교 요약

| 방법 | 시간 | 추가 공간 | 입력 보존 |
|------|------|-----------|-----------|
| 원본 (이중 루프) | **O(n²)** | O(1) | ✅ |
| HashMap | **O(n)** | O(n) | ✅ |
| 정렬 후 비교 | **O(n log n)** | O(1) | ❌ |

**권장:** 대부분의 경우 HashMap 방법(O(n))이 최적.
