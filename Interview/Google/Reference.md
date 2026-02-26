
## Recommendation:
### [Google L3 Interview Expierence | Full Questions](https://leetcode.com/discuss/post/6438710/google-l3-interview-expierence-full-ques-4yxt/)
- [Google Interview L3 Question: Phone](https://leetcode.com/discuss/post/5720580/google-swe-2-l3-bangalorehyd-aug-24-by-a-7kr8/)
- [Google Interview L3 Question: Round 2](https://leetcode.com/discuss/post/6409785/google-l3-full-question-by-anonymous_use-qlsk/)
- [Google Interview L3 Question: Round 3](https://leetcode.com/discuss/post/6437651/google-interview-l3-question-by-anonymou-lntb/)

**Googlyness**:
- Conflict with co-worker or team
- Challenge at work
- When did you promote an inclusive work environment? (I fumbled during this question)
- Time when you had to adapt
- Imagine your Team is very close to releasing a critical bug fix for your application. However, prior to releasing the application, you recognize that there are some issues with certain user subgroups of the application. What are you going to do?

**Here are a couple of helpful tips for the preparation**:
- Even though people say Google doesn't repeat questions, I encountered three questions during my interviews which were already leaked in the discuss section (phone screen, onsite 2 and onsite 3).
- Most people just focus on generic DSA content like `Neetcode` and `Strivers A2Z` for Interview preparation, but I truly believe that the most efficient way to practice is by solving questions that were actually asked by a company in past interviews.
- For instance, the interviewer for onsite 3 mentioned that barely 10% of all candidates are able to solve this question. Quite honestly, if I didn't encounter this question previously, I would've also stepped out of the interviewer.
- For my preparation I've compiled all Google questions from the **LeetCode** `Interview Questions` and `Interview Experience` section in an excel sheet and solved almost all of them. I looked at all questions from 2022 till 2024 and felt very well prepared by doing that. Of course it's a huge time investment, but landing a six figure salary is certainly worth the investment (at least in my country).
- I also used this link [Google Interview Questions Compilation](https://leetcode.com/discuss/interview-question/6185127/Google-Interview-Question-Compilation) for my prep journey and found it really helpful

### [Google Interview questions and preparation](https://leetcode.com/discuss/post/5547675/google-interview-questions-and-preparati-u9tg/)
#### Data Structures and Algorithms:
- [307. Range Sum Query - Mutable](https://leetcode.com/problems/range-sum-query-mutable/)
- [542. 01 Matrix](https://leetcode.com/problems/01-matrix/description/)
- **String Validator**: Given a tree where each node is a character, you can only move to its children, not the parent. With a stream of input characters, check if the character is a child of the current node, move to it if yes, or throw an `InvalidStringException` if no. If at a leaf node, move back to the root. Implement this as an iterator class with `hasNext()` and `next()` methods.
- **Find Median without sorting array in constant space and linear Time**: Use Quick Select. Google love this algorithm. This is initially put as design question, so be aware of the return type in the `findMedian()` method (should be float).
- **Validate Equations**: Given equations like `a < b`, `b = c`, `d > c`, validate for contradictions (e.g., `c < a`). Use Kahn's algorithm of BFS for cycle detection in directed graphs.
- **Favorite Cities**: Given directed connections (weighted edges) between cities and names of favorite cities, find the nearest favorite city from a specific city using Dijkstra's algorithm implemented with BFS.
- **Friend Requests**: You are given a list of friend requests, each represented as `a, b`, meaning `a` and `b` are now friends. After processing all these requests, you need to determine the number of individual groups of people. This can be solved using the Disjoint Set data structure.

#### System Design
Design an API to determine whether to show an ad based on its AdId. Each ad has a total budget that decreases each time it is shown. If the budget is non-zero, the ad can be shown. The system must be highly available and scalable with low latency in decision-making. Consistency in budget updates can have some error margin.

#### Preparation Tips
1. Brush up on all common algorithms and data structures.
   - **Sorting**: Learn Merge Sort as it is also a common solution for problems like Count of smaller numbers after self.
   - **Binary** Tree: Implement a class with insert, delete, find, and various traversals (inorder, postorder, preorder, left view, right view, top view, bottom view, BFS).
   - **Binary** Search Tree: Insert, delete, find, convert between tree and list.
   - **Union** Find: Implement union, and find.
   - **Segment** Tree: Implement get range answer, and update.
   - **Trie**: Implement insert and read.
   - **Heap**: Implement heapify, add, and remove.
   - **Matrix**: Solve shortest path problems like 01-matrix, rotten oranges, flood fill, rat in a maze etc.
   - **Graph**: Solve shortest path problems using Dijkstra.
2. Once you are clear with the basics, I recommend this cheatsheet URL1. It's not mine, but I'm thankful I randomly found it in a LinkedIn post. It covers a wide variety of questions and specific topic tabs. Practice a few questions on important topics like Tree, BST, Graph, Shortest Path, Trie, BFS, Disjoint Set, Matrix, etc.
3. Remove topic tags from questions and shuffle them to practice identifying the appropriate data structure during an interview.
4. I always used DFS to detect cycles in directed graphs and failed many interviews because of it. A turning point came when I accidentally learned about Kahn's algorithm while looking at the solution for Course Schedule II URL2.
Since then, I've solved every graph, tree, and matrix problem using BFS, and it has been the most efficient approach. While some problems may require DFS, those are rarely asked in interviews. Focus on mastering BFS; it will elevate your problem-solving skills to the next level. Trust me, you will rarely need to use DFS in interviews!
5. Only practice medium and hard questions. Forget about easy ones. Even if medium questions feel challenging, you will only improve by attempting them. Easy questions are never asked in interviews.
6. Google interview questions are unique and can't be found elsewhere. There is a fixed pool of confidential questions created each year from which the interview panel selects, including follow-up questions.
7. If you've completed all your preparations and need more questions to practice, I strongly suggest visiting URL3. In the problems section, filter for hard and medium problems with the company tag "Google." The problems there are unique, challenging, and can also be found on LeetCode.
8. Systems Design:
- First, understand the intuition behind the CAP theorem. You should grasp the meanings of consistency and availability. For example, I always imagine a bank with one ATM. If you prioritize consistency, the ATM will only be available when the bank’s server is online, ensuring that the balance is always correct. Conversely, if you prioritize availability, the ATM will process withdrawals even if the server is down, allowing money to be withdrawn without verifying the balance. This is why banking systems are designed to prioritize consistency.
- Start with learning about unique ID generators. Although this has become a cliché and isn’t asked as often, it’s still valuable to understand every variation, including the Twitter Snowflake technique, to see how system design concepts evolve.
- Next, learn how to implement a distributed key-value database for both consistent and available use cases.
- Study how to implement distributed locking.
  - While common problems like Ola, YouTube, and TinyURL are frequently discussed online, interviews present new challenges. For example, designing a Top-K leaderboard. Recommended resources:
  - **System Design Interview – An Insider's Guide by Alex Xu Volume 1:** Good for understanding basic concepts and classic problems.
  - **System Design Interview – An Insider's Guide by Alex Xu Volume 2**: Covers a new set of less common problems.
  - YouTube Channel **@jordanhasnolife5163** : Offers videos on unique and less common system design problems. The explanations may be fast, but they are useful for understanding diverse scenarios.
  - It's useful to understand key MSA patterns like API Gateway, BFF, Circuit Breaker, Database per Service, Event-Driven Architecture, and Saga. Focus on the Saga pattern (using a queue) to handle rollbacks and avoiding race conditions in multi-service transactions like order placement, ensuring consistency and atomicity—either all changes go through or none.
  - YouTube Channel **@interviewpen**: It covers many questions that SDEs often wonder about but rarely have time to research, such as OAuth and SQL query optimization. While these topics might not provide immediate help for an upcoming interview, they will certainly benefit you in the long run.
- Don’t overlook Low-Level Design (LLD). Although High-Level Design (HLD) questions are popular, LLD is just as important for interviews. You can find relevant questions on LeetCode, such as external merge sort, elevator algorithms, parking lots, or Twitter feeds.
- Einstein said, “If you can’t explain it simply, you don’t understand it well enough.” In HLD, we often add unnecessary components for minor objectives. For example, a queue, CDN or Elasticsearch might not be needed, but we tend to include them to impress the interviewer. However, seasoned system architects see this as over-complicating the system, which increases operational costs.
- Avoid overcommitting. Don’t add features or requirements that weren’t explicitly asked for by the interviewer, as it can make your design overly complex and difficult to implement.

#### [[2024] Google Interview Questions Compilation](https://leetcode.com/discuss/post/6185127/2024-google-interview-questions-compilat-mjrf/)