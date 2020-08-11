# keep-learning
每日写写代码练练手

## 经验 
```
1 读题清晰
2 确定边界
3 数组确定是索引还是通过索引取得的值
4 排序问题注意边界

递归问题：
    1 返回值
    2 递归做了什么
    3 终止条件
滑动窗口：
    1、当移动right扩大窗口，即加入字符时，应该更新哪些数据？
    
    2、什么条件下，窗口应该暂停扩大，开始移动left缩小窗口？
    
    3、当移动left缩小窗口，即移出字符时，应该更新哪些数据？
    
    4、我们要的结果应该在扩大窗口时还是缩小窗口时进行更新？

动态规划：
     1、确定 base case
     2、确定「状态」，也就是原问题和子问题中会变化的变量。
     3、确定「选择」，也就是导致「状态」产生变化的行为。
     4、明确 dp 函数/数组的定义。

回溯问题：
    1、路径：也就是已经做出的选择。
    2、选择列表：也就是你当前可以做的选择。
    3、结束条件：也就是到达决策树底层，无法再做选择的条件。

    def backtrack(路径, 选择列表):
        if 满足结束条件:
            result.add(路径)
            return
    
        for 选择 in 选择列表:
            做选择
            backtrack(路径, 选择列表)
            撤销选择

```

## todolist
[最长回文字符串](https://leetcode-cn.com/problems/longest-palindromic-substring/solution/zui-chang-hui-wen-zi-chuan-by-leetcode-solution/)

```
1 动态规划： coins v1 v2 v3  还需要在练习
```
