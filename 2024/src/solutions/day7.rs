use regex::Regex;

pub struct Solver(pub String);

impl Solver {
    fn parse_input(&self) -> Vec<(u128, Vec<u128>)> {
        let re = Regex::new(r"(\d+): ((?:\d+ ?)+)").unwrap();
        let captures = re.captures_iter(self.0.as_str());
        let mut operations: Vec<(u128, Vec<u128>)> = vec![];
        for c in captures {
            let (g1, g2) = (c.get(1).unwrap(), c.get(2).unwrap());
            let mut o = (g1.as_str().parse().unwrap(), vec![]);
            for n in g2.as_str().split_whitespace() {
                o.1.push(n.parse().unwrap());
            }
            operations.push(o);
        }
        operations
    }
}

impl super::lib::Puzzle<u128> for Solver {
    async fn part_one(&self) -> u128 {
        let operations = self.parse_input();
        let mut total = 0;
        for (res, vals) in operations {
            let mut res_vec = vec![vals[0]];
            for i in 1..vals.len() {
                let mut new_vec = vec![];
                for v in res_vec {
                    new_vec.push(v * vals[i]);
                    new_vec.push(v + vals[i]);
                }
                res_vec = new_vec;
            }
            let valid_count = res_vec.iter().filter(|&v| *v == res).count();
            if valid_count > 0 {
                total += res
            }
        }
        total
    }

    async fn part_two(&self) -> u128 {
        let operations = self.parse_input();
        let mut total = 0;
        for (res, vals) in operations {
            total += vals
                .iter()
                .skip(1)
                .fold(vec![vals[0]], |acc, &b| {
                    acc.iter()
                        .flat_map(|&a| vec![a * b, a + b, format!("{}{}", a, b).parse().unwrap()])
                        .filter(|&v| v <= res)
                        .collect()
                })
                .iter()
                .find(|&v| *v == res)
                .unwrap_or(&0);
        }
        total
    }
}
