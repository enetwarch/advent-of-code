#include <stdio.h>
#include <stdlib.h>

#define LINES 1000

int list_comparator(const void *x, const void *y) {
    return *(int *)x - *(int *)y;
}

int main(void) {
    FILE *file = fopen("input.txt", "r");
    if (file == NULL) {
        perror("There was an error in opening the file");
        return 1;
    }

    int left_list[LINES];
    int right_list[LINES];
    for (unsigned int i = 0; i < LINES; i++) {
        fscanf(file, "%d%d", &left_list[i], &right_list[i]);
    }
    fclose(file);

    qsort(left_list, LINES, sizeof(left_list[0]), list_comparator);
    qsort(right_list, LINES, sizeof(right_list[0]), list_comparator);

    int answer = 0;
    for (unsigned int i = 0; i < LINES; i++) {
        answer += abs(left_list[i] - right_list[i]);
    }

    printf("Day 1: %d\n", answer);
    return 0;
}
