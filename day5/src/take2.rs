use std::{fmt::Debug, ops::Deref};

use itertools::Itertools;
use nom::{
    branch::alt,
    bytes::complete::{tag, take},
    character::complete::digit1,
    combinator::{all_consuming, map, map_res},
    multi::separated_list1,
    sequence::{delimited, preceded, tuple},
    Finish, IResult,
};

pub fn part1(input: &str) -> String {
    let (stacks, instructions) = crate::split_input(input);
    let mut piles = Piles::from(stacks);
    let instructions = parse_instruction_lines(instructions);

    piles.apply_instructions(instructions);
    piles.tops()
}

pub fn part2(_input: &str) -> String {
    todo!()
}

#[derive(Debug, PartialEq, Clone, Copy)]
pub struct Crate(char);

impl std::fmt::Display for Crate {
    fn fmt(&self, f: &mut std::fmt::Formatter<'_>) -> std::fmt::Result {
        Debug::fmt(self, f)
    }
}

fn parse_crate(i: &str) -> IResult<&str, Crate> {
    let first_char = |s: &str| Crate(s.chars().next().unwrap());
    let f = delimited(tag("["), take(1_usize), tag("]"));
    map(f, first_char)(i)
}

fn parse_hole(i: &str) -> IResult<&str, ()> {
    map(tag("   "), drop)(i)
}

fn parse_hole_or_crate(i: &str) -> IResult<&str, Option<Crate>> {
    alt((map(parse_crate, Some), map(parse_hole, |_| None)))(i)
}

fn parse_crate_line(i: &str) -> IResult<&str, Vec<Option<Crate>>> {
    separated_list1(tag(" "), parse_hole_or_crate)(i)
}

pub fn parse_crate_lines(i: &str) -> Vec<Vec<Crate>> {
    let mut crate_lines = vec![];
    i.lines().for_each(|line| {
        if let Ok((_rest, crate_line)) = all_consuming(parse_crate_line)(line).finish() {
            crate_lines.push(crate_line)
        }
    });

    transpose_rev(crate_lines)
}

fn transpose_rev<T>(v: Vec<Vec<Option<T>>>) -> Vec<Vec<T>> {
    assert!(!v.is_empty());
    let len = v.iter().map(|inner| inner.len()).max().unwrap();
    let mut iters = v.into_iter().map(|n| n.into_iter()).collect::<Vec<_>>();
    (0..len)
        .map(|_| {
            iters
                .iter_mut()
                .rev()
                .filter_map(|n| n.next().unwrap_or_default())
                .collect::<Vec<T>>()
        })
        .collect()
}

#[derive(Debug, PartialEq)]
struct Instruction {
    count: usize,
    from: usize,
    to: usize,
}

fn parse_number(i: &str) -> IResult<&str, usize> {
    map_res(digit1, |s: &str| s.parse::<usize>())(i)
}

fn parse_piles_number(i: &str) -> IResult<&str, usize> {
    map(parse_number, |i| i - 1)(i)
}

fn parse_instruction(i: &str) -> IResult<&str, Instruction> {
    map(
        tuple((
            preceded(tag("move "), parse_number),
            preceded(tag(" from "), parse_piles_number),
            preceded(tag(" to "), parse_piles_number),
        )),
        |(count, from, to)| Instruction { count, from, to },
    )(i)
}

fn parse_instruction_lines(i: &str) -> Vec<Instruction> {
    let mut instructions = Vec::new();
    i.lines().for_each(|line| {
        if let Ok((_rest, instruction)) = parse_instruction(line).finish() {
            instructions.push(instruction);
        }
    });
    instructions
}

struct Piles(Vec<Vec<Crate>>);

impl From<&str> for Piles {
    fn from(value: &str) -> Self {
        Piles(parse_crate_lines(value))
    }
}

impl Deref for Piles {
    type Target = Vec<Vec<Crate>>;

    fn deref(&self) -> &Self::Target {
        &self.0
    }
}

impl Debug for Piles {
    fn fmt(&self, f: &mut std::fmt::Formatter<'_>) -> std::fmt::Result {
        for (i, pile) in self.0.iter().enumerate() {
            writeln!(f, "Pile {i}: {:?}", pile)?;
        }
        Ok(())
    }
}

impl Piles {
    fn apply_instructions(&mut self, instructions: Vec<Instruction>) {
        for instr in &instructions {
            for _ in 0..instr.count {
                let crate_ = self.0[instr.from].pop().unwrap();
                self.0[instr.to].push(crate_);
            }
        }
    }

    fn tops(&self) -> String {
        self.0.iter().map(|pile| pile.last().unwrap()).join("")
    }
}

#[cfg(test)]
mod test {
    use super::*;

    const INPUT: &str = "
    [D]
[N] [C]
[Z] [M] [P]
 1   2   3

move 1 from 2 to 1
move 3 from 1 to 3
move 2 from 2 to 1
move 1 from 1 to 2";

    const INSTRUCTIONS: &str = "\
move 1 from 2 to 1
move 3 from 1 to 3
move 2 from 2 to 1
move 1 from 1 to 2";

    macro_rules! crates {
        ($($x:expr),*) => {
            vec![$(Crate($x)),*]
        };
    }

    #[test]
    fn parse_crates() {
        let crate_stacks = parse_crate_lines(INPUT);
        let expected = vec![crates!['Z', 'N'], crates!['M', 'C', 'D'], crates!['P']];
        assert_eq!(crate_stacks, expected);
    }

    #[test]
    fn parse_instructions() {
        let result = parse_instruction_lines(INSTRUCTIONS);

        assert_eq!(
            result,
            vec![
                Instruction {
                    count: 1,
                    from: 1,
                    to: 0
                },
                Instruction {
                    count: 3,
                    from: 0,
                    to: 2
                },
                Instruction {
                    count: 2,
                    from: 1,
                    to: 0
                },
                Instruction {
                    count: 1,
                    from: 0,
                    to: 1
                },
            ],
        );
    }
}
