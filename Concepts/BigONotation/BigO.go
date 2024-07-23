package BigONotation

/*
# Big O Notation

Ref:
- The Complete Guide to Big O Notation & Complexity Analysis for Algorithms: Part 1: https://www.youtube.com/watch?v=HfIH3czXc-8
- The Complete Guide to Big O Notation & Complexity Analysis for Algorithms: Part 2: https://www.youtube.com/watch?v=zo7YFqw5hNw

## Simplifying Big O
- **Product Rule**: If the Big O is the product of multiple terms, drop the constant terms
- **Sum Rule**: If the Big O is the sum of multiple terms, keep the largest term then drop the rest

## 7 Common Complexity Classes
- Constant O(n): The number of steps doesn't depend on the input size
- Logarithmic O(log(n)): The number of steps can be expressed as a logarithm on the input size
	- A log is the opposite of an exponent
	- An exponent is a repeated multiplication; a log is a repeated division
- Linear O(n)
- Loglinear O(n*log(n)):
	- Has linear behaviour nested in log steps
	- Bigger than O(n) but smaller than O(n^2)
- Polynomial O(n^c)
	- c is some constants
	- Include O(n^2) quadratic & O(n^3) cubic, etc
- Exponential O(c^n)
	- n is the size of the input & c is some constants
	- Include O(2^n) & O(3^n), etc
- Factorial O(n!)
*/
