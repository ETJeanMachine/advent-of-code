use std::collections::HashSet;

pub struct Solver(pub String);

type Map = Vec<Vec<char>>;
type Coord = (usize, usize);
impl Solver {
    fn parse_input(&self) -> (Map, Coord) {
        let mut map = Map::new();
        let mut init = (0, 0);
        let mut r = 0;
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

#[derive(Clone, Copy)]
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

impl super::lib::Puzzle<usize> for Solver {
    async fn part_one(&self) -> usize {
        let (map, init) = self.parse_input();
        let mut distinct_set: HashSet<Coord> = HashSet::from([init]);
        let mut curr_pos = init;
        let mut curr_dir = Direction::N;
        while let Some((pos, dir)) = step(&map, &curr_pos, &curr_dir) {
            (curr_pos, curr_dir) = (pos, dir);
            distinct_set.insert(curr_pos);
        }
        distinct_set.len()
    }

    async fn part_two(&self) -> usize {
        0
    }
}
