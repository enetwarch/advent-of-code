#include <stdio.h>
#include <stdlib.h>
#include <string.h>
#include <stdbool.h>

#define FILE_NAME "input.txt"
#define MAX_LINE_LENGTH 24
#define MAX_INT_PER_LINE 8

int part_one(char *file_name);
bool is_unsafe_part_one(int *input, int length);
int part_two(char *file_name);

int main(int argc, char **argv) {
    printf("Part 1: %d\n", part_one(FILE_NAME));
    printf("Part 2: %d\n", part_two(FILE_NAME));

    return 0;
}

int part_one(char *file_name) {
    FILE *file = fopen(file_name, "r");
    if (file == NULL) {
        perror("Error in opening the file");
        exit(1);
    }

    int safe_reports = 0;
    char line[MAX_LINE_LENGTH + 2];
    while (fgets(line, sizeof(line), file) != NULL) {
        int input[MAX_INT_PER_LINE], length = 0;
        char *token = strtok(line, " ");
        while (token != NULL) {
            input[length++] = atoi(token);
            token = strtok(NULL, " ");
        }

        if (is_unsafe_part_one(input, length)) continue;
        safe_reports++;
    }

    fclose(file);
    return safe_reports;
}

bool is_unsafe_part_one(int *input, int length) {
    if (length < 2) return true;

    bool is_ascending = false, is_descending = false;
    for (int i = 0; i < length - 1; i++) {
        int diff = input[i] - input[i + 1];

        if (!is_ascending && !is_descending) {
            if (diff < 0) is_ascending = true;
            else if (diff > 0) is_descending = true;
            else return true;
        }

        if (is_ascending && (diff < -3 || diff >= 0)) return true;
        else if (is_descending && (diff > 3 || diff <= 0)) return true;
    }

    return false;
}

int part_two(char *file_name) {
    return -1;
}
