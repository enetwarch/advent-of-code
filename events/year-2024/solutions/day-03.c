#include <ctype.h>
#include <stdbool.h>
#include <stddef.h>
#include <stdio.h>
#include <stdlib.h>
#include <string.h>

#include "../metadata.h"

int y2024_d03_p1(char *file_name, int max_instruction_length);
int y2024_d03_p2(char *file_name, int max_instruction_length);

int multiply(char *instruction, int result);
bool is_numeric(char *token);
bool enable(char *instruction, bool enabled);
void validate_file(FILE *file);

int main(void) {
    char *file_name = Y2024_D03_INPUT_FILE_NAME;
    int max_instruction_length = Y2024_D03_MAX_INSTRUCTION_LENGTH;

    int p1_answer = y2024_d03_p1(file_name, max_instruction_length);
    int p2_answer = y2024_d03_p2(file_name, max_instruction_length);

    PRINT_ANSWER_INT(Y2024_D03_P1_LABEL, p1_answer);
    PRINT_ANSWER_INT(Y2024_D03_P2_LABEL, p2_answer);

    return EXIT_SUCCESS;
}

int y2024_d03_p1(char *file_name, int max_instruction_length) {
    FILE *file = fopen(file_name, "r");
    validate_file(file);

    int result = 0;
    int buffer = 0;
    char instruction[max_instruction_length + 1];
    while ((buffer = fgetc(file)) != EOF) {
        if (buffer != 'm') continue;

        // The file just reads ahead a bit and goes back.
        // This is what these 4 complex looking lines do.

        fseek(file, -1, SEEK_CUR);
        long location = ftell(file);

        fgets(instruction, (int)sizeof(instruction), file);
        fseek(file, location + 1, SEEK_SET);

        result = multiply(instruction, result);
    }

    fclose(file);
    return result;
}

int y2024_d03_p2(char *file_name, int max_instruction_length) {
    FILE *file = fopen(file_name, "r");
    validate_file(file);

    int result = 0;
    int buffer = 0;
    bool enabled = true;
    char instruction[max_instruction_length + 1];
    while ((buffer = fgetc(file)) != EOF) {
        if (buffer != 'm' && buffer != 'd') continue;
        if (buffer == 'm' && !enabled) continue;

        fseek(file, -1, SEEK_CUR);
        long location = ftell(file);

        fgets(instruction, (int)sizeof(instruction), file);
        fseek(file, location + 1, SEEK_SET);

        if (buffer == 'm')
            result = multiply(instruction, result);
        else if (buffer == 'd')
            enabled = enable(instruction, enabled);
    }

    fclose(file);
    return result;
}

int multiply(char *instruction, int result) {
    if (strncmp(instruction, "mul(", 4) != 0) return result;

    size_t instruction_length = strlen(instruction + 4);
    char buffer[instruction_length + 1];

    // sscanf() will lead to shorter code, but this makes it more manual.
    // NOLINTNEXTLINE: Unsafe functions are fine for AOC problems.
    memcpy(buffer, instruction + 4, instruction_length + 1);

    char *token = strtok(buffer, ",)");
    int x = is_numeric(token) ? atoi(token) : 0;

    token = strtok(NULL, ",)");
    int y = is_numeric(token) ? atoi(token) : 0;

    return result + (x * y);
}

bool is_numeric(char *token) {
    if (token == NULL || *token == '\0') return false;
    while (*token != '\0') {
        if (!isdigit(*token)) return false;
        token++;
    }

    return true;
}

bool enable(char *instruction, bool enabled) {
    if (strncmp(instruction, "do()", 4) == 0) return true;
    if (strncmp(instruction, "don't()", 7) == 0) return false;
    return enabled;
}

void validate_file(FILE *file) {
    if (file != NULL) return;

    perror(FOPEN_ERROR_MESSAGE);
    exit(EXIT_FAILURE);
}
