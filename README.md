# Equation Evaluator

## Problem Description
Your task is to write a command-line program to evaluate a set of equations, each specified on separate lines. An equation is defined by:

```
<LHS> = <RHS>
```

`<LHS>` is the left-hand side of the equation and is always a variable name. A variable name can only be composed of letters from the alphabet (e.g. for which isalpha(c) is 1). `<RHS>` is the right hand side of the equation and can be composed of variables, unsigned integers, and the + operator.

Here is one example set of equations:

```
offset = 4 + random + 1
location = 1 + origin + offset
origin = 3 + 5
random = 2
```

Your program should take a filename as input. The file contains a set of equations, like the above. It should evaluate the set of equations and print the unsigned integer value of each variable.

```
<variable name> = <unsigned integer value>
```

The output should be sorted by in ascending order by variable name.

The output for the example above would be:

```
location = 16
offset = 7
origin = 8
random = 2
```

You may assume the following: You may assume the input is well formed. There will be one or more white spaces between each token. You may use C++, the Standard C libraries and the Standard Template Library (STL). You may use std::sort and qsort. All variables in the equation set will have a definition. You may also assume a variable is only defined once and will have a valid integer solution.

## To run this code
 - You will need a working Go environment. Tested with `go1.9.2 linux/amd64`.
 - This code only uses standard packages, so it is not confined to a Go workspace.
 - `go run main.go input.txt` where input.txt contains the well-formed equations. Results will be printed to std out.
 - to use my tests, please use test.sh from a bash environment.

 ## Thoughts
  - I did not include much error checking beyond file operations, because I assume that the input is well formed.
  - For the parser, while the typical approach is to build a tree completely describing the syntax, given that only positive integers would be used and that order-of-operations was irrelavant, I could simplify expressions to a list of variables and the sum of constants. For the same reason, I choose to store variable values in a map, rather than storing pointers to different elements.
  - I assume that cases like `variable = variable` are not well-formed input; given that a variable *will* have valid integer solution, and a variable can only be defined once, it would be impossible to assign a value to `variable`.
  - The solve function could possibly be optimized by sorting equations in ascending order by the number of unknowns. For small outputs (that I could type out by hand), the performance difference would not matter.