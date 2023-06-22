use anyhow::anyhow;
use aoc22::day_inputs;

#[derive(Debug, Clone)]
enum Move {
    Rock,
    Paper,
    Scissor,
}

impl Move {
    const MOVES: [Self; 3] = [Self::Rock, Self::Paper, Self::Scissor];

    fn points_for_move(&self) -> usize {
        match self {
            Self::Rock => 1,
            Self::Paper => 2,
            Self::Scissor => 3,
        }
    }

    fn beats(&self, opponent: &Self) -> bool {
        matches!(
            (self, opponent),
            (Self::Paper, Self::Rock) | (Self::Rock, Self::Scissor) | (Self::Scissor, Self::Paper)
        )
    }

    fn outcome(&self, opponent: &Self) -> Outcome {
        if self.beats(opponent) {
            Outcome::Win
        } else if opponent.beats(self) {
            Outcome::Lose
        } else {
            Outcome::Draw
        }
    }

    fn winning_move(&self) -> &Self {
        Self::MOVES
            .iter()
            .by_ref()
            .find(|&m| m.beats(self))
            .unwrap()
    }

    fn losing_move(&self) -> &Self {
        Self::MOVES
            .iter()
            .by_ref()
            .find(|&m| self.beats(m))
            .unwrap()
    }

    fn outcome_counter(&self, outcome: &Outcome) -> Self {
        match outcome {
            Outcome::Lose => self.losing_move().clone(),
            Outcome::Draw => self.clone(),
            Outcome::Win => self.winning_move().clone(),
        }
    }
}

impl TryFrom<char> for Move {
    type Error = anyhow::Error;

    fn try_from(value: char) -> Result<Self, Self::Error> {
        match value {
            'A' | 'X' => Ok(Self::Rock),
            'B' | 'Y' => Ok(Self::Paper),
            'C' | 'Z' => Ok(Self::Scissor),
            _ => Err(anyhow!("invalid move: {value:?}")),
        }
    }
}

#[derive(Debug)]
enum Outcome {
    Lose,
    Draw,
    Win,
}

impl Outcome {
    fn points(&self) -> usize {
        match self {
            Outcome::Win => 6,
            Outcome::Lose => 0,
            Outcome::Draw => 3,
        }
    }
}

impl TryFrom<char> for Outcome {
    type Error = anyhow::Error;

    fn try_from(value: char) -> Result<Self, Self::Error> {
        match value {
            'X' => Ok(Self::Lose),
            'Y' => Ok(Self::Draw),
            'Z' => Ok(Self::Win),
            _ => Err(anyhow!("invalid move: {value:?}")),
        }
    }
}

#[derive(Debug)]
struct Round {
    player1: Move,
    player2: Move,
}

impl Round {
    fn p2_points(&self) -> usize {
        self.player2.points_for_move() + self.player2.outcome(&self.player1).points()
    }

    fn from_part1(s: &str) -> anyhow::Result<Self> {
        let mut chars = s.trim().chars();
        let (Some(player1), _, Some(player2)) = (chars.next(), chars.next(), chars.next()) else {
            return Err(anyhow!("bad game {s:?}"));
        };

        Ok(Self {
            player1: player1.try_into()?,
            player2: player2.try_into()?,
        })
    }

    fn from_part2(s: &str) -> anyhow::Result<Self> {
        let mut chars = s.trim().chars();
        let (Some(player1), _, Some(outcome)) = (chars.next(), chars.next(), chars.next()) else {
            return Err(anyhow!("bad game {s:?}"));
        };

        let player1 = Move::try_from(player1)?;
        let outcome = Outcome::try_from(outcome)?;
        let player2 = player1.outcome_counter(&outcome);

        Ok(Self { player1, player2 })
    }
}

fn main() {
    let input = day_inputs(2);
    println!("{}", part1(&input));
    println!("{}", part2(&input));
}

fn part1(input: &str) -> usize {
    input
        .lines()
        .map(|line| Round::from_part1(line).unwrap().p2_points())
        .sum()
}

fn part2(input: &str) -> usize {
    input
        .lines()
        .map(|line| Round::from_part2(line).unwrap().p2_points())
        .sum()
}

#[cfg(test)]
mod test {
    const SAMPLE: &str = "\
A Y
B X
C Z";

    #[test]
    fn part1() {
        assert_eq!(crate::part1(SAMPLE), 15)
    }

    #[test]
    fn part2() {
        assert_eq!(crate::part2(SAMPLE), 12)
    }
}
