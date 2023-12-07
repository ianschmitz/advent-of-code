use std::fs;

fn main() {
    let contents = fs::read_to_string("src/bin/2_2/input.txt")
        .expect("Should have been able to read the file");

    let solution = solve(&contents);

    println!("The answer is: {}", solution);
}

fn solve(input: &str) -> u32 {
    input
        .trim()
        .lines()
        .map(|line| -> u32 {
            let unsorted_color_counts: Vec<_> = line
                .split(':')
                .skip(1)
                // Returns
                .flat_map(|grab| -> Vec<(u32, String)> {
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
                                [count, color, ..] => {
                                    (count.parse::<u32>().unwrap_or(0), color.to_string())
                                }
                                // Handle the case where there are less than two parts
                                _ => (0, "".to_string()),
                            }
                        })
                        .collect()
                })
                .collect();

            let mut sorted_color_counts = unsorted_color_counts.clone();
            sorted_color_counts.sort_by(|(a, _), (b, _)| b.partial_cmp(a).unwrap());

            let mut min_red = 1;
            let mut min_green = 1;
            let mut min_blue = 1;

            for color_count in sorted_color_counts {
                let (count, color) = color_count;
                if color == "red" && count > min_red {
                    min_red = count;
                } else if color == "green" && count > min_green {
                    min_green = count;
                } else if color == "blue" && count > min_blue {
                    min_blue = count;
                }
            }

            min_red * min_green * min_blue
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
Game 1: 3 blue, 4 red; 1 red, 2 green, 6 blue; 2 green
Game 2: 1 blue, 2 green; 3 green, 4 blue, 1 red; 1 green, 1 blue
Game 3: 8 green, 6 blue, 20 red; 5 blue, 4 red, 13 green; 5 green, 1 red
Game 4: 1 green, 3 red, 6 blue; 3 green, 6 red; 3 green, 15 blue, 14 red
Game 5: 6 red, 1 blue, 3 green; 2 blue, 1 red, 2 green
";

        assert_eq!(solve(input,), 2286);
    }
}
