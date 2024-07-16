package DynamicProgramming

// https://leetcode.com/problems/minimum-ascii-delete-sum-for-two-strings/description/
/*
Given two strings s1 and s2, return the lowest ASCII sum of deleted characters to make two strings equal.
*/
func minimumDeleteSum(s1 string, s2 string) int {
	return 0
}

// DP - Bottom-Up
/*
s1 = "sea", s2 = "eat"
1. ADS(sea, eat) => ADS(ea, eat) + s, ADS(sea, at) + e
2. ADS(ea, eat)  => ADS(a, at)
...
n. ADS("", "") => 0
```
   s    e    a    -
e  st   t    ...  eat
a  ...  ...  t    at
t  seat eat  at   t
-  sea  ea   a    0
```
dp[i][j] = ADS(s1[i:], s2[j:])
1. if dp[i] != dp[j]	=> dp[i+1][j] + s1[i]
                 		=> dp[i][j+1] + s2[j]
2. if dp[i] == dp[j] => dp[i+1][j+1]

f(string s1, int idx1, string s2, int idx2):
1. if s1[idx1] != s2[idx2]
	return min(f(s1, idx1+1, s2, idx2) + s1[idx1], f(s1, idx, s2, idx2+1) + s2[idx])
2. if s1[idx1] == s2[idx2]
	return f(s1, idx1+1, s2, idx2+1);

class Solution {
public:
  int minimumDeleteSum(string s1, string s2) {
    int m = s1.size(), n = s2.size();
    vector<vector<int>> dp(m+1, vector<int>(n+1, -1));
    dp[m][n] = 0;
    for (int i = m-1; i >= 0; i--)
      dp[i][n] = dp[i+1][n] + s1[i];
    for (int j = n-1; j >= 0; j--)
      dp[m][j] = dp[m][j+1] + s2[j];
    for (int i = m-1; i >= 0; i--) {
      for (int j = n-1; j >= 0; j--) {
        if (s1[i] != s2[j])
          dp[i][j] = min(dp[i+1][j] + s1[i], dp[i][j+1] + s2[j]);
        else
          dp[i][j] = dp[i+1][j+1];
      }
    }
    return dp[0][0];
  }
};
*/

// DP - Top-Down
/*
s1 = "sea", s2 = "eat"
"sea"
  ^
"eat"
    ^

f(string s1, int idx1, string s2, int idx2):
1. if s1[idx1] != s2[idx2]
    return min(f(s1, idx1+1, s2, idx2) + s1[idx1], f(s1, idx, s2, idx2+1) + s2[idx])
2. if s1[idx1] == s2[idx2]
    return f(s1, idx1+1, s2, idx2+1);

class Solution {
public:
  int backtracking(const string& s1, int idx1, const string& s2, int idx2, vector<vector<int>>& dp) {
    if (dp[idx1][idx2] != -1) return dp[idx1][idx2];
    if (idx1 >= s1.size() && idx2 >= s2.size()) return 0; // Base Case
    int sum1 = 0, sum2 = 0;
    if (idx1 >= s1.size()) {
      for (int i = idx2; i < s2.size(); i++) {
        sum2 += s2[i];
      }
      dp[idx1][idx2] = sum2;
      return sum2;
    }
    if (idx2 >= s2.size()) {
      for (int i = idx1; i < s1.size(); i++) {
        sum1 += s1[i];
      }
      dp[idx1][idx2] = sum1;
      return sum1;
    }
    if (s1[idx1] != s2[idx2])
      dp[idx1][idx2] = min(backtracking(s1, idx1+1, s2, idx2, dp) + s1[idx1], backtracking(s1, idx1, s2, idx2+1, dp) + s2[idx2]);
    else
      dp[idx1][idx2] = backtracking(s1, idx1+1, s2, idx2+1, dp);
    return dp[idx1][idx2];
  }

  int minimumDeleteSum(string s1, string s2) {
    vector<vector<int>> dp(s1.size() + 1, vector<int>(s2.size() + 1, -1));
    return backtracking(s1, 0, s2, 0, dp);
  }
};
*/
