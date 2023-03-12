fn part1(input: String) -> usize {
    input
        .trim()
        .split("\n\n")
        .map(|elf| {
            elf.lines()
                .map(|cal| cal.parse::<usize>().unwrap())
                .sum::<usize>()
        })
        .max()
        .unwrap()
}

const INPUT: &str = r#"1000
2000
3000

4000

5000
6000

7000
8000
9000

10000"#;

#[cfg(test)]
mod tests {
    use super::*;

    #[test]
    fn test_part1() {
        let result = part1(INPUT.to_string());
        assert_eq!(result, 24000);
        // println!("{}", part1(INPUT.to_string()));
    }
}

