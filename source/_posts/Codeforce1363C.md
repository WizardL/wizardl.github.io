---
title: Codeforce 1363C
description: "Codeforce 1363C math graph game"
---

<!-- toc -->

# Codeforce 1363C Game On Leaves
Today we are gonna take a look of problem Codeforce 1363C
[Problem Link](https://codeforces.com/problemset/problem/1363/C)

> **Problem**
Ayush and Ashish play a game on an unrooted tree consisting of n nodes numbered 1 to n. Players make the following move in turns:
>* Select any leaf node in the tree and remove it together with any edge which has this node as one of its endpoints. A leaf node is a node with degree less than or equal to 1.
>
>A tree is a connected undirected graph without cycles.
There is a special node numbered x. The player who removes this node wins the game.
Ayush moves first. Determine the winner of the game if each player plays optimally.
**Input**
The first line of the input contains a single integer t (1≤t≤10) — the number of testcases. The description of the test cases follows.
The first line of each testcase contains two integers n and x (1≤n≤1000,1≤x≤n) — the number of nodes in the tree and the special node respectively.
Each of the next n−1 lines contain two integers u, v (1≤u,v≤n, u≠v), meaning that there is an edge between nodes u and v in the tree.
**Output**
For every test case, if Ayush wins the game, print "Ayush", otherwise print "Ashish" (without quotes).

### Idea
Imagine how a game proceed: Provided that there's $k$ subtrees connected to vertex $x$. Players are ought to make $x$ a leaf, $i.e.$ removing subtrees till there's only $1$ subtree($i.e.$ the graph is a straight line).
Imagine the moment before the graph is going to become a straight line, one subtree has 1 one left(after removing this, there's only one subtree left), and another subtree has $s$ nodes left. If $s\neq1$, since no one will remove that 1 node left on the subtree that is about to be removed(or another player will win instantly), the players will start to remove that $s$ nodes.
$i.e.$ The game will definately become $node_1-node_X-node_2$ if there's at least 2 subtrees connect to vertex $x$, and then it's clear that who have to move first $now$ lose.
Now it left to determine who have to move first when things come down to $node_1-node_X-node_2$. We need to remove $n-3$ nodes before this situation.
(Yes, we don't need much information about how the edges connected except for the number of edges connected to $x$)
### Code:
```c++
#include <bits/stdc++.h>
using namespace std;
int t, n, x, u, v;
main(void) {
  cin.tie(0);
  ios_base::sync_with_stdio(0);
  cin >> t;
  while (t--) {
    cin >> n >> x;
    int cnt = 0;
    for (int i = 0; i < n - 1; i++) {
      cin >> u >> v;
      if (u == x or v == x) cnt++;
    }
    if (cnt < 2)
      cout << "Ayush\n";
    else {
      n -= 3;
      cout << ((n & 1) ? "Ayush\n" : "Ashish\n");
    }
  }
  return 0;
}
```
[Submission](https://codeforces.com/contest/1363/submission/87274342)