pub struct Solver(pub String);

fn scan_block(block: [[char; 4]; 4]) -> u8 {
    let mut substrings: Vec<String> = vec![];

    let diag_one = [block[0][0], block[1][1], block[2][2], block[3][3]];
    let diag_two = [block[0][3], block[1][2], block[2][1], block[3][0]];
    substrings.push(diag_one.iter().collect());
    substrings.push(diag_two.iter().collect());

    for i in 0..block.len() {
        substrings.push(block[i].iter().collect());
        let col = [block[0][i], block[1][i], block[2][i], block[3][i]];
        substrings.push(col.iter().collect());
    }
    substrings.iter().fold(0, |acc, e| {
        acc + match e.as_str() {
            "XMAS" | "SMAX" => 1,
            _ => 0,
        }
    })
}

impl super::lib::Puzzle<i32> for Solver {
    async fn part_one(&self) -> i32 {
        let mut word_search = vec![vec![]];
        let mut row = 0;

        for c in self.0.chars() {
            if c == '\n' {
                word_search.push(vec![]);
                row += 1;
            } else {
                word_search[row].push(c);
            }
        }

        for i in 0..word_search.len() - 4 {
            for j in 0..word_search[i].len() - 4 {
                let mut block = [['.'; 4]; 4];
            }
        }

        println!("{:?}", word_search);

        0
    }

    async fn part_two(&self) -> i32 {
        0
    }
}
