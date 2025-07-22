#ifndef METADATA_H
#define METADATA_H

// Metadata like file name, rows, columns, etc. are all stored here.
// File names will be written relative to the root of the repository.
// Macros, error messages and other useful constants are included.

#define Y2023_SOLUTIONS_DIRECTORY "./events/year-2023/solutions/"
#define Y2023_INPUT_DIRECTORY "./events/year-2023/input/"

// Problem: https://adventofcode.com/2024/day/1
#define Y2023_D01_P1_LABEL "Year 2023 Day 1 Part 1"
#define Y2023_D01_P2_LABEL "Year 2023 Day 1 Part 2"
#define Y2023_D01_INPUT_FILE_NAME Y2023_INPUT_DIRECTORY "day-01.txt"
#define Y2023_D01_MAX_LINE_LENGTH 60
#define Y2023_D01_P1_ANSWER 54877
#define Y2023_D01_P2_ANSWER 54100

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
