use itertools::Itertools;
use std::ops::{Deref, DerefMut};

pub fn part1(input: &str) -> String {
    let (stacks, instructions) = crate::split_input(input);
    let mut stacks = CargoStacks::from(stacks);
    let instructions = Instructions::from(instructions);

    CrateMover9000::move_crates(&mut stacks, &instructions);
    stacks.tops()
}

pub fn part2(input: &str) -> String {
    let (stacks, instructions) = crate::split_input(input);
    let mut stacks = CargoStacks::from(stacks);
    let instructions = Instructions::from(instructions);

    CrateMover9001::move_crates(&mut stacks, &instructions);
    stacks.tops()
}

#[derive(Debug, PartialEq)]
struct CargoStacks(Vec<Vec<char>>);

impl CargoStacks {
    fn stack_count(value: &str) -> Option<usize> {
        value
            .lines()
            .rev()
            .next()
            .and_then(|line| line.split_whitespace().last())
            .and_then(|num_str| num_str.parse::<usize>().ok())
    }

    fn tops(&self) -> String {
        let mut ret = String::with_capacity(self.0.len());
        for stack in self.0.iter() {
            if let Some(top) = stack.last() {
                ret.push(*top);
            }
        }
        return ret;
    }
}

impl Deref for CargoStacks {
    type Target = Vec<Vec<char>>;

    fn deref(&self) -> &Self::Target {
        &self.0
    }
}

impl DerefMut for CargoStacks {
    fn deref_mut(&mut self) -> &mut Self::Target {
        &mut self.0
    }
}

impl From<&str> for CargoStacks {
    fn from(value: &str) -> Self {
        let stack_count = CargoStacks::stack_count(value).unwrap();
        let mut stacks: Vec<Vec<char>> = vec![Vec::new(); stack_count];

        value.lines().rev().skip(1).for_each(|line| {
            line.chars()
                .chunks(4)
                .into_iter()
                .enumerate()
                .map(|(idx, chunk)| (idx, chunk.collect::<Vec<_>>()))
                .filter(|(_, chunk)| chunk[1].is_alphabetic())
                .for_each(|(idx, chunk)| stacks[idx].push(chunk[1]))
        });

        Self(stacks)
    }
}

#[derive(Debug, PartialEq)]
struct Instruction {
    count: usize,
    from: usize,
    to: usize,
}

#[derive(Debug)]
struct Instructions(Vec<Instruction>);

impl From<&str> for Instructions {
    fn from(value: &str) -> Self {
        let mut instructions = Vec::new();
        value.lines().for_each(|line| {
            let parts = line
                .split_whitespace()
                .filter_map(|s| s.parse::<usize>().ok())
                .collect::<Vec<_>>();

            instructions.push(Instruction {
                count: parts[0],
                from: parts[1],
                to: parts[2],
            })
        });

        Self(instructions)
    }
}

impl Deref for Instructions {
    type Target = Vec<Instruction>;

    fn deref(&self) -> &Self::Target {
        &self.0
    }
}

trait CrateMover {
    fn move_crates(cargo: &mut CargoStacks, instructions: &Instructions);
}

struct CrateMover9000;

impl CrateMover for CrateMover9000 {
    fn move_crates(cargo: &mut CargoStacks, instructions: &Instructions) {
        for instr in instructions.iter() {
            for _ in 0..instr.count {
                let crate_ = cargo[instr.from - 1].pop().unwrap();
                cargo[instr.to - 1].push(crate_);
            }
        }
    }
}

struct CrateMover9001;

impl CrateMover for CrateMover9001 {
    fn move_crates(cargo: &mut CargoStacks, instructions: &Instructions) {
        for instr in instructions.iter() {
            let from = &mut cargo[instr.from - 1];
            let popped = from
                .drain(from.len() - instr.count..from.len())
                .collect::<Vec<_>>();
            cargo[instr.to - 1].extend(popped);
        }
    }
}

