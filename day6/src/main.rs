use std::collections::HashSet;

use aoc22::day_inputs;


fn main() {
    let input = day_inputs(6);
    println!("Part 1: {}", naive(&input, 4).unwrap());
    println!("Part 2: {}", naive(&input, 14).unwrap());
}

fn naive(input: &str, win_size: usize) -> Option<usize> {
    input
        .as_bytes()
        .windows(win_size)
        .position(|window| window.iter().collect::<HashSet<_>>().len() == win_size)
        .map(|pos| pos + win_size)
}

#[cfg(test)]
mod tests {
    use super::*;

    #[test]
    fn part1_naive() {
        assert_eq!(Some(7), naive("mjqjpqmgbljsphdztnvjfqwrcgsmlb", 4));
        assert_eq!(Some(5), naive("bvwbjplbgvbhsrlpgdmjqwftvncz", 4));
        assert_eq!(Some(6), naive("nppdvjthqldpwncqszvftbrmjlhg", 4));
        assert_eq!(Some(10), naive("nznrnfrfntjfmvfwmzdfjlvtqnbhcprsg", 4));
        assert_eq!(Some(11), naive("zcfzfwzzqfrljwzlrfnpqdbhtmscgvjw", 4));
    }

    #[test]
    fn part2_naive() {
        assert_eq!(Some(19), naive("mjqjpqmgbljsphdztnvjfqwrcgsmlb", 14));
        assert_eq!(Some(23), naive("bvwbjplbgvbhsrlpgdmjqwftvncz", 14));
        assert_eq!(Some(23), naive("nppdvjthqldpwncqszvftbrmjlhg", 14));
        assert_eq!(Some(29), naive("nznrnfrfntjfmvfwmzdfjlvtqnbhcprsg", 14));
        assert_eq!(Some(26), naive("zcfzfwzzqfrljwzlrfnpqdbhtmscgvjw", 14));
    }
}
