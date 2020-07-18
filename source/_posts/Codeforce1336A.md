---
title: Codeforce 1336A
description: "Codeforce 1336A solution explanation dp tree-dp sort nth_element dfs graph greedy"
---

<!-- toc -->

# Codeforce 1336A Linova and Kingdom
Today we are gonna take a look of problem Codeforce 1336A
[Problem Link](https://codeforces.com/problemset/problem/1336/A)

> **Problem**
Writing light novels is the most important thing in Linova's life. Last night, Linova dreamed about a fantastic kingdom. She began to write a light novel for the kingdom as soon as she woke up, and of course, she is the queen of it.
![](https://i.imgur.com/ERbyCyh.png)
There are $n$ cities and $n−1$ two-way roads connecting pairs of cities in the kingdom. From any city, you can reach any other city by walking through some roads. The cities are numbered from $1$ to $n$, and the city $1$ is the capital of the kingdom. So, the kingdom has a tree structure.
As the queen, Linova plans to choose exactly $k$ cities developing industry, while the other cities will develop tourism. The capital also can be either industrial or tourism city.
A meeting is held in the capital once a year. To attend the meeting, each industry city sends an envoy. All envoys will follow the shortest path from the departure city to the capital (which is unique).
Traveling in tourism cities is pleasant. For each envoy, his happiness is equal to the number of tourism cities on his path.
In order to be a queen loved by people, Linova wants to choose $k$ cities which can maximize the sum of happinesses of all envoys. Can you calculate the maximum sum for her?
**Input**
The first line contains two integers $n$ and $k$ ($2\le n\le2\cdot 10^5$, $1\le k<n$)  — the number of cities and industry cities respectively.
Each of the next $n−1$ lines contains two integers $u$ and $v$ ($1\le u,v\le n$), denoting there is a road connecting city $u$ and city $v$.
It is guaranteed that from any city, you can reach any other city by the roads.
**Output**
Print the only line containing a single integer  — the maximum possible sum of happinesses of all envoys.
### Idea
First, as the problem states, "It is guaranteed that from any city, you can reach any other city by the roads.", and there's exactly $n-1$ edges, so we can sure that the structure of the graph is a $good$ tree with no cycles and is connected.
Since we know the graph is a tree, to find the shortest path for any vertex to vertex 1, we can try regard vertex 1 as root of the tree(every vertex can be regarded as root), then things gets clear now, the shortest path from $v$ to $1$ is the depth of the vertex $v$(depth of $1$ is $0$).
Now, we can relatively easy to find that we can choose industry cities greedily, choose the maximum depth vertex, and then gradually decrease...
To calculate the maximum needs some more effort, since we know, if the newly chosen industry city $v$ has child nodes, then the happiness of these nodes(has greater depth of $v$) will decrease.
Now we need to first realize one thing: We will first choose greater depth nodes $then$ lower depth nodes.
So we can do the following: We assign a value "happy" to each node, which is the happiness the node will $contribute$ when the node is chosen. Note that we use the term $contribute$ since when a node being chosen, we need to consider the decrease of the happiness of child nodes.
When a node $v$ being chosen, the happiness contribute is $depth[v]-number\ of\ childs$. With happiness being calculated, we can sort the happy array in decresing order and sum up first $k$ elements, which is the answer. Note that in the sorted array, the order is also from deeper depth to shallow depth since depth is decresing as we choose the nodes, and the $negtive$ of number of child is also decresing as we choose the nodes, so the happiness we assigned have the same order as the depth.
To calculate the desired values, we use dfs. When the child tree is being calculated, we can then calculate the values for current node.
### Code:
```c++
#include <bits/stdc++.h>
using namespace std;
const int _n = 2e5 + 10;
int n, k, uu, vv, num[_n], happy[_n];
vector<int> G[_n];
void dfs(int v, int d, int fa) {
  for (int i = 0; i < G[v].size(); i++)
    if (G[v][i] != fa) dfs(G[v][i], d + 1, v);
  num[v] = 1, happy[v] = d;
  for (int i = 0; i < G[v].size(); i++) {
    if (G[v][i] != fa) {
      num[v] += num[G[v][i]];
      happy[v] -= num[G[v][i]];
    }
  }
}
main(void) {
  cin.tie(0);
  ios_base::sync_with_stdio(0);
  cin >> n >> k;
  for (int i = 0; i < n - 1; i++) {
    cin >> uu >> vv;
    G[uu].push_back(vv), G[vv].push_back(uu);
  }
  dfs(1, 0, 1);
  nth_element(happy + 1, happy + k + 1, happy + n + 1, greater<int>());
  long long sum = 0;
  for (int i = 1; i <= k; i++) sum += happy[i];
  cout << sum << '\n';
  return 0;
}
```
Note that I use nth-element in substitute of sort function as the tutorial of codeforce. nth-element has $O(n)$ average complexity but no guarentee on worst complexity. In most implementation, the worst complexity is $O(n\log n)$, and in some implementation, the worst complexity is $O(n^2)$
(As I tested, the sort function version is quicker than nth-element version by 16ms)
[Submission](https://codeforces.com/contest/1336/submission/87233490)