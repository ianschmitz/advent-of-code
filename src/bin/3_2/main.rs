use std::fs;

fn main() {
    let contents = fs::read_to_string("src/bin/3_2/input.txt")
        .expect("Should have been able to read the file");

    let solution = solve(&contents);

    println!("The answer is: {}", solution);
}

fn solve(input: &str) -> u32 {
    let num_regex = regex::Regex::new(r"\d+").unwrap();
    let gear_symbol_regex = regex::Regex::new(r"\*").unwrap();

    let lines: Vec<&str> = input.trim().lines().collect();

    lines
        .iter()
        .enumerate()
        .flat_map(|(index, line)| -> Vec<u32> {
            let prev_line = if index > 0 {
                Some(lines[index - 1])
            } else {
                None
            };

            let next_line = if index + 1 < lines.len() {
                Some(lines[index + 1])
            } else {
                None
            };

            gear_symbol_regex
                .find_iter(line)
                .map(|gear_symbol| {
                    let current_line_matches = num_regex
                        .find_iter(line)
                        .filter(|num| {
                            return gear_symbol.start() == num.end()
                                || gear_symbol.end() == num.start();
                        })
                        .map(|num| num.as_str())
                        .collect();

                    let adjacent_prev_line_matches = if let Some(prev_line) = prev_line {
                        num_regex
                            .find_iter(prev_line)
                            .filter(|num| {
                                return gear_symbol.start() >= num.start().saturating_sub(1)
                                    && gear_symbol.start() <= num.end();
                            })
                            .map(|num| num.as_str())
                            .collect()
                    } else {
                        vec![]
                    };

                    let adjacent_next_line_matches = if let Some(next_line) = next_line {
                        num_regex
                            .find_iter(next_line)
                            .filter(|num| {
                                return gear_symbol.start() >= num.start().saturating_sub(1)
                                    && gear_symbol.start() <= num.end();
                            })
                            .map(|num| num.as_str())
                            .collect()
                    } else {
                        vec![]
                    };

                    let mut found_nums: Vec<_> = ([
                        adjacent_prev_line_matches,
                        current_line_matches,
                        adjacent_next_line_matches,
                    ])
                    .concat()
                    .iter()
                    .map(|num| num.parse::<u32>().unwrap_or(0))
                    .collect();

                    found_nums.sort();
                    found_nums.reverse();

                    let first = found_nums.get(0);
                    let second = found_nums.get(1);

                    match (first, second) {
                        (Some(first), Some(second)) => first * second,
                        _ => 0,
                    }
                })
                .collect()
        })
        .sum::<u32>()
}

#[cfg(test)]
mod tests {
    // Note this useful idiom: importing names from outer (for mod tests) scope.
    use super::*;

    #[test]
    fn test() {
        let input = "
467..114..
...*......
..35..633.
......#...
617*......
.....+.58.
..592.....
......755.
...$.*....
.664.598..
";

        assert_eq!(solve(input,), 467835);
    }
}
