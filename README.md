# Advent of Code Solutions

This repository contains **my own personal solutions** to [Advent of Code](https://adventofcode.com/) (AOC) problems. I aim to improve my fundamentals in programming and computer understanding by completing each problem, which is why **my language choice is [`C`](https://www.w3schools.com/c/c_intro.php)**. **There will be little to no testing** in this repository as AOC already provides it in their website.

## Repository Structure

```bash
advent-of-code/ # repository
|--year-*/ # problem year
|  |--input/ # problem input files
|  |  |--day-*.txt # day of the problem input
|  |--solutions/ # solution files
|  |  |--day-*.* # day of the solution file
|  |--target/ # binary files
|  |  |--day-*.* # day of the binary file
|  |--metadata.h # reusable constants
|  |--Makefile # workspace-specific automation tool
|  |--... # other files...
|--Makefile # universal automation tool
|--README.md # this documentation file
|--... # other files...
```

The directory structure for this repository is very simple and short. Every year has a dedicated `year-*/` directory that contains a predetermined structure as shown above. Inside `input/`, `solutions/`, and `target/` directories, each file will be named `day-*` followed by their file extension. The `Makefile` contains all the useful build tools and utility scripts for formatting, and linting. To learn more, read the comments in the `Makefile` for more details.

## License

This project is licensed under the [MIT license](./LICENSE).
