#include <stdbool.h>
#include <stddef.h>
#include <stdio.h>
#include <stdlib.h>
#include <string.h>

#define FILE_NAME "./events/year-2024/day-4/input.txt"
#define ROWS 140
#define COLS 140

#define PART_ONE_ANSWER 2549
#define PART_TWO_ANSWER 0

int y2024_d04_p1(char *file_name, int rows, int cols);
int y2024_d04_p2(char *file_name, int rows, int cols);

int count_surrounding_xmas(char *input, int index, int rows, int cols);
bool is_valid_step(int from_index, int to_index, int rows, int cols);
char *parse_file(char *file_name, int rows, int cols);

int main(void) {
    char *file_name = FILE_NAME;
    int rows = ROWS;
    int cols = COLS;

    const int part_one = y2024_d04_p1(file_name, rows, cols);
    const int part_two = y2024_d04_p2(file_name, rows, cols);

    printf("Year 2024 Day 4 Part 1: %d\n", part_one);
    printf("Year 2024 Day 4 Part 2: %d\n", part_two);

    return EXIT_SUCCESS;
}

int y2024_d04_p1(char *file_name, int rows, int cols) {
    char *input = parse_file(file_name, rows, cols);

    int xmas_count = 0;
    for (int row = 0; row < rows; row++) {
        for (int col = 0; col < cols; col++) {
            int index = (row * cols) + col;
            char character = input[index];

            if (character != 'X') continue;
            xmas_count += count_surrounding_xmas(input, index, rows, cols);
        }
    }

    free(input);
    return xmas_count;
}

int y2024_d04_p2(char *file_name, int rows, int cols) {
    char *input = parse_file(file_name, rows, cols);

    free(input);
    return 0;
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

// NOLINTNEXTLINE(bugprone-easily-swappable-parameters)
int count_surrounding_xmas(char *input, int index, int rows, int cols) {
    int (*directions[8])(int, int) = {n, ne, e, se, s, sw, w, nw};

    char xmas_string[] = "XMAS";
    size_t xmas_length = strlen(xmas_string);

    int xmas_count = 0;
    for (size_t i = 0; i < (sizeof(directions) / sizeof(directions[0])); i++) {
        int (*direction)(int, int) = directions[i];
        int current_index = index;

        int matching_characters = 0;
        for (size_t j = 0; j < xmas_length; j++) {
            int next_index = j > 0 ? direction(current_index, cols) : index;
            if (!is_valid_step(current_index, next_index, rows, cols)) break;

            current_index = next_index;
            if (input[current_index] != xmas_string[j]) break;

            matching_characters++;
        }

        if (matching_characters != (int)xmas_length) continue;
        xmas_count++;
    }

    return xmas_count;
}

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
        perror("An error occured when opening the file");
        exit(EXIT_FAILURE);
    }

    char *input = malloc(sizeof(char) * rows * cols);
    if (input == NULL) {
        perror("An error occured when executing malloc");
        exit(EXIT_FAILURE);
    }

    char buffer[cols + 2];
    for (int row = 0; row < rows; row++) {
        fgets(buffer, sizeof(buffer), file);

        for (int col = 0; col < cols; col++) {
            input[(row * cols) + col] = buffer[col];
        }
    }

    fclose(file);
    return input;
}
