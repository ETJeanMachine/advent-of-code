use std::collections::{HashMap, HashSet};

pub struct Solver(pub String);

type Map = Vec<Vec<char>>;
type Coord = (usize, usize);
impl Solver {
    fn parse_input(&self) -> (Map, Coord) {
        let mut map = Map::new();
        let mut init = (0, 0);
        let mut r = 0;
        let _test = "....#.....\n.........#\n..........\n..#.......\n.......#..\n..........\n.#..^.....\n........#.\n#.........\n......#...";
        for line in self.0.split("\n") {
            map.push(line.chars().collect());
            match map[r].iter().position(|&c| c == '^') {
                Some(c) => init = (r, c),
                None => r += 1,
            }
        }
        return (map, init);
    }
}

#[derive(Clone, Copy, Hash, PartialEq, Eq)]
enum Direction {
    N,
    E,
    S,
    W,
}

impl Direction {
    fn rotate_cw(&self) -> Self {
        match self {
            Direction::N => Direction::E,
            Direction::E => Direction::S,
            Direction::S => Direction::W,
            Direction::W => Direction::N,
        }
    }
}

fn step(map: &Map, pos: &Coord, dir: &Direction) -> Option<(Coord, Direction)> {
    // out of bounds
    if pos.0 == 0 || pos.1 == 0 || pos.0 >= map.len() - 1 || pos.1 >= map[0].len() - 1 {
        return None;
    }
    let new_pos = match dir {
        Direction::N => (pos.0 - 1, pos.1),
        Direction::E => (pos.0, pos.1 + 1),
        Direction::S => (pos.0 + 1, pos.1),
        Direction::W => (pos.0, pos.1 - 1),
    };
    match map[new_pos.0][new_pos.1] {
        '#' => Some((*pos, dir.rotate_cw())),
        _ => Some((new_pos, dir.clone())),
    }
}

/// Gets the number of distinct positions in the map; or returns None if it's a cycle.
fn distinct_pos(
    map: &Map,
    init_pos: &Coord,
    init_dir: &Direction,
) -> Option<HashMap<Coord, HashSet<Direction>>> {
    let (mut curr_pos, mut curr_dir) = (*init_pos, *init_dir);
    let mut visited = HashMap::new();
    visited.insert(*init_pos, HashSet::from([*init_dir]));
    while let Some((next_pos, next_dir)) = step(map, &curr_pos, &curr_dir) {
        (curr_pos, curr_dir) = (next_pos, next_dir);
        // if we try to insert the current direction for this position, and it's already there,
        // it's a cycle, so we return None.
        if let Some(set) = visited.get_mut(&curr_pos) {
            if !set.insert(curr_dir) {
                return None;
            }
        } else {
            visited.insert(curr_pos, HashSet::from([curr_dir]));
        }
    }
    Some(visited)
}

impl super::lib::Puzzle<usize> for Solver {
    async fn part_one(&self) -> usize {
        let (map, init_pos) = self.parse_input();
        let distinct = distinct_pos(&map, &init_pos, &Direction::N).unwrap();
        return distinct.len();
    }

    // this is bad and runs super slow. ill fix it prommy :3
    async fn part_two(&self) -> usize {
        let (map, init_pos) = self.parse_input();
        let visited = distinct_pos(&map, &init_pos, &Direction::N).unwrap();
        let mut v_iter = visited.iter();
        let mut obstacle_set = HashSet::new();
        while let Some((pos, dirs)) = v_iter.next() {
            for dir in dirs {
                let obstacle_pos = match step(&map, pos, dir) {
                    Some((p, _)) => p,
                    None => continue,
                };
                let mut obstacle_map = map.clone();
                let map_loc = &mut obstacle_map[obstacle_pos.0][obstacle_pos.1];
                *map_loc = match *map_loc != '#' {
                    true => '#',
                    false => continue,
                };
                // This isn't great, but it runs okay-ish (about 20s)
                if distinct_pos(&obstacle_map, &init_pos, &Direction::N).is_none() {
                    obstacle_set.insert(obstacle_pos);
                }
            }
        }
        obstacle_set.len()
    }
}
