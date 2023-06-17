use aoc22::day_inputs;

fn main() {
    let input = day_inputs(1);
    println!("PART1: {}", part1(&input));
}

fn part1(input: &str) -> usize {
    input
        .trim()
        .split("\n\n")
        .map(|elf| {
            elf.split('\n')
                .map(|cal| cal.parse::<usize>().expect("calories should be numbers"))
                .sum::<usize>()
        })
        .max()
        .expect("input should have a size")
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
}
