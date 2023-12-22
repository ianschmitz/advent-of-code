use std::fs;

fn main() {
    let contents = fs::read_to_string("src/bin/3_1/input.txt")
        .expect("Should have been able to read the file");

    let solution = solve(&contents);

    println!("The answer is: {}", solution);
}

fn solve(input: &str) -> u32 {
    let num_regex = regex::Regex::new(r"\d+").unwrap();
    let symbol_regex = regex::Regex::new(r"[^\.\d\w\s]").unwrap();

    let lines: Vec<&str> = input.trim().lines().collect();

    lines
        .iter()
        .enumerate()
        .map(|(index, line)| -> u32 {
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

            num_regex.find_iter(line).fold(0, |acc, num| {
                let part_number = num.as_str().parse::<u32>().unwrap();

                let prev_line_slice = if let Some(prev_line) = prev_line {
                    &prev_line[num.start().saturating_sub(1)
                        ..(std::cmp::min(num.end() + 1, line.len() - 1))]
                } else {
                    ""
                };

                let next_line_slice = if let Some(next_line) = next_line {
                    &next_line[num.start().saturating_sub(1)
                        ..(std::cmp::min(num.end() + 1, line.len() - 1))]
                } else {
                    ""
                };

                let line_chars: Vec<char> = line.chars().collect();

                let preceding_char = if num.start() > 0 {
                    line_chars[num.start() - 1]
                } else {
                    '.'
                };

                let following_char = if num.end() < line_chars.len() {
                    line_chars[num.end()]
                } else {
                    '.'
                };

                let all_surrounding_chars = format!(
                    "{} {} {} {}",
                    prev_line_slice, preceding_char, following_char, next_line_slice,
                );

                let is_part_number = symbol_regex.is_match(&all_surrounding_chars);

                match is_part_number {
                    true => acc + part_number,
                    false => acc,
                }
            })
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

        assert_eq!(solve(input,), 4361);
    }
}
