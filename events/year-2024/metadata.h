#ifndef METADATA_H
#define METADATA_H

// Metadata like file name, rows, columns, etc. are all stored here.
// File names will be written relative to the root of the repository.
// Error messages and other useful constants are also included.

#define Y2024_SOLUTION_DIRECTORY "./events/year-2024/solutions/"
#define Y2024_INPUT_DIRECTORY "./events/year-2024/input/"

#define Y2024_D01_LABEL "Year 2024 Day 1"
#define Y2024_D01_P1_LABEL Y2024_D01_LABEL " Part 1"
#define Y2024_D01_P2_LABEL Y2024_D01_LABEL " Part 2"
#define Y2024_D01_SOLUTION_FILE_NAME Y2024_SOLUTION_DIRECTORY "day-01.c"
#define Y2024_D01_INPUT_FILE_NAME Y2024_INPUT_DIRECTORY "day-01.txt"
#define Y2024_D01_LINES 1000
#define Y2024_D01_LINE_LENGTH 13
#define Y2024_D01_P1_ANSWER 1222801
#define Y2024_D01_P2_ANSWER 22545250

#define Y2024_D02_LABEL "Year 2024 Day 2"
#define Y2024_D02_P1_LABEL Y2024_D02_LABEL " Part 1"
#define Y2024_D02_P2_LABEL Y2024_D02_LABEL " Part 2"
#define Y2024_D02_SOLUTION_FILE_NAME Y2024_SOLUTION_DIRECTORY "day-02.c"
#define Y2024_D02_INPUT_FILE_NAME Y2024_INPUT_DIRECTORY "day-02.txt"
#define Y2024_D02_LINES 1000
#define Y2024_D02_MAX_LINE_LENGTH 24
#define Y2024_D02_MAX_INT_PER_LINE 8
#define Y2024_D02_P1_ANSWER 252
#define Y2024_D02_P2_ANSWER 324

#define Y2024_D03_LABEL "Year 2024 Day 3"
#define Y2024_D03_P1_LABEL Y2024_D03_LABEL " Part 1"
#define Y2024_D03_P2_LABEL Y2024_D03_LABEL " Part 2"
#define Y2024_D03_SOLUTION_FILE_NAME Y2024_SOLUTION_DIRECTORY "day-03.c"
#define Y2024_D03_INPUT_FILE_NAME Y2024_INPUT_DIRECTORY "day-03.txt"
#define Y2024_D03_MAX_INSTRUCTION_LENGTH 12
#define Y2024_D03_P1_ANSWER 175015740
#define Y2024_D03_P2_ANSWER 112272912

#define Y2024_D04_LABEL "Year 2024 Day 4"
#define Y2024_D04_P1_LABEL Y2024_D04_LABEL " Part 1"
#define Y2024_D04_P2_LABEL Y2024_D04_LABEL " Part 2"
#define Y2024_D04_SOLUTION_FILE_NAME Y2024_SOLUTION_DIRECTORY "day-04.c"
#define Y2024_D04_INPUT_FILE_NAME Y2024_INPUT_DIRECTORY "day-04.txt"
#define Y2024_D04_ROWS 140
#define Y2024_D04_COLUMNS 140
#define Y2024_D04_P1_ANSWER 2549
#define Y2024_D04_P2_ANSWER 2003

// The following are constants for error messages.
// Each error message will refer to a specific standard library function.

#define FOPEN_ERROR_MESSAGE "An error occured when opening the file"
#define MALLOC_ERROR_MESSAGE "An error occured when executing malloc"
#define CALLOC_ERROR_MESSAGE "An error occured when executing calloc"

#endif
