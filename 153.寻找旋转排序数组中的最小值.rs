/*
 * @lc app=leetcode.cn id=153 lang=rust
 *
 * [153] 寻找旋转排序数组中的最小值
 */

struct Solution;

// @lc code=start
impl Solution {
    pub fn find_min(nums: Vec<i32>) -> i32 {
        let offset = find_offset(&nums);
        nums[offset as usize]
    }
}

fn find_offset(nums: &[i32]) -> i32 {
    let (mut l, mut r) = (0, nums.len() as i32 - 1);
    while l <= r {
        let m = (l + r) / 2;

        if nums[m as usize] > nums[r as usize] {
            l = m + 1;
        } else if nums[m as usize] < nums[l as usize] {
            r = m;
        } else {
            break;
        }
    }

    return l;
}
// @lc code=end

#[test]
fn test() {
    struct Desc {
        input: Vec<i32>,
        output: i32,
    }

    for desc in [
        Desc {
            input: vec![3, 4, 5, 1, 2],
            output: 1,
        },
        Desc {
            input: vec![4, 5, 6, 7, 0, 1, 2],
            output: 0,
        },
        Desc {
            input: vec![11, 13, 15, 17],
            output: 11,
        },
        Desc {
            input: vec![3, 1, 2],
            output: 1,
        },
    ] {
        assert_eq!(Solution::find_min(desc.input), desc.output)
    }
}
