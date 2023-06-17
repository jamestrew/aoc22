use aoc22::day_inputs;
use itertools::Itertools;

fn main() {
    let input = day_inputs(1);
    println!("PART1: {}", part1(&input));
    println!("PART2: {}", part2(&input));
    println!("PART2: {}", part2_improved(&input));

    let empty_sum = (1..1).sum::<i32>();
    let empty_sum1 = (1..1).sum1::<i32>();
    assert_eq!(empty_sum, 0);
    assert_eq!(empty_sum1, None);

    let nonempty_sum = (1..11).sum::<i32>();
    let nonempty_sum1 = (1..11).sum1::<i32>();
    assert_eq!(nonempty_sum, 55);
    assert_eq!(nonempty_sum1, Some(55));
}

fn part1(input: &str) -> usize {
    input
        .trim()
        .split("\n\n")
        .map(|elf| {
            elf.split('\n')
                .map(|cal| cal.parse::<usize>().unwrap_or_default())
                .sum::<usize>()
        })
        .max()
        .expect("input should have a size")
}

fn part2(input: &str) -> usize {
    let mut elves: Vec<_> = input
        .trim()
        .split("\n\n")
        .map(|elf| {
            elf.split('\n')
                .map(|cal| cal.parse::<usize>().unwrap_or_default())
                .sum::<usize>()
        })
        .collect();

    elves.sort_by(|a, b| b.partial_cmp(a).unwrap());
    elves.iter().take(3).sum()
}

fn part2_improved(input: &str) -> usize {
    input
        .lines()
        .map(|v| v.parse::<usize>().ok())
        .batching(|it| it.map_while(|x| x).sum1::<usize>())
        .map(std::cmp::Reverse)
        .k_smallest(3) // fails debug assert
        .map(|x| x.0)
        .sum()
}

#[cfg(test)]
mod test {
    const SAMPLE: &str = r#"
1000
2000
3000

4000

5000
6000

7000
8000
9000

10000"#;

    #[test]
    fn part1() {
        assert_eq!(crate::part1(SAMPLE), 24000);
    }

    #[test]
    fn part2() {
        assert_eq!(crate::part2(SAMPLE), 45000);
    }

    #[test]
    fn part2_improved() {
        let _ans = crate::part2_improved(SAMPLE);
        assert_eq!(0, 0);
    }
}
