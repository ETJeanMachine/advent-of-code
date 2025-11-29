use regex::Regex;

const W: usize = 101;
const W_MID: usize = 51;
const H: usize = 103;
const H_MID: usize = 52;

pub struct Robot {
    position: (usize, usize),
    velocity: (i32, i32),
}

impl Robot {
    pub fn move_robot(&mut self, time_seconds: u32) {
        let (velocity_x, velocity_y) = self.velocity;
        let (start_x, start_y) = self.position;
        let new_x =
            ((start_x as i32 + (velocity_x * time_seconds as i32)).rem_euclid(W as i32)) as usize;
        let new_y =
            ((start_y as i32 + (velocity_y * time_seconds as i32)).rem_euclid(H as i32)) as usize;
        self.position = (new_x, new_y);
    }
}

pub struct Solver(pub String);

impl Solver {
    fn parse_input(&self) -> Vec<Robot> {
        let mut robots = vec![];
        let re = Regex::new(r"p=(-?\d+),(-?\d+) v=(-?\d+),(-?\d+)").unwrap();
        let captures = re.captures_iter(self.0.as_str());
        for capture in captures {
            let position: Vec<_> = vec![capture.get(1), capture.get(2)]
                .iter()
                .map(|c| c.unwrap().as_str().parse::<usize>().unwrap())
                .collect();
            let position = (position[0], position[1]);
            let velocity: Vec<_> = vec![capture.get(3), capture.get(4)]
                .iter()
                .map(|c| c.unwrap().as_str().parse::<i32>().unwrap())
                .collect();
            let velocity = (velocity[0], velocity[1]);
            robots.push(Robot { position, velocity });
        }
        robots
    }
}

fn safety_factor(robots: Vec<Robot>) -> u32 {
    let (mut q1, mut q2, mut q3, mut q4) = (0, 0, 0, 0);
    for r in robots {
        match r.position {
            (0..W_MID, 0..H_MID) => q1 += 1,
            (W_MID..W, 0..H_MID) => q2 += 1,
            (0..W_MID, H_MID..H) => q3 += 1,
            (W_MID..W, H_MID..H) => q4 += 1,
            _ => unreachable!(),
        }
    }
    q1 * q2 * q3 * q4
}

impl super::lib::Puzzle<u32> for Solver {
    async fn part_one(&self) -> u32 {
        let mut robots = self.parse_input();
        robots.iter_mut().for_each(|r| r.move_robot(100));
        safety_factor(robots)
    }

    async fn part_two(&self) -> u32 {
        0
    }
}
