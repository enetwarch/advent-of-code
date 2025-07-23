#ifndef METADATA_H
#define METADATA_H

// Metadata like file name, rows, columns, etc. are all stored here.
// File names will be written relative to the root of the repository.
// Macros, error messages and other useful constants are included.

#define Y2024_SOLUTIONS_DIRECTORY "./events/year-2024/solutions/"
#define Y2024_INPUT_DIRECTORY "./events/year-2024/input/"

// Problem: https://adventofcode.com/2024/day/1
#define Y2024_D01_P1_LABEL "Year 2024 Day 1 Part 1"
#define Y2024_D01_P2_LABEL "Year 2024 Day 1 Part 2"
#define Y2024_D01_INPUT_FILE_NAME Y2024_INPUT_DIRECTORY "day-01.txt"
#define Y2024_D01_LINES 1000
#define Y2024_D01_LINE_LENGTH 13
#define Y2024_D01_P1_ANSWER 1222801
#define Y2024_D01_P2_ANSWER 22545250

// Problem: https://adventofcode.com/2024/day/2
#define Y2024_D02_P1_LABEL "Year 2024 Day 2 Part 1"
#define Y2024_D02_P2_LABEL "Year 2024 Day 2 Part 2"
#define Y2024_D02_INPUT_FILE_NAME Y2024_INPUT_DIRECTORY "day-02.txt"
#define Y2024_D02_LINES 1000
#define Y2024_D02_MAX_LINE_LENGTH 24
#define Y2024_D02_MAX_INT_PER_LINE 8
#define Y2024_D02_P1_ANSWER 252
#define Y2024_D02_P2_ANSWER 324

// Problem: https://adventofcode.com/2024/day/3
#define Y2024_D03_P1_LABEL "Year 2024 Day 3 Part 1"
#define Y2024_D03_P2_LABEL "Year 2024 Day 3 Part 2"
#define Y2024_D03_INPUT_FILE_NAME Y2024_INPUT_DIRECTORY "day-03.txt"
#define Y2024_D03_MAX_INSTRUCTION_LENGTH 12
#define Y2024_D03_P1_ANSWER 175015740
#define Y2024_D03_P2_ANSWER 112272912

// Problem: https://adventofcode.com/2024/day/4
#define Y2024_D04_P1_LABEL "Year 2024 Day 4 Part 1"
#define Y2024_D04_P2_LABEL "Year 2024 Day 4 Part 2"
#define Y2024_D04_INPUT_FILE_NAME Y2024_INPUT_DIRECTORY "day-04.txt"
#define Y2024_D04_ROWS 140
#define Y2024_D04_COLUMNS 140
#define Y2024_D04_P1_ANSWER 2549
#define Y2024_D04_P2_ANSWER 2003

// Problem: https://adventofcode.com/2024/day/5
#define Y2024_D05_P1_LABEL "Year 2024 Day 5 Part 1"
#define Y2024_D05_P2_LABEL "Year 2024 Day 5 Part 2"
#define Y2024_D05_INPUT_FILE_NAME Y2024_INPUT_DIRECTORY "day-05.txt"
#define Y2024_D05_RULE_LINES 1176
#define Y2024_D05_MAX_INT_PER_LINE 23
#define Y2024_D05_P1_ANSWER 5713
#define Y2024_D05_P2_ANSWER 5180

// The following are constants for error messages.
// Each error message will refer to a specific standard library function.

#define FOPEN_ERROR_MESSAGE "An error occured when executing fopen"
#define MALLOC_ERROR_MESSAGE "An error occured when executing malloc"
#define CALLOC_ERROR_MESSAGE "An error occured when executing calloc"
#define REALLOC_ERROR_MESSAGE "An error occured when executing realloc"

// The following are helpful macros to help reduce boilerplate.
// Most of them will just involve printf or minor calculations.

#define PRINT_ANSWER_INT(label, answer) printf("%s: %d\n", label, answer)
#define PRINT_ANSWER_STRING(label, answer) printf("%s: %s\n", label, answer)

#endif
