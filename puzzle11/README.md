# Solution for [Puzzle 11](https://adventofcode.com/2023/day/11)

This is a path fiding challenge. Because I already implemented A star last year, so I am going to reuse it here. I used it in part 1 because I suspected that
part 2 may be a more complicated path finding challenge. However this turned out not to be the case and A* is NOT the best option at all. The challenge was correctly expanding the map. For both part 1 and 2 you can use the much simpler and quick 'manhattan distance' formula.

Part 1 was slow to complete using A* naievly. A lot of distances to work out. Using goroutines it was tollerable. But its nonsenical to use it.

For part 2 the working out the expansion properly sent me down a few rabbit holes because the distances werent working out. Took WAY longer to correct that it should have.
