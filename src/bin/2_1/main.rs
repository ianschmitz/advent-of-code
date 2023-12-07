use std::fs;

fn main() {
    let contents = fs::read_to_string("src/bin/2_1/input.txt")
        .expect("Should have been able to read the file");

    let solution = solve(&contents, (12, 13, 14));

    println!("The answer is: {}", solution);
}

fn solve(input: &str, (red_cubes, green_cubes, blue_cubes): (u32, u32, u32)) -> u32 {
    input
        .trim()
        .lines()
        .enumerate()
        .map(|(index, line)| -> (u32, bool) {
            let invalid_grabs: Vec<_> = line
                .split(':')
                .skip(1)
                // Returns
                .flat_map(|grab| -> Vec<(String, String)> {
                    // we don't care about the individual grabs, just if any amount of
                    // cubes is over the limit.
                    grab.replace(';', ",")
                        .split(',')
                        .map(|num_color| {
                            let parts: Vec<String> = num_color
                                .trim()
                                .split(' ')
                                .map(|part| part.trim().to_string())
                                .collect();

                            match parts.as_slice() {
                                [count, color, ..] => (count.to_string(), color.to_string()),
                                // Handle the case where there are less than two parts
                                _ => ("".to_string(), "".to_string()),
                            }
                        })
                        .collect()
                })
                // Returns true if it exceeds the allowed number of cubes and is thus invalid.
                .filter(|(count, color)| {
                    let parsed_count = (*count).parse::<u32>().unwrap();

                    match color.as_str() {
                        "red" => parsed_count > red_cubes,
                        "green" => parsed_count > green_cubes,
                        "blue" => parsed_count > blue_cubes,
                        _ => true,
                    }
                })
                .collect();

            // If there are any counts that exceeded the allowed number of cubes, this game is
            // invalid.
            let is_valid = invalid_grabs.is_empty();

            // Games start at index 1.
            ((index as u32) + 1, is_valid)
        })
        .filter(|(_index, valid)| *valid)
        .map(|(index, _valid)| index)
        .sum::<u32>()
}

#[cfg(test)]
mod tests {
    // Note this useful idiom: importing names from outer (for mod tests) scope.
    use super::*;

    #[test]
    fn test() {
        let input = "
Game 1: 3 blue, 4 red; 1 red, 2 green, 6 blue; 2 green
Game 2: 1 blue, 2 green; 3 green, 4 blue, 1 red; 1 green, 1 blue
Game 3: 8 green, 6 blue, 20 red; 5 blue, 4 red, 13 green; 5 green, 1 red
Game 4: 1 green, 3 red, 6 blue; 3 green, 6 red; 3 green, 15 blue, 14 red
Game 5: 6 red, 1 blue, 3 green; 2 blue, 1 red, 2 green
";

        assert_eq!(solve(input, (12, 13, 14)), 8);
    }
}
