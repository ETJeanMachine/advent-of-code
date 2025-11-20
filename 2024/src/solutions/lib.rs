use std::time::Instant;

use crate::solutions as sol;

type Solution = (String, u128);

pub trait Puzzle<T: ToString> {
    fn part_one(&self) -> impl Future<Output = T>;
    fn part_two(&self) -> impl Future<Output = T>;
}

async fn run_day<T: ToString>(puzzle: &impl Puzzle<T>) -> (Solution, Solution) {
    let now = Instant::now();
    let res1 = puzzle.part_one().await;
    let time1 = now.elapsed().as_nanos();

    let now = Instant::now();
    let res2 = puzzle.part_two().await;
    let time2 = now.elapsed().as_nanos();

    ((res1.to_string(), time1), (res2.to_string(), time2))
}

pub async fn run(day: u8, input: String) -> (Solution, Solution) {
    match day {
        1 => run_day(&sol::day1::Solver(input)).await,
        2 => run_day(&sol::day2::Solver(input)).await,
        3 => run_day(&sol::day3::Solver(input)).await,
        _ => unimplemented!(),
    }
}
