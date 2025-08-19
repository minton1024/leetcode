# LeetCode 算法题解 - Go 语言实现
本仓库使用 Go 语言实现 LeetCode 算法题的解题思路和代码，包含详细的题解说明、复杂度分析和测试用例。每个题目都有独立的 Go 文件实现，并附有清晰的 README 说明。

## 目录结构

```
.
├── README.md                   # 仓库总说明
├── LICENSE                     # 开源许可证
├── hashtable/                   # 哈希表相关题目
│   ├── twoSum/                # 两数之和
│   ├── groupAnagrams/         # 字母异位词分组
│   └── longestConsecutive/    # 最长连续序列
├── twoPointers/               # 双指针相关题目
│   ├── moveZeroes/            # 移动零
│   ├── containerWithMostWater/   # 盛最多水的容器
│   ├── threeSum/              # 三数之和
│   └── trappingRainWater/    # 接雨水
├── slidingWindow/             # 滑动窗口相关题目
│   ├── longestSubstringWithoutRepeating/ # 无重复字符的最长子串
│   ├── findAllAnagrams/      # 找到字符串中所有字母异位词
│   ├── subarraySumEqualsK/  # 和为 K 的子数组
│   ├── slidingWindowMaximum/ # 滑动窗口最大值
│   └── minimumWindowSubstring/ # 最小覆盖子串
├── array/                      # 数组相关题目
│   ├── maximumSubarray/       # 最大子数组和
│   ├── mergeIntervals/        # 合并区间
│   ├── rotateArray/           # 轮转数组
│   ├── productExceptSelf/    # 除自身以外数组的乘积
│   └── firstMissingPositive/ # 缺失的第一个正数
├── matrix/                     # 矩阵相关题目
│   ├── setMatrixZeroes/      # 矩阵置零
│   ├── spiralMatrix/          # 螺旋矩阵
│   ├── rotateImage/           # 旋转图像
│   └── search2DMatrixII/   # 搜索二维矩阵 II
├── linkedList/                # 链表相关题目
│   ├── intersectionTwoLinkedLists/ # 相交链表
│   ├── reverseLinkedList/    # 反转链表
│   ├── palindromeLinkedList/ # 回文链表
│   ├── linkedListCycle/      # 环形链表
│   ├── linkedListCycleII/   # 环形链表 II
│   ├── mergeTwoSortedLists/ # 合并两个有序链表
│   ├── addTwoNumbers/        # 两数相加
│   ├── removeNthNodeFromEnd/ # 删除链表的倒数第 N 个结点
│   ├── swapNodesInPairs/    # 两两交换链表中的节点
│   ├── reverseNodesInKGroup/ # K 个一组翻转链表
│   ├── copyListWithRandomPointer/ # 随机链表的复制
│   ├── sortList/              # 排序链表
│   ├── mergeKSortedLists/   # 合并 K 个升序链表
│   └── lruCache/              # LRU 缓存
├── binaryTree/                # 二叉树相关题目
│   ├── inorderTraversal/      # 二叉树的中序遍历
│   ├── maximumDepth/           # 二叉树的最大深度
│   ├── invertBinaryTree/     # 翻转二叉树
│   ├── symmetricTree/         # 对称二叉树
│   ├── diameterOfBinaryTree/ # 二叉树的直径
│   ├── levelOrderTraversal/  # 二叉树的层序遍历
│   ├── sortedArrayToBST/    # 将有序数组转换为二叉搜索树
│   ├── validateBST/           # 验证二叉搜索树
│   ├── kthSmallestInBST/    # 二叉搜索树中第 K 小的元素
│   ├── binaryTreeRightView/ # 二叉树的右视图
│   ├── flattenBinaryTree/    # 二叉树展开为链表
│   ├── constructBinaryTree/  # 从前序与中序遍历序列构造二叉树
│   ├── pathSumIII/           # 路径总和 III
│   ├── lowestCommonAncestor/ # 二叉树的最近公共祖先
│   └── binaryTreeMaxPathSum/ # 二叉树中的最大路径和
├── graph/                      # 图论相关题目
│   ├── numberOfIslands/      # 岛屿数量
│   ├── rottingOranges/        # 腐烂的橘子
│   ├── courseSchedule/        # 课程表
│   └── implementTrie/         # 实现 Trie (前缀树)
├── backtracking/               # 回溯相关题目
│   ├── permutations/           # 全排列
│   ├── subsets/                # 子集
│   ├── letterCombinations/    # 电话号码的字母组合
│   ├── combinationSum/        # 组合总和
│   ├── generateParentheses/   # 括号生成
│   ├── wordSearch/            # 单词搜索
│   ├── palindromePartitioning/ # 分割回文串
│   └── nQueens/               # N 皇后
├── binarySearch/              # 二分查找相关题目
│   ├── searchInsertPosition/ # 搜索插入位置
│   ├── search2DMatrix/       # 搜索二维矩阵
│   ├── findFirstLastPosition/ # 在排序数组中查找元素的第一个和最后一个位置
│   ├── searchInRotatedSortedArray/ # 搜索旋转排序数组
│   ├── findMinInRotatedSortedArray/ # 寻找旋转排序数组中的最小值
│   └── medianOfTwoSortedArrays/ # 寻找两个正序数组的中位数
├── stack/                      # 栈相关题目
│   ├── validParentheses/      # 有效的括号
│   ├── minStack/              # 最小栈
│   ├── decodeString/          # 字符串解码
│   ├── dailyTemperatures/     # 每日温度
│   └── largestRectangleInHistogram/ # 柱状图中最大的矩形
├── heap/                       # 堆相关题目
│   ├── kthLargestElement/    # 数组中的第K个最大元素
│   ├── topKFrequentElements/ # 前 K 个高频元素
│   └── findMedianFromDataStream/ # 数据流的中位数
├── greedy/                     # 贪心算法相关题目
│   ├── bestTimeToBuyAndSellStock/ # 买卖股票的最佳时机
│   ├── jumpGame/              # 跳跃游戏
│   ├── jumpGameII/           # 跳跃游戏 II
│   └── partitionLabels/        # 划分字母区间
├── dp/                         # 动态规划相关题目
│   ├── climbingStairs/        # 爬楼梯
│   ├── pascalsTriangle/       # 杨辉三角
│   ├── houseRobber/           # 打家劫舍
│   ├── perfectSquares/        # 完全平方数
│   ├── coinChange/            # 零钱兑换
│   ├── wordBreak/             # 单词拆分
│   ├── longestIncreasingSubsequence/ # 最长递增子序列
│   ├── maximumProductSubarray/ # 乘积最大子数组
│   ├── partitionEqualSubsetSum/ # 分割等和子集
│   └── longestValidParentheses/ # 最长有效括号
├── multiDp/                   # 多维动态规划相关题目
│   ├── uniquePaths/           # 不同路径
│   ├── minimumPathSum/       # 最小路径和
│   ├── longestPalindromicSubstring/ # 最长回文子串
│   ├── longestCommonSubsequence/ # 最长公共子序列
│   └── editDistance/          # 编辑距离
└── tricks/                     # 技巧类题目
    ├── singleNumber/          # 只出现一次的数字
    ├── majorityElement/       # 多数元素
    ├── sortColors/            # 颜色分类
    ├── nextPermutation/       # 下一个排列
    └── findDuplicate/         # 寻找重复数
```

## 使用说明

1. **查找题目**：根据题目类别在目录结构中查找对应的题目
2. **查看题解**：每个题目目录下都有详细的 README 说明，包含：
    - 题目描述
    - 解题思路
    - 复杂度分析
    - Go 代码实现
    - 测试用例
3. **运行代码**：可以直接运行每个题目目录下的 Go 文件

## 贡献指南

欢迎贡献你的 LeetCode 题解！请遵循以下步骤：

1. Fork 本仓库
2. 创建你的分支 (`git checkout -b feature/your-solution`)
3. 提交你的修改 (`git commit -am 'Add some solution'`)
4. 推送你的分支 (`git push origin feature/your-solution`)
5. 创建 Pull Request

## 许可证

本项目采用 MIT 许可证 - 详情请参阅 LICENSE 文件。

## 致谢

感谢 LeetCode 提供优秀的算法题库，感谢 Go 语言社区的支持！