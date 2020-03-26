# JenRik

JenRik is a simple but powerful testing tool.

The main idea was to write a generic binary testing tool.\
JenRik simply parse a [toml](https://github.com/toml-lang/toml)
file containing the tests and run them.

## Installation

#### Dependencies:
The dependencies are the python parser for toml and termcolor. Install both with:
```
pip install toml
pip install termcolor
```

#### Installation from sources:
```bash
git clone https://github.com/Yohannfra/JenRik
cd JenRik
sudo ./install.sh
```

## Quick Start

##### Let's say we need to test this basic python script.

my_prog.py :
```python
#!/usr/bin/python3
import sys

if len(sys.argv) == 1:
    print("No arguments given")
    exit(1)
else:
    print(sys.argv[1])
```

##### First we need to initialize a test file for our my_prog.py

```bash
$ jenerik init ./my_prog.py
```

It will create a *test_my_prog.toml* file with this content:
```toml
binary_path = "my_prog.py"

# A sample test
[test1]
args = ["-h"]
status = 0
stdout=""
stderr=""
```

The first line ```binary_path = "my_prog.py"``` indicates the path to the binary to test.

The line ```[test1]``` define a test and the following values are parts of it:
```toml
args = ["-h"] # the command line arguments
status = 0    # the expected exit status
stdout=""     # the expected stdout (not tested if empty)
stderr=""     # the expected stderr (not tested if empty)
```

Now we can write the tests for our my_prog.py. First delete the sample test and write this instead
```toml
[test_no_arguments]
args = []
status = 1
stdout="No arguments given\n"

[with_args]
args = ["Hello"]
status = 0
stdout="Hello\n"
stderr=""
```

⚠ **Each test requires at least the args and status values !**

There are many other available commands:
- **pre** : run a shell command before executing the test
- **post** : run a shell command after executing the test
- **stderr_file** : compare your program stderr with the content of a given file
- **stdout_file** : compare your program stdout with the content of a given file
- **pipe_stdout** : redirect your program stdout to a specified shell command before checking it
- **pipe_stderr** : redirect your program stderr to a specified shell command before checking it
- **should_fail** : make the test success if it fails
- **timeout** : make the test fail if it times out, after killing it (SIGTERM) (the time is given in seconds)

**pre** is usefull if you need to prepare a file needed by your programm for a test\
**post** is mainly usefull to cleanup after a test\
**stderr_file** and **stdout_file** are usefull if the output of you program is on multiples lines or if it's a lot of text and you don't want it written in you test file.
**should_fail** is used if you want a test fail to be its success

⚠ **Don't forget that the paths are all relatives to the test file.**

If you want more examples on how to write tests you should see this [file](test_JenRik.toml)

Here is a quick example of how to use all availables commands

```toml
# args
args = []
args = ["-h"]
args = ["1", "2", "3"]

# status
status = 1

# stdout
stdout="Hello\n"

# stderr
stderr="Hello err\n"

# pre
pre = "touch test.txt && echo 'hello' > test.txt"

# post
post = "rm -f test.txt"

# stderr_file
stderr_file = "./my_file.txt"

# stdout_file
stdout_file = "./my_file.txt"

# pipe_stdout
pipe_stdout = "| grep 'Usage'"

# pipe_stderr
pipe_stderr = "| cut -d ':' -f1"

# should_fail (true or false)
should_fail = true

# timeout (in seconds)
timeout = 0.4

```

See [Usage](#Usage) to run the tests

## Usage
Once you have written the test file you just have to :
```
jenerik test_my_prog.toml
```

The output will look like that
```
test_no_arguments : OK
with_args : OK

Summary ./my_prog.py: 2 tests ran
2 : OK
0 : KO
```

## Tests

Jenrik tests itself.
```
JenRik test_JenRik.toml
```

## Roadmap
- [x] the possibility to diff the output with an existing file
- [x] Add a pre and a post command
- [x] Add pipe_stdout and pipe_stderr
- [x] Add a timeout feature

## Licence
    This project is licensed under the terms of the MIT license.
