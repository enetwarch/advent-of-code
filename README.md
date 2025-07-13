# Advent of Code Solutions

This repository contains **my own personal solutions** to [Advent of Code](https://adventofcode.com/) (AOC) problems. I aim to improve my fundamentals in programming and computer understanding by completing each problem, which is why **my language choice is [`C`](https://www.w3schools.com/c/c_intro.php)**. A `README.md` file with the instructions and problem link will live alongside the solution for quick problem to solution view. **There will be little to no testing** in this repository as AOC already provides it in their website.

## Repository Structure

```bash
advent-of-code/ # repository
|--.github/ # github specific files and ci/cd pipelines
|--events/ # organized solutions directory
|--LICENSE # mit license
|--README.md # this documentation file
|--Makefile # automation tool
|--... # other files
```

The directory structure for this repository is very simple and short. All the solutions go to the [`events/`](./events/) directory which has its own `README.md` file for documentation. The [`Makefile`](./Makefile) contains all the useful build tools and utility scripts for formatting, and linting. To learn more, read the comments in the `Makefile` for more details. Each major directory will have their own documentation file as well to avoid cluttering this document.

## License

This project is licensed under the [MIT license](./LICENSE).
