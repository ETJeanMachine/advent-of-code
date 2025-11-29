use std::time::Instant;

use crate::solutions::*;

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
        1 => run_day(&day1::Solver(input)).await,
        2 => run_day(&day2::Solver(input)).await,
        3 => run_day(&day3::Solver(input)).await,
        4 => run_day(&day4::Solver(input)).await,
        5 => run_day(&day5::Solver(input)).await,
        6 => run_day(&day6::Solver(input)).await,
        7 => run_day(&day7::Solver(input)).await,
        8 => run_day(&day8::Solver(input)).await,
        9 => run_day(&day9::Solver(input)).await,
        10 => run_day(&day10::Solver(input)).await,
        11 => run_day(&day11::Solver(input)).await,
        12 => run_day(&day12::Solver(input)).await,
        13 => run_day(&day13::Solver(input)).await,
        14 => run_day(&day14::Solver(input)).await,
        ..=25 => unimplemented!(),
        _ => unreachable!(),
    }
}
