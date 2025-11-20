use std::collections::HashMap;

pub struct Solver(pub String);

impl super::lib::Puzzle<i32> for Solver {
    async fn part_one(&self) -> i32 {
        let mut list1 = vec![];
        let mut list2 = vec![];
        // Consumes the iterator, returns an (Optional) String
        for line in self.0.split("\n") {
            // Splitting the list into two separate strings
            let mut split = line.split_whitespace();
            if let (Some(s1), Some(s2)) = (split.next(), split.next()) {
                let n = s1.parse::<i32>().unwrap();
                let m = s2.parse::<i32>().unwrap();
                list1.push(n);
                list2.push(m);
            }
        }
        list1.sort();
        list2.sort();
        let mut sum = 0;
        for i in 0..list1.len() {
            sum += (list1[i] - list2[i]).abs();
        }
        sum
    }

    async fn part_two(&self) -> i32 {
        let mut left = vec![];
        let mut right = HashMap::new();
        // Consumes the iterator, returns an (Optional) String
        for line in self.0.split("\n") {
            // Splitting the list into two separate strings
            let mut split = line.split_whitespace();
            if let (Some(s1), Some(s2)) = (split.next(), split.next()) {
                let n = s1.parse::<i32>().unwrap();
                let m = s2.parse::<i32>().unwrap();
                left.push(n);
                if let Some(i) = right.get_mut(&m) {
                    *i += 1;
                } else {
                    right.insert(m, 1);
                }
            }
        }
        let mut sum = 0;
        for n in left {
            sum += n * right.get(&n).unwrap_or(&0);
        }
        sum
    }
}
