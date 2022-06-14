/*
 * @lc app=leetcode.cn id=33 lang=rust
 *
 * [33] 搜索旋转排序数组
 */

use crate::structs::*;

// @lc code=start
use std::ops::Index;

impl Solution {
    pub fn search(nums: Vec<i32>, target: i32) -> i32 {
        RotatedArr::new(nums).binary_search(target)
    }
}

struct RotatedArr {
    offset: i32,
    arr: Vec<i32>,
}

impl RotatedArr {
    #[inline(always)]
    pub fn addr(&self, i: i32) -> usize {
        (i + self.offset) as usize % self.arr.len()
    }

    pub fn new(arr: Vec<i32>) -> Self {
        let (mut l, mut r) = (0, arr.len() as i32 - 1);

        while arr[l as usize] > arr[r as usize] {
            let m = (l + r) / 2;
            if arr[m as usize] > arr[r as usize] {
                l = m + 1;
                continue;
            }
            if arr[l as usize] > arr[m as usize] {
                r = m;
                continue;
            }
            unreachable!("Array should be rotated in this scope")
        }

        RotatedArr { offset: l, arr }
    }

    pub fn binary_search(&self, target: i32) -> i32 {
        let (mut l, mut r) = (0, self.arr.len() as i32 - 1);

        while l <= r {
            let m = (l + r) / 2;
            if self[m] == target {
                return self.addr(m) as i32;
            }
            if target < self[m] {
                r = m - 1
            } else {
                l = m + 1
            }
        }

        return -1;
    }
}

impl Index<i32> for RotatedArr {
    type Output = i32;
    #[inline(always)]
    fn index(&self, index: i32) -> &Self::Output {
        &self.arr[self.addr(index)]
    }
}

// @lc code=end

#[test]
fn test() {
    assert_eq!(Solution::search(vec![5, 1, 3], 5), 0)
}