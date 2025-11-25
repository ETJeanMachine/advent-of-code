use std::{
    collections::{HashSet, VecDeque},
    fmt::Display,
    str::FromStr,
};

pub struct TopoMap {
    map: Vec<Vec<u8>>,
    trailheads: Vec<(usize, usize)>,
}

impl FromStr for TopoMap {
    type Err = String;

    fn from_str(s: &str) -> Result<Self, Self::Err> {
        let mut map = vec![];
        let mut trailheads = vec![];
        for (row, line) in s.split("\n").enumerate() {
            map.push(
                line.chars()
                    .map(|c| c.to_string().parse().unwrap())
                    .collect(),
            );
            trailheads.extend(
                line.chars()
                    .enumerate()
                    .filter(|(_, x)| *x == '0')
                    .map(|(col, _)| (row, col)),
            );
        }
        Ok(TopoMap { map, trailheads })
    }
}

impl Display for TopoMap {
    fn fmt(&self, f: &mut std::fmt::Formatter<'_>) -> std::fmt::Result {
        let fmt_str = self.map.iter().fold(String::new(), |acc, v| {
            format!(
                "{}\n{}",
                acc,
                v.iter()
                    .fold(String::new(), |acc2, x| format!("{}{}", acc2, x))
            )
        });
        write!(f, "{}", fmt_str)
    }
}

impl TopoMap {
    fn steps(&self, row: usize, col: usize) -> Vec<(usize, usize)> {
        let mut steps = vec![];
        let height = self.map[row][col] + 1;
        if let Some(north_row) = row.checked_sub(1)
            && self.map[north_row][col] == height
        {
            steps.push((north_row, col));
        }
        if let Some(west_col) = col.checked_sub(1)
            && self.map[row][west_col] == height
        {
            steps.push((row, west_col))
        }
        let (south_row, east_col) = (row + 1, col + 1);
        if south_row < self.map.len() && self.map[south_row][col] == height {
            steps.push((south_row, col));
        }
        if east_col < self.map[row].len() && self.map[row][east_col] == height {
            steps.push((row, east_col))
        }
        steps
    }

    fn trailhead_score(&self, row: usize, col: usize) -> u32 {
        let mut queue = VecDeque::from([(row, col)]);
        let mut explored = HashSet::from([(row, col)]);
        let mut total_score = 0;
        while let Some((curr_row, curr_col)) = queue.pop_front() {
            if self.map[curr_row][curr_col] == 9 {
                total_score += 1;
            }
            for adjacent in self.steps(curr_row, curr_col) {
                if explored.insert(adjacent) {
                    queue.push_back(adjacent);
                }
            }
        }
        total_score
    }

    pub fn total_score(&self) -> u32 {
        self.trailheads
            .iter()
            .fold(0, |acc, (row, col)| acc + self.trailhead_score(*row, *col))
    }

    fn trailhead_rating(&self, row: usize, col: usize) -> u32 {
        if self.map[row][col] == 9 {
            return 1;
        }
        self.steps(row, col)
            .iter()
            .fold(0, |acc, (next_row, next_col)| {
                acc + self.trailhead_rating(*next_row, *next_col)
            })
    }

    pub fn total_rating(&self) -> u32 {
        self.trailheads
            .iter()
            .fold(0, |acc, (row, col)| acc + self.trailhead_rating(*row, *col))
    }
}

pub struct Solver(pub String);

impl super::lib::Puzzle<u32> for Solver {
    async fn part_one(&self) -> u32 {
        let map = TopoMap::from_str(self.0.as_str()).unwrap();
        map.total_score()
    }

    async fn part_two(&self) -> u32 {
        let map = TopoMap::from_str(self.0.as_str()).unwrap();
        map.total_rating()
    }
}
