# 字母异位词分组 (Group Anagrams) - LeetCode 题解

## 题目描述

给你一个字符串数组，请你将字母异位词组合在一起。可以按任意顺序返回结果列表。

字母异位词是指由相同字母重新排列形成的不同单词。

**示例 1：**
```
输入: strs = ["eat", "tea", "tan", "ate", "nat", "bat"]
输出: [["bat"],["nat","tan"],["ate","eat","tea"]]
解释：
- "bat" 没有其他字母异位词
- "nat" 和 "tan" 是字母异位词
- "ate"、"eat" 和 "tea" 是字母异位词
```

**示例 2：**
```
输入: strs = [""]
输出: [[""]]
```

**示例 3：**
```
输入: strs = ["a"]
输出: [["a"]]
```

## 解题思路

### 方法一：排序 + 哈希表（推荐）

1. **排序键法**：
    - 对每个字符串进行排序，将排序后的字符串作为哈希表的键
    - 将原始字符串作为值存储在对应键的列表中
    - 这样所有字母异位词都会被分到同一组

2. **复杂度分析**：
    - 时间复杂度：O(n * k log k)，其中 n 是字符串数量，k 是最长字符串长度
    - 空间复杂度：O(n * k)，需要存储所有字符串

### 方法二：计数 + 哈希表（优化空间）

1. **计数键法**：
    - 统计每个字符串中各个字母的出现次数
    - 将计数结果转换为字符串作为哈希表的键（如 "a1b2c3"）
    - 这种方法避免了排序，但构建键的过程可能较慢

2. **复杂度分析**：
    - 时间复杂度：O(n * k)，其中 n 是字符串数量，k 是最长字符串长度
    - 空间复杂度：O(n * k)

## Go 代码实现

```go
func groupAnagrams(strs []string) [][]string {
    groups := make(map[string][]string)
    for _, str := range strs {
        // 将字符串转换为字符切片并排序
        s := strings.Split(str, "")
        sort.Strings(s)
        sortedStr := strings.Join(s, "")
        // 将原始字符串添加到对应的分组中
        groups[sortedStr] = append(groups[sortedStr], str)
    }
    
    // 将分组结果转换为二维切片
    result := make([][]string, 0, len(groups))
    for _, group := range groups {
        result = append(result, group)
    }
    return result
}

func main() {
    // 示例测试
    strs := []string{"eat", "tea", "tan", "ate", "nat", "bat"}
    result := groupAnagrams(strs)
    fmt.Println(result) // 输出: [["bat"] ["nat" "tan"] ["ate" "eat" "tea"]]
}
```

## 测试用例

```go
func TestGroupAnagrams(t *testing.T) {
    tests := []struct {
        input []string
        want  [][]string
    }{
        {
            input: []string{"eat", "tea", "tan", "ate", "nat", "bat"},
            want:  [][]string{{"bat"}, {"nat", "tan"}, {"ate", "eat", "tea"}},
        },
        {
            input: []string{""},
            want:  [][]string{{""}},
        },
        {
            input: []string{"a"},
            want:  [][]string{{"a"}},
        },
    }

    for _, tt := range tests {
        got := groupAnagrams(tt.input)
        if !isEqual(got, tt.want) {
            t.Errorf("groupAnagrams(%v) = %v, want %v", tt.input, got, tt.want)
        }
    }
}

// 辅助函数：比较两个二维字符串切片是否相等（顺序无关）
func isEqual(a, b [][]string) bool {
    if len(a) != len(b) {
        return false
    }
    
    aMap := make(map[string]bool)
    for _, group := range a {
        sort.Strings(group)
        key := strings.Join(group, ",")
        aMap[key] = true
    }
    
    for _, group := range b {
        sort.Strings(group)
        key := strings.Join(group, ",")
        if !aMap[key] {
            return false
        }
    }
    return true
}
```

## 复杂度分析

1. **时间复杂度**：
    - 对每个字符串进行排序：O(k log k)，其中 k 是字符串长度
    - 总共需要处理 n 个字符串：O(n * k log k)

2. **空间复杂度**：
    - 需要存储所有字符串：O(n * k)
    - 哈希表存储分组信息：O(n)

## 优化思路

1. **计数法优化**：
    - 可以使用长度为26的数组记录每个字母的出现次数
    - 将计数数组转换为字符串作为哈希键（如 "#1#2#0..."）
    - 这种方法避免了排序，时间复杂度降为 O(n * k)

2. **并行处理**：
    - 对于大规模数据，可以考虑并行处理字符串排序和分组

## 总结

这道题考察了对哈希表的灵活运用，关键在于如何设计合适的哈希键来识别字母异位词。排序法直观易懂，计数法则在性能上可能更有优势。掌握这种分组问题的解决思路对处理类似问题（如分组统计、分类等）很有帮助。

## 扩展思考

1. 如果字符串包含Unicode字符而非仅小写字母，该如何修改解决方案？
2. 如何优化算法以处理超长字符串（如长度超过1000）？
3. 如果要求分组结果按组大小排序，该如何实现？