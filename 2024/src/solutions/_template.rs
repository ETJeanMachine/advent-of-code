use std::time::Instant;

async fn part_one(input: String) -> i32 {
    0
}

async fn part_two(input: String) -> i32 {
    0
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
