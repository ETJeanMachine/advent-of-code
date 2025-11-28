use std::{collections::HashSet, str::FromStr};

pub struct Farm {
    farm_map: Vec<Vec<char>>,
    height: usize,
    width: usize,
}

impl FromStr for Farm {
    type Err = String;

    fn from_str(s: &str) -> Result<Self, Self::Err> {
        let mut farm_map: Vec<Vec<char>> = vec![];
        s.split("\n")
            .for_each(|line| farm_map.push(line.chars().collect()));
        let (height, width) = (farm_map.len(), farm_map[0].len());
        Ok(Self {
            farm_map,
            height,
            width,
        })
    }
}

#[derive(PartialEq, Eq, Clone, Copy, Hash)]
enum Direction {
    NORTH,
    SOUTH,
    EAST,
    WEST,
}

impl Farm {
    fn neighbouring_plants(
        &self,
        plant: char,
        pos: (usize, usize),
    ) -> Vec<((usize, usize), Direction)> {
        let mut neighbours = vec![];
        let farm_map = &self.farm_map;
        let (row, col) = pos;
        if row > 0 && farm_map[row - 1][col] == plant {
            neighbours.push(((row - 1, col), Direction::NORTH));
        }
        if row < self.height - 1 && farm_map[row + 1][col] == plant {
            neighbours.push(((row + 1, col), Direction::SOUTH));
        }
        if col > 0 && farm_map[row][col - 1] == plant {
            neighbours.push(((row, col - 1), Direction::WEST));
        }
        if col < self.width - 1 && farm_map[row][col + 1] == plant {
            neighbours.push(((row, col + 1), Direction::EAST));
        }
        neighbours
    }

    fn next_undiscovered(
        &self,
        curr_row: usize,
        discovered: &HashSet<(usize, usize)>,
    ) -> Option<(usize, usize)> {
        for row in curr_row..self.height {
            for col in 0..self.width {
                if !discovered.contains(&(row, col)) {
                    return Some((row, col));
                }
            }
        }
        None
    }

    fn region_area_and_perimeter(
        &self,
        discovered: &mut HashSet<(usize, usize)>,
        curr: (usize, usize),
    ) -> (u32, u32) {
        discovered.insert(curr);
        let plant = self.farm_map[curr.0][curr.1];
        let neighbours = self.neighbouring_plants(plant, curr);
        let (mut perimeter, mut area) = (4 - neighbours.len() as u32, 1);
        for (neighbour, _) in neighbours {
            if !discovered.contains(&neighbour) {
                let (next_perimeter, next_area) =
                    self.region_area_and_perimeter(discovered, neighbour);
                perimeter += next_perimeter;
                area += next_area;
            }
        }
        (perimeter, area)
    }

    pub fn total_region_cost_by_perimeter(&self) -> u32 {
        let mut total_cost = 0;
        let mut discovered = HashSet::new();
        let mut curr_row = 0;
        while let Some(curr) = self.next_undiscovered(curr_row, &discovered) {
            curr_row = curr.0;
            let (perimeter, area) = self.region_area_and_perimeter(&mut discovered, curr);
            total_cost += perimeter * area;
        }
        total_cost
    }

    fn region_area_and_sides(
        &self,
        discovered: &mut HashSet<(usize, usize)>,
        sides: &mut HashSet<(usize, Direction)>,
        curr: (usize, usize),
    ) -> u32 {
        discovered.insert(curr);
        let (row, col) = curr;
        let plant = self.farm_map[row][col];
        let neighbours = self.neighbouring_plants(plant, curr);
        let mut area = 1;
        let non_side_dirs: Vec<_> = neighbours.iter().map(|(_, dir)| *dir).collect();
        let side_dirs: Vec<_> = vec![
            Direction::NORTH,
            Direction::SOUTH,
            Direction::EAST,
            Direction::WEST,
        ]
        .iter()
        .filter(|dir| !non_side_dirs.contains(dir))
        .map(|dir| *dir)
        .collect();
        for dir in side_dirs {
            sides.insert(match dir {
                Direction::NORTH | Direction::SOUTH => (row, dir),
                Direction::EAST | Direction::WEST => (col, dir),
            });
        }
        for (neighbour, _) in neighbours {
            if !discovered.contains(&neighbour) {
                let next_area = self.region_area_and_sides(discovered, sides, neighbour);
                area += next_area;
            }
        }
        area
    }

    pub fn total_region_cost_by_sides(&self) -> u32 {
        let mut total_cost = 0;
        let mut discovered = HashSet::new();
        let mut curr_row = 0;
        while let Some(curr) = self.next_undiscovered(curr_row, &discovered) {
            curr_row = curr.0;
            let mut sides = HashSet::new();
            let area = self.region_area_and_sides(&mut discovered, &mut sides, curr);
            total_cost += sides.len() as u32 * area
        }
        total_cost
    }
}

pub struct Solver(pub String);

impl super::lib::Puzzle<u32> for Solver {
    async fn part_one(&self) -> u32 {
        let farm = Farm::from_str(self.0.as_str()).unwrap();
        farm.total_region_cost_by_perimeter()
    }

    async fn part_two(&self) -> u32 {
        let farm = Farm::from_str(self.0.as_str()).unwrap();
        farm.total_region_cost_by_sides()
    }
}
