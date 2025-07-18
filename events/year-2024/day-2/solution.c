#include <stdbool.h>
#include <stdio.h>
#include <stdlib.h>
#include <string.h>

#define FILE_NAME "./events/year-2024/day-2/input.txt"
#define FILE_LINES 1000
#define MAX_LINE_LENGTH 24
#define MAX_INT_PER_LINE 8

#define PART1_ANSWER 252
#define PART2_ANSWER 324

int year2024_day2_part1(char *file_name);
int year2024_day2_part2(char *file_name);

bool is_safe(int *input, int length);
bool is_safe_with_dampener(int *input, int length);
void validate_file(FILE *file);

int main(void) {
    printf("Year 2024 Day 2 Part 1: %d\n", year2024_day2_part1(FILE_NAME));
    printf("Year 2024 Day 2 Part 2: %d\n", year2024_day2_part2(FILE_NAME));

    return 0;
}

int year2024_day2_part1(char *file_name) {
    FILE *file = fopen(file_name, "r");
    validate_file(file);

    int safe_reports = 0;
    char line[MAX_LINE_LENGTH + 2];
    while (fgets(line, sizeof(line), file) != NULL) {
        // Read through each line and parses each 2-digit token to int.
        // Solving directly is more efficient than mallocing to memory.

        int input[MAX_INT_PER_LINE], length = 0;
        char *token = strtok(line, " ");
        while (token != NULL) {
            input[length++] = atoi(token);
            token = strtok(NULL, " ");
        }

        if (is_safe(input, length)) safe_reports++;
    }

    fclose(file);
    return safe_reports;
}

int year2024_day2_part2(char *file_name) {
    FILE *file = fopen(file_name, "r");
    validate_file(file);

    int safe_reports = 0;
    char line[MAX_LINE_LENGTH + 2];
    while (fgets(line, sizeof(line), file) != NULL) {
        int input[MAX_INT_PER_LINE], length = 0;
        char *token = strtok(line, " ");
        while (token != NULL) {
            input[length++] = atoi(token);
            token = strtok(NULL, " ");
        }

        if (is_safe_with_dampener(input, length)) safe_reports++;
    }

    fclose(file);
    return safe_reports;
}

bool is_safe(int *input, int length) {
    if (length < 2) return false;

    bool is_ascending = false, is_descending = false;
    for (int i = 0; i < length - 1; i++) {
        int difference = input[i] - input[i + 1];

        if (!is_ascending && !is_descending) {
            is_ascending = difference < 0;
            is_descending = difference > 0;

            if (!is_ascending && !is_descending) return false;
        }

        if (is_ascending && (difference < -3 || difference >= 0)) return false;
        if (is_descending && (difference > 3 || difference <= 0)) return false;
    }

    return true;
}

// This is just a wrapper for the is_safe function.
// It tries to check all cases where 1 level is removed.
bool is_safe_with_dampener(int *input, int length) {
    if (length < 2) return false;
    if (is_safe(input, length)) return true;

    for (int i = 0; i < length; i++) {
        int dampened_input[length - 1];
        for (int j = 0, k = 0; j < length; j++) {
            if (i != j) dampened_input[k++] = input[j];
        }

        if (is_safe(dampened_input, length - 1)) return true;
    }

    return false;
}

void validate_file(FILE *file) {
    if (file != NULL) return;

    perror("An error occured when opening the file");
    exit(1);
}
