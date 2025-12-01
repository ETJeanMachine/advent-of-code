use regex::Regex;
use std::io::{self, Write, stdin, stdout};

const W: usize = 101;
const H: usize = 103;

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
            (0..=49, 0..=50) => q1 += 1,
            (51..W, 0..=50) => q2 += 1,
            (0..=49, 52..H) => q3 += 1,
            (51..W, 52..H) => q4 += 1,
            _ => (),
        }
    }
    q1 * q2 * q3 * q4
}

fn to_arr(robots: &Vec<Robot>) -> [[u16; W]; H] {
    let mut robot_arr = [[0_u16; W]; H];
    for r in robots {
        let (col, row) = r.position;
        robot_arr[row][col] += 1;
    }
    robot_arr
}

fn xmas_tree() -> [[u16; W]; H] {
    let mut tree_arr = [[0_u16; W]; H];
    for (row, tree_row) in tree_arr.iter_mut().enumerate() {
        if row < W / 2 {
            let (l_col, r_col) = (W / 2 - row, W / 2 + row);
            tree_row[l_col] = 1;
            tree_row[r_col] = 1;
        } else if row == W / 2 {
            tree_row.iter_mut().for_each(|v| *v = 1);
        }
    }
    tree_arr
}

fn print_arr(arr: [[u16; W]; H]) {
    for row in arr.iter() {
        let row_str: String = row
            .iter()
            .map(|x| if *x == 0 { '.' } else { '*' })
            .collect();
        println!("{}", row_str);
    }
}

impl super::lib::Puzzle<u32> for Solver {
    async fn part_one(&self) -> u32 {
        let mut robots = self.parse_input();
        robots.iter_mut().for_each(|r| r.move_robot(100));
        safety_factor(robots)
    }

    async fn part_two(&self) -> u32 {
        let mut robots = self.parse_input();
        let mut seconds_elapsed = 0;
        // loop {
        //     robots.iter_mut().for_each(|r| r.move_robot(1));
        //     seconds_elapsed += 1;
        //     println!("\rSeconds Elapsed: {}", seconds_elapsed);
        //     let arr = to_arr(&robots);
        //     print_arr(arr);
        //     stdout().flush().unwrap();
        //     std::thread::sleep(std::time::Duration::from_millis(5));
        // }
        seconds_elapsed
    }
}
