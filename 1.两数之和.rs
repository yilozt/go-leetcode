/*
 * @lc app=leetcode.cn id=1 lang=rust
 *
 * [1] 两数之和
 */

use crate::structs::*;

// @lc code=start
impl Solution {
    pub fn two_sum(nums: Vec<i32>, target: i32) -> Vec<i32> {
        let mut hash_map = std::collections::HashMap::<i32, i32>::new();

        for (i, num) in nums.into_iter().enumerate() {
            if let Some(rest_i) = hash_map.get(&(target - num)) {
                return vec![i as i32, *rest_i];
            }
            hash_map.insert(num, i as i32);
        }

        unreachable!("Only one answer")
    }
}
// @lc code=end
