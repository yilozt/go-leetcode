use std::{cell::RefCell, rc::Rc};
pub struct Solution;

#[derive(Debug, PartialEq, Eq)]
pub struct TreeNode {
    pub val: i32,
    pub left: Option<Rc<RefCell<TreeNode>>>,
    pub right: Option<Rc<RefCell<TreeNode>>>,
}

impl TreeNode {
    #[inline]
    pub fn new(val: i32) -> Self {
        TreeNode {
            val,
            left: None,
            right: None,
        }
    }
}

/// build binary tree from **lever-order** traversal, `None` present `NULL`.
///
/// # Examples
///
/// ```
/// //        1
/// //       / \
/// //      2   3
/// //     /    \
/// //    4      5
/// use lv7_binarytree::build_tree;
/// let t = build_tree(&[Some(1), Some(2), Some(3), Some(4), None, None, 5]);
/// println!("{:#?}", t);
/// ```
pub fn build_tree(nodes: &[Option<i32>]) -> Option<Rc<RefCell<TreeNode>>> {
    if nodes.is_empty() {
        return None;
    }

    let get_node = |i: usize| {
        if i >= nodes.len() || nodes[i].is_none() {
            None
        } else {
            Some(Rc::new(RefCell::new(TreeNode::new(nodes[i].unwrap()))))
        }
    };

    let mut deque = std::collections::VecDeque::new();
    let root = get_node(0);
    deque.push_back(root.clone());

    let mut i = 1;

    while !deque.is_empty() {
        let mut nd = deque.pop_front().unwrap();
        let left = get_node(i);
        let right = get_node(i + 1);
        let mut nd = nd.as_mut().unwrap().borrow_mut();
        nd.left = left.clone();
        nd.right = right.clone();
        i += 2;
        if left.is_some() {
            deque.push_back(left);
        }
        if right.is_some() {
            deque.push_back(right);
        }
    }

    root
}

/// Helper macro for test
///
/// Build binary tree from `lever` order
///
/// ## Example
///
///```
/// use lv7_binarytree::tree;
///
/// //        1
/// //       / \
/// //      2   3
/// //     /    \
/// //    4      5
/// let t = tree!(1, 2, 3, 4, null, null, 5);
/// println!("{:#?}", t);
/// ```
#[macro_export]
macro_rules! tree {
  () => ($crate::build_tree(&[]));
  ($($x:tt)*) => ($crate::build_tree(&macros::tree_input!($($x)*)));
}
