#include <ctype.h>
#include <stdbool.h>
#include <stddef.h>
#include <stdio.h>
#include <stdlib.h>
#include <string.h>

#include "../metadata.h"

int y2023_d02_p1(char *file_name, int max_line_length);
int y2023_d02_p2(char *file_name, int max_line_length);

bool is_game_possible(char *line_buffer);
int get_power_of_minimum_cubes(char *line_buffer);
void validate_file(FILE *file);

int main(void) {
    char *file_name = Y2023_D02_INPUT_FILE_NAME;
    int max_line_length = Y2023_D02_MAX_LINE_LENGTH;

    int p1_answer = y2023_d02_p1(file_name, max_line_length);
    int p2_answer = y2023_d02_p2(file_name, max_line_length);

    PRINT_ANSWER_INT(Y2023_D02_P1_LABEL, p1_answer);
    PRINT_ANSWER_INT(Y2023_D02_P2_LABEL, p2_answer);

    return EXIT_SUCCESS;
}

int y2023_d02_p1(char *file_name, int max_line_length) {
    FILE *file = fopen(file_name, "r");
    validate_file(file);

    int answer = 0;
    char line_buffer[max_line_length + 2];
    while (fgets(line_buffer, (int)sizeof(line_buffer), file) != NULL) {
        if (!is_game_possible(line_buffer)) continue;

        int game_id;
        // NOLINTNEXTLINE(clang-analyzer-security.insecureAPI.DeprecatedOrUnsafeBufferHandling)
        sscanf(line_buffer, "Game %d:", &game_id);

        answer += game_id;
    }

    fclose(file);
    return answer;
}

int y2023_d02_p2(char *file_name, int max_line_length) {
    FILE *file = fopen(file_name, "r");
    validate_file(file);

    int answer = 0;
    char line_buffer[max_line_length + 2];
    while (fgets(line_buffer, (int)sizeof(line_buffer), file) != NULL) {
        answer += get_power_of_minimum_cubes(line_buffer);
    }

    fclose(file);
    return answer;
}

bool is_game_possible(char *line_buffer) {
    enum { RGB_CUBE_VARIANTS = 3 };
    char *rgb_cube_colors[RGB_CUBE_VARIANTS] = {"red", "green", "blue"};
    int rgb_cube_limit[RGB_CUBE_VARIANTS] = {12, 13, 14};

    // Initial strtok is here to skip Game %d:
    // NOLINTNEXTLINE(clang-analyzer-deadcode.DeadStores)
    char *cube = strtok(line_buffer, ":,;");
    while ((cube = strtok(NULL, ":,;")) != NULL) {
        int cube_amount;
        char cube_color[16];

        // NOLINTNEXTLINE(clang-analyzer-security.insecureAPI.DeprecatedOrUnsafeBufferHandling)
        sscanf(cube, "%d %s", &cube_amount, cube_color);
        for (int i = 0; i < RGB_CUBE_VARIANTS; i++) {
            if (strcmp(cube_color, rgb_cube_colors[i]) != 0) continue;
            if (cube_amount > rgb_cube_limit[i]) return false;
        }
    }

    return true;
}

int get_power_of_minimum_cubes(char *line_buffer) {
    enum { RGB_CUBE_VARIANTS = 3 };
    char *rgb_cube_colors[RGB_CUBE_VARIANTS] = {"red", "green", "blue"};
    int rgb_minimum_cube_count[RGB_CUBE_VARIANTS] = {0, 0, 0};

    // Initial strtok is here to skip Game %d:
    // NOLINTNEXTLINE(clang-analyzer-deadcode.DeadStores)
    char *cube = strtok(line_buffer, ":,;");
    while ((cube = strtok(NULL, ":,;")) != NULL) {
        int cube_amount;
        char cube_color[16];

        // NOLINTNEXTLINE(clang-analyzer-security.insecureAPI.DeprecatedOrUnsafeBufferHandling)
        sscanf(cube, "%d %s", &cube_amount, cube_color);
        for (int i = 0; i < RGB_CUBE_VARIANTS; i++) {
            if (strcmp(cube_color, rgb_cube_colors[i]) != 0) continue;
            if (cube_amount <= rgb_minimum_cube_count[i]) break;

            rgb_minimum_cube_count[i] = cube_amount;
            break;
        }
    }

    int power_of_minimum_cubes = rgb_minimum_cube_count[0];
    for (int i = 1; i < RGB_CUBE_VARIANTS; i++) {
        power_of_minimum_cubes *= rgb_minimum_cube_count[i];
    }

    return power_of_minimum_cubes;
}

void validate_file(FILE *file) {
    if (file != NULL) return;

    perror(FOPEN_ERROR_MESSAGE);
    exit(EXIT_FAILURE);
}
