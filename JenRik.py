#!/usr/bin/python3

import sys
import toml

def main(argc, argv):
    f = open(argv[1], 'r')

    c = f.read()

    parsed = toml.loads(c)

    for i in parsed.keys():
        print(i)

if __name__ == '__main__':
    main(len(sys.argv), sys.argv)
