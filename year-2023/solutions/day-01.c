#include <ctype.h>
#include <stddef.h>
#include <stdio.h>
#include <stdlib.h>
#include <string.h>

#include "../metadata.h"

int y2023_d01_p1(char *file_name, int max_line_length);
int y2023_d01_p2(char *file_name, int max_line_length);

void get_edge_digits(char *digits, char *line_buffer);
void digitify_digit_strings(char *line_buffer);
void validate_file(FILE *file);

int main(void) {
    char *file_name = Y2023_D01_INPUT_FILE_NAME;
    int max_line_length = Y2023_D01_MAX_LINE_LENGTH;

    int p1_answer = y2023_d01_p1(file_name, max_line_length);
    int p2_answer = y2023_d01_p2(file_name, max_line_length);

    PRINT_ANSWER_INT(Y2023_D01_P1_LABEL, p1_answer);
    PRINT_ANSWER_INT(Y2023_D01_P2_LABEL, p2_answer);

    return EXIT_SUCCESS;
}

int y2023_d01_p1(char *file_name, int max_line_length) {
    FILE *file = fopen(file_name, "r");
    validate_file(file);

    int answer = 0;
    char line_buffer[max_line_length + 2];
    while (fgets(line_buffer, (int)sizeof(line_buffer), file) != NULL) {
        char digits[3] = {-1, -1, '\0'};

        get_edge_digits(digits, line_buffer);
        answer += atoi(digits);
    }

    fclose(file);
    return answer;
}

int y2023_d01_p2(char *file_name, int max_line_length) {
    FILE *file = fopen(file_name, "r");
    validate_file(file);

    int answer = 0;
    char line_buffer[max_line_length + 2];
    while (fgets(line_buffer, (int)sizeof(line_buffer), file) != NULL) {
        char digits[3] = {-1, -1, '\0'};

        digitify_digit_strings(line_buffer);
        get_edge_digits(digits, line_buffer);
        answer += atoi(digits);
    }

    fclose(file);
    return answer;
}

void get_edge_digits(char *digits, char *line_buffer) {
    for (size_t i = 0; i < strlen(line_buffer); i++) {
        if (!isdigit(line_buffer[i])) continue;
        digits[digits[0] == -1 ? 0 : 1] = line_buffer[i];
    }

    // NOLINTNEXTLINE(bugprone-narrowing-conversions)
    digits[1] = digits[1] == -1 ? digits[0] : digits[1];
}

void digitify_digit_strings(char *line_buffer) {
    char *strings[10] = {"zero", "one", "two",   "three", "four",
                         "five", "six", "seven", "eight", "nine"};

    for (int i = 0; i < (int)(sizeof(strings) / sizeof(strings[0])); i++) {
        char *matching_string = strstr(line_buffer, strings[i]);
        if (matching_string == NULL) continue;

        int ascii_offset = 48;

        // NOLINTNEXTLINE(bugprone-narrowing-conversions)
        matching_string[(int)(strlen(strings[i]) / 2)] = i + ascii_offset;
        digitify_digit_strings(line_buffer);
    }
}

void validate_file(FILE *file) {
    if (file != NULL) return;

    perror(FOPEN_ERROR_MESSAGE);
    exit(EXIT_FAILURE);
}
