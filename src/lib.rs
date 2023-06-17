use std::path::PathBuf;
use std::process::{Command, Stdio};

pub fn get_root_dir() -> PathBuf {
    let output = Command::new("git")
        .args(&["rev-parse", "--show-toplevel"])
        .stdout(Stdio::piped())
        .output()
        .expect("Failed to execute git command");

    let root_path = String::from_utf8(output.stdout).unwrap().trim().to_string();
    let path = PathBuf::from(root_path);
    path
}

pub fn day_inputs(day: usize) -> String {
    let root = get_root_dir();
    let path = root.join("inputs").join(day.to_string());
    std::fs::read_to_string(path)
        .expect(format!("failed to read input file from day {}", day).as_str())
}
