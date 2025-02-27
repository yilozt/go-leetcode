## 仓库介绍

此为 [yilozt](https://github.com/yilozt) 的刷题题解仓库，fork 自 [amtoaer/go-leetcode](https://github.com/amtoaer/go-leetcode)。自`2022/6/13`开始进入每日更新状态（大概）。

## 题目组成

1. 文件名有序号

    题目来自[力扣](https://leetcode-cn.com)官方题库，文件序号和名称为力扣题目的编号和标题。

2. 文件名无序号

    包含一些常用算法的实现及各种测试（如`字节跳动7天刷题挑战`）中的独有题目。

## 刷题顺序

开始时按照力扣官方分类刷题，后期主要参考[算法模板](https://greyireland.gitbook.io/algorithm-pattern/)中的推荐刷题顺序。

## 补充说明

虽然该仓库内题解均通过了测试，但其中大部分只是个人的解决思路，并没有在参考官方题解后做相应修改。即：**只保证正确，不保证最优**。

- **编辑器**：Vscode
  插件：[LeetCode](https://marketplace.visualstudio.com/items?itemName=LeetCode.vscode-leetcode) + Go + rust-analyzer 
- **Debug**
  1. go: 需要自己添加 main 函数，之后使用 `go run` 运行或者按下 F5
  2. rust: 添加 `#[test]` 测试用例，在 `src/lib.rs` 里添加模块
