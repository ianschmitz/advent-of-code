use std::{collections::HashSet, fs};

fn main() {
    let contents = fs::read_to_string("src/bin/4_1/input.txt")
        .expect("Should have been able to read the file");

    let solution = solve(&contents);

    println!("The answer is: {}", solution);
}

fn solve(input: &str) -> u32 {
    let lines: Vec<&str> = input.trim().lines().collect();

    lines
        .iter()
        .map(|line| {
            let game_delimiter_index = line.find(':').unwrap();

            let list_numbers: Vec<_> = line[game_delimiter_index + 1..]
                .split('|')
                .map(|list| {
                    list.trim()
                        // Replace double space with single to make split cleaner
                        .replace("  ", " ")
                        .split(' ')
                        .map(|num| num.parse::<u32>().unwrap())
                        .collect::<Vec<u32>>()
                })
                .collect();

            let winning_numbers: HashSet<u32> = list_numbers[0].clone().into_iter().collect();
            let player_numbers: HashSet<u32> = list_numbers[1].clone().into_iter().collect();

            let winning_number_count = winning_numbers.intersection(&player_numbers).count();

            match winning_number_count {
                0 => 0,
                _ => {
                    2u32.pow((winning_number_count - 1).try_into().unwrap())
                }
            }
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
Card 1: 41 48 83 86 17 | 83 86  6 31 17  9 48 53
Card 2: 13 32 20 16 61 | 61 30 68 82 17 32 24 19
Card 3:  1 21 53 59 44 | 69 82 63 72 16 21 14  1
Card 4: 41 92 73 84 69 | 59 84 76 51 58  5 54 83
Card 5: 87 83 26 28 32 | 88 30 70 12 93 22 82 36
Card 6: 31 18 13 56 72 | 74 77 10 23 35 67 36 11
";

        assert_eq!(solve(input,), 13);
    }
}
