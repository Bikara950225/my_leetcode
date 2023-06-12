# 652. 寻找重复的子树

https://leetcode.cn/problems/find-duplicate-subtrees/description/

思路：遍历树的节点，通过一些字符串规则解析成string（注意区分好左右），通过Map判断重复序列，重复出现2次即判断为重复子树，记录当前的节点到结果集中