#include <stdbool.h>
#include <stdio.h>
#include <stdlib.h>
#include <string.h>

#include "../metadata.h"

int y2024_d02_p1(char *file_name, int max_line_length, int max_int_per_line);
int y2024_d02_p2(char *file_name, int max_line_length, int max_int_per_line);

bool is_safe(int *input, int length);
bool is_safe_with_dampener(int *input, int length);
void validate_file(FILE *file);

int main(void) {
    char *file_name = Y2024_D02_INPUT_FILE_NAME;
    int max_line_length = Y2024_D02_MAX_LINE_LENGTH;
    int max_int_per_line = Y2024_D02_MAX_INT_PER_LINE;

    int p1_answer = y2024_d02_p1(file_name, max_line_length, max_int_per_line);
    int p2_answer = y2024_d02_p2(file_name, max_line_length, max_int_per_line);

    PRINT_ANSWER_INT(Y2024_D02_P1_LABEL, p1_answer);
    PRINT_ANSWER_INT(Y2024_D02_P2_LABEL, p2_answer);

    return EXIT_SUCCESS;
}

// NOLINTNEXTLINE(bugprone-easily-swappable-parameters)
int y2024_d02_p1(char *file_name, int max_line_length, int max_int_per_line) {
    FILE *file = fopen(file_name, "r");
    validate_file(file);

    int safe_reports = 0;
    char line_buffer[max_line_length + 2];
    while (fgets(line_buffer, (int)sizeof(line_buffer), file) != NULL) {
        // Read through each line and parses each 2-digit token to int.
        // Solving directly is more efficient than mallocing to memory.

        int input_index = 0;
        int input[max_int_per_line];

        char *token = strtok(line_buffer, " ");
        while (token != NULL) {
            input[input_index++] = atoi(token);
            token = strtok(NULL, " ");
        }

        if (is_safe(input, input_index)) safe_reports++;
    }

    fclose(file);
    return safe_reports;
}

// NOLINTNEXTLINE(bugprone-easily-swappable-parameters)
int y2024_d02_p2(char *file_name, int max_line_length, int max_int_per_line) {
    FILE *file = fopen(file_name, "r");
    validate_file(file);

    int safe_reports = 0;
    char line_buffer[max_line_length + 2];
    while (fgets(line_buffer, (int)sizeof(line_buffer), file) != NULL) {
        int input_index = 0;
        int input[max_int_per_line];

        char *token = strtok(line_buffer, " ");
        while (token != NULL) {
            input[input_index++] = atoi(token);
            token = strtok(NULL, " ");
        }

        if (is_safe_with_dampener(input, input_index)) safe_reports++;
    }

    fclose(file);
    return safe_reports;
}

bool is_safe(int *input, int length) {
    if (length < 2) return false;

    bool is_ascending = false;
    bool is_descending = false;
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

    perror(FOPEN_ERROR_MESSAGE);
    exit(EXIT_FAILURE);
}
