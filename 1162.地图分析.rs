/*
 * @lc app=leetcode.cn id=1162 lang=rust
 *
 * [1162] 地图分析
 */

struct Solution;

// @lc code=start
struct Data {
    w: i32,
    h: i32,
    visited: Vec<Vec<bool>>,
}

const DIRS: [(i32, i32); 4] = [(-1, 0), (0, 1), (1, -0), (0, -1)];

impl Solution {
    pub fn max_distance(grid: Vec<Vec<i32>>) -> i32 {
        let mut ans = -1;
        let w = grid[0].len() as i32;
        let h = grid.len() as i32;
        let visited = vec![vec![false; grid[0].len()]; grid.len()];
        let mut s = Data { w, h, visited };
        for y in 0..s.h {
            for x in 0..s.w {
                if grid[y as usize][x as usize] == 0 {
                    ans = ans.max(s.bfs(&grid, y, x));
                }
            }
        }
        ans
    }
}

impl Data {
    fn bfs(&mut self, grid: &Vec<Vec<i32>>, y: i32, x: i32) -> i32 {
        for i in self.visited.iter_mut() {
            i.fill(false)
        }

        let mut queue = std::collections::VecDeque::new();
        queue.push_back((y, x, 0));
        self.visited[y as usize][x as usize] = true;

        while queue.len() != 0 {
            let (y, x, dist) = queue.pop_front().unwrap();
            for (dy, dx) in DIRS {
                let (ny, nx) = (y + dy, x + dx);
                if !(nx >= 0 && ny >= 0 && nx < self.w && ny < self.h) {
                    continue;
                }
                if !self.visited[ny as usize][nx as usize] {
                    if grid[ny as usize][nx as usize] == 1 {
                        return dist + 1;
                    }
                    self.visited[ny as usize][nx as usize] = true;
                    queue.push_back((ny, nx, dist + 1))
                }
            }
        }

        return -1;
    }
}
// @lc code=end

#[test]
fn test() {
    struct Problem {
        grid: Vec<Vec<i32>>,
        res: i32,
    }

    let desc = [
        Problem {
            grid: vec![vec![1, 0, 1], vec![0, 0, 0], vec![1, 0, 1]],
            res: 2,
        },
        Problem {
            grid: vec![vec![1, 0, 0], vec![0, 0, 0], vec![0, 0, 0]],
            res: 4,
        },
        Problem {
            grid: vec![vec![0, 0, 0], vec![0, 0, 0]],
            res: -1,
        },
        Problem {
            grid: vec![vec![1, 1, 1], vec![1, 1, 1]],
            res: -1,
        },
    ];

    for des in desc {
        let got = Solution::max_distance(des.grid);
        let want = des.res;

        assert_eq!(got, want)
    }
}
