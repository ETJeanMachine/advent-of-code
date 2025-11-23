use itertools::Itertools;
use std::{
    collections::{HashMap, HashSet},
    str::FromStr,
};

type Point = (usize, usize);
#[derive(Clone, Debug)]
pub struct RoofMap {
    map: HashMap<char, Vec<Point>>,
    height: usize,
    width: usize,
}

impl FromStr for RoofMap {
    type Err = String;

    fn from_str(s: &str) -> Result<Self, Self::Err> {
        let mut map: HashMap<char, Vec<Point>> = HashMap::new();
        let lines = s.split("\n");
        let height = lines.clone().count();
        let width = lines
            .clone()
            .last()
            .ok_or("Zero-width map".to_string())?
            .len();
        for (row, line) in lines.enumerate() {
            for (col, c) in line.chars().enumerate() {
                if c != '.' {
                    match map.get_mut(&c) {
                        Some(v) => v.push((row, col)),
                        None => {
                            map.insert(c, vec![(row, col)]);
                        }
                    }
                }
            }
        }
        Ok(RoofMap::new(map, height, width))
    }
}

impl RoofMap {
    pub fn new(map: HashMap<char, Vec<Point>>, height: usize, width: usize) -> Self {
        Self { map, height, width }
    }

    fn checked_row_add(&self, row: usize, rise: usize) -> Option<usize> {
        if row + rise < self.height {
            Some(row + rise)
        } else {
            None
        }
    }

    fn checked_col_add(&self, col: usize, run: usize) -> Option<usize> {
        if col + run < self.width {
            Some(col + run)
        } else {
            None
        }
    }

    fn next_bot(&self, pt: Point, rise: usize, run: usize, slope: i32) -> Option<Point> {
        let new_r = self.checked_row_add(pt.0, rise);
        let new_c = match slope {
            -1 => pt.1.checked_sub(run),
            1 => self.checked_col_add(pt.1, run),
            _ => unreachable!(),
        };
        if let (Some(r), Some(c)) = (new_r, new_c) {
            return Some((r, c));
        }
        None
    }

    fn next_top(&self, pt: Point, rise: usize, run: usize, slope: i32) -> Option<Point> {
        let new_r = pt.0.checked_sub(rise);
        let new_c = match slope {
            -1 => self.checked_col_add(pt.1, run),
            1 => pt.1.checked_sub(run),
            _ => unreachable!(),
        };
        if let (Some(r), Some(c)) = (new_r, new_c) {
            return Some((r, c));
        }
        None
    }

    fn all_pts(&self, a: Point, b: Point, limit: Option<usize>) -> Vec<Point> {
        let mut antinodes = vec![];
        let (rise, run) = (a.0.abs_diff(b.0), a.1.abs_diff(b.1));
        let slope = ((a.0 as f32 - b.0 as f32) / (a.1 as f32 - b.1 as f32)).signum() as i32;
        let (mut top_pt, mut bot_pt) = if a.0 < b.0 {
            (Some(a), Some(b))
        } else {
            (Some(b), Some(a))
        };

        let mut max_len = false;
        let mut add_attempt = 0;
        top_pt = self.next_top(top_pt.unwrap(), rise, run, slope);
        bot_pt = self.next_bot(bot_pt.unwrap(), rise, run, slope);
        while (top_pt.is_some() || bot_pt.is_some()) && !max_len {
            if let Some(pt) = top_pt {
                antinodes.push(pt);
                top_pt = self.next_top(pt, rise, run, slope);
            }
            if let Some(pt) = bot_pt {
                antinodes.push(pt);
                bot_pt = self.next_bot(pt, rise, run, slope);
            }
            add_attempt += 2;
            max_len = if let Some(n) = limit {
                add_attempt >= n
            } else {
                false
            };
        }

        antinodes
    }

    /// Solves for all antinode positions on the roof, with `limit` designating how
    /// many antinodes can exist for a given pair of similar sattelite points.
    pub fn antinodes(&self, limit: Option<usize>) -> HashSet<Point> {
        let mut set = HashSet::new();
        for sat in self.map.keys() {
            let points = self.map.get(sat).unwrap().to_owned();
            // All unique pairs of points.
            let pairs = points.iter().combinations(2).map(|v| (*v[0], *v[1]));
            for (a, b) in pairs {
                set.extend(self.all_pts(a, b, limit));
            }
        }
        return set;
    }
}

pub struct Solver(pub String);

impl super::lib::Puzzle<usize> for Solver {
    async fn part_one(&self) -> usize {
        let roof = RoofMap::from_str(self.0.as_str()).unwrap();
        let set = roof.antinodes(Some(2));
        set.len()
    }

    async fn part_two(&self) -> usize {
        let roof = RoofMap::from_str(self.0.as_str()).unwrap();
        let set = roof.antinodes(None);
        set.len()
    }
}
