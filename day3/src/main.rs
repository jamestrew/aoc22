use std::{
    collections::HashSet,
    fmt::{Debug, Write},
};

use aoc22::day_inputs;
use itertools::Itertools;

#[repr(transparent)]
#[derive(PartialEq, Clone, Eq, Hash)]
struct Item(u8);

impl Debug for Item {
    fn fmt(&self, f: &mut std::fmt::Formatter<'_>) -> std::fmt::Result {
        f.write_char(self.0 as char)
    }
}

impl Item {
    pub fn priority(&self) -> usize {
        let ret = match self.0 {
            b'a'..=b'z' => self.0 - b'a' + 1,
            b'A'..=b'Z' => self.0 - b'A' + 27,
            _ => unreachable!("only allow A-Za-z"),
        };
        ret as usize
    }
}

impl From<u8> for Item {
    fn from(value: u8) -> Self {
        match value {
            b'a'..=b'z' | b'A'..=b'Z' => Self(value),
            _ => unreachable!("{value:?} is not a valid item"),
        }
    }
}

struct Rucksack(Vec<Item>);

impl Rucksack {
    fn common_item(&self) -> &Item {
        let (first, second) = self.0.split_at(self.0.len() / 2);
        second
            .iter()
            .find_map(|item2| first.iter().find(|&item1| item1 == item2))
            .unwrap()
    }
}

impl From<&str> for Rucksack {
    fn from(value: &str) -> Self {
        Self(value.bytes().map(|item| Item::from(item)).collect())
    }
}

fn main() {
    let input = day_inputs(3);
    println!("{}", part1(&input));
    println!("{}", part2(&input));
}

fn part1(input: &str) -> usize {
    input
        .lines()
        .map(Rucksack::from)
        .map(|rs| rs.common_item().priority())
        .sum()
}

fn part2(input: &str) -> usize {
    input
        .lines()
        .map(|line| line.bytes().map(Item::from).collect::<HashSet<_>>())
        .tuples()
        .map(|(a, b, c)| {
            a.iter()
                .find(|i| b.contains(i) && c.contains(i))
                .map(|i| i.priority())
                .unwrap_or_default()
        })
        .sum()
}

#[cfg(test)]
mod test {
    const SAMPLE: &str = "\
vJrwpWtwJgWrhcsFMMfFFhFp
jqHRNqRjqzjGDLGLrsFMfFZSrLrFZsSL
PmmdzqPrVvPwwTWBwg
wMqvLMZHhHMvwLHjbvcjnnSBnvTQFn
ttgJtRGJQctTZtZT
CrZsJsPPZsGzwwsLwLmpwMDw";

    #[test]
    fn part1() {
        assert_eq!(crate::part1(SAMPLE), 157);
    }

    #[test]
    fn part2() {
        assert_eq!(crate::part2(SAMPLE), 70);
    }
}
