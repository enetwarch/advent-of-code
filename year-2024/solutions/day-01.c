#include <stddef.h>
#include <stdio.h>
#include <stdlib.h>
#include <string.h>

#include "../metadata.h"

int y2024_d01_p1(char *file_name, int lines, int line_length);
int y2024_d01_p2(char *file_name, int lines, int line_length);

void bubble_sort_list(int *list, int length);
int *parse_file(char *file_name, int lines, int line_length, char position);

int main(void) {
    char *file_name = Y2024_D01_INPUT_FILE_NAME;
    int lines = Y2024_D01_LINES;
    int line_length = Y2024_D01_LINE_LENGTH;

    int p1_answer = y2024_d01_p1(file_name, lines, line_length);
    int p2_answer = y2024_d01_p2(file_name, lines, line_length);

    PRINT_ANSWER_INT(Y2024_D01_P1_LABEL, p1_answer);
    PRINT_ANSWER_INT(Y2024_D01_P2_LABEL, p2_answer);

    return EXIT_SUCCESS;
}

int y2024_d01_p1(char *file_name, int lines, int line_length) {
    int *left_list = parse_file(file_name, lines, line_length, 'l');
    int *right_list = parse_file(file_name, lines, line_length, 'r');

    bubble_sort_list(left_list, lines);
    bubble_sort_list(right_list, lines);

    int answer = 0;
    for (int i = 0; i < lines; i++) {
        answer += abs(left_list[i] - right_list[i]);
    }

    free(left_list);
    free(right_list);
    return answer;
}

int y2024_d01_p2(char *file_name, int lines, int line_length) {
    int *left_list = parse_file(file_name, lines, line_length, 'l');
    int *right_list = parse_file(file_name, lines, line_length, 'r');

    int answer = 0;
    for (int i = 0; i < lines; i++) {
        int left_number = left_list[i];

        int similarity_score = 0;
        for (int j = 0; j < lines; j++) {
            int right_number = right_list[j];
            if (left_number == right_number) similarity_score++;
        }

        answer += left_list[i] * similarity_score;
    }

    free(left_list);
    free(right_list);
    return answer;
}

void bubble_sort_list(int *list, int length) {
    for (int i = 0; i < length; i++) {
        for (int j = 0; j < length - 1; j++) {
            // NOLINTNEXTLINE(clang-analyzer-core.UndefinedBinaryOperatorResult)
            if (list[j] <= list[j + 1]) continue;

            int temporary_value = list[j];
            list[j] = list[j + 1];
            list[j + 1] = temporary_value;
        }
    }
}

// I know position is not type-safe and an enum will be better.
// I just want to avoid boilerplate from making enums at day 1.
// NOLINTNEXTLINE(bugprone-easily-swappable-parameters)
int *parse_file(char *file_name, int lines, int line_length, char position) {
    FILE *file = fopen(file_name, "r");
    if (file == NULL) {
        perror(FOPEN_ERROR_MESSAGE);
        exit(EXIT_FAILURE);
    }

    int *list = malloc(sizeof(int) * lines);
    if (list == NULL) {
        perror(MALLOC_ERROR_MESSAGE);
        exit(EXIT_FAILURE);
    }

    char line_buffer[line_length + 2];
    for (int i = 0; i < lines; i++) {
        if (fgets(line_buffer, (int)sizeof(line_buffer), file) == NULL) break;

        if (position != 'l' && position != 'r') break;
        char *token = strtok(line_buffer, "   ");
        list[i] = atoi(token);

        if (position != 'r') continue;
        token = strtok(NULL, "   ");
        list[i] = atoi(token);
    }

    fclose(file);
    return list;
}
