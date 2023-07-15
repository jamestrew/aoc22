pub mod take1;
pub mod take2;

pub fn split_input(input: &str) -> (&str, &str) {
    let mut split = input.split("\n\n");
    (split.next().unwrap(), split.next().unwrap())
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

    #[test]
    fn test_split_input() {
        let (stack, instructions) = split_input(INPUT);
        assert_eq!(stack, CARGO);
        assert_eq!(instructions, INSTRUCTIONS);
    }
}
