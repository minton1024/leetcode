# 盛最多水的容器 (Container With Most Water) - LeetCode 题解

## 题目描述

给定一个长度为 `n` 的整数数组 `height`，其中 `n` 条垂线的两个端点分别是 `(i, 0)` 和 `(i, height[i])`。找出两条线，使得它们与 x 轴共同构成的容器可以容纳最多的水，返回容器可以储存的最大水量。

**注意**：容器不能倾斜。

**示例 1：**
```
输入：[1,8,6,2,5,4,8,3,7]
输出：49
解释：图中垂直线代表输入数组 [1,8,6,2,5,4,8,3,7]。在此情况下，容器能够容纳水的最大值为 49。
```
![图1](question_11.jpg)

**示例 2：**
```
输入：height = [1,1]
输出：1
```

## 解题思路

### 方法一：双指针法（最优解）

1. **初始化指针**：
    - `left` 指针从数组开头（0）开始
    - `right` 指针从数组末尾（len(height)-1）开始

2. **计算面积**：
    - 当前面积 = min(height[left], height[right]) * (right - left)
    - 更新最大面积

3. **移动指针**：
    - 移动高度较小的指针（因为移动高度较大的指针不可能得到更大的面积）
    - 如果两指针高度相同，可以移动任意一个

4. **终止条件**：
    - 当 `left` 和 `right` 指针相遇时结束

5. **复杂度分析**：
    - 时间复杂度：O(n)，只需一次遍历
    - 空间复杂度：O(1)，只使用了常数空间

### 方法二：暴力法（不推荐）

1. **双重循环**：
    - 遍历所有可能的线对组合
    - 计算每对线构成的容器面积
    - 记录最大面积

2. **缺点**：
    - 时间复杂度：O(n²)
    - 空间复杂度：O(1)
    - 不适用于大规模数据

## Go 代码实现

### 双指针法实现

```go
func maxArea(height []int) int {
    maxArea := 0
    left, right := 0, len(height)-1
    
    for left < right {
        h := min(height[left], height[right])
        width := right - left
        area := h * width
        
        if area > maxArea {
            maxArea = area
        }
        
        if height[left] < height[right] {
            left++
        } else {
            right--
        }
    }
    
    return maxArea
}

func min(a, b int) int {
    if a < b {
        return a
    }
    return b
}

func main() {
    height := []int{1, 8, 6, 2, 5, 4, 8, 3, 7}
    result := maxArea(height)
    println("Max area:", result) // 输出: 49
}
```

### 暴力法实现（对比）

```go
func maxAreaBruteForce(height []int) int {
    maxArea := 0
    for i := 0; i < len(height); i++ {
        for j := i + 1; j < len(height); j++ {
            h := min(height[i], height[j])
            width := j - i
            area := h * width
            if area > maxArea {
                maxArea = area
            }
        }
    }
    return maxArea
}
```

## 测试用例

```go
func TestMaxArea(t *testing.T) {
    tests := []struct {
        input  []int
        expect int
    }{
        {[]int{1, 8, 6, 2, 5, 4, 8, 3, 7}, 49},
        {[]int{1, 1}, 1},
        {[]int{4, 3, 2, 1, 4}, 16},
        {[]int{1, 2, 1}, 2},
        {[]int{}, 0},
        {[]int{1}, 0},
    }

    for _, tt := range tests {
        got := maxArea(tt.input)
        if got != tt.expect {
            t.Errorf("maxArea(%v) = %d, want %d", tt.input, got, tt.expect)
        }
    }
}
```

## 复杂度分析

1. **双指针法**：
    - 时间复杂度：O(n)，只需一次遍历
    - 空间复杂度：O(1)，只使用了常数空间

2. **暴力法**：
    - 时间复杂度：O(n²)，双重循环
    - 空间复杂度：O(1)

## 算法正确性证明

双指针法的正确性基于以下观察：

1. **面积公式**：面积 = min(height[left], height[right]) * (right - left)
2. **移动策略**：每次移动高度较小的指针
    - 因为移动高度较大的指针不可能得到更大的面积（宽度减小，高度受限于较小值）
    - 只有移动较小指针才有可能找到更高的边界

## 优化思路

1. **提前终止**：
    - 当剩余宽度 * 最大可能高度 <= 当前最大面积时，可以提前终止

2. **跳过重复计算**：
    - 如果移动指针后高度不比之前高，可以继续移动直到找到更高的边界

## 总结

这道题是典型的双指针应用问题，展示了如何通过巧妙的指针移动策略将O(n²)的暴力解法优化为O(n)的高效解法。关键在于理解为什么可以安全地移动高度较小的指针而不遗漏可能的更大面积。

掌握这种双指针技巧对解决类似的数组和容器类问题（如接雨水问题、两数之和等）非常有帮助。

## 扩展思考

1. 如果要求找出所有可能的容器组合（而不仅仅是最大面积），该如何修改算法？
2. 如何修改算法以处理高度可能为负数的情况？
3. 如果容器可以倾斜（即可以形成梯形），该如何计算最大面积？