#include <ctype.h>
#include <stdbool.h>
#include <stddef.h>
#include <stdio.h>
#include <stdlib.h>
#include <string.h>

#include "../metadata.h"

int y2023_d03_p1(char *file_name, int rows, int cols);
int y2023_d03_p2(char *file_name, int rows, int cols);

int get_number_starting_index(char *input, int index);
int get_number_ending_index(char *input, int index);
bool is_adjacent_to_a_sign(char *input, int index, int rows, int cols);
bool is_valid_step(int from_index, int to_index, int rows, int cols);
char *parse_file(char *file_name, int rows, int cols);

int main(void) {
    char *file_name = Y2023_D03_INPUT_FILE_NAME;
    int rows = Y2023_D03_ROWS;
    int cols = Y2023_D03_COLUMNS;

    int p1_answer = y2023_d03_p1(file_name, rows, cols);
    int p2_answer = y2023_d03_p2(file_name, rows, cols);

    PRINT_ANSWER_INT(Y2023_D03_P1_LABEL, p1_answer);
    PRINT_ANSWER_INT(Y2023_D03_P2_LABEL, p2_answer);

    return EXIT_SUCCESS;
}

int y2023_d03_p1(char *file_name, int rows, int cols) {
    char *input = parse_file(file_name, rows, cols);
    int answer = 0;

    for (int index = 0; index < rows * cols; index++) {
        if (!isdigit(input[index])) continue;
        if (!is_adjacent_to_a_sign(input, index, rows, cols)) continue;

        int starting_index = get_number_starting_index(input, index);
        index = get_number_ending_index(input, index);

        answer += atoi(input + starting_index);
    }

    free(input);
    return answer;
}

int y2023_d03_p2(char *file_name, int rows, int cols) {
    char *input = parse_file(file_name, rows, cols);
    int answer = 0;

    free(input);
    return answer;
}

// Small directional calculator functions.

int n(int index, int cols) { return index - cols; }
int ne(int index, int cols) { return index - cols + 1; }

// NOLINTNEXTLINE(bugprone-easily-swappable-parameters)
int e(int index, int cols) { return index + 1 + (cols - cols); };
int se(int index, int cols) { return index + cols + 1; }
int s(int index, int cols) { return index + cols; }
int sw(int index, int cols) { return index + cols - 1; }

// NOLINTNEXTLINE(bugprone-easily-swappable-parameters)
int w(int index, int cols) { return index - 1 - (cols - cols); }
int nw(int index, int cols) { return index - cols - 1; }

int get_number_starting_index(char *input, int index) {
    int previous_index = index - 1;
    while (isdigit(input[previous_index])) {
        previous_index--;
    }

    return previous_index + 1;
}

int get_number_ending_index(char *input, int index) {
    int next_index = index + 1;
    while (isdigit(input[next_index])) {
        next_index++;
    }

    return next_index - 1;
}

bool is_adjacent_to_a_sign(char *input, int index, int rows, int cols) {
    int (*directions[8])(int, int) = {n, ne, e, se, s, sw, w, nw};

    for (size_t i = 0; i < (sizeof(directions) / sizeof(directions[0])); i++) {
        int index_to_check = directions[i](index, cols);
        if (!is_valid_step(index, index_to_check, rows, cols)) continue;

        char cell = input[index_to_check];
        if (!isdigit(cell) && cell != '.') return true;
    }

    return false;
}

// Makes sure the next index is within the 2d array boundary.
bool is_valid_step(int from_index, int to_index, int rows, int cols) {
    if (from_index == to_index) return true;
    if (to_index < 0 || to_index >= (rows * cols)) return false;

    int from_row = (int)(from_index / cols);
    int to_row = (int)(to_index / cols);
    int from_col = (int)(from_index % cols);
    int to_col = (int)(to_index % cols);

    return (abs(from_row - to_row) <= 1) && (abs(from_col - to_col) <= 1);
}

char *parse_file(char *file_name, int rows, int cols) {
    FILE *file = fopen(file_name, "r");
    if (file == NULL) {
        perror(FOPEN_ERROR_MESSAGE);
        exit(EXIT_FAILURE);
    }

    char *input = malloc(sizeof(char) * cols * rows);
    if (input == NULL) {
        perror(MALLOC_ERROR_MESSAGE);
        exit(EXIT_FAILURE);
    }

    char line_buffer[cols + 2];
    for (int row = 0; row < rows; row++) {
        fgets(line_buffer, (int)sizeof(line_buffer), file);

        for (int col = 0; col < cols; col++) {
            input[(row * cols) + col] = line_buffer[col];
        }
    }

    fclose(file);
    return input;
}
