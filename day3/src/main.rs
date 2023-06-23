use aoc22::day_inputs;

#[derive(Debug, PartialEq, Eq)]
struct Item(u8);

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

pub(crate) struct Rucksack(Vec<Item>);

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
}

fn part1(input: &str) -> usize {
    input
        .lines()
        .map(|rs| Rucksack::from(rs))
        .map(|rs| rs.common_item().priority())
        .sum()
}

fn part2(input: &str) -> usize {
    todo!()
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
