/*
 * @lc app=leetcode.cn id=162 lang=rust
 *
 * [162] 寻找峰值
 */

use crate::Solution;

// @lc code=start
impl Solution {
    fn val(nums: &[i32], i: i32) -> i64 {
        if i < 0 || i >= nums.len() as i32 {
            i64::MIN
        } else {
            *nums.get(i as usize).unwrap() as i64
        }
    }

    fn find(nums: &[i32], l: i32, r: i32) -> i32 {
        if l > r {
            return -1;
        }

        let m = (l + r) / 2;

        if Solution::val(nums, m as i32 - 1) < Solution::val(nums, m)
            && Solution::val(nums, m) > Solution::val(nums, m as i32 + 1)
        {
            return m as i32;
        }

        let i_l = Solution::find(&nums, l, m - 1);
        let i_r = Solution::find(&nums, m + 1, r);

        if i_l != -1 {
            i_l
        } else if i_r != -1 {
            i_r
        } else {
            -1
        }
    }

    pub fn find_peak_element(nums: Vec<i32>) -> i32 {
        Solution::find(&nums, 0, nums.len() as i32 - 1)
    }
}
// @lc code=end
