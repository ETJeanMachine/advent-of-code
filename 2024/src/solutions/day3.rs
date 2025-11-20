use regex::Regex;

pub struct Solver(pub String);

impl super::lib::Puzzle<i32> for Solver {
    async fn part_one(&self) -> i32 {
        let input = self.0.to_owned();
        let re = Regex::new(r"mul\((\d+),(\d+)\)").unwrap();
        let mut sum = 0;
        for c in re.captures_iter(input.as_str()) {
            let (_, [m1, m2]) = c.extract();
            sum += m1.parse::<i32>().unwrap() * m2.parse::<i32>().unwrap();
        }
        sum
    }

    async fn part_two(&self) -> i32 {
        let input = self.0.to_owned();
        let re = Regex::new(r"(do|don't)(\(\))|mul\((\d+),(\d+)\)").unwrap();
        let mut sum = 0;
        let mut enabled = true;
        for c in re.captures_iter(input.as_str()) {
            let (full, [m1, m2]) = c.extract();
            match full {
                "do()" => enabled = true,
                "don't()" => enabled = false,
                _ => {
                    sum += match enabled {
                        true => m1.parse::<i32>().unwrap() * m2.parse::<i32>().unwrap(),
                        false => 0,
                    }
                }
            };
        }
        sum
    }
}
