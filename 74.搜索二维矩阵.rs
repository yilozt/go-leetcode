/*
 * @lc app=leetcode.cn id=74 lang=rust
 *
 * [74] 搜索二维矩阵
 */

use crate::*;

// @lc code=start
trait Index {
    fn index(&self, i: i32) -> i32;
}

impl Index for Vec<Vec<i32>> {
    fn index(&self, i: i32) -> i32 {
        let y = i / self[0].len() as i32;
        let x = i % self[0].len() as i32;
        self[y as usize][x as usize]
    }
}


impl Solution {
    pub fn search_matrix(matrix: Vec<Vec<i32>>, target: i32) -> bool {
        let (mut l, mut r) = (0, (matrix.len() * matrix[0].len() - 1) as i32);
        while l <= r {
            let m = (l + r) / 2;
            if matrix.index(m) == target {
                return true;
            }
            if target < matrix.index(m) {
                r = m - 1;
            } else {
                l = m + 1;
            }
        }
        false
    }
}
// @lc code=end

#[test]
fn test() {
    assert_eq!(Solution::search_matrix(vec![vec![1]], 2), false)
}