#[cfg(test)]
mod test {
    const INPUT: &str = "
    [D]
[N] [C]
[Z] [M] [P]
 1   2   3

move 1 from 2 to 1
move 3 from 1 to 3
move 2 from 2 to 1
move 1 from 1 to 2";

    const CARGO: &str = "
    [D]
[N] [C]
[Z] [M] [P]
 1   2   3";

    const INSTRUCTIONS: &str = "\
move 1 from 2 to 1
move 3 from 1 to 3
move 2 from 2 to 1
move 1 from 1 to 2";

    use super::*;

    #[test]
    fn stack_count() {
        assert_eq!(CargoStacks::stack_count(CARGO), Some(3));
    }

    #[test]
    fn parse_cargo() {
        assert_eq!(
            *CargoStacks::from(CARGO),
            vec![vec!['Z', 'N'], vec!['M', 'C', 'D'], vec!['P'],]
        );
    }

    #[test]
    fn parse_instructions() {
        assert_eq!(
            Instructions::from(INSTRUCTIONS).0,
            vec![
                Instruction {
                    count: 1,
                    from: 2,
                    to: 1
                },
                Instruction {
                    count: 3,
                    from: 1,
                    to: 3
                },
                Instruction {
                    count: 2,
                    from: 2,
                    to: 1
                },
                Instruction {
                    count: 1,
                    from: 1,
                    to: 2
                },
            ],
        );
    }

    #[test]
    fn cratemover9000_first_instruction() {
        let mut stacks = CargoStacks(vec![vec!['Z', 'N'], vec!['M', 'C', 'D'], vec!['P']]);
        let instructions = Instructions(vec![Instruction {
            count: 1,
            from: 2,
            to: 1,
        }]);

        CrateMover9000::move_crates(&mut stacks, &instructions);

        assert_eq!(
            *stacks,
            vec![vec!['Z', 'N', 'D'], vec!['M', 'C'], vec!['P']]
        );
    }

    #[test]
    fn cratemover9000_whole_sample() {
        let mut stacks = CargoStacks::from(CARGO);
        let instructions = Instructions::from(INSTRUCTIONS);
        CrateMover9000::move_crates(&mut stacks, &instructions);

        assert_eq!(
            *stacks,
            vec![vec!['C'], vec!['M'], vec!['P', 'D', 'N', 'Z']]
        );
    }

    #[test]
    fn cratemover9001_first_instruction() {
        let mut stacks = CargoStacks(vec![vec!['Z', 'N'], vec!['M', 'C', 'D'], vec!['P']]);
        let instructions = Instructions(vec![Instruction {
            count: 1,
            from: 2,
            to: 1,
        }]);

        CrateMover9000::move_crates(&mut stacks, &instructions);

        assert_eq!(
            *stacks,
            vec![vec!['Z', 'N', 'D'], vec!['M', 'C'], vec!['P']]
        );
    }

    #[test]
    fn cratemover9001_first_two_instructions() {
        let mut stacks = CargoStacks(vec![vec!['Z', 'N'], vec!['M', 'C', 'D'], vec!['P']]);
        let instructions = Instructions(vec![
            Instruction {
                count: 1,
                from: 2,
                to: 1,
            },
            Instruction {
                count: 3,
                from: 1,
                to: 3,
            },
        ]);

        CrateMover9001::move_crates(&mut stacks, &instructions);

        assert_eq!(
            *stacks,
            vec![vec![], vec!['M', 'C'], vec!['P', 'Z', 'N', 'D']]
        );
    }

    #[test]
    fn cratemover9001_whole_sample() {
        let mut stacks = CargoStacks::from(CARGO);
        let instructions = Instructions::from(INSTRUCTIONS);
        CrateMover9001::move_crates(&mut stacks, &instructions);

        assert_eq!(
            *stacks,
            vec![vec!['M'], vec!['C'], vec!['P', 'Z', 'N', 'D']]
        );
    }
}
