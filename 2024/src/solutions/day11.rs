use std::{fmt::Display, num::ParseIntError, str::FromStr, time::Instant};

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
        for i in 0..0 {
            let blink_instant = Instant::now();
            stones.blink();
            println!("Blink {}: {}ms", i, blink_instant.elapsed().as_millis());
        }
        stones.len()
    }
}
