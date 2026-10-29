# HDE-Challenge-2016

A Go solution to HDE Challenge 003, a programming challenge set by HDE, Inc. (now HENNGE) as part of its hiring process in 2016: sum the squares of the positive numbers in each test case, without using a `for` statement.

## Overview

The challenge brief is not included in the repository. From the header comment in `HDEChallenge.go`, the program "calculates the sum of squares given by standard input" and "does not use FOR statement". Since `for` is Go's only loop keyword, every loop is written as a recursive function instead.

## How it works

1. Read `N`, the number of test cases (the code notes 1 <= N <= 100).
2. For each test case, read `X`, the count of integers, then the `X` integers.
3. Sum the squares of the integers, ignoring any that are zero or negative.
4. After all input has been read, print each test case's sum on its own line.

Reading the cases, reading each case's integers, summing and printing are each done by a separate recursive function (`populateSums`, `fillArray`, `sumOfSquaresArray`, `printArray`).

## Input and output

Example input:

```
2
4
3 -1 1 14
5
9 6 -53 32 16
```

Output:

```
206
1397
```

The integers in each case are read into `int8` values, so each one must be between -128 and 127, and indices are `uint8`, so `N` and `X` must be at most 255.

## Getting started

### Prerequisites

- A Go toolchain. The program uses only the standard library and needs no `go.mod`.

### Run

```sh
go run HDEChallenge.go < input.txt
```

Or build a binary:

```sh
go build -o hde HDEChallenge.go
./hde < input.txt
```

## Project structure

```
HDEChallenge.go   The complete solution
```
