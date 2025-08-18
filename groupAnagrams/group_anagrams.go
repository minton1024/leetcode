package main

import (
	"fmt"
	"sort"
	"strings"
)

/*
*
给你一个字符串数组，请你将 字母异位词 组合在一起。可以按任意顺序返回结果列表。

示例 1:
输入: strs = ["eat", "tea", "tan", "ate", "nat", "bat"]
输出: [["bat"],["nat","tan"],["ate","eat","tea"]]
解释：
在 strs 中没有字符串可以通过重新排列来形成 "bat"。
字符串 "nat" 和 "tan" 是字母异位词，因为它们可以重新排列以形成彼此。
字符串 "ate" ，"eat" 和 "tea" 是字母异位词，因为它们可以重新排列以形成彼此。

示例 2:
输入: strs = [""]
输出: [[""]]

示例 3:
输入: strs = ["a"]
输出: [["a"]]
*/
func groupAnagrams(strs []string) [][]string {
	groups := make(map[string][]string)
	for _, str := range strs {
		// Sort the string to find its anagram group
		s := strings.Split(str, "")
		sort.Strings(s)
		sortedStr := strings.Join(s, "")
		groups[sortedStr] = append(groups[sortedStr], str)
	}
	result := make([][]string, 0, len(groups))
	for _, group := range groups {
		result = append(result, group)
	}
	return result
}

func main() {
	// 示例测试
	strs1 := []string{"eat", "tea", "tan", "ate", "nat", "bat"}
	result1 := groupAnagrams(strs1)
	fmt.Println(result1)
}
