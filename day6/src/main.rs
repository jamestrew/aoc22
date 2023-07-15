use aoc22::day_inputs;
use itertools::Itertools;

fn main() {
    let input = day_inputs(6);
    println!("Part 1: {}", blazingly_fast(&input, 4).unwrap());
    println!("Part 2: {}", blazingly_fast(&input, 14).unwrap());
}

#[allow(dead_code)]
fn naive(input: &str, win_size: usize) -> Option<usize> {
    input
        .as_bytes()
        .windows(win_size)
        .position(|window| window.iter().unique().count() == win_size)
        .map(|pos| pos + win_size)
}

fn blazingly_fast(input: &str, win_size: usize) -> Option<usize> {
    let mut idx = 0;
    while let Some(slice) = input.as_bytes().get(idx..idx + win_size) {
        let mut state = 0u32;

        if let Some(pos) = slice.iter().rposition(|byte| {
            let bit_idx = byte % 32;
            let is_dup = state & (1 << bit_idx) != 0;
            state |= 1 << bit_idx;
            is_dup
        }) {
            idx += pos + 1;
        } else {
            return Some(idx + win_size);
        }
    }
    None
}

#[cfg(test)]
mod tests {
    use super::*;
    use test_case::test_case;

    #[test_case(7, "mjqjpqmgbljsphdztnvjfqwrcgsmlb")]
    #[test_case(5, "bvwbjplbgvbhsrlpgdmjqwftvncz")]
    #[test_case(6, "nppdvjthqldpwncqszvftbrmjlhg")]
    #[test_case(10, "nznrnfrfntjfmvfwmzdfjlvtqnbhcprsg")]
    #[test_case(11, "zcfzfwzzqfrljwzlrfnpqdbhtmscgvjw")]
    fn part1(expected: usize, input: &str) {
        assert_eq!(Some(expected), naive(input, 4));
        assert_eq!(Some(expected), blazingly_fast(input, 4));
    }

    #[test_case(19, "mjqjpqmgbljsphdztnvjfqwrcgsmlb")]
    #[test_case(23, "bvwbjplbgvbhsrlpgdmjqwftvncz")]
    #[test_case(23, "nppdvjthqldpwncqszvftbrmjlhg")]
    #[test_case(29, "nznrnfrfntjfmvfwmzdfjlvtqnbhcprsg")]
    #[test_case(26, "zcfzfwzzqfrljwzlrfnpqdbhtmscgvjw")]
    fn part2(expected: usize, input: &str) {
        assert_eq!(Some(expected), naive(input, 14));
        assert_eq!(Some(expected), blazingly_fast(input, 14));
    }
}
