# JenRik

JenRik is a simple but powerfull enough testing script.

## Use case

Before thinking about using it you should read this section.

It's very simple and fast to setup JenRik and write tests for it.\
 But..\
It can test only two things:
- exit status
- standart output

Why only that ? because it's enough to thest a majority of simple scripts and programs.\

JenRik is great to test error cases and not complicated output.\
If you need more you should probably look for something else.

If it's enough for you you're at the perfect place !
What's great with JeneRik is that it can test any language, it only needs an executable.

## Installation

```bash
git clone https://github.com/Yohannfra/JenRik
cd JenRik
sudo ./install.sh
```

## Quick Start

Let's say we need to test this basic shell script.

my_prog.sh : 
```bash
#!/bin/bash

if [[ $1 = "1" ]]; then
    echo "it's 1 !"
    exit 0
else
    echo "it's not 1"
    exit 1
fi
```

First we need to init a test file for our my_prog.sh.\
**A JeneRik test file ends with .jrk**
```bash
$ jenerik init ./my_prog.sh # you must give the path and not just the binary name
```

It will create a *my_prog.sh.jrk* file with this content:
```
BINARY_PATH=./my_prog.sh

# test name |args|exit_status|stdout|
sample test |-h|0||
```

The first line ```BINARY_PATH=./my_prog.sh``` indicates the path to the binary to test.

Lines that start with a "#" are commentaries, they are ignored.
The commentaty here explains how to create a test:
1. The name of the test
2. The command line arguments for this test
3. The expected exit status 
4. The expected standart output

The line 
```
sample test |-h|0||
```
Is an example test, it will start your program with '-h' as argument and expect it to exit 0.\
As you can see the field of the standart output is empty. It means that you won't test test it.\
You can also give an empty arguments field.

## USAGE
Once you have written the test file you just have to :
```
jenerik ./test_file.jrk
```

## Limitations
- For now the parser is quite basic so the .jrk syntax is very limited by it.
- You can't use '|' (pipe) caracters in you tests because it's used by the parser to delimit fields.

## Roadmap
-   Improve the parser and make it more flexible

## Licence
    This project is licensed under the terms of the MIT license.