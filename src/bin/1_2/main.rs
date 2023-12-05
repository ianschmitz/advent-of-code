use regex::Regex;
use std::fs;

fn main() {
    let contents = fs::read_to_string("src/bin/1_2/input.txt")
        .expect("Should have been able to read the file");

    let solution = solve(&contents);

    println!("The answer is: {}", solution);
}

fn word_to_num(word: &str) -> Result<u32, &'static str> {
    match word {
        "one" => Ok(1),
        "two" => Ok(2),
        "three" => Ok(3),
        "four" => Ok(4),
        "five" => Ok(5),
        "six" => Ok(6),
        "seven" => Ok(7),
        "eight" => Ok(8),
        "nine" => Ok(9),
        _ => Err("No matching number found"),
    }
}

fn solve(input: &str) -> u32 {
    let num_regex = Regex::new(r"\d|one|two|three|four|five|six|seven|eight|nine").unwrap();

    input
        .trim()
        .lines()
        .map(|line| {
            let matches: Vec<&str> = num_regex.find_iter(line).map(|x| x.as_str()).collect();

            let default = "0";
            let first = matches.first().unwrap_or(&default);
            let last = matches.last().unwrap_or(&default);

            let first_num = match first.parse::<u32>() {
                Ok(num) => num,
                Err(_) => word_to_num(first).unwrap(),
            };

            let last_num = match last.parse::<u32>() {
                Ok(num) => num,
                Err(_) => word_to_num(last).unwrap(),
            };

            format!("{}{}", first_num, last_num)
        })
        .map(|x| x.parse::<u32>().unwrap_or(0))
        .sum::<u32>()
}

#[cfg(test)]
mod tests {
    // Note this useful idiom: importing names from outer (for mod tests) scope.
    use super::*;

    #[test]
    fn test() {
        let input = "
two1nine
eightwothree
abcone2threexyz
xtwone3four
4nineeightseven2
zoneight234
7pqrstsixteen";

        assert_eq!(solve(input), 281);
    }
}
