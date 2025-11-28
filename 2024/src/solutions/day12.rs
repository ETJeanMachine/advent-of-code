use std::{
    collections::{BinaryHeap, HashMap, HashSet},
    str::FromStr,
};

#[derive(PartialEq, Eq, Clone, Copy, Hash)]
enum Direction {
    NORTH,
    SOUTH,
    EAST,
    WEST,
}

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

impl Farm {
    fn neighbouring_plants(
        &self,
        plant: char,
        pos: (usize, usize),
    ) -> Vec<((usize, usize), Direction)> {
        let mut neighbours = vec![];
        let farm_map = &self.farm_map;
        let (row, col) = pos;
        if row < self.height - 1 && farm_map[row + 1][col] == plant {
            neighbours.push(((row + 1, col), Direction::SOUTH));
        }
        if col < self.width - 1 && farm_map[row][col + 1] == plant {
            neighbours.push(((row, col + 1), Direction::EAST));
        }
        if row > 0 && farm_map[row - 1][col] == plant {
            neighbours.push(((row - 1, col), Direction::NORTH));
        }
        if col > 0 && farm_map[row][col - 1] == plant {
            neighbours.push(((row, col - 1), Direction::WEST));
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

    fn add_to_sides(
        &self,
        sides: &mut HashMap<(usize, Direction), BinaryHeap<usize>>,
        curr: (usize, usize),
        neighbours: &Vec<((usize, usize), Direction)>,
    ) {
        let (row, col) = curr;
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
            let (side, index) = match dir {
                Direction::NORTH | Direction::SOUTH => (row, col),
                Direction::EAST | Direction::WEST => (col, row),
            };
            if let Some(heap) = sides.get_mut(&(side, dir)) {
                heap.push(index);
            } else {
                sides.insert((side, dir), BinaryHeap::from([index]));
            }
        }
    }

    fn region_area_and_sides(
        &self,
        discovered: &mut HashSet<(usize, usize)>,
        sides: &mut HashMap<(usize, Direction), BinaryHeap<usize>>,
        curr: (usize, usize),
    ) -> u32 {
        discovered.insert(curr);
        let (row, col) = curr;
        let plant = self.farm_map[row][col];
        let neighbours = self.neighbouring_plants(plant, curr);
        let mut area = 1;
        self.add_to_sides(sides, curr, &neighbours);
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
            let mut sides = HashMap::new();
            let area = self.region_area_and_sides(&mut discovered, &mut sides, curr);
            let mut total_sides = 0;
            for heap in sides.values_mut() {
                total_sides += 1;
                let mut last_index = heap.pop().unwrap();
                while let Some(index) = heap.pop() {
                    if index.abs_diff(last_index) != 1 {
                        total_sides += 1;
                    }
                    last_index = index;
                }
            }
            total_cost += total_sides * area
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
