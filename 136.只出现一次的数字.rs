/*
 * @lc app=leetcode.cn id=136 lang=rust
 *
 * [136] 只出现一次的数字
 */

use crate::structs::*;

// @lc code=start
impl Solution {
    pub fn single_number(nums: Vec<i32>) -> i32 {
        let mut res = 0;
        for n in nums {
            res ^= n;
        }
        res
    }
}
// @lc code=end

