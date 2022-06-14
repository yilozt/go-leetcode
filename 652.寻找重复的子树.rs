/*
 * @lc app=leetcode.cn id=652 lang=rust
 *
 * [652] 寻找重复的子树
 */

use crate::*;

// @lc code=start
// Definition for a binary tree node.
// #[derive(Debug, PartialEq, Eq)]
// pub struct TreeNode {
//   pub val: i32,
//   pub left: Option<Rc<RefCell<TreeNode>>>,
//   pub right: Option<Rc<RefCell<TreeNode>>>,
// }
//
// impl TreeNode {
//   #[inline]
//   pub fn new(val: i32) -> Self {
//     TreeNode {
//       val,
//       left: None,
//       right: None
//     }
//   }
// }
use std::cell::RefCell;
use std::collections::HashMap;
use std::rc::Rc;
type Tree = Option<Rc<RefCell<TreeNode>>>;

impl Solution {
    pub fn find_duplicate_subtrees(
        root: Option<Rc<RefCell<TreeNode>>>,
    ) -> Vec<Option<Rc<RefCell<TreeNode>>>> {
        let mut keys = HashMap::new();

        Self::trace(root, &mut keys);

        keys.into_values().filter(|n| n.is_some()).collect()
    }

    // Tip：二叉树序列化
    pub fn trace(root: Option<Rc<RefCell<TreeNode>>>, keys: &mut HashMap<String, Tree>) -> String {
        match root.clone() {
            None => "n".to_string(),
            Some(rc) => {
                let n = rc.borrow();
                let key = format!(
                    "{} {} {}",
                    Self::trace(n.left.clone(), keys),
                    Self::trace(n.right.clone(), keys),
                    n.val,
                );

                // 记录重复节点
                if let Some(n) = keys.get_mut(&key) {
                    *n = root;
                } else {
                    keys.insert(key.clone(), None);
                }
                key
            }
        }
    }
}
// @lc code=end

#[test]
fn test() {
    fn trace_tree(t: Option<Rc<RefCell<TreeNode>>>) -> String {
        match t {
            Some(rc) => {
                let n = rc.borrow();
                let hash = format!(
                    "{}->{}->{}",
                    trace_tree(n.left.clone()),
                    trace_tree(n.right.clone()),
                    n.val,
                );
                println!("{}: {}", n.val, hash);
                hash
            }
            None => String::from("nil"),
        }
    }

    let tree = tree!(1, 2, 3, 4, null, 2, 4, null, null, 4);
    trace_tree(tree);
}
