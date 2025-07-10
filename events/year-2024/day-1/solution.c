#include <stdio.h>
#include <stdlib.h>
#include <stddef.h>

#define FILE_NAME "input.txt"
#define LINES 1000

int compare_list(const void *one, const void *two) {
    int int_one  = *(int *)one;
    int int_two = *(int *)two;

    if (int_one > int_two) return 1;
    if (int_one < int_two) return -1;
    return 0;
}

int part_one() {
    FILE *file = fopen(FILE_NAME, "r");
    const int lines = LINES;

    int left_list[lines];
    int right_list[lines];
    for (size_t i = 0; i < lines; i++) {
        fscanf(file, "%d%d", &left_list[i], &right_list[i]);
    }

    fclose(file);

    qsort(left_list, lines, sizeof(left_list[0]), compare_list);
    qsort(right_list, lines, sizeof(right_list[0]), compare_list);

    int answer = 0;
    for (size_t i = 0; i < lines; i++) {
        answer += abs(left_list[i] - right_list[i]);
    }

    return answer;
}

int part_two() {
    FILE *file = fopen(FILE_NAME, "r");
    const int lines = LINES;

    int left_list[lines];
    int right_list[lines];
    for (size_t i = 0; i < lines; i++) {
        fscanf(file, "%d%d", &left_list[i], &right_list[i]);
    }

    fclose(file);

    int answer = 0;
    for (size_t i = 0; i < lines; i++) {
        int left_number = left_list[i];
        
        int similarity_score = 0;
        for (size_t j = 0; j < lines; j++) {
            int right_number = right_list[j];
            if (left_number == right_number) similarity_score++;
        }

        answer += left_list[i] * similarity_score;
    }

    return answer;
}

int main(void) {
    printf("Part 1: %d\n", part_one());
    printf("Part 2: %d\n", part_two());

    return 0;
}
