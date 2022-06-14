/*
 * @lc app=leetcode.cn id=88 lang=rust
 *
 * [88] 合并两个有序数组
 */

use crate::*;

// @lc code=start
impl Solution {
    pub fn merge(nums1: &mut Vec<i32>, m: i32, nums2: &mut Vec<i32>, n: i32) {
        let res = nums1;
        let nums1 = Vec::from(&res[0..m as usize]);

        let (mut it_res, mut it_nums1, mut it_nums2) = (res.iter_mut(), nums1.iter(), nums2.iter());

        let (mut n_1, mut n_2) = (it_nums1.next(), it_nums2.next());

        loop {
            *it_res.next().unwrap() = match (n_1, n_2) {
                (None, None) => return,
                (None, Some(n2)) => {
                    n_2 = it_nums2.next();
                    *n2
                }
                (Some(n1), None) => {
                    n_1 = it_nums1.next();
                    *n1
                }
                (Some(n1), Some(n2)) => {
                    if n1 < n2 {
                        n_1 = it_nums1.next();
                        *n1
                    } else {
                        n_2 = it_nums2.next();
                        *n2
                    }
                }
            }
        }
    }
}
// @lc code=end
