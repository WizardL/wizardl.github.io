---
title: Codeforce 377A
description: "Codeforce 377A dfs graph"
---

<!-- toc -->

# Codeforce 377A Maze
Today we are gonna take a look of problem Codeforce 377A
[Problem Link](https://codeforces.com/problemset/problem/377/A)

> **Problem**
Pavel loves grid mazes. A grid maze is an $n × m$ rectangle maze where each cell is either empty, or is a wall. You can go from one cell to another only if both cells are empty and have a common side.
Pavel drew a grid maze with all empty cells forming a connected area. That is, you can go from any empty cell to any other one. Pavel doesn't like it when his maze has too little walls. He wants to turn exactly $k$ empty cells into walls so that all the remaining cells still formed a connected area. Help him.
**Input**
The first line contains three integers $n$, $m$, $k$ ($1 ≤ n, m ≤ 500$, $0 ≤ k < s$), where $n$ and $m$ are the maze's height and width, correspondingly, $k$ is the number of walls Pavel wants to add and letter $s$ represents the number of empty cells in the original maze.
Each of the next $n$ lines contains m characters. They describe the original maze. If a character on a line equals ".", then the corresponding cell is empty and if the character equals "#", then the cell is a wall.
**Output**
Print $n$ lines containing $m$ characters each: the new maze that fits Pavel's requirements. Mark the empty cells that you transformed into walls as "X", the other cells must be left without changes (that is, "." and "#").
It is guaranteed that a solution exists. If there are multiple solutions you can output any of them.
### Idea
After observing some test data, I found that if each time we choose the $leaf$ of the empty cell tree, then there will be no road to empty cell being blocked since there's no more empty tree along this road.
So we dfs into the tree, and when the dfs function is about to return, we convert that cell.
### Code:
```c++
#include <bits/stdc++.h>
using namespace std;
int n, m, k;
char mp[510][510];
bool vis[510][510];
void dfs(int x, int y) {
  vis[x][y] = 1;
  if (!k) return;
  if (!vis[x - 1][y] and mp[x - 1][y] == '.') dfs(x - 1, y);
  if (!k) return;
  if (!vis[x + 1][y] and mp[x + 1][y] == '.') dfs(x + 1, y);
  if (!k) return;
  if (!vis[x][y - 1] and mp[x][y - 1] == '.') dfs(x, y - 1);
  if (!k) return;
  if (!vis[x][y + 1] and mp[x][y + 1] == '.') dfs(x, y + 1);
  if (k) k--, mp[x][y] = 'X';
}
main(void) {
  cin.tie(0);
  ios_base::sync_with_stdio(0);
  cin >> n >> m >> k;
  for (int x = 0; x <= m; x++) mp[x][n + 1] = mp[x][0] = '#';
  for (int y = 0; y <= n; y++) mp[m + 1][y] = mp[0][y] = '#';
  for (int y = 1; y <= n; y++)
    for (int x = 1; x <= m; x++) cin >> mp[x][y];
  for (int y = 1; y <= n; y++)
    for (int x = 1; x <= m; x++) {
      if (mp[x][y] == '.') {
        dfs(x, y);
        goto A;
      }
    }
A:
  for (int y = 1; y <= n; y++) {
    for (int x = 1; x <= m; x++) {
      cout << mp[x][y];
    }
    cout << '\n';
  }
  return 0;
}
```
Note that the bunch of "if (!k) return;" is for stopping the dfs when there's no more cell need to convert.
[Submission](https://codeforces.com/problemset/submission/377/87199577)