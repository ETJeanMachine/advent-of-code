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
    fn first_empty(&self) -> Option<usize> {
        self.blocks.iter().position(|x| x.0.is_none())
    }

    fn last_file(&self) -> Option<usize> {
        if let Some(pos) = self.blocks.iter().rev().position(|x| x.0.is_some()) {
            return Some(self.blocks.len() - pos - 1);
        }
        None
    }

    pub fn compress(&mut self) {
        while let Some(empty_idx) = self.first_empty() {
            let empty_size = self.blocks[empty_idx].1;
            if empty_size == 0 {
                self.blocks.remove(empty_idx);
            } else if let Some(file_idx) = self.last_file() {
                let (file_id, file_size) = self.blocks[file_idx];
                let file_id = file_id.unwrap();
                if file_size <= empty_size {
                    self.blocks[empty_idx].1 -= file_size;
                    self.blocks.remove(file_idx);
                    self.blocks.insert(empty_idx, (Some(file_id), file_size));
                } else {
                    self.blocks[empty_idx].1 = 0;
                    self.blocks[file_idx].1 -= empty_size;
                    self.blocks.insert(empty_idx, (Some(file_id), empty_size));
                }
            }
        }
    }

    fn first_empty_fits(&self, size: usize) -> Option<usize> {
        self.blocks
            .iter()
            .position(|(block, block_size)| block.is_none() && *block_size >= size)
    }

    fn file_loc(&self, id: usize) -> Option<usize> {
        self.blocks.iter().position(|(block, _)| {
            if let Some(file_id) = block {
                *file_id == id
            } else {
                false
            }
        })
    }

    fn clear_space(&mut self, idx: usize) {
        let mut idx = idx;
        self.blocks[idx] = (None, self.blocks[idx].1);
        if let Some(n) = idx.checked_sub(1)
            && self.blocks[n].0.is_none()
        {
            self.blocks[idx].1 += self.blocks[n].1;
            self.blocks.remove(n);
            idx -= 1;
        }
        if idx + 1 < self.blocks.len() && self.blocks[idx + 1].0.is_none() {
            self.blocks[idx].1 += self.blocks[idx + 1].1;
            self.blocks.remove(idx + 1);
        }
    }

    pub fn compact(&mut self) {
        let final_idx = self.last_file().unwrap();
        let final_id = self.blocks[final_idx].0.unwrap();
        let mut file_id = final_id;
        loop {
            if let Some(file_idx) = self.file_loc(file_id) {
                file_id = self.blocks[file_idx].0.unwrap();
                let file_size = self.blocks[file_idx].1;
                if let Some(empty_idx) = self.first_empty_fits(file_size)
                    && empty_idx < file_idx
                {
                    self.blocks[empty_idx].1 -= file_size;
                    self.clear_space(file_idx);
                    self.blocks.insert(empty_idx, (Some(file_id), file_size));
                }
            }
            file_id = match file_id.checked_sub(1) {
                Some(n) => n,
                None => break,
            };
        }
    }

    pub fn checksum(&self) -> usize {
        let mut final_pos = 0;
        let mut checksum = 0;
        for (id, size) in self.blocks.iter() {
            let init_pos = final_pos;
            final_pos = init_pos + size;
            if let Some(id) = id {
                checksum += id * (init_pos..final_pos).fold(0, |acc, x| acc + x);
            }
        }
        checksum
    }
}

pub struct Solver(pub String);

impl super::lib::Puzzle<usize> for Solver {
    async fn part_one(&self) -> usize {
        let mut diskmap = DiskMap::from_str(self.0.as_str()).unwrap();
        diskmap.compress();
        diskmap.checksum()
    }

    async fn part_two(&self) -> usize {
        let _test = "2333133121414131402";
        let mut diskmap = DiskMap::from_str(self.0.as_str()).unwrap();
        diskmap.compact();
        diskmap.checksum()
    }
}
