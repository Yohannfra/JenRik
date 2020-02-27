# JenRik

JenRik is a simple but powerfull testing script.

The main idea was to write a generic binary testing script.\
JenRik simply parse a [toml](https://github.com/toml-lang/toml)
file containing the tests and run them.

## Installation

#### Dependencies:
The only dependency is the python parser for toml. Install it with:
```
pip install toml
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

##### First we need to init a test file for our my_prog.py

```bash
$ jenerik init ./my_prog.py # you must give the path and not just the binary name
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

⚠ Each test require at least the args and status values !\
stdout and stderr can be omitted.

If you want more examples on how to write tests you should see this [file](test_JenRik.toml)

See [#USAGE](Usage) to run the tests

## USAGE
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


## Roadmap
- Add the possibility to diff the output with an existing file
- Add a pre and a post command

## Licence
    This project is licensed under the terms of the MIT license.
