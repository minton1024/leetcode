# 两数之和 (Two Sum) - LeetCode 题解

## 题目描述

给定一个整数数组 `nums` 和一个整数目标值 `target`，请你在该数组中找出和为目标值 `target` 的那两个整数，并返回它们的数组下标。

你可以假设每种输入只会对应一个答案，并且你不能使用两次相同的元素。你可以按任意顺序返回答案。

**示例 1：**
```
输入：nums = [2,7,11,15], target = 9
输出：[0,1]
解释：因为 nums[0] + nums[1] == 9 ，返回 [0, 1] 。
```

**示例 2：**
```
输入：nums = [3,2,4], target = 6
输出：[1,2]
```

**示例 3：**
```
输入：nums = [3,3], target = 6
输出：[0,1]
```

## 解题思路

### 方法一：暴力枚举
最直观的方法是使用双重循环遍历数组中的每一个数 `nums[i]`，然后查找是否存在另一个数 `nums[j]` 使得 `nums[i] + nums[j] == target`。

- 时间复杂度：O(n²)
- 空间复杂度：O(1)

### 方法二：哈希表（推荐）
我们可以使用哈希表来优化查找过程，将时间复杂度降低到 O(n)：

1. 创建一个哈希表 `numMap` 用于存储数组元素及其对应的索引
2. 遍历数组 `nums`，对于每个元素 `nums[i]`：
    - 计算补数 `complement = target - nums[i]`
    - 检查 `complement` 是否存在于哈希表中
        - 如果存在，返回 `[numMap[complement], i]`
        - 如果不存在，将当前元素及其索引存入哈希表 `numMap[nums[i]] = i`

这种方法利用了哈希表的快速查找特性，将查找时间从 O(n) 降低到 O(1)。

- 时间复杂度：O(n)
- 空间复杂度：O(n)

## Go 代码实现

```go
func twoSum(nums []int, target int) []int {
    numMap := make(map[int]int)
    for i := 0; i < len(nums); i++ {
        complement := target - nums[i]
        if index, found := numMap[complement]; found {
            return []int{index, i}
        }
        numMap[nums[i]] = i
    }
    return nil
}

func main() {
    nums := []int{2, 7, 11, 15}
    target := 9
    result := twoSum(nums, target)
    fmt.Println(result) // 输出: [0 1]
}
```

## 复杂度分析

- **时间复杂度**：O(n)，我们只遍历了包含 n 个元素的列表一次。在哈希表中查找补数的时间复杂度是 O(1)。
- **空间复杂度**：O(n)，哈希表最多需要存储 n 个元素。

## 测试用例

```go
func TestTwoSum(t *testing.T) {
    tests := []struct {
        nums   []int
        target int
        want   []int
    }{
        {[]int{2, 7, 11, 15}, 9, []int{0, 1}},
        {[]int{3, 2, 4}, 6, []int{1, 2}},
        {[]int{3, 3}, 6, []int{0, 1}},
        {[]int{1, 2, 3, 4}, 10, nil},
    }

    for _, tt := range tests {
        got := twoSum(tt.nums, tt.target)
        if !reflect.DeepEqual(got, tt.want) {
            t.Errorf("twoSum(%v, %d) = %v, want %v", tt.nums, tt.target, got, tt.want)
        }
    }
}
```

## 总结

这道题是 LeetCode 的第一题，也是经典的哈希表应用问题。通过使用哈希表，我们可以将时间复杂度从 O(n²) 优化到 O(n)，这在处理大规模数据时尤为重要。掌握这种空间换时间的思路对解决其他算法问题也很有帮助。

## 扩展思考

1. 如果题目要求返回所有可能的解（而不仅是一个），该如何修改代码？
2. 如果数组已经排序，是否有更优的解法？（可以使用双指针法）
3. 如何修改代码以处理输入数据中可能有重复元素的情况？