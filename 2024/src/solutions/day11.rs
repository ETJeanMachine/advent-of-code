use std::{fmt::Display, num::ParseIntError, str::FromStr, sync::Arc, time::Instant};
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
    pub async fn blink(&self) -> Self {
        fn task(chunk: &[u64]) -> Vec<u64> {
            let mut new_stones = vec![];
            for stone in chunk.get(0) {
                if *stone == 0 {
                    new_stones.push(1);
                } else if (stone.ilog10() + 1) % 2 == 0 {
                    let base = (10 as u64).pow((stone.ilog10() + 1) / 2);
                    new_stones.push(*stone / base);
                    new_stones.push(*stone % base);
                } else {
                    new_stones.push(*stone * 2024);
                }
            }
            new_stones
        }

        let mut new_vec = vec![];
        let mut task_set = JoinSet::new();
        let stones_arc = Arc::new(self).clone();
        let chunk_size = self.len() / 4;
        for chunk in stones_arc.0.chunks(chunk_size) {
            task_set.spawn_blocking(move || task(chunk));
        }
        let join_results = task_set.join_all().await;
        for result in join_results {
            new_vec.extend(result);
        }
        Self(new_vec)
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
            stones = stones.blink().await;
        }
        stones.len()
    }

    async fn part_two(&self) -> usize {
        let mut stones = Stones::from_str(self.0.as_str()).unwrap();
        for i in 0..0 {
            let blink_instant = Instant::now();
            stones.blink().await;
            println!("Blink {}: {}ms", i, blink_instant.elapsed().as_millis());
        }
        stones.len()
    }
}
