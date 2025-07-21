#include <stdbool.h>
#include <stddef.h>
#include <stdio.h>
#include <stdlib.h>
#include <string.h>

#include "../metadata.h"

int y2024_d05_p1(char *file_name, int rule_lines, int max_int_per_line);
int y2024_d05_p2(char *file_name, int rule_lines, int max_int_per_line);

bool is_in_right_order(int *l, int *r, int lr_size, int *input, int i_size);
void parse_file(FILE *file, int *left_list, int *right_list, int lines);

int main(void) {
    char *file_name = Y2024_D05_INPUT_FILE_NAME;
    int rule_lines = Y2024_D05_RULE_LINES;
    int max_int_per_line = Y2024_D05_MAX_INT_PER_LINE;

    int p1_answer = y2024_d05_p1(file_name, rule_lines, max_int_per_line);
    int p2_answer = y2024_d05_p2(file_name, rule_lines, max_int_per_line);

    PRINT_ANSWER_INT(Y2024_D05_P1_LABEL, p1_answer);
    PRINT_ANSWER_INT(Y2024_D05_P2_LABEL, p2_answer);

    return EXIT_SUCCESS;
}

// NOLINTNEXTLINE(bugprone-easily-swappable-parameters)
int y2024_d05_p1(char *file_name, int rule_lines, int max_int_per_line) {
    FILE *file = fopen(file_name, "r");
    int *left = malloc(sizeof(int) * rule_lines);
    int *right = malloc(sizeof(int) * rule_lines);

    parse_file(file, left, right, rule_lines);
    if (fgetc(file) != '\n') return -1;

    int answer = 0;
    int max_line_length = (max_int_per_line * 3) - 1;
    char line_buffer[max_line_length + 2];
    while (fgets(line_buffer, (int)sizeof(line_buffer), file) != NULL) {
        int input[max_int_per_line];
        int i_size = 0;

        char *token = strtok(line_buffer, ",");
        while (token != NULL) {
            input[i_size++] = atoi(token);
            token = strtok(NULL, ",");
        }

        int lr_size = rule_lines;
        if (!is_in_right_order(left, right, lr_size, input, i_size)) continue;

        // NOLINTNEXTLINE(clang-analyzer-core.uninitialized.Assign)
        answer += input[(int)(i_size / 2)];
    }

    fclose(file);
    free(left);
    free(right);
    return answer;
}

// NOLINTNEXTLINE(bugprone-easily-swappable-parameters)
int y2024_d05_p2(char *file_name, int rule_lines, int max_int_per_line) {
    FILE *file = fopen(file_name, "r");
    int *left = malloc(sizeof(int) * rule_lines);
    int *right = malloc(sizeof(int) * rule_lines);
    parse_file(file, left, right, rule_lines);

    fclose(file);
    free(left);
    free(right);
    return 0;
}

// The arguments are hard to reason with and may seem ambiguous.
// I can make a struct just for this use case, but I decided not to.
// Just shortening the argument names is enough.
// In exchange, variables have descriptive names when accessing the pointers.
// NOLINTNEXTLINE(bugprone-easily-swappable-parameters)
bool is_in_right_order(int *l, int *r, int lr_size, int *input, int i_size) {
    for (int i = 0; i < i_size - 1; i++) {
        for (int j = 0; j < lr_size; j++) {
            int right_list_element = r[j];
            int current_order_element = input[i];

            if (right_list_element != current_order_element) continue;

            for (int k = i + 1; k < i_size; k++) {
                int left_list_element = l[j];
                int upcoming_order_element = input[k];

                if (upcoming_order_element == left_list_element) return false;
            }
        }
    }

    return true;
}

void parse_file(FILE *file, int *left_list, int *right_list, int lines) {
    if (file == NULL) {
        perror(FOPEN_ERROR_MESSAGE);
        exit(EXIT_FAILURE);
    }

    if (left_list == NULL || right_list == NULL) {
        perror(MALLOC_ERROR_MESSAGE);
        exit(EXIT_FAILURE);
    }

    for (int i = 0; i < lines; i++) {
        // NOLINTNEXTLINE(clang-analyzer-security.insecureAPI.DeprecatedOrUnsafeBufferHandling)
        fscanf(file, "%d|%d", &left_list[i], &right_list[i]);
    }
}
