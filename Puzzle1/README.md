## Solution for [Puzzle 1](https://adventofcode.com/2023/day/1)

A little mind bender for the first puzzle. But if you are good at string wrangling you can figure out many ways to tackle and solve this one.

Basic strategy for part one and part two was to search for substrings from the start and end of each line to find the first and last occurrences of the digits.

Part two required some thinking time! I didnt bother trying to find digits in the string and just used substring searchs to find the first and last occurrences as required. Got stumped when my final value was to high. I was only looking for the FIRST occurrence of the substring in the input line, which meant if a digit was repeated in a string I would miscount. Fixed by finding both the first and last occurrence of each digit substring in the input string.