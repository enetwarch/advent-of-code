#include <ctype.h>
#include <stddef.h>
#include <stdio.h>
#include <stdlib.h>
#include <string.h>

#include "../metadata.h"

int y2023_d01_p1(char *file_name, int max_line_length);
int y2023_d01_p2(char *file_name, int max_line_length);

void validate_file(FILE *file);

void bubble_sort_list(int *list, int length);
int *parse_file(char *file_name, int lines, int line_length, char position);

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

        for (int i = 0; line_buffer[i] != '\0'; i++) {
            if (!isdigit(line_buffer[i])) continue;
            digits[digits[0] == -1 ? 0 : 1] = line_buffer[i];
        }

        // NOLINTNEXTLINE(bugprone-narrowing-conversions)
        digits[1] = digits[1] == -1 ? digits[0] : digits[1];
        answer += atoi(digits);
    }

    fclose(file);
    return answer;
}

int y2023_d01_p2(char *file_name, int max_line_length) {
    FILE *file = fopen(file_name, "r");
    validate_file(file);

    (void)(max_line_length);

    fclose(file);
    return 0;
}

void validate_file(FILE *file) {
    if (file != NULL) return;

    perror(FOPEN_ERROR_MESSAGE);
    exit(EXIT_FAILURE);
}
