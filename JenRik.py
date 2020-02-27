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
        print("You must provide valid toml file")
        exit(1)
    try:
        f = open(fp, 'r')
    except:
        print(f"Could not open file {fp}")
        exit(1)
    c = f.read()
    content = toml.loads(c) # Parse the toml file
    f.close()
    return content

def init_file(fp):
    """ create a default test file """
    test_file_name = 'test_' + fp + '.toml'

    if os.path.exists(test_file_name):
        print(f"{test_file_name} already exists, can't init the file")
        exit(1)
    try:
        f = open(test_file_name, 'w')
    except:
        print(f"Could not create file {test_file_name}")
        exit(1)
    f.write(f"binary_path = \"{fp}\"\n\n")
    f.write("# A sample test\n")
    f.write("[test1]\n")
    f.write("args = [\"-h\"]\n")
    f.write("status = 0\n")
    f.write("stdout=\"\"\n")
    f.write("stderr=\"\"\n")
    f.close()
    print(f"Initialized {test_file_name} with success")


def check_file_validity(content):
    for i in content.keys():
        print(i)
        for k in content[i]:
            print("| ", end="")
            print("", k)


def main(argc, argv):
    if argc == 1 or argc > 3 or argc == 3 and argv[1] != 'init':
        print_help(argv[0])
        exit(1)
    if argc == 3:
        init_file(argv[2])
        exit(0)
    elif argc == 2:
        content = open_file(argv[1])
        check_file_validity(content)
        pass

if __name__ == '__main__':
    main(len(sys.argv), sys.argv)
