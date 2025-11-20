use std::{cmp::Ordering, collections::HashMap};

pub struct Solver(pub String);

type Rules = HashMap<i32, Vec<i32>>;
type Updates = Vec<Vec<i32>>;
impl Solver {
    fn parse_input(&self) -> (Rules, Updates) {
        let mut rules: Rules = HashMap::new();
        let mut updates = vec![];
        let split: Vec<&str> = self.0.split("\n\n").collect();
        for line in split[0].split("\n") {
            let vars: Vec<i32> = line.split("|").map(|x| x.parse::<i32>().unwrap()).collect();
            match rules.get_mut(&vars[0]) {
                Some(arr) => arr.push(vars[1]),
                None => {
                    rules.insert(vars[0], vec![vars[1]]);
                }
            };
        }
        for line in split[1].split("\n") {
            let vars: Vec<i32> = line.split(",").map(|x| x.parse::<i32>().unwrap()).collect();
            updates.push(vars);
        }
        (rules, updates)
    }
}

impl super::lib::Puzzle<i32> for Solver {
    async fn part_one(&self) -> i32 {
        let (rules, updates) = self.parse_input();
        let mut middle_sum = 0;
        for update in updates {
            let is_sorted = update.is_sorted_by(|a, b| match rules.get(a) {
                Some(v) => v.contains(b),
                None => true,
            });
            if is_sorted {
                middle_sum += update[update.len() / 2];
            }
        }
        middle_sum
    }

    async fn part_two(&self) -> i32 {
        let (rules, updates) = self.parse_input();
        let mut middle_sum = 0;

        for update in updates {
            let is_sorted = update.is_sorted_by(|a, b| match rules.get(a) {
                Some(v) => v.contains(b),
                None => true,
            });
            if !is_sorted {
                let mut sorted = update.clone();
                sorted.sort_by(|a, b| match rules.get(a) {
                    Some(v) => {
                        if v.contains(b) {
                            Ordering::Less
                        } else {
                            Ordering::Greater
                        }
                    }
                    None => Ordering::Equal,
                });
                middle_sum += sorted[sorted.len() / 2];
            }
        }
        middle_sum
    }
}
