/*
 * @lc app=leetcode.cn id=217 lang=rust
 *
 * [217] 存在重复元素
 */

use crate::Solution;

// @lc code=start
impl Solution {
    pub fn contains_duplicate(mut nums: Vec<i32>) -> bool {
        nums.sort();
        for i in 0..nums.len().checked_sub(1).unwrap_or(0) {
            if nums[i] == nums[i + 1] {
                return true;
            }
        }
        false
    }
}
// @lc code=end

