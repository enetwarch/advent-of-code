#include <stddef.h>
#include <stdio.h>
#include <stdlib.h>
#include <string.h>

#define FILE_NAME "./events/year-2024/day-1/input.txt"
#define FILE_LINES 1000
#define MAX_LINE_LENGTH 13

#define PART1_ANSWER 1222801
#define PART2_ANSWER 22545250

int year2024_day1_part1(char *file_name);
int year2024_day1_part2(char *file_name);

void parse_file(char *file_name, int *left_list, int *right_list);
void sort_list(int *list, int length);

int main(void) {
    printf("Year 2024 Day 1 Part 1: %d\n", year2024_day1_part1(FILE_NAME));
    printf("Year 2024 Day 1 Part 2: %d\n", year2024_day1_part2(FILE_NAME));

    return 0;
}

int year2024_day1_part1(char *file_name) {
    int *left_list = calloc(FILE_LINES, sizeof(int));
    int *right_list = calloc(FILE_LINES, sizeof(int));
    parse_file(file_name, left_list, right_list);

    sort_list(left_list, FILE_LINES);
    sort_list(right_list, FILE_LINES);

    int answer = 0;
    for (int i = 0; i < FILE_LINES; i++) {
        answer += abs(left_list[i] - right_list[i]);
    }

    free(left_list);
    free(right_list);
    return answer;
}

int year2024_day1_part2(char *file_name) {
    int *left_list = calloc(FILE_LINES, sizeof(int));
    int *right_list = calloc(FILE_LINES, sizeof(int));
    parse_file(file_name, left_list, right_list);

    int answer = 0;
    for (int i = 0; i < FILE_LINES; i++) {
        int left_number = left_list[i];

        int similarity_score = 0;
        for (int j = 0; j < FILE_LINES; j++) {
            int right_number = right_list[j];
            if (left_number == right_number) similarity_score++;
        }

        answer += left_list[i] * similarity_score;
    }

    free(left_list);
    free(right_list);
    return answer;
}

void parse_file(char *file_name, int *left_list, int *right_list) {
    if (left_list == NULL || right_list == NULL) {
        perror("An error occured when executing calloc");
        exit(1);
    }

    FILE *file = fopen(file_name, "r");
    if (file == NULL) {
        perror("An error occured when opening the file");
        exit(1);
    }

    char line[MAX_LINE_LENGTH + 2];
    for (int i = 0; i < FILE_LINES; i++) {
        if (fgets(line, sizeof(line), file) == NULL) break;

        char *token = strtok(line, "   ");
        left_list[i] = atoi(token);

        token = strtok(NULL, "   ");
        right_list[i] = atoi(token);
    }

    fclose(file);
}

void sort_list(int *list, int length) {
    for (int i = 0; i < length; i++) {
        for (int j = 0; j < length - 1; j++) {
            if (list[j] <= list[j + 1]) continue;

            int temporary_value = list[j];
            list[j] = list[j + 1];
            list[j + 1] = temporary_value;
        }
    }
}
