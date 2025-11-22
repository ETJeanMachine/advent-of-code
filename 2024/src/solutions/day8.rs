use std::collections::{HashMap, HashSet};

pub struct Solver(pub String);

type Point = (usize, usize);
impl Solver {
    fn parse_input(&self) -> (HashMap<char, Vec<Point>>, usize, usize) {
        let mut map: HashMap<char, Vec<Point>> = HashMap::new();
        let lines = self.0.split("\n");
        let height = lines.clone().count();
        let width = lines.clone().last().unwrap().len();
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
        (map, height, width)
    }
}

fn antinodes(a: &Point, b: &Point, height: usize, width: usize) -> (Option<Point>, Option<Point>) {
    let (rise, run) = (a.0.abs_diff(b.0), a.1.abs_diff(b.1));
    let mut antinode_a = (None, None);
    let mut antinode_b = (None, None);
    if a.0 < b.0 {
        antinode_a.0 = a.0.checked_sub(rise);
        antinode_b.0 = if b.0 + rise < height {
            Some(b.0 + rise)
        } else {
            None
        }
    } else {
        antinode_b.0 = b.0.checked_sub(rise);
        antinode_a.0 = if a.0 + rise < height {
            Some(a.0 + rise)
        } else {
            None
        }
    }
    if a.1 < b.1 {
        antinode_a.1 = a.1.checked_sub(run);
        antinode_b.1 = if b.1 + run < width {
            Some(b.1 + run)
        } else {
            None
        }
    } else {
        antinode_b.1 = b.1.checked_sub(run);
        antinode_a.1 = if a.1 + run < height {
            Some(a.1 + run)
        } else {
            None
        }
    }
    let antinode_a = if let (Some(r), Some(c)) = antinode_a {
        Some((r, c))
    } else {
        None
    };
    let antinode_b = if let (Some(r), Some(c)) = antinode_b {
        Some((r, c))
    } else {
        None
    };
    return (antinode_a, antinode_b);
}

impl super::lib::Puzzle<usize> for Solver {
    async fn part_one(&self) -> usize {
        let (map, height, width) = self.parse_input();
        let mut node_set: HashSet<Point> = HashSet::new();
        for k in map.keys() {
            let points = map.get(k).unwrap();
            for a in points {
                for b in points {
                    if a == b {
                        continue;
                    }
                    let (antinode_a, antinode_b) = antinodes(a, b, height, width);
                    if let Some(antinode_a) = antinode_a {
                        node_set.insert(antinode_a);
                    }
                    if let Some(antinode_b) = antinode_b {
                        node_set.insert(antinode_b);
                    }
                }
            }
        }
        node_set.len()
    }

    async fn part_two(&self) -> usize {
        0
    }
}
