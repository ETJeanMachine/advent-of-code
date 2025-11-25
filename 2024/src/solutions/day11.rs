use std::{
    fmt::Display,
    num::ParseIntError,
    str::FromStr,
    sync::{Arc, Mutex},
    time::Instant,
};

use rayon::{iter::ParallelIterator, slice::ParallelSliceMut};
use tokio::task::JoinSet;

pub struct Stones(Vec<u64>);

impl FromStr for Stones {
    type Err = ParseIntError;

    fn from_str(s: &str) -> Result<Self, Self::Err> {
        let parsed = s
            .split_whitespace()
            .map_while(|substr| substr.parse().ok())
            .collect();
        Ok(Self(parsed))
    }
}

impl Display for Stones {
    fn fmt(&self, f: &mut std::fmt::Formatter<'_>) -> std::fmt::Result {
        let fmt_str = self
            .0
            .iter()
            .fold(String::new(), |acc, x| format!("{} {}", acc, x));
        write!(f, "{}", fmt_str.trim())
    }
}

impl Stones {
    pub fn blink(&mut self) {
        let mut new_stones = vec![];
        for stone in self.0.iter_mut() {
            if *stone == 0 {
                *stone = 1;
            } else if (stone.ilog10() + 1) % 2 == 0 {
                let base = (10 as u64).pow((stone.ilog10() + 1) / 2);
                new_stones.push(*stone / base);
                *stone %= base;
            } else {
                *stone *= 2024;
            }
        }
        self.0.extend(new_stones)
    }

    pub fn blink_parallel(&mut self) {
        let n = if self.len() > 202019 { 16 } else { 4 };
        let chunk_size = self.len() / n;
        let new_stones_vec: Vec<_> = self
            .0
            .par_chunks_mut(chunk_size)
            .map(move |chunk| {
                chunk.iter_mut().fold(vec![], move |mut new_stones, stone| {
                    if *stone == 0 {
                        *stone = 1;
                    } else if (stone.ilog10() + 1) % 2 == 0 {
                        let base = (10 as u64).pow((stone.ilog10() + 1) / 2);
                        new_stones.push(*stone / base);
                        *stone %= base;
                    } else {
                        *stone *= 2024;
                    }
                    new_stones
                })
            })
            .collect();
        for new_stones in new_stones_vec {
            self.0.extend(new_stones);
        }
    }

    pub fn len(&self) -> usize {
        self.0.len()
    }
}

pub struct Solver(pub String);

impl super::lib::Puzzle<usize> for Solver {
    async fn part_one(&self) -> usize {
        let mut stones = Stones::from_str(self.0.as_str()).unwrap();
        for _ in 0..25 {
            stones.blink();
        }
        stones.len()
    }

    async fn part_two(&self) -> usize {
        let mut stones = Stones::from_str(self.0.as_str()).unwrap();
        for i in 0..75 {
            let blink_instant = Instant::now();
            stones.blink_parallel();
            println!("Blink {}: {}ms", i, blink_instant.elapsed().as_millis());
            println!("Blink {}: {}", i, stones.len());
        }
        stones.len()
    }
}
