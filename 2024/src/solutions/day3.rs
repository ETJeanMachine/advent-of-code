use regex::Regex;
use std::time::Instant;

async fn part_one(input: String) -> i32 {
    let re = Regex::new(r"mul\((\d+),(\d+)\)").unwrap();
    let mut sum = 0;
    for c in re.captures_iter(input.as_str()) {
        let (_, [m1, m2]) = c.extract();
        sum += m1.parse::<i32>().unwrap() * m2.parse::<i32>().unwrap();
    }
    sum
}

async fn part_two(input: String) -> i32 {
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

type Solution = (String, u128);

pub async fn solve(input: String) -> (Solution, Solution) {
    async fn run_part(n: u8, input: String) -> Solution {
        let now = Instant::now();
        let res = match n {
            1 => part_one(input).await,
            2 => part_two(input).await,
            _ => unreachable!(),
        };
        let elapsed_ns = now.elapsed().as_nanos();
        return (res.to_string(), elapsed_ns);
    }
    let res1 = run_part(1, input.to_owned()).await;
    let res2 = run_part(2, input.to_owned()).await;
    return (res1, res2);
}
