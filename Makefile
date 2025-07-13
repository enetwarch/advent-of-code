# Compiler and flags
CC = clang
CFLAGS = -lm -Wall -Wextra -O2

# Finds all files that matches "*.c" regex
# Attempts to convert it to binary (no extension)
SRCS := $(shell find . -name "*.c")
OBJS := $(SRCS:.c=)

all: $(OBJS)
	@echo All C files compiled!

%: %.c
	$(CC) $(CFLAGS) -o $@ $<

clean:
	rm -rf $(OBJS)
	@echo All executables removed!
