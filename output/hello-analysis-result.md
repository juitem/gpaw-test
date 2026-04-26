## `ExpandHome` 분석

### 잠재적 문제점

**1. 오류 무시**
```go
home, _ := os.UserHomeDir()
```
`UserHomeDir()`가 실패하면 `home`은 빈 문자열(`""`) — 결과가 `/original/path` 대신 `/path`처럼 잘못된 경로가 됨. 호출자가 오류를 알 방법 없음.

**2. Windows 비호환**
`"~/"` 하드코딩은 Windows의 `~\` 경로를 처리 못 함. `filepath.Join` 미사용도 OS별 구분자 문제 야기.

**3. `~` 단독 처리 안 됨**
`p == "~"`인 경우 홈 디렉토리를 반환해야 하지만 그냥 `"~"` 그대로 반환.

---

### 개선안

```go
func ExpandHome(p string) (string, error) {
    if p != "~" && !strings.HasPrefix(p, "~/") {
        return p, nil
    }
    home, err := os.UserHomeDir()
    if err != nil {
        return "", fmt.Errorf("expand home: %w", err)
    }
    if p == "~" {
        return home, nil
    }
    return filepath.Join(home, p[2:]), nil
}
```

**변경 요약**
| 항목 | 기존 | 개선 |
|------|------|------|
| 오류 처리 | 무시 (`_`) | `error` 반환 |
| `~` 단독 | 미처리 | 홈 반환 |
| 경로 결합 | 문자열 `+` | `filepath.Join` |
| 패턴 검사 | `p[:2]` 슬라이싱 | `strings.HasPrefix` |
