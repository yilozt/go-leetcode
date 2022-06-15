/*
 * @lc app=leetcode.cn id=719 lang=rust
 *
 * [719] 找出第 K 小的数对距离
 */

use core::num;

struct Solution;

// @lc code=start

impl Solution {
    fn lower_bound(nums: &[i32], k: i32) -> i32 {
        let len = nums.len();
        let (mut l, mut r) = (0 as i32, len as i32- 1);

        while l <= r {
            let mid = (l + r) / 2;
            if nums[mid as usize] < k {
                l = mid + 1;
            } else {
                r = mid - 1;
            }
        }
        return l;
    }

    pub fn smallest_distance_pair(mut nums: Vec<i32>, k: i32) -> i32 {
        nums.sort();

        let len = nums.len();
        let (mut l, mut r) = (0, nums[len - 1]);

        while l <= r {
            let mid = (l + r) / 2;

            // 搜索距离小于等于 mid 的数对个数
            let mut cnt = 0;
            for j in 0..len {
                let i = Self::lower_bound(&nums[0..=j], nums[j] - mid);
                cnt += j as i32 - i;
            }

            if cnt >= k {
                r = mid - 1;
            } else {
                l = mid + 1;
            }
        }

        l
    }
}
// @lc code=end
