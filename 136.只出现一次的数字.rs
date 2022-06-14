/*
 * @lc app=leetcode.cn id=136 lang=rust
 *
 * [136] 只出现一次的数字
 */

use crate::*;

// @lc code=start
impl Solution {
    pub fn single_number(nums: Vec<i32>) -> i32 {
        let mut hash_map = std::collections::HashMap::new();

        for n in nums {
            match hash_map.get(&n) {
                Some(_) => hash_map.remove(&n),
                None => hash_map.insert(n, 1),
            };
        }

        for (k, _) in hash_map {
            return k;
        }

        unreachable!()
    }
}
// @lc code=end

