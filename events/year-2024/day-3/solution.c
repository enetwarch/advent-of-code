#include <stdbool.h>
#include <stddef.h>
#include <stdio.h>
#include <stdlib.h>
#include <string.h>

#define FILE_NAME "./input.txt"
#define MAX_INSTRUCTION_LENGTH 12

int year2024_day3_part1(char *file_name);
int year2024_day3_part2(char *file_name);

int multiply(char *instruction, int result);
bool enable(char *instruction, bool enabled);
void validate_file(FILE *file);

int main(void) {
    printf("Part 1: %d\n", year2024_day3_part1(FILE_NAME));
    printf("Part 2: %d\n", year2024_day3_part2(FILE_NAME));

    return 0;
}

int year2024_day3_part1(char *file_name) {
    FILE *file = fopen(file_name, "r");
    validate_file(file);

    int buffer = 0, result = 0;
    char instruction[MAX_INSTRUCTION_LENGTH + 1];
    while ((buffer = fgetc(file)) != EOF) {
        if (buffer != 'm') continue;

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
    char closing;
    int x = 0, y = 0;
    if (sscanf(instruction, "mul(%3d,%3d%1c", &x, &y, &closing) == 3) {
        if (closing == ')') return result + (x * y);
    }

    return result;
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
