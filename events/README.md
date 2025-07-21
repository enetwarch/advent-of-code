# Events

```bash
advent-of-code/ # repository
|--events/ # this directory
|  |--year-*/ # problem year
|  |  |--input/ # problem input files
|  |  |  |--day-*.txt # day of the problem input
|  |  |--solutions/ # solution files
|  |  |  |--day-*.* # day of the solution file
|  |  |--target/ # binary files
|  |  |  |--day-* # day of the binary file
|  |  |--metadata.h # reusable constants
|  |--README.md # this documentation file
```

The [`events/`](./) directory is the most important directory that contains the solutions in this repository. Every year has a dedicated [`year-*/`](./year-2024/) directory that contains a predetermined structure as shown above. Inside [`input/`](./year-2024/input/), [`solutions/`](./year-2024/solutions/), and `target/` directories, each file will be named `day-*` followed by their file extension.

## Executables

In the `Makefile`, there is a command that allows you to compile every source code and run all of the executables. Every `solution/day-*.*` file has a `main()` function that prints out the results and a predefined `FILE_NAME` constant stored in `metadata.h` based from the root of the repository. By simply running `make run`, you can see all the answers in this directory.
