# 最长连续序列 (Longest Consecutive Sequence) - LeetCode 题解

## 题目描述

给定一个未排序的整数数组 `nums`，找出数字连续的最长序列（不要求序列元素在原数组中连续）的长度。要求设计并实现时间复杂度为 O(n) 的算法解决此问题。

**示例 1：**
```
输入：nums = [100,4,200,1,3,2]
输出：4
解释：最长数字连续序列是 [1, 2, 3, 4]。它的长度为 4。
```

**示例 2：**
```
输入：nums = [0,3,7,2,5,8,4,6,0,1]
输出：9
解释：最长数字连续序列是 [0,1,2,3,4,5,6,7,8]。它的长度为 9。
```

**示例 3：**
```
输入：nums = [1,0,1,2]
输出：3
解释：最长数字连续序列是 [0,1,2]。它的长度为 3。
```

## 解题思路

### 方法一：哈希集合（推荐）

1. **哈希集合预处理**：
    - 将所有数字存入哈希集合中，实现O(1)时间复杂度的查找
    - 去重处理，避免重复数字的影响

2. **寻找连续序列**：
    - 遍历数组中的每个数字
    - 对于每个数字，检查它是否是某个连续序列的起点（即num-1不在集合中）
    - 如果是起点，则向后查找连续的数字，计算当前连续序列的长度
    - 更新最长连续序列的长度

3. **复杂度分析**：
    - 时间复杂度：O(n)，每个数字最多被访问两次（一次在哈希集合中，一次在连续序列查找中）
    - 空间复杂度：O(n)，用于存储哈希集合

### 方法二：排序法（不满足O(n)要求）

1. **排序数组**：
    - 先对数组进行排序
    - 然后遍历排序后的数组，寻找最长连续序列

2. **缺点**：
    - 时间复杂度为O(n log n)，不满足题目要求
    - 仅作为对比思路提及

## Go 代码实现

```go
func longestConsecutive(nums []int) int {
    if len(nums) == 0 {
        return 0
    }
    
    // 创建哈希集合存储所有数字
    numSet := make(map[int]bool)
    for _, num := range nums {
        numSet[num] = true
    }
    
    longestStreak := 0
    
    // 遍历哈希集合中的每个数字
    for num := range numSet {
        // 检查当前数字是否是某个连续序列的起点
        if !numSet[num-1] {
            currentNum := num
            currentStreak := 1
            
            // 向后查找连续的数字
            for numSet[currentNum+1] {
                currentNum++
                currentStreak++
            }
            
            // 更新最长连续序列长度
            if currentStreak > longestStreak {
                longestStreak = currentStreak
            }
        }
    }
    
    return longestStreak
}

func main() {
    // 示例测试
    nums1 := []int{100, 4, 200, 1, 3, 2}
    result1 := longestConsecutive(nums1)
    fmt.Println("Longest consecutive sequence length:", result1) // 输出: 4
    
    nums2 := []int{0, 3, 7, 2, 5, 8, 4, 6, 0, 1}
    result2 := longestConsecutive(nums2)
    fmt.Println("Longest consecutive sequence length:", result2) // 输出: 9
    
    nums3 := []int{1, 0, 1, 2}
    result3 := longestConsecutive(nums3)
    fmt.Println("Longest consecutive sequence length:", result3) // 输出: 3
}
```

## 测试用例

```go
func TestLongestConsecutive(t *testing.T) {
    tests := []struct {
        input []int
        want  int
    }{
        {[]int{100, 4, 200, 1, 3, 2}, 4},
        {[]int{0, 3, 7, 2, 5, 8, 4, 6, 0, 1}, 9},
        {[]int{1, 0, 1, 2}, 3},
        {[]int{}, 0},
        {[]int{1}, 1},
        {[]int{1, 3, 5, 7, 9}, 1},
        {[]int{1, 2, 3, 4, 5}, 5},
        {[]int{5, 4, 3, 2, 1}, 5},
    }

    for _, tt := range tests {
        got := longestConsecutive(tt.input)
        if got != tt.want {
            t.Errorf("longestConsecutive(%v) = %d, want %d", tt.input, got, tt.want)
        }
    }
}
```

## 复杂度分析

1. **时间复杂度**：O(n)
    - 创建哈希集合：O(n)
    - 遍历哈希集合：每个数字最多被访问两次（一次在哈希集合中，一次在连续序列查找中）
    - 总体时间复杂度为线性

2. **空间复杂度**：O(n)
    - 需要额外的哈希集合存储所有数字

## 优化思路

1. **并行处理**：
    - 对于大规模数据，可以考虑将数组分割后并行处理
    - 需要处理边界条件和合并结果

2. **内存优化**：
    - 如果数字范围有限，可以使用位图代替哈希集合
    - 减少内存使用，提高缓存命中率

3. **流式处理**：
    - 对于无法一次性加载到内存的超大数据集
    - 可以使用外部排序和分段处理技术

## 总结

这道题考察了对哈希数据结构的灵活运用，关键在于如何高效地判断连续序列的起点和计算序列长度。通过使用哈希集合，我们可以在O(n)时间复杂度内解决问题，这比排序法更高效。掌握这种利用空间换时间的思路对解决类似问题很有帮助。

## 扩展思考

1. 如果要求返回最长的连续序列本身而不仅仅是长度，该如何修改代码？
2. 如何修改算法以处理包含重复数字的情况（虽然当前解法已经处理了）？
3. 如果数字范围非常大但稀疏，如何进一步优化空间复杂度？