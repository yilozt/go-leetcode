/*
 * @lc app=leetcode.cn id=34 lang=rust
 *
 * [34] 在排序数组中查找元素的第一个和最后一个位置
 */

use crate::structs::*;

// @lc code=start
impl Solution {
    pub fn search_range(nums: Vec<i32>, target: i32) -> Vec<i32> {
        if nums.len() == 0 {
            return vec![-1, -1];
        }

        vec![
            search_left_bound(&nums, target),
            search_right_bound(&nums, target),
        ]
    }
}

fn search_left_bound(nums: &Vec<i32>, target: i32) -> i32 {
    let (mut l, mut r) = (0 as i32, nums.len() as i32 - 1);
    loop {
        let mid = (l + r) / 2;
        if target == nums[mid as usize] && (mid == l || nums[mid as usize - 1] != target) {
            return mid;
        }
        if l > r {
            return -1;
        }
        if target <= nums[mid as usize] {
            r = mid - 1;
        } else {
            l = mid + 1;
        }
    }
}
fn search_right_bound(nums: &Vec<i32>, target: i32) -> i32 {
    let (mut l, mut r) = (0 as i32, nums.len() as i32 - 1);
    loop {
        let mid = (l + r) / 2;
        if target == nums[mid as usize] && (mid == r || nums[mid as usize + 1] != target) {
            return mid;
        }
        if l > r {
            return -1;
        }
        if target < nums[mid as usize] {
            r = mid - 1;
        } else {
            l = mid + 1;
        }
    }
}
// @lc code=end

#[test]
fn test() {
    use crate::structs::*;
    let got = Solution::search_range(vec![1], 1);
    let want = [0, 0];

    assert_eq!(got, want)
}