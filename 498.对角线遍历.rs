/*
 * @lc app=leetcode.cn id=498 lang=rust
 *
 * [498] 对角线遍历
 */

use crate::*;

// @lc code=start
struct Pos {
    x: i32,
    y: i32,
}

enum Dir {
    LeftDown,
    RightUp,
}

impl Dir {
    fn toggle(&mut self) {
        *self = match self {
            Dir::LeftDown => Dir::RightUp,
            Dir::RightUp => Dir::LeftDown,
        }
    }

    fn next_pos(&mut self, pos: Pos, bound: &Pos) -> Pos {
        let mut out_of_bound = false;
        let mut next_pos = match self {
            Dir::LeftDown => Pos {
                x: pos.x - 1,
                y: pos.y + 1,
            },
            Dir::RightUp => Pos {
                x: pos.x + 1,
                y: pos.y - 1,
            },
        };

        // 检查边界并修正位置
        if next_pos.y < 0 {
            next_pos.x = pos.x + 1;
            next_pos.y = 0;
            out_of_bound = true;
        }
        if next_pos.x < 0 {
            next_pos.x = 0;
            next_pos.y = pos.y + 1;
            out_of_bound = true;
        }
        if next_pos.x == bound.x {
            next_pos.x = bound.x - 1;
            next_pos.y = pos.y + 1;
            out_of_bound = true;
        }
        if next_pos.y == bound.y {
            next_pos.x = pos.x + 1;
            next_pos.y = bound.y - 1;
            out_of_bound = true
        }

        if out_of_bound {
            self.toggle()
        }

        return next_pos;
    }
}

impl Solution {
    pub fn find_diagonal_order(mat: Vec<Vec<i32>>) -> Vec<i32> {
        let bound = Pos {
            x: mat[0].len() as i32,
            y: mat.len() as i32,
        };
        let mut pos = Pos { x: 0, y: 0 };
        let mut dir = Dir::RightUp;

        let mut res = Vec::with_capacity((bound.x * bound.y) as usize);

        while pos.x < bound.x && pos.y < bound.y {
            res.push(mat[pos.y as usize][pos.x as usize]);
            pos = dir.next_pos(pos, &bound)
        }
        res
    }
}
// @lc code=end
