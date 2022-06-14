/*
 * @lc app=leetcode.cn id=53 lang=rust
 *
 * [53] 最大子数组和
 */

use crate::structs::*;

// @lc code=start
impl Solution {
    pub fn max_sub_array(nums: Vec<i32>) -> i32 {
        let (mut prev, mut max_ans) = (0, nums[0]);
        for num in nums {
            prev = num.max(prev + num);
            max_ans = max_ans.max(prev);
        }
        max_ans
    }
}
// @lc code=end

