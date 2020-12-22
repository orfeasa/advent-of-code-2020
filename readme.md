# 🎄 Advent of Code 2020 🎄

![AoC2020 logo](https://raw.githubusercontent.com/orfeasa/advent-of-code-2020/main/header.png)

## Summary

[Advent of Code](http://adventofcode.com/) is an annual Advent calendar of programming puzzles.

This year I am doing it in Go.

## Running the code

To run the code of a specific day from the root directory run the following, replacing `XX` with the day number, `01` - `25`:

```sh
go run day_XX/main.go
```

Make sure you [have Go installed](https://golang.org/doc/install).

To run the code of all days run the script:

```sh
./run_all.sh
```

Make sure you have given permission to execute (`chmod +x run_all.sh`).

## Overview

| Day                                        | Name                    | Stars |
| ------------------------------------------ | ----------------------- | ----- |
| [01](https://adventofcode.com/2020/day/1)  | Report Repair           | ⭐⭐    |
| [02](https://adventofcode.com/2020/day/2)  | Password Philosophy     | ⭐⭐    |
| [03](https://adventofcode.com/2020/day/3)  | Toboggan Trajectory     | ⭐⭐    |
| [04](https://adventofcode.com/2020/day/4)  | Passport Processing     | ⭐⭐    |
| [05](https://adventofcode.com/2020/day/5)  | Binary Boarding         | ⭐⭐    |
| [06](https://adventofcode.com/2020/day/6)  | Custom Customs          | ⭐⭐    |
| [07](https://adventofcode.com/2020/day/7)  | Handy Haversacks        | ⭐⭐    |
| [08](https://adventofcode.com/2020/day/8)  | Handheld Halting        | ⭐⭐    |
| [09](https://adventofcode.com/2020/day/9)  | Encoding Error          | ⭐⭐    |
| [10](https://adventofcode.com/2020/day/10) | Adapter Array           | ⭐⭐    |
| [11](https://adventofcode.com/2020/day/11) | Seating System          | ⭐⭐    |
| [12](https://adventofcode.com/2020/day/12) | Rain Risk               | ⭐⭐    |
| [13](https://adventofcode.com/2020/day/13) | Shuttle Search          | ⭐⭐    |
| [14](https://adventofcode.com/2020/day/14) | Docking Data            | ⭐⭐    |
| [15](https://adventofcode.com/2020/day/15) | Rambunctious Recitation | ⭐⭐    |
| [16](https://adventofcode.com/2020/day/16) | Ticket Translation      | ⭐⭐    |
| [17](https://adventofcode.com/2020/day/17) | Conway Cubes            | ⭐⭐    |
| [18](https://adventofcode.com/2020/day/18) | Operation Order         | ⭐⭐    |
| [19](https://adventofcode.com/2020/day/19) | Monster Messages        | ⭐⭐    |
| [20](https://adventofcode.com/2020/day/20) | Jurassic Jigsaw         | ⭐     |
| [21](https://adventofcode.com/2020/day/21) | Allergen Assessment     | ⭐⭐    |
| [22](https://adventofcode.com/2020/day/22) | Crab Combat             | ⭐⭐    |

## Linting

```sh
gofmt -s -w .
git ls-files | grep .go | xargs golint
```
