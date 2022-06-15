/*
 * @lc app=leetcode.cn id=240 lang=rust
 *
 * [240] 搜索二维矩阵 II
 */

struct Solution;

// @lc code=start
#[derive(Clone)]
struct Pos {
    col: i32,
    row: i32,
}

impl Solution {
    fn search(matrix: &Vec<Vec<i32>>, start: Pos, end: Pos, target: i32) -> bool {
        if start.col > end.col || start.row > end.row {
            return false;
        }

        let mid = Pos {
            col: (start.col + end.col) / 2,
            row: (start.row + end.row) / 2,
        };

        let mid_val = matrix[mid.row as usize][mid.col as usize];

        if mid_val == target {
            true
        } else if start.col == end.col && end.row == start.row {
            false
        } else if target < mid_val {
            Solution::search(
                matrix,
                start.clone(),
                Pos {
                    col: mid.col - 1,
                    row: end.row,
                },
                target,
            ) || Solution::search(
                matrix,
                Pos {
                    col: mid.col,
                    row: start.row,
                },
                Pos {
                    col: end.col,
                    row: mid.row - 1,
                },
                target,
            )
        } else {
            Solution::search(
                matrix,
                Pos {
                    col: mid.col + 1,
                    row: start.row,
                },
                end.clone(),
                target,
            ) || Solution::search(
                matrix,
                Pos {
                    col: start.col,
                    row: mid.row + 1,
                },
                Pos {
                    col: mid.col,
                    row: end.row,
                },
                target,
            )
        }
    }

    pub fn search_matrix(matrix: Vec<Vec<i32>>, target: i32) -> bool {
        Solution::search(
            &matrix,
            Pos { col: 0, row: 0 },
            Pos {
                col: matrix[0].len() as i32 - 1,
                row: matrix.len() as i32 - 1,
            },
            target,
        )
    }
}
// @lc code=end

#[test]
fn test() {
    let matrix = vec![
        vec![1, 4, 7, 11, 15],
        vec![2, 5, 8, 12, 19],
        vec![3, 6, 9, 16, 22],
        vec![10, 13, 14, 17, 24],
        vec![18, 21, 23, 26, 30],
    ];

    assert_eq!(Solution::search_matrix(matrix.clone(), 5), true);
    assert_eq!(Solution::search_matrix(matrix.clone(), 20), false);
    assert_eq!(Solution::search_matrix(vec![vec![-5]], -10), false);
    assert_eq!(Solution::search_matrix(vec![vec![5], vec![6]], 6), true);

    let matrix = vec![
        vec![1, 2, 3, 4, 5],
        vec![6, 7, 8, 9, 10],
        vec![11, 12, 13, 14, 15],
        vec![16, 17, 18, 19, 20],
        vec![21, 22, 23, 24, 25],
    ];
    assert_eq!(Solution::search_matrix(matrix, 5), true);

    let matrix = vec![vec![1], vec![3], vec![5]];
    assert_eq!(Solution::search_matrix(matrix, 1), true);
}
