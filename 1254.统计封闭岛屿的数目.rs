/*
 * @lc app=leetcode.cn id=1254 lang=rust
 *
 * [1254] 统计封闭岛屿的数目
 */

struct Solution;

// @lc code=start
impl Solution {
    fn travel_island(grid: &mut Vec<Vec<i32>>, col: i32, row: i32) -> i32 {
        let mut closed = 0;

        if grid[row as usize][col as usize] != 0 {
            return closed;
        }

        closed = 1;

        let row_len = grid.len() as i32;
        let col_len = grid[0].len() as i32;
        let dirs = [(1, 0), (-1, 0), (0, -1), (0, 1)];

        let mut q = std::collections::VecDeque::new();
        q.push_back(col);
        q.push_back(row);

        while !q.is_empty() {
            let col = q.pop_front().unwrap();
            let row = q.pop_front().unwrap();

            if grid[row as usize][col as usize] == 1 {
                continue;
            }

            for (dir_col, dir_row) in dirs {
                let next_col = dir_col + col;
                let next_row = dir_row + row;

                if next_col < 0 || next_row < 0 || next_col >= col_len || next_row >= row_len {
                    closed = 0;
                    continue;
                }

                if grid[next_row as usize][next_col as usize] == 0 {
                    q.push_back(next_col);
                    q.push_back(next_row);
                }
            }

            grid[row as usize][col as usize] = 1;
        }
        closed
    }

    pub fn closed_island(mut grid: Vec<Vec<i32>>) -> i32 {
        let mut count = 0;
        for i in 0..grid.len() {
            for j in 0..grid[0].len() {
                count += Solution::travel_island(&mut grid, j as i32, i as i32)
            }
        }
        count
    }
}
// @lc code=end
