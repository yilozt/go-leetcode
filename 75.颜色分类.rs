/*
 * @lc app=leetcode.cn id=75 lang=rust
 *
 * [75] 颜色分类
 */

struct Solution;

// @lc code=start
impl Solution {
    pub fn sort_colors(nums: &mut Vec<i32>) {
        let mut bucket = [0; 3];
        for i in 0..nums.len() {
            bucket[nums[i] as usize] += 1;
        }

        let mut i = 0;
        for color in 0..bucket.len() {
            let count = bucket[color];
            for _ in 0..count {
                nums[i] = color as i32;
                i += 1;
            }
        }
    }
}
// @lc code=end

