# JenRik

JenRik is a simple but powerful testing tool.

![](https://github.com/Yohannfra/JenRik/workflows/JenRik/badge.svg)

The main idea was to write a generic binary testing tool.\
JenRik simply parse a [toml](https://github.com/toml-lang/toml)
file containing the tests and run them.

## Installation

#### Dependencies:
The dependencies are the python parser for toml and termcolor. Install both with:
```
pip install -r requirements.txt
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
$ jenrik init ./my_prog.py
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
- **stdin** : write in the stdin of the process
- **stdin_file** : write in the stdin of the process from the content of a file
- **env** : change environment variable(s) (replace the value with the given one)
- **add_env** : change environment variable(s) (append the given value to environment value)
- **repeat** : repeat the test x times

### Example of how to use some commands

- **pre** is usefull if you need to prepare a file needed by your programm for a test
- **post** is mainly usefull to cleanup after a test
- **stderr_file** and **stdout_file** are usefull if the output of you program is on multiples lines or if it's a lot of text and you don't want it written in you test file.
- **should_fail** is used if you want a test fail to be its success
- **add_env** is mainly used for environment variables like PATH (when you want to append a value to the existing one)
- **repeat** is super usefull if you want to test a proram that relies on some random data. It makes it easier to run many tests to check if it's always working

⚠ **Don't forget that the paths are all relatives to the test file.**

If you want more examples on how to write tests you should see this [file](test_jenrik.toml)

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

# stdin
stdin = ["Hello", "World"]

# stdin_file
stdin_file = "my_stdin.txt"

# env
env.USER = "toto"
env.TERM = "xterm"

# add_env
add_env.PATH= ":~/.my_folder"

# repeat
repeat = 12
```

See [Usage](#Usage) to run the tests

## Usage
Once you have written the test file you just have to :
```
jenrik test_my_prog.toml
```

The output will look like that
```
test_no_arguments : OK
with_args : OK

Summary ./my_prog.py: 2 tests ran
2 : OK
0 : KO
```


### Build

If you want to build your programm when calling jenrik you can use the **build_command** option.\
It will run it before executing the tests. eg:

```toml
binary_path = "./my_program"

build_command = "make"

[test_example]
args = []
status = 1
# ...
```

### Exit status

If a parsing error or a configuration error occurs then the exit status is 1\
Otherwise the exit status is the number of failed tests

## Tests

JenRik tests itself.
```
jenrik test_jenrik.toml
```

You can also run JenRik tests within a docker
```
$ docker build -t image_jenrik .
$ docker run image_jenrik:latest
```

## Current version
```
v 1.07
```

## Licence
    This project is licensed under the terms of the MIT license.
