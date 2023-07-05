use aoc22::day_inputs;
use day5::take1::*;

fn main() {
    let input = day_inputs(5);
    println!("Part 1: {}", part1(&input));
    // println!("Part 2: {}", part2(&input));
}

fn part1(input: &str) -> String {
    let (stacks, instructions) = split_input(input);
    let mut stacks = CargoStacks::from(stacks);
    let instructions = Instructions::from(instructions);

    CrateMove9000::move_crates(&mut stacks, &instructions);
    stacks.tops()
}

fn part2(input: &str) -> String {
    todo!()
}

// #[cfg(test)]
// mod test {
//     const SAMPLE: &str = "\
//     [D]
// [N] [C]
// [Z] [M] [P]
//  1   2   3
//
// move 1 from 2 to 1
// move 3 from 1 to 3
// move 2 from 2 to 1
// move 1 from 1 to 2";
//
//     #[test]
//     fn part1() {
//         assert_eq!(crate::part1(SAMPLE), "CMZ");
//     }
// }
