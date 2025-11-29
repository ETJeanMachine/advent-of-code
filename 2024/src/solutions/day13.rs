use fraction::Fraction;
use regex::Regex;
use std::{num::ParseIntError, str::FromStr};

#[derive(Hash, Clone, Copy)]
pub struct ClawMachine {
    button_a: (u64, u64),
    button_b: (u64, u64),
    prize: (u64, u64),
}

impl FromStr for ClawMachine {
    type Err = ParseIntError;

    fn from_str(s: &str) -> Result<Self, Self::Err> {
        let re = Regex::new(r".+: X.?(\d+), Y.?(\d+)").unwrap();
        let mut captures = re.captures_iter(s);
        let button_a_capture = captures.next().unwrap();
        let button_a = (
            button_a_capture.get(1).unwrap().as_str().parse()?,
            button_a_capture.get(2).unwrap().as_str().parse()?,
        );
        let button_b_capture = captures.next().unwrap();
        let button_b = (
            button_b_capture.get(1).unwrap().as_str().parse()?,
            button_b_capture.get(2).unwrap().as_str().parse()?,
        );
        let prize_capture = captures.next().unwrap();
        let prize = (
            prize_capture.get(1).unwrap().as_str().parse()?,
            prize_capture.get(2).unwrap().as_str().parse()?,
        );
        Ok(Self {
            button_a,
            button_b,
            prize,
        })
    }
}

impl ClawMachine {
    pub fn min_tokens(&self, prize_add: u64) -> Option<u64> {
        let mut r1: Vec<_> = vec![self.button_a.0, self.button_b.0, self.prize.0 + prize_add]
            .iter()
            .map(|x| Fraction::from(*x))
            .collect();
        let mut r2: Vec<_> = vec![self.button_a.1, self.button_b.1, self.prize.1 + prize_add]
            .iter()
            .map(|x| Fraction::from(*x))
            .collect();
        let multiplier = -(r2[0] / r1[0]);
        r1.iter_mut().for_each(|val| *val *= multiplier);
        r2.iter_mut()
            .enumerate()
            .for_each(|(idx, val)| *val += r1[idx]);
        let r1_mult = Fraction::from(1) / r1[0];
        let r2_mult = Fraction::from(1) / r2[1];
        r1.iter_mut().for_each(|val| *val *= r1_mult);
        r2.iter_mut().for_each(|val| *val *= r2_mult);
        let b_pushes = r2[2];
        let a_pushes = r1[2] - (r1[1] * b_pushes);
        if *a_pushes.denom().unwrap() == 1 && *b_pushes.denom().unwrap() == 1 {
            return Some((*a_pushes.numer().unwrap() * 3) + *b_pushes.numer().unwrap());
        }
        None
    }
}

pub struct Solver(pub String);

impl Solver {
    fn parse_input(&self) -> Vec<ClawMachine> {
        let mut claw_machines = vec![];
        for group in self.0.split("\n\n") {
            claw_machines.push(ClawMachine::from_str(group).unwrap());
        }
        claw_machines
    }
}

impl super::lib::Puzzle<u64> for Solver {
    async fn part_one(&self) -> u64 {
        let claw_machines = self.parse_input();
        claw_machines.iter().fold(0, |acc, machine| {
            acc + match machine.min_tokens(0) {
                Some(val) => val,
                None => 0,
            }
        })
    }

    async fn part_two(&self) -> u64 {
        let claw_machines = self.parse_input();
        claw_machines.iter().fold(0, |acc, machine| {
            acc + match machine.min_tokens(10000000000000) {
                Some(val) => val,
                None => 0,
            }
        })
    }
}
