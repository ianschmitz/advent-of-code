use regex::Regex;
use std::fs;

fn main() {
    let contents = fs::read_to_string("src/bin/1_1/input.txt")
        .expect("Should have been able to read the file");

    let solution = solve(&contents);

    println!("The answer is: {}", solution);
}

fn solve(input: &str) -> u32 {
    let num_regex = Regex::new(r"\d").unwrap();

    input
        .trim()
        .lines()
        .map(|line| {
            let matches: Vec<&str> = num_regex.find_iter(line).map(|x| x.as_str()).collect();

            let first = matches.first().unwrap().to_string();
            let last = matches.last().unwrap().to_string();

            first + &last
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
1abc2
pqr3stu8vwx
a1b2c3d4e5f
treb7uchet";

        assert_eq!(solve(input), 142);
    }
}
