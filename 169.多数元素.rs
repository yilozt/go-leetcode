/*
 * @lc app=leetcode.cn id=169 lang=rust
 *
 * [169] 多数元素
 */

use crate::structs::*;

// @lc code=start
impl Solution {
    pub fn majority_element(nums: Vec<i32>) -> i32 {
        let len = nums.len();
        let mut hash_map = std::collections::HashMap::<i32, i32>::with_capacity(len);
        for n in nums {
            let c = hash_map.entry(n).or_insert(0);
            *c += 1;
            if *c > len as i32 / 2 {
                return n;
            }
        }
        unreachable!()
    }
}
// @lc code=end

#[test]
fn test() {
    assert_eq!(Solution::majority_element(vec![3, 2, 3]), 3)
}
