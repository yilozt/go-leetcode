/*
 * @lc app=leetcode.cn id=200 lang=rust
 *
 * [200] 岛屿数量
 */

use crate::structs::*;

// @lc code=start
impl Solution {
    pub fn num_islands(mut grid: Vec<Vec<char>>) -> i32 {
        let (lc, lr) = (grid[0].len() as i32, grid.len() as i32);

        let mut found = 0;
        for c in 0..lc {
            for r in 0..lr {
                if Solution::fill_land(&mut grid, r, c) {
                    found += 1
                }
            }
        }

        found
    }

    fn fill_land(grid: &mut Vec<Vec<char>>, row: i32, col: i32) -> bool {
        let color = grid[row as usize][col as usize];

        if color == '0' || color == 'x' {
            return false;
        }

        let (lc, lr) = (grid[0].len() as i32, grid.len() as i32);
        let mut q = std::collections::VecDeque::new();

        q.push_back(row);
        q.push_back(col);

        while !q.is_empty() {
            let row = q.pop_front().unwrap();
            let col = q.pop_front().unwrap();

            for (dir_r, dir_c) in [(0, -1), (0, 1), (1, 0), (-1, 0)] {
                let n_row = row + dir_r;
                let n_col = col + dir_c;

                if n_col < 0 || n_row < 0 || n_col >= lc || n_row >= lr {
                    continue;
                }

                let color = grid[row as usize][col as usize];
                let n_color = grid[n_row as usize][n_col as usize];

                if n_color == color && n_color != 'x' {
                    q.push_back(n_row);
                    q.push_back(n_col);
                }
            }

            grid[row as usize][col as usize] = 'x';
        }

        return true;
    }
}
// @lc code=end
