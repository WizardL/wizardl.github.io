---
title: Codeforce 1370C
description: "Codeforce 1370C solution explanation number theory"
---

<!-- toc -->

# Codeforce 1370C Number Game
Today we are gonna take a look of problem Codeforce 1370C
[Problem Link](https://codeforces.com/problemset/problem/1370/C)

> **Problem**
Ashishgup and FastestFinger play a game.
They start with a number n and play in turns. In each turn, a player can make any one of the following moves:
>* Divide n by any of its odd divisors greater than 1.
>* Subtract 1 from n if n is greater than 1.
>
>Divisors of a number include the number itself.
The player who is unable to make a move loses the game.
Ashishgup moves first. Determine the winner of the game if both of them play optimally.
**Input**
The first line contains a single integer t ($1≤t≤100$)  — the number of test cases. The description of the test cases follows.
The only line of each test case contains a single integer  — n ($1≤n≤10^9$).
**Output**
For each test case, print "Ashishgup" if he wins, and "FastestFinger" otherwise (without quotes).

### Idea
At first glance, you might think of alpha-beta pruning or dp...
But the numbers are too big (~1e9, a reasonable number of steps an online judge can process is around 1e5~1e6), I guess this'd be a math problem.
Look some example data, we may come up with these observations:
1. Player got number 1 lose
2. Player got number 2 win (do minus one)
3. Player got odd number win (divide by this odd number)

And then, since we can't possibly calculate what are the factors of inputed numbers(the complexity is too big for number~1e9), I thought we might need to use the number "$2$", since "$2$" has very good properties:
1. It's even
2. All numbers $n$ can be represented as $n=2^km$ where $k\in\mathbb{N}\cup\{0\}$ and $m\in\mathbb{N}$ odd.

Keeping this in mind, we found more observations: For number $n=2^km$
1. Player got number $n$ s.t. $k\ge2$, $m=1$ lose
2. Player got number $n$ s.t. $k\ge2$, $m\neq1$ lose
3. Player got number $n$ s.t. $k=1$, $m=1$ win (n=2)
4. Provided $k=1$, $m\neq1$
    * $m$ is prime: lose
    * $m$ is not prime: win

Eventually, we can deal with all numbers
### Code:
```c++
#include <bits/stdc++.h>
using namespace std;
int t, n, k, m;
inline void f() {
  int nn = n;
  k = 0;
  while (((nn >> 1) << 1) == nn) k++, nn >>= 1;
  m = n / (1 << k);
}
inline bool p() {
  int hf = sqrt(m);
  for (int i = 2; i <= hf; i++) {
    if (m % i == 0) return false;
  }
  return true;
}
main(void) {
  cin.tie(0);
  ios_base::sync_with_stdio(0);
  cin >> t;
  while (t--) {
    cin >> n;
    f();
    if (n != 1 and (n == 2 or (n & 1) or (k >= 2 and m != 1) or (k == 1 and m != 1 and !p())))
      cout << "Ashishgup\n";
    else
      cout << "FastestFinger\n";
  }
  return 0;
}
```
[Submission](https://codeforces.com/contest/1370/submission/86878205)