import re
import os


def fix_file(file_path):
    with open(file_path, "r") as f:
        content = f.read()

    # Fix pattern: var xxx Type = {}
    pattern1 = r"var\s+([A-Za-z0-9_]+)\s+([A-Za-z0-9_]+)\s*=\s*{}"
    content = re.sub(pattern1, r"var \1 \2", content)

    # Fix pattern: var xxx Type = false
    pattern2 = r"var\s+([A-Za-z0-9_]+)\s+([A-Za-z0-9_]+)\s*=\s*false"
    content = re.sub(pattern2, r"var \1 \2", content)

    with open(file_path, "w") as f:
        f.write(content)


def main():
    for root, _, files in os.walk("."):
        for file in files:
            if file.endswith(".go"):
                file_path = os.path.join(root, file)
                fix_file(file_path)
    print("Fix openapi vars done.")


if __name__ == "__main__":
    main()
