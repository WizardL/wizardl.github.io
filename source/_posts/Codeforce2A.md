---
title: Codeforce 2A
description: "Codeforce 2A solution explanation hashing implementation"
---

<!-- toc -->

# Codeforce 2A Winner
Today we are gonna take a look of problem Codeforce 2A
[Problem Link](https://codeforces.com/problemset/problem/2/A)

> **Problem**
The winner of the card game popular in Berland "Berlogging" is determined according to the following rules. If at the end of the game there is only one player with the maximum number of points, he is the winner. The situation becomes more difficult if the number of such players is more than one. During each round a player gains or loses a particular number of points. In the course of the game the number of points is registered in the line "name score", where name is a player's name, and score is the number of points gained in this round, which is an integer number. If score is negative, this means that the player has lost in the round. So, if two or more players have the maximum number of points (say, it equals to $m$) at the end of the game, than wins the one of them who scored at least m points first. Initially each player has 0 points. It's guaranteed that at the end of the game at least one player has a positive number of points.
**Input**
The first line contains an integer number n (1  ≤  n  ≤  1000), n is the number of rounds played. Then follow n lines, containing the information about the rounds in "name score" format in chronological order, where name is a string of lower-case Latin letters with the length from 1 to 32, and score is an integer number between -1000 and 1000, inclusive.
**Output**
Print the name of the winner.

### Idea
At first, to store the score of a given player, we need a technique "hashing", convert string into array index.
C++ std::map or std::unordered_map had done this for us(literally, unordered_map has no some member function such as upper_bound), so we can just use _unordered\_map<string, int> mp;_ to store scores.
Now there's only some cumbersome steps such as get the maximum final score and find the first player has score greater equal to maximum final score.
To deal with these, we store the input data into _int_ _sc[\_n];_ and _string_ _s[\_n];_ such that we can easily get these things.
### Code:
```c++
#include <bits/stdc++.h>
using namespace std;
const int _n = 1010;
int n, sc[_n];
string s[_n];
unordered_map<string, int> mp;
vector<string> candidate;
unordered_map<string, bool> is_candidate;
main(void) {
  cin.tie(0);
  ios_base::sync_with_stdio(0);
  cin >> n;
  for (int i = 0; i < n; i++) {
    cin >> s[i] >> sc[i];
    if (!mp.count(s[i])) mp[s[i]] = 0;
    mp[s[i]] += sc[i];
  }
  int maxx = -1;
  for (auto it = mp.begin(); it != mp.end(); it++) maxx = max(maxx, it->second);
  for (int i = 0; i < n; i++)
    if (mp[s[i]] == maxx) candidate.push_back(s[i]), is_candidate[s[i]] = 1;
  string ans = candidate[0];
  if (candidate.size() > 1) {
    mp.clear();
    for (int i = 0; i < n; i++) {
      if (!mp.count(s[i])) mp[s[i]] = 0;
      mp[s[i]] += sc[i];
      if (mp[s[i]] >= maxx and is_candidate.count(s[i]) == 1 and is_candidate[s[i]]) {
        ans = s[i];
        break;
      }
    }
  }
  cout << ans << '\n';
  return 0;
}
```
[Submission](https://codeforces.com/contest/2/submission/86959676)