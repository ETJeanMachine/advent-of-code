use regex::Regex;
use std::{num::ParseIntError, str::FromStr};

#[derive(Hash, Clone, Copy)]
pub struct ClawMachine {
    button_a: (u32, u32),
    button_b: (u32, u32),
    prize: (u32, u32),
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
    pub fn min_tokens(&self) -> Option<u32> {
        let mut r1: Vec<_> = vec![self.button_a.0, self.button_b.0, self.prize.0]
            .iter()
            .map(|x| *x as f64)
            .collect();
        let mut r2: Vec<_> = vec![self.button_a.1, self.button_b.1, self.prize.1]
            .iter()
            .map(|x| *x as f64)
            .collect();
        let multiplier = -(r2[0] / r1[0]);
        r1.iter_mut().for_each(|val| *val *= multiplier);
        r2.iter_mut()
            .enumerate()
            .for_each(|(idx, val)| *val += r1[idx]);
        let r1_mult = 1.0 / r1[0];
        let r2_mult = 1.0 / r2[1];
        r1.iter_mut().for_each(|val| *val *= r1_mult);
        r2.iter_mut().for_each(|val| *val *= r2_mult);
        let b_pushes = r2[2];
        let a_pushes = r1[2] - (r1[1] * b_pushes);
        if a_pushes.fract() == 0.0 && b_pushes.fract() == 0.0 {
            return Some((a_pushes as u32 * 3) + b_pushes as u32);
        }
        None
    }
}

pub struct Solver(pub String);

impl Solver {
    fn parse_input(&self) -> Vec<ClawMachine> {
        let _test = "Button A: X+94, Y+34\nButton B: X+22, Y+67\nPrize: X=8400, Y=5400\n\nButton A: X+26, Y+66\nButton B: X+67, Y+21\nPrize: X=12748, Y=12176\n\nButton A: X+17, Y+86\nButton B: X+84, Y+37\nPrize: X=7870, Y=6450\n\nButton A: X+69, Y+23\nButton B: X+27, Y+71\nPrize: X=18641, Y=10279";
        let mut claw_machines = vec![];
        for group in _test.split("\n\n") {
            claw_machines.push(ClawMachine::from_str(group).unwrap());
        }
        claw_machines
    }
}

impl super::lib::Puzzle<u32> for Solver {
    async fn part_one(&self) -> u32 {
        let claw_machines = self.parse_input();
        claw_machines.iter().fold(0, |acc, machine| {
            acc + match machine.min_tokens() {
                Some(val) => val,
                None => 0,
            }
        })
    }

    async fn part_two(&self) -> u32 {
        0
    }
}
