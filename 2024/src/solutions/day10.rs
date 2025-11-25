use std::{fmt::Display, str::FromStr};

pub struct TopoMap {
    map: Vec<Vec<u8>>,
    trailheads: Vec<(usize, usize)>,
}

impl FromStr for TopoMap {
    type Err = String;

    fn from_str(s: &str) -> Result<Self, Self::Err> {
        let mut map = vec![];
        let mut trailheads = vec![];
        let mut r = 0;
        for line in s.split("\n") {
            map.push(vec![]);
            let mut c = 0;
            for d in line.chars() {
                let d = d.to_string().parse().unwrap();
                map[r].push(d);
                if d == 0 {
                    trailheads.push((r, c));
                }
                c += 1;
            }
            r += 1;
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
        if let Some(r) = row.checked_sub(1)
            && self.map[r][col] == height
        {
            steps.push((r, col));
        }
        if let Some(c) = col.checked_sub(1)
            && self.map[row][c] == height
        {
            steps.push((row, c))
        }
        let (r, c) = (row + 1, col + 1);
        if r < self.map.len() && self.map[r][col] == height {
            steps.push((r, col));
        }
        if c < self.map[row].len() && self.map[r][col] == height {
            steps.push((row, c))
        }
        steps
    }

    fn trailhead_score(&self, row: usize, col: usize) -> u32 {
        0
    }

    pub fn walkable_trails(&self) -> u32 {
        self.trailheads
            .iter()
            .fold(0, |acc, (row, col)| acc + self.trailhead_score(*row, *col))
    }
}

pub struct Solver(pub String);

impl super::lib::Puzzle<u32> for Solver {
    async fn part_one(&self) -> u32 {
        let map = TopoMap::from_str(self.0.as_str()).unwrap();
        map.walkable_trails()
    }

    async fn part_two(&self) -> u32 {
        0
    }
}
