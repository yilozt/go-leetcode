/*
 * @lc app=leetcode.cn id=706 lang=rust
 *
 * [706] 设计哈希映射
 */

// @lc code=start
struct MyHashMap {
    datas: Vec<Vec<(i32, i32)>>,
}

/**
 * `&self` means the method takes an immutable reference.
 * If you need a mutable reference, change it to `&mut self` instead.
 */
impl MyHashMap {
    fn new() -> Self {
        Self {
            datas: vec![vec![]; 1000],
        }
    }

    fn put(&mut self, key: i32, value: i32) {
        let i = key as usize % self.datas.len();

        let mut found = false;
        for entry in self.datas[i].iter_mut() {
            if entry.0 == key {
                entry.1 = value;
                found = true;
                break;
            }
        }
        if !found {
            self.datas[i].push((key, value))
        }
    }

    fn get(&self, key: i32) -> i32 {
        let i = key as usize % self.datas.len();

        for entry in self.datas[i].iter() {
            if entry.0 == key {
                return entry.1;
            }
        }

        -1
    }

    fn remove(&mut self, key: i32) {
        let i = key as usize % self.datas.len();

        let mut found = None;
        for j in 0..self.datas[i].len() {
            if self.datas[i][j].0 == key {
                found = Some(j);
                break;
            }
        }

        if let Some(r) = found {
            self.datas[i].remove(r);
        }
    }
}

/**
 * Your MyHashMap object will be instantiated and called as such:
 * let obj = MyHashMap::new();
 * obj.put(key, value);
 * let ret_2: i32 = obj.get(key);
 * obj.remove(key);
 */
struct A;
// @lc code=end
