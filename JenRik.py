#!/usr/bin/python3

import sys
import toml
import os


def print_help(binary_name):
    """ Print a basic help showing how to use Jenerik """
    print(f"USAGE : {binary_name} file.jrk | init path_to_binary")
    print("\tinit\tcreate a basic test file for the given binary")


def open_file(fp):
    """ Open the toml file and parse it """
    if not fp.endswith(".toml"):
        sys.exit("You must provide valid toml file")
    try:
        f = open(fp, 'r')
    except:
        sys.exit(f"Could not open file {fp}")
    c = f.read()
    content = toml.loads(c)  # Parse the toml file
    f.close()
    return content


def init_file(fp):
    """ create a default test file """
    test_file_name = 'test_' + fp + '.toml'

    default_file_content = [
        f"binary_path = \"{fp}\"\n\n",
        "# A sample test\n",
        "[test1]\n",
        "args = [\"-h\"]\n",
        "status = 0\n",
        "stdout=\"\"\n",
        "stderr=\"\"\n",
    ]

    if os.path.exists(test_file_name):
        sys.exit(f"{test_file_name} already exists, can't init the file")
    try:
        f = open(test_file_name, 'w')
    except:
        sys.exit(f"Could not create file {test_file_name}")
    for line in default_file_content:
        f.write(line)
    f.close()
    print(f"Initialized {test_file_name} with success")


def check_binary_validity(binary_path):
    """ check if the binary path is a valid executable file """
    if os.path.exists(binary_path):
        if not os.access(binary_path, os.X_OK):
            sys.exit(f"{binary_path} : is not executable")
    else:
        sys.exit(f"{binary_path} : file not found")


def check_tests_validity(test_name, values):
    pass


def check_file_validity(content, fp):
    binary_path = ""
    known_tests_keys = ['args', 'status', 'stdout', 'stderr']

    for key in content.keys():
        if key == "binary_path":
            binary_path = content[key]
        else:
            check_tests_validity(key, content[key])

    if binary_path == "":
        sys.exit(f"Could not find binary_path key in {fp}")

    check_binary_validity(binary_path)
    print(binary_path)


def main(argc, argv):
    if argc == 1 or argc > 3 or argc == 3 and argv[1] != 'init':
        print_help(argv[0])
        exit(1)
    if argc == 3:
        init_file(argv[2])
        exit(0)
    elif argc == 2:
        content = open_file(argv[1])
        check_file_validity(content, argv[1])


if __name__ == '__main__':
    main(len(sys.argv), sys.argv)
