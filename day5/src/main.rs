use aoc22::day_inputs;
use day5::take1;

fn main() {
    let input = day_inputs(5);
    println!("Part 1: {}", part1(&input));
    println!("Part 2: {}", part2(&input));
}

fn part1(input: &str) -> String {
    take1::part1(input)
}

fn part2(input: &str) -> String {
    take1::part2(input)
}
