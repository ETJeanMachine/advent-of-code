use clap::Parser;
use dotenvy::dotenv;
use std::{error::Error, process};
use tokio;

pub mod solutions;

#[derive(Parser, Debug)]
struct Args {
    /// Solution day we are running.
    #[arg(value_parser = clap::value_parser!(u8).range(0..=25))]
    day: u8,
}

async fn get_input(day: u8) -> Result<String, Box<dyn Error>> {
    dotenv().ok();
    let session = std::env::var("SESSION_COOKIE").expect("SESSION_COOKIE not set");
    let session_cookie = format!("session={}", session);

    let url = format!("https://adventofcode.com/2024/day/{}/input", day);

    let client = reqwest::Client::new();
    let response = client
        .get(url.to_owned())
        .header("Cookie", session_cookie)
        .send()
        .await?
        .text()
        .await?;
    return Ok(response);
}

#[tokio::main]
async fn main() {
    let args = Args::parse();
    let input = match get_input(args.day).await {
        Ok(t) => t,
        _ => panic!(),
    };

    let ((res_one, time_one), (res_two, time_two)) = match args.day {
        1 => solutions::day1::solve(input).await,
        2 => solutions::day2::solve(input).await,
        3 => solutions::day3::solve(input).await,
        _ => {
            eprintln!("Day {} not implemented yet!", args.day);
            process::exit(0);
        }
    };

    println!("Advent of Code 2024 Day {}:", args.day);
    println!("Part One: {}", res_one);
    println!("Time One: {:.2}ms\n", time_one as f64 / 1e6);
    println!("Part Two: {}", res_two);
    println!("Time Two: {:.2}ms\n", time_two as f64 / 1e6);
}
