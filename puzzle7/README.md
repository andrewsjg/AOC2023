# Solution for [Puzzle 7](https://adventofcode.com/2023/day/7)

I enjoyed this one. Once again taking the naive approach first. This can be simplified a lot.

Part 2 tripped me up as there are a lot of cases to deal with when there are wild cards in the hand. I kept finding cases that I hadnt accounted for. I think the code that deals with them can be simplified a LOT as I was just adding fixes for cases as I found them.

This is actually a BAD approach as every patch has a side effect you may not anticipate. I wasted a lot of time on this. A much better approach is to realise what the state of the cards can be without the wildcards and find all the hands that are made when a wildcards are added.
