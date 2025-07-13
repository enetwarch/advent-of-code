# Events

```bash
advent-of-code/ # repository
|--events/ # this directory
|  |--year-*/ # problem year
|  |  |--day-*/ # problem day
|  |  |  |--input.txt # problem input file
|  |  |  |--solution.c # solution source code
|  |  |  |--README.md # problem instructions
|  |--README.md # this documentation file
```

The [`events/`](./) directory is the most important directory that contains the solutions in this repository. It has a very simple structure with `year-*/` and `day-*/` being the only subdirectories you need to keep in mind. The `year-*/` subdirectory refers to the year of the problem like the [`year-2024/`](./year-2024/) directory, and likewise, the `day-*/` subdirectory has [`day-1/`](./year-2024/day-1/) as an example. 

## File Naming

Three essential files will always be inside the `day-*/` subdirectory, which are: [`input.txt`](./year-2024/day-1/input.txt) as the input file, [`solution.c`](./year-2024/day-1/solution.c) as the solution, and [`README.md`](./year-2024/day-1/README.md) for the problem instructions. Other files may appear, but these essential files will always be there at the minimum. Solution files will always follow their respective language's naming convention and capitalization. However, `solution.c` will always be the main entry point of the program. 

## Executables

In the `Makefile`, there is a command that allows you to compile every source code and run all of the executables. Every `solution.c` file has a `main()` function that prints out the results and a predefined `FILE_NAME` constant based from the root of the repository. By simply running `make run`, you can see all the answers in this directory.
