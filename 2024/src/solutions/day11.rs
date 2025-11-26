use std::collections::HashMap;

pub struct Solver(pub String);

impl Solver {
    fn parse_input(&self) -> Vec<u64> {
        self.0
            .split_whitespace()
            .map(|substr| substr.parse().unwrap())
            .collect()
    }
}

fn blink(stone: &u64, memoized_stones: &mut Vec<HashMap<u64, u64>>, blinks: &usize) -> u64 {
    if *blinks == 0 {
        return 1;
    } else if let Some(count) = memoized_stones[*blinks - 1].get(stone) {
        return *count;
    }
    let mut new_stones = vec![];
    if *stone == 0 {
        new_stones.push(1);
    } else if (stone.ilog10() + 1) % 2 == 0 {
        let base = (10 as u64).pow((stone.ilog10() + 1) / 2);
        new_stones.push(*stone / base);
        new_stones.push(*stone % base);
    } else {
        new_stones.push(*stone * 2024);
    }
    let result = new_stones.iter().fold(0, |acc, new_stone| {
        acc + blink(new_stone, memoized_stones, &(blinks - 1))
    });
    memoized_stones[*blinks - 1].insert(*stone, result);
    result
}

impl super::lib::Puzzle<u64> for Solver {
    async fn part_one(&self) -> u64 {
        let stones = self.parse_input();
        let mut memoized_stones = (0..25).map(|_| HashMap::new()).collect();
        stones.iter().fold(0, |acc, stone| {
            acc + blink(stone, &mut memoized_stones, &25)
        })
    }

    async fn part_two(&self) -> u64 {
        let stones = self.parse_input();
        let mut memoized_stones = (0..75).map(|_| HashMap::new()).collect();
        stones.iter().fold(0, |acc, stone| {
            acc + blink(stone, &mut memoized_stones, &75)
        })
    }
}
