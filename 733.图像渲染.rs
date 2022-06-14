/*
 * @lc app=leetcode.cn id=733 lang=rust
 *
 * [733] 图像渲染
 */

use crate::structs::*;

// @lc code=start

impl Solution {
    pub fn flood_fill(image: Vec<Vec<i32>>, sr: i32, sc: i32, new_color: i32) -> Vec<Vec<i32>> {
        // 类似与二叉树的层序遍历，使用队列来缓存元素 （广度优先）
        let mut q = std::collections::VecDeque::new();
        let mut image = image;
        let (lc, lr) = (image[0].len() as i32, image.len() as i32);
        q.push_back((sc, sr));

        while !q.is_empty() {
            let (sc, sr) = q.pop_front().unwrap();
            let color = image[sr as usize][sc as usize];
            for (dir_c, dir_r) in [(0, -1), (0, 1), (1, 0), (-1, 0)] {
                let (sc_n, sr_n) = (sc + dir_c, sr + dir_r);
                
                let check_bounds = sc_n >= 0 && sc_n < lc && sr_n >= 0 && sr_n < lr;
                if !check_bounds {
                    continue;
                }

                let color_n = image[sr_n as usize][sc_n as usize];
                if color_n != new_color && color_n == color {
                    q.push_back((sc_n, sr_n));
                }
            }

            image[sr as usize][sc as usize] = new_color;
        }

        image
    }
}
// @lc code=end

#[test]
fn test() {
    assert_eq!(
        Solution::flood_fill(vec![vec![1, 1, 1], vec![1, 1, 0], vec![1, 0, 1]], 1, 1, 2),
        [[2, 2, 2], [2, 2, 0], [2, 0, 1]]
    );
    assert_eq!(
        Solution::flood_fill(vec![vec![0, 0, 0], vec![0, 0, 0]], 0, 0, 2),
        [[2, 2, 2], [2, 2, 2]]
    );
}
