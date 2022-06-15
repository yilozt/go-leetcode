struct Solution;

impl Solution {
    pub fn find_repeat_number(nums: Vec<i32>) -> i32 {
        let mut map = std::collections::HashSet::new();

        for n in nums {
            if map.get(&n).is_some() {
                return n;
            }
            map.insert(n);
        }

        unreachable!()
    }
}

#[test]
fn test() {
    assert_eq!(Solution::find_repeat_number(vec![2, 3, 1, 0, 2, 5, 3]), 2)
}
