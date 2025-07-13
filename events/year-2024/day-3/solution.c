#include <ctype.h>
#include <stdbool.h>
#include <stddef.h>
#include <stdio.h>
#include <stdlib.h>
#include <string.h>

#define FILE_NAME "./input.txt"
#define MAX_INSTRUCTION_LENGTH 12

#define PART1_ANSWER 175015740
#define PART2_ANSWER 112272912

int year2024_day3_part1(char *file_name);
int year2024_day3_part2(char *file_name);

int multiply(char *instruction, int result);
bool is_numeric(char *token);
bool enable(char *instruction, bool enabled);
void validate_file(FILE *file);

int main(void) {
    printf("Year 2024 Day 3 Part 1: %d\n", year2024_day3_part1(FILE_NAME));
    printf("Year 2024 Day 3 Part 2: %d\n", year2024_day3_part2(FILE_NAME));

    return 0;
}

int year2024_day3_part1(char *file_name) {
    FILE *file = fopen(file_name, "r");
    validate_file(file);

    int buffer = 0, result = 0;
    char instruction[MAX_INSTRUCTION_LENGTH + 1];
    while ((buffer = fgetc(file)) != EOF) {
        if (buffer != 'm') continue;

        // The file just reads ahead a bit and goes back.
        // This is what these 4 complex looking lines do.

        fseek(file, -1, SEEK_CUR);
        long location = ftell(file);

        fgets(instruction, sizeof(instruction), file);
        fseek(file, location + 1, SEEK_SET);

        result = multiply(instruction, result);
    }

    fclose(file);
    return result;
}

int year2024_day3_part2(char *file_name) {
    FILE *file = fopen(file_name, "r");
    validate_file(file);

    bool enabled = true;
    int buffer = 0, result = 0;
    char instruction[MAX_INSTRUCTION_LENGTH + 1];
    while ((buffer = fgetc(file)) != EOF) {
        if (buffer != 'm' && buffer != 'd') continue;
        if (buffer == 'm' && !enabled) continue;

        fseek(file, -1, SEEK_CUR);
        long location = ftell(file);

        fgets(instruction, sizeof(instruction), file);
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

    perror("An error occured when opening the file");
    exit(1);
}
