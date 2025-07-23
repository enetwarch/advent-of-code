COMPILER = clang
COMPILER_FLAGS = -lm -Wall -Wextra -O2

SOURCE_DIRECTORY = solutions
TARGET_DIRECTORY = target

SOURCE_FILES := $(shell find $(SOURCE_DIRECTORY) -name "*.c" | sort)
TARGET_FILES := $(patsubst $(SOURCE_DIRECTORY)/%.c, $(TARGET_DIRECTORY)/%, $(SOURCE_FILES))

$(TARGET_DIRECTORY)/%: $(SOURCE_DIRECTORY)/%.c
	@mkdir -p $(dir $@)
	@$(COMPILER) $(COMPILER_FLAGS) -o $@ $<

.PHONY: all run clean

all: $(TARGET_FILES)
	@echo All files compiled!

run: $(TARGET_FILES)
	@for file in $(TARGET_FILES); do \
		./$$file; \
	done

clean:
	@rm -rf $(TARGET_DIRECTORY)
