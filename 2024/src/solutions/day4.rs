pub struct Solver(pub String);

type Point = (usize, usize);

fn scan_block(block: Vec<&[char]>) -> Vec<(Point, Point)> {
    for r in block.as_slice() {
        println!("{:?}", r);
    }
    println!();
    let mut substrings: Vec<(Point, Point, String)> = vec![];

    let diag_one = [block[0][0], block[1][1], block[2][2], block[3][3]];
    let diag_two = [block[0][3], block[1][2], block[2][1], block[3][0]];
    substrings.push(((0, 0), (3, 3), diag_one.iter().collect()));
    substrings.push(((0, 3), (3, 0), diag_two.iter().collect()));

    for i in 0..block.len() {
        substrings.push(((i, 0), (i, 3), block[i].iter().collect()));
        let col = [block[0][i], block[1][i], block[2][i], block[3][i]];
        substrings.push(((0, i), (3, i), col.iter().collect()));
    }
    let mut ranges = vec![];
    for sub in substrings.clone() {
        match sub.2.as_str() {
            "XMAS" | "SAMX" => ranges.push((sub.0, sub.1)),
            _ => (),
        }
    }
    return ranges;
}

impl super::lib::Puzzle<usize> for Solver {
    async fn part_one(&self) -> usize {
        let mut word_search = vec![vec![]];
        let mut row = 0;

        let test = "MMMSXXMASM\nMSAMXMSMSA\nAMXSXMAAMM\nMSAMASMSMX\nXMASAMXAMM\nXXAMMXXAMA\nSMSMSASXSS\nSAXAMASAAA\nMAMMMXMMMM\nMXMXAXMASX";

        for c in test.chars() {
            if c == '\n' {
                word_search.push(vec![]);
                row += 1;
            } else {
                word_search[row].push(c);
            }
        }

        let mut total = 0;

        for i in 0..word_search.len() - 4 {
            for j in 0..word_search[i].len() - 4 {
                let rows = &word_search[i..i + 4];
                let block: Vec<&[char]> = rows.iter().map(|v| &v[j..j + 4]).collect();
                let ranges = scan_block(block);
                total += ranges.len();
            }
        }

        total
    }

    async fn part_two(&self) -> usize {
        0
    }
}
