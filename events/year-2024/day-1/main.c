#include <stdio.h>

int main() {
    FILE* file = fopen("input.txt", "r");
    if (file == NULL) {
        perror("There was an error in opening the file");
        return 1;
    }

    char line[64];
    while (fgets(line, sizeof(line), file)) {
        printf("%s", line);
    }

    return 0;
}

