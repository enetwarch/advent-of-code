#include <stdio.h>
#include <stdlib.h>
#include <stddef.h>

#define FILE_NAME "./input.txt"
#define FILE_LINES 1000
#define MAX_LINE_LENGTH 13

int year2024_day1_part1(char *file_name);
int year2024_day1_part2(char *file_name);

void parse_file(char *file_name, int *left_list, int *right_list);
void sort_list(int *list, int length);

int main(void) {
    printf("Part 1: %d\n", year2024_day1_part1(FILE_NAME));
    printf("Part 2: %d\n", year2024_day1_part2(FILE_NAME));

    return 0;
}

int year2024_day1_part1(char *file_name) {
    int *left_list = malloc(sizeof(int) * FILE_LINES);
    int *right_list = malloc(sizeof(int) * FILE_LINES);
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
    int *left_list = malloc(sizeof(int) * FILE_LINES);
    int *right_list = malloc(sizeof(int) * FILE_LINES);
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
        perror("An error occured when executing malloc");
        exit(1);
    }

    FILE *file = fopen(file_name, "r");
    if (file == NULL) {
        perror("An error occured when opening the file");
        exit(1);
    }

    for (int i = 0; i < FILE_LINES; i++) {
        fscanf(file, "%d   %d", &left_list[i], &right_list[i]);
    }

    fclose(file);
}

void sort_list(int *list, int length) {
    for (int i = 0; i < length; i++) {
        for (int j = 0; j < length - 1; j++) {
            if (!(list[j] > list[j + 1])) continue;

            int temp_value = list[j];
            list[j] = list[j + 1];
            list[j + 1] = temp_value;
        }
    }
}
