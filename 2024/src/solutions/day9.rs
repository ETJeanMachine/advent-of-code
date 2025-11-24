use std::{fmt, str::FromStr};

#[derive(Debug, Clone)]
pub struct DiskMap {
    blocks: Vec<(Option<usize>, usize)>,
    size: usize,
    free_space: usize,
}

impl FromStr for DiskMap {
    type Err = String;

    fn from_str(s: &str) -> Result<Self, Self::Err> {
        let mut blocks = vec![];
        let (mut size, mut free_space) = (0, 0);
        let mut id = 0;
        let mut is_file = true;
        for c in s.chars() {
            let block_size = c.to_digit(10).ok_or("Non-digit in input")? as usize;
            if is_file {
                blocks.push((Some(id), block_size));
                id += 1;
            } else {
                blocks.push((None, block_size));
                free_space += 1;
            }
            size += block_size;
            is_file = !is_file;
        }
        Ok(DiskMap {
            blocks,
            size,
            free_space,
        })
    }
}

impl fmt::Display for DiskMap {
    fn fmt(&self, f: &mut fmt::Formatter<'_>) -> fmt::Result {
        let file_count = self.blocks.iter().filter(|x| x.0.is_some()).count();
        write!(
            f,
            "{{File Count: {}, Disk Size: {}, Free Space: {}}}",
            file_count, self.size, self.free_space
        )
    }
}

impl DiskMap {
    fn fill_first_empty(&mut self) -> bool {
        let first_empty = self.blocks.iter().position(|x| x.0.is_none());
        if let Some(i) = first_empty {
            let mut file_iter = self.blocks.iter_mut().rev().filter(|x| x.0.is_some());
            while self.blocks[i].0.is_none() {
                let mut next_file = file_iter.next().unwrap();
            }
        }
        false
    }

    pub fn compress(&mut self) {}

    pub fn checksum(&self) -> usize {
        0
    }
}

pub struct Solver(pub String);

impl super::lib::Puzzle<usize> for Solver {
    async fn part_one(&self) -> usize {
        let mut diskmap = DiskMap::from_str(self.0.as_str()).unwrap();
        println!("{}", diskmap);
        diskmap.compress();
        diskmap.checksum()
    }

    async fn part_two(&self) -> usize {
        let diskmap = DiskMap::from_str(self.0.as_str()).unwrap();
        0
    }
}
