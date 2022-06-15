/*
 * @lc app=leetcode.cn id=56 lang=rust
 *
 * [56] 合并区间
 */

struct Solution;

// @lc code=start
impl Solution {
    pub fn merge(mut intervals: Vec<Vec<i32>>) -> Vec<Vec<i32>> {
        intervals.sort_by_key(|i| i[0]);
        
        let mut res = vec![];

        for iterval in intervals {
            if res.len() == 0 {
                res.push(iterval);
            } else {
                let last = res.last_mut().unwrap();
                
                // 合并
                if last[1] >= iterval[0] {
                    last[1] = iterval[1].max(last[1]);
                    last[0] = iterval[0].min(last[0]);
                } else {
                    res.push(iterval)
                }
            }
        }

        res
    }
}
// @lc code=end

