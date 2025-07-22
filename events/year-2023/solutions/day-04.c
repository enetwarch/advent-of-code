#include <ctype.h>
#include <stdbool.h>
#include <stddef.h>
#include <stdio.h>
#include <stdlib.h>
#include <string.h>

#include "../metadata.h"

int y2023_d04_p1(char *file_name, int winning_amount, int own_amount);
int y2023_d04_p2(char *file_name, int winning_amount, int own_amount);

int calculate_points(char *line_buffer, int winning_amount, int own_amount);
void parse_numbers(char *line_buffer, int *numbers, int amount, char *del);
void validate_file(FILE *file);

int main(void) {
    char *file_name = Y2023_D04_INPUT_FILE_NAME;
    int winning_amount = Y2023_D04_WINNING_AMOUNT;
    int own_amount = Y2023_D04_OWN_AMOUNT;

    int p1_answer = y2023_d04_p1(file_name, winning_amount, own_amount);
    int p2_answer = y2023_d04_p2(file_name, winning_amount, own_amount);

    PRINT_ANSWER_INT(Y2023_D04_P1_LABEL, p1_answer);
    PRINT_ANSWER_INT(Y2023_D04_P2_LABEL, p2_answer);

    return EXIT_SUCCESS;
}

int y2023_d04_p1(char *file_name, int winning_amount, int own_amount) {
    FILE *file = fopen(file_name, "r");
    validate_file(file);

    int points = 0;
    char line_buffer[(winning_amount * 3) + (own_amount * 3) + 20];
    while (fgets(line_buffer, (int)sizeof(line_buffer), file) != NULL) {
        points += calculate_points(line_buffer, winning_amount, own_amount);
    }

    fclose(file);
    return points;
}

// NOLINTNEXTLINE(bugprone-easily-swappable-parameters)
int y2023_d04_p2(char *file_name, int winning_amount, int own_amount) {
    FILE *file = fopen(file_name, "r");
    validate_file(file);

    (void)winning_amount;
    (void)own_amount;
    int points = 0;

    fclose(file);
    return points;
}

// NOLINTNEXTLINE(bugprone-easily-swappable-parameters)
int calculate_points(char *line_buffer, int winning_amount, int own_amount) {
    int winning_numbers[winning_amount];
    int own_numbers[own_amount];

    parse_numbers(line_buffer, winning_numbers, winning_amount, ":|");
    parse_numbers(line_buffer, own_numbers, own_amount, "|");

    int points = 0;
    for (int i = 0; i < winning_amount; i++) {
        for (int j = 0; j < own_amount; j++) {
            if (winning_numbers[i] != own_numbers[j]) continue;

            points = points == 0 ? 1 : points * 2;
            break;
        }
    }

    return points;
}

void parse_numbers(char *line_buffer, int *numbers, int amount, char *del) {
    char buffer_copy[strlen(line_buffer) + 2];
    // NOLINTNEXTLINE(clang-analyzer-security.insecureAPI.strcpy)
    strcpy(buffer_copy, line_buffer);

    // This strtok is here to skip to the second half of the line buffer.
    // ":|" delimeter will parse winning numbers, "|" will parse own numbers.
    strtok(buffer_copy, del);
    char *number_characters = strtok(NULL, del);

    char *number_char = strtok(number_characters, " ");
    for (int i = 0; i < amount; i++) {
        numbers[i] = atoi(number_char);
        number_char = strtok(NULL, " ");
    }
}

void validate_file(FILE *file) {
    if (file != NULL) return;

    perror(FOPEN_ERROR_MESSAGE);
    exit(EXIT_FAILURE);
}
