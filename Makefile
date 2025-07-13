# The variables defining the compiler, formatter, and linter are important.
# They need to be installed in your local machine in order to work.

COMPILER = clang
COMPILER_FLAGS = -lm -Wall -Wextra -O2

FORMATTER = clang-format
FORMATTER_CONFIG = {BasedOnStyle: Google, IndentWidth: 4}
FORMATTER_FLAGS = -i -style='$(FORMATTER_CONFIG)'

LINTER = clang-tidy
LINTER_FLAGS = -checks=bugprone-* -quiet -extra-arg=-fno-caret-diagnostics
LINTER_COMPILER_FLAGS = -Iinclude -Wall

SOURCE := $(shell find . -name "*.c" | sort)
HEADER := $(shell find . -name "*.h" | sort)

# Attempts to convert all *.c files to * (binary)
TARGET := $(SOURCE:.c=)

# Makes these commands runnable even if a file with the same name exists.
.PHONY: all clean check format lint

all: $(TARGET)
	@echo All files compiled!

# Predefined rule for the make all dependency.
%: %.c
	@$(COMPILER) $(COMPILER_FLAGS) -o $@ $<

run: all
	@for binary in $(TARGET); do \
		$$binary; \
	done

clean:
	@rm -rf $(TARGET)
	@echo All executables removed!

check: format lint
	@echo All files checked!

format:
	@$(FORMATTER) $(FORMATTER_FLAGS) $(SOURCE) $(HEADER)
	@echo All files formatted!

lint:
	@for file in $(SOURCE); do \
		$(LINTER) $(LINTER_FLAGS) $$file -- $(LINTER_COMPILER_FLAGS); \
	done
	@echo All files linted!
