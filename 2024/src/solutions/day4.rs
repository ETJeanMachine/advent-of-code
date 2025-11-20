pub struct Solver(pub String);

impl Solver {
    fn parse_input(&self) -> Vec<Vec<char>> {
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
        word_search
    }
}

impl super::lib::Puzzle<usize> for Solver {
    async fn part_one(&self) -> usize {
        let word_search = self.parse_input();

        let mut total = 0;
        for i in 0..word_search.len() - 3 {
            for j in 0..word_search[i].len() - 3 {
                let mut substrings: Vec<String> = vec![];
                // The top row of the scan
                substrings.push(word_search[i][j..j + 4].iter().collect());
                // If we're at the bottom, we add the last 4 rows.
                if i == word_search.len() - 4 {
                    for k in i + 1..i + 4 {
                        substrings.push(word_search[k][j..j + 4].iter().collect());
                    }
                }
                // The left-most column of the scan
                substrings.push(word_search[i..i + 4].iter().map(|v| v[j]).collect());
                // If we're at the right-most edge, we add the last 4 columns
                if j == word_search[i].len() - 4 {
                    for k in j + 1..j + 4 {
                        substrings.push(word_search[i..i + 4].iter().map(|v| v[k]).collect());
                    }
                }
                // Adding the diagonals
                let (mut d1, mut d2) = (j, j + 3);
                let diag_one: String = word_search[i..i + 4]
                    .iter()
                    .map(|e| {
                        let res = e[d1];
                        d1 += 1;
                        res
                    })
                    .collect();
                let diag_two: String = word_search[i..i + 4]
                    .iter()
                    .map(|e| {
                        let res = e[d2];
                        if d2 != 0 {
                            d2 -= 1;
                        }
                        res
                    })
                    .collect();
                substrings.extend(vec![diag_one, diag_two]);
                total += substrings
                    .iter()
                    .filter(|&s| s == "XMAS" || s == "SAMX")
                    .count();
            }
        }

        total
    }

    async fn part_two(&self) -> usize {
        let word_search = self.parse_input();

        let mut total = 0;
        for i in 0..word_search.len() - 2 {
            for j in 0..word_search[i].len() - 2 {
                let mut substrings: Vec<String> = vec![];
                // The X's in X-MAS
                let (mut d1, mut d2) = (j, j + 2);
                let diag_one: String = word_search[i..i + 3]
                    .iter()
                    .map(|e| {
                        let res = e[d1];
                        d1 += 1;
                        res
                    })
                    .collect();
                let diag_two: String = word_search[i..i + 3]
                    .iter()
                    .map(|e| {
                        let res = e[d2];
                        if d2 != 0 {
                            d2 -= 1;
                        }
                        res
                    })
                    .collect();
                substrings.extend(vec![diag_one, diag_two]);
                let x_mas = substrings
                    .iter()
                    .filter(|&s| s == "MAS" || s == "SAM")
                    .count();
                if x_mas == 2 {
                    total += 1
                }
            }
        }

        total
    }
}
