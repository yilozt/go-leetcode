/*
 * @lc app=leetcode.cn id=15 lang=rust
 *
 * [15] 三数之和
 */

use crate::structs::Solution;

// @lc code=start
impl Solution {
    pub fn three_sum(mut nums: Vec<i32>) -> Vec<Vec<i32>> {
        if nums.len() < 3 {
            return vec![];
        }

        nums.sort();
        let mut res = vec![];
        let mut set = std::collections::HashSet::new();

        for n in &nums {
            set.insert(n);
        }

        let mut i = 0;
        while i < nums.len() - 2 {
            let mut j = i + 1;
            while j < nums.len() - 1 {
                let rest = 0 - (nums[i] + nums[j]);

                // 剪裁
                if rest < nums[j + 1] {
                    break;
                }

                if set.get(&rest).is_some() {
                    res.push(vec![nums[i], nums[j], rest]);
                }

                // 忽略重复元素
                j += 1;
                while j < nums.len() && nums[j] == nums[j - 1] {
                    j += 1;
                }
            }

            i += 1;
            while i < nums.len() && nums[i] == nums[i - 1] {
                i += 1;
            }
        }

        res
    }
}

// @lc code=end

#[test]
fn test() {
    assert_eq!(
        Solution::three_sum(vec![-1, 0, 1, 2, -1, -4]),
        [[-1, -1, 2], [-1, 0, 1]]
    );
    assert!(Solution::three_sum(vec![]).len() == 0);
    assert_eq!(Solution::three_sum(vec![0, 0, 0]), [[0, 0, 0]]);
}
