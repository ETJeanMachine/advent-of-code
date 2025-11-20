pub struct Solver(pub String);

impl super::lib::Puzzle<i32> for Solver {
    async fn part_one(&self) -> i32 {
        let mut reports = vec![];

        // Consumes the iterator, returns an (Optional) String
        for line in self.0.split("\n") {
            if line.is_empty() {
                break;
            }
            let mut split = line.split_whitespace();
            reports.push(vec![]);
            while let Some(s) = split.next() {
                let n = s.parse::<i32>().unwrap();
                reports.last_mut().unwrap().push(n);
            }
        }

        let mut safe_count = reports.len() as i32;
        for v in reports {
            let mut dir = 0;
            for i in 0..v.len() - 1 {
                let step = v[i + 1] - v[i];
                // setting the initial direction based on the first step
                if dir == 0 {
                    dir = step.signum();
                }
                if dir != step.signum() || step.abs() > 3 || step.abs() == 0 {
                    safe_count -= 1;
                    break;
                }
            }
        }
        safe_count
    }

    async fn part_two(&self) -> i32 {
        let mut reports = vec![];

        // Consumes the iterator, returns an (Optional) String
        for line in self.0.split("\n") {
            let split = line.split_whitespace();
            if line.is_empty() {
                break;
            }
            reports.push(vec![]);
            for s in split {
                let n = s.parse::<i32>().unwrap();
                reports.last_mut().unwrap().push(n);
            }
        }

        let mut safe_count = reports.len() as i32;
        for v in reports {
            let slopes: Vec<i32> = v.windows(2).map(|n| n[1] - n[0]).collect();
            let sign = slopes.iter().sum::<i32>().signum();
            let valid = |n: i32| -> bool { n.signum() == sign && (1..=3).contains(&n.abs()) };
            let problem_idx = slopes
                .iter()
                .enumerate()
                .find(|&(_, &n)| !valid(n))
                .map(|(i, _)| i);
            if let Some(i) = problem_idx {
                let mut shift_right = slopes.clone();
                let right_step = shift_right.remove(i);
                if i < shift_right.len() {
                    shift_right[i] += right_step;
                }
                let mut shift_left = slopes.clone();
                let left_step = shift_left.remove(i);
                if i > 0 {
                    shift_left[i - 1] += left_step;
                }
                let right_problem = shift_right.iter().find(|&&n| !valid(n));
                let left_problem = shift_left.iter().find(|&&n| !valid(n));
                if right_problem.is_some() && left_problem.is_some() {
                    safe_count -= 1;
                }
            }
        }
        safe_count
    }
}
