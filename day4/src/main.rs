use std::ops::RangeInclusive;

use aoc22::day_inputs;

// #[derive(Debug)]
// struct Section {
//     start: usize,
//     end: usize,
// }
//
// impl From<&str> for Section {
//     fn from(value: &str) -> Self {
//         let mut split = value.split("-");
//         let (Some(start), Some(end)) = (split.next(), split.next()) else {
//             unreachable!("failed to split section")
//         };
//
//         Section {
//             start: start.parse().expect("number"),
//             end: end.parse().expect("number"),
//         }
//     }
// }

// #[derive(Debug)]
// struct Pair {
//     first: Section,
//     second: Section,
// }

// impl From<&str> for Pair {
//     fn from(value: &str) -> Self {
//         let mut split = value.split(",");
//         let (Some(first), Some(second)) = (split.next(), split.next()) else {
//             unreachable!("failed to split pairs")
//         };
//
//         Pair {
//             first: first.into(),
//             second: second.into(),
//         }
//     }
// }

// impl Pair {
//     fn is_self_contained(&self) -> bool {
//         self.first_encompasses_second() || self.second_encompasses_first()
//     }
//
//     fn first_encompasses_second(&self) -> bool {
//         self.first.start <= self.second.start && self.first.end >= self.second.end
//     }
//
//     fn second_encompasses_first(&self) -> bool {
//         self.first.start >= self.second.start && self.first.end <= self.second.end
//     }
//
//     fn has_overlaps(&self) -> bool {
//         !(self.first_before_second() || self.second_before_first())
//     }
//
//     fn first_before_second(&self) -> bool {
//         self.first.start < self.second.start && self.first.end < self.second.start
//     }
//
//     fn second_before_first(&self) -> bool {
//         self.first.start > self.second.start && self.first.start > self.second.end
//     }
// }

#[derive(Debug)]
struct Pair {
    first: RangeInclusive<usize>,
    second: RangeInclusive<usize>,
}

impl Pair {
    fn gen_section_range(section: &str) -> RangeInclusive<usize> {
        let mut split = section.split("-");
        let (Some(start), Some(end)) = (split.next(), split.next()) else {
            unreachable!("failed to split section")
        };
        RangeInclusive::new(start.parse().unwrap(), end.parse().unwrap())
    }

    fn is_self_contained(&self) -> bool {
        self.first.contains_range(&self.second) || self.second.contains_range(&self.first)
    }

    fn has_overlaps(&self) -> bool {
        self.first.contains(self.second.start()) || self.second.contains(self.first.start())
    }
}

impl From<&str> for Pair {
    fn from(value: &str) -> Self {
        let mut split = value.split(",");
        let (Some(first), Some(second)) = (split.next(), split.next()) else {
            unreachable!("failed to split pairs")
        };

        Self {
            first: Self::gen_section_range(first),
            second: Self::gen_section_range(second),
        }
    }
}

trait InclusiveRangeExt {
    fn contains_range(&self, other: &Self) -> bool;
}

impl<T> InclusiveRangeExt for RangeInclusive<T>
where
    T: PartialOrd,
{
    fn contains_range(&self, other: &Self) -> bool {
        self.contains(other.start()) && self.contains(other.end())
    }
}

fn main() {
    let input = day_inputs(4);
    println!("Part 1: {}", part1(&input));
    println!("Part 2: {}", part2(&input));
}

fn part1(input: &str) -> usize {
    input
        .lines()
        .map(|line| Pair::from(line))
        .filter(|pair| pair.is_self_contained())
        .count()
}

fn part2(input: &str) -> usize {
    input
        .lines()
        .map(|line| Pair::from(line))
        .filter(|pair| pair.has_overlaps())
        .count()
}

#[cfg(test)]
mod test {
    const SAMPLE: &str = "\
2-4,6-8
2-3,4-5
5-7,7-9
2-8,3-7
6-6,4-6
2-6,4-8";

    #[test]
    fn part1() {
        assert_eq!(crate::part1(SAMPLE), 2);
    }

    #[test]
    fn part2() {
        assert_eq!(crate::part2(SAMPLE), 4);
    }
}
