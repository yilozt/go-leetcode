/// This macro use to convert args contains null to `[Option<T>]` array.
///
/// ## Exam
///
/// ```
/// use macros::tree_input;
/// let a = 1;
///  assert_eq!(
///     tree_input!(a, null, 2, 3, -1, null, 5,),
///     [Some(1), None, Some(2), Some(3), Some(-1), None, Some(5)]
///  );
/// ```
///
/// Leetcode use **layer-order traversal** array to represent binary tree, when you are solving
/// a problem about binary tree, you may see the input like this:
///
/// ```txt
/// [3,9,20,null,null,15,7]
///
///          3
///         / \
///        9  20
///           / \
///         15   7
/// ```
#[proc_macro]
pub fn tree_input(input: proc_macro::TokenStream) -> proc_macro::TokenStream {
    let mut res_token = "[".to_string();

    for i in input.to_string().split(", ") {
        res_token.push_str(&match i {
            "null" => "None,".to_string(),
            data => format!("Some({}),", data.replace(" ", "")),
        });
    }
    res_token.push_str("]");

    res_token.parse().unwrap()
}