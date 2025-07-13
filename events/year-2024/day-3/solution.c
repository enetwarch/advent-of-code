#include <stddef.h>
#include <stdio.h>
#include <stdlib.h>
#include <string.h>

#define FILE_NAME "./input.txt"

int year2024_day3_part1(char *file_name);
int year2024_day3_part2(char *file_name);

int main(void) {
    printf("Part 1: %d\n", year2024_day3_part1(FILE_NAME));
    printf("Part 2: %d\n", year2024_day3_part2(FILE_NAME));

    return 0;
}

int year2024_day3_part1(char *file_name) {
    FILE *file = fopen(file_name, "r");
    if (file == NULL) {
        perror("Error opening file");
        exit(1);
    }

    fseek(file, 0, SEEK_END);
    long file_characters = ftell(file);
    fseek(file, 0, SEEK_SET);

    char *buffer = malloc((sizeof(char) * file_characters) + 1);
    if (buffer == NULL) {
        perror("An error occured when executing malloc");
        exit(1);
    }

    fread(buffer, sizeof(char), file_characters, file);
    buffer[file_characters] = '\0';

    char closing, *p_buffer = buffer;
    int x = 0, y = 0, result = 0;
    while ((p_buffer = strstr(p_buffer, "mul(")) != NULL) {
        if (sscanf(p_buffer, "mul(%3d,%3d%1c", &x, &y, &closing) == 3) {
            if (closing == ')') result += (x * y);
        }

        p_buffer++;
    }

    free(buffer);
    fclose(file);
    return result;
}

int year2024_day3_part2(char *file_name) {
    FILE *file = fopen(file_name, "r");
    if (file == NULL) {
        perror("An error occured when opening the file");
        exit(1);
    }

    fclose(file);
    return -1;
}
