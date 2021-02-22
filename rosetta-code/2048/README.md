### Task
Implement a 2D sliding block puzzle game where blocks with numbers are combined to add their values.

The rules are that on each turn the player must choose a direction (up, down, left or right) and all tiles move as far as possible in that direction, some more than others. Two adjacent tiles (in that direction only) with matching numbers combine into one bearing the sum of those numbers. A move is valid when at least one tile can be moved, if only by combination. A new tile with the value of 2 is spawned at the end of each turn at a randomly chosen empty square, if there is one. To win, the player must create a tile with the number 2048. The player loses if no valid moves are possible.

The name comes from the popular open-source implementation of this game mechanic, 2048.

### Requirements:
- "Non-greedy" movement. The tiles that were created by combining other tiles should not be combined again during the same turn (move). That is to say that moving the tile row of
```
[2][2][2][2]
```
to the right should result in
```
 ......[4][4]
```
and not
```
 .........[8]
```
- "Move direction priority". If more than one variant of combining is possible, move direction shall indicate which combination will take effect. For example, moving the tile row of
```
...[2][2][2]
```
to the right should result in
```
......[2][4]
```
and not
```
......[4][2]
```
- Adding a new tile on a blank space. Most of the time, a new "2" is to be added and occasionally (10% of the time) - a "4".
- Check for valid moves. The player shouldn't be able to skip their turn by trying a move that doesn't change the board.
- Win condition.
- Lose condition.
