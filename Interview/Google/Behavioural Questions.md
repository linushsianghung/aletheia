# Interview Questions

## Reference:
- Google
  - [Customer and Partner Solutions Engineer interview prep guide](https://www.google.com/about/careers/applications/candidate-prep/cpse)
  - [What is Your Career Identity?](https://www.youtube.com/watch?v=_xbT4qMrot4)
- Dan Lok
  - [Tell Me About Yourself - A Good Answer To This Interview Question](https://www.youtube.com/watch?v=5v-wyR5emRw)
  - [Interview Question: “What Are Your Weaknesses?” And You Say, “...”](https://www.youtube.com/watch?v=VwzFWmNX7GI)
  - [Interview Question: “Why Did You Leave Your Last Job?”](https://www.youtube.com/watch?v=NyAZmZZaNPs)
  - [“What Kind Of Work Environment Do You Like?” Best Answer To This Interview Question](https://www.youtube.com/watch?v=yhCYdp5KxMs)
  - ["Do You Work Well Under Pressure?"](https://www.youtube.com/watch?v=lfY-Glap8DI)
  - [How To Be Confident In Interviews](https://www.youtube.com/watch?v=wBJ0MUkA1cA)
- Others
  - [How To Answer Behavioral Interview Questions](https://www.youtube.com/watch?v=sq3pyauZRhI)

- Philosophy
  - Simplicity is paramount
  - Complex system is difficult to be rationalised and would be failed in complex way
  - If you can't explain it simply, you don't understand it well enough.
  - Asking the right person
  - Rephrase in 2 ways, one is rephrasing in positive sentence and the other is in negative sentence.



----------
## Behavioural Questions
A behavioral question looks at how candidates have handled a specific challenge in the past to assess if they’ll be a good match for the role.

Usually start with phrases such as:
- Tell me about a time when ...
- Give me an example of ...
- Describe a decision you made ...

- **CARL** approach:
  - **Context**
  - **Action**
  - **Result**
  - **Learning**

- **STAR** approach:
  - **Situation**: Describe the situation you were in or the task you needed to accomplish
  - **Task**: Explain the goal you were working toward
  - **Activity**: Detail the specific steps you took and the role(s) you played
  - **Result**: Describe your accomplishments and the overall outcome

> Specific for Customer Solutions Engineer (CSE):
> 
> The focus of this part of the interview will be on your stakeholder management and consulting skills. Your interviewer will want to see your consulting skills and understand how you work with customers and partners through both hypothetical and behavioral questions: how you uncover and assess stakeholder needs through conversations in order to make recommendations. They’ll want to see that you effectively engage, collaborate with, and influence stakeholders by forging collaborative relationships. Show how you manage expectations and agreed objectives while remaining engaged and authentic, regardless of personality, technical ability, or background.

### Questions:
#### Tell me how you handled a difficult situation?
- Shows **ownership**: You didn’t wait; you created a system
- Shows **strategy**: Targeted learning, not random effort
- Shows **impact**: Faster ramp-up + team-wide benefit
- Shows **scalability**: Your solution helped others too

> One of the most challenging situations I faced was when I first joined a healthcare project with very limited domain knowledge. It’s a highly specialized industry, ~~and many of my colleagues had over 20 years of experience,~~ so I found it difficult to handle new requirements initially. To address this, I took a structured and proactive approach to try to ramp up quickly. 
> 
> First, ~~instead of relying on others passively,~~ I studied internal documentation and existing systems to build a baseline understanding. Then, I scheduled focused discussions with experienced colleagues — rather than just asking broad questions, I came prepared with well-researched questions. This made our conversations much more efficient and valuable. At the same time, I documented everything I learned — especially from a beginner’s perspective, including key concepts and common pitfalls.
> 
> As a result, I was able to significantly shorten my ramp-up time and start contributing to development with more confidence. My documentation was also adopted by the team and helped to improve onboarding efficiency for new engineers.
> 
> This experience taught me how to quickly adapt in an unfamiliar domain and turn a difficult situation into a scalable improvement for the team.

#### Tell me about a time ...... conflict with a teammate / your manager
The points are:
- `Be the last one to speak`
- Candidate can listen to and accept other people's point of views even if you have already form his/her own opinion.
  - So the Core part of the answer should include asking the other party to expand on the reason as to why they believe they are correct. Then, rephrase their words and repeat it back to them to make sure I fully understand their opinion.
- Candidate is able to put the needs of business above his/her own when a conflict does arise in the workplace.
  - Pinpoint the exact moment where all parties involves are at the most risk of losing face. Then show how I de-escalated the situation by providing a platform for everyone to step down on.
  - Acknowledging the other direction or idea has more merit than my own and potentially be a better solution
  - Show how the broader team or business benefited as a result of your actions

> In one project, a teammate and I disagreed on how to implement a new feature in a backend service. He preferred reusing the existing library to minimize the risk and maintain stability, while I was concerned it wouldn’t handle the new requirements well, especially around performance and extensibility.
> 
> ~~Instead of pushing back immediately,~~ I first tried to understand his perspective. His main concern was reducing the production risk, which I totally agreed with. So I reframed my suggestion around this shared goal — I proposed using a newer library, but with safeguards. More specifically, I suggested we validate the approach through performance testing and limit the rollout behind a feature toggle. I also walked through a concrete example showing how the new library would simplify the codebase and improve performance and maintainability. My purpose is to shift the discussion from personal opinions to objective trade-offs.
> 
> Then, we both agreed to try the new approach with validation which showed no regression issue. We moved forward with it, and it runs well in production while also reducing code complexity.
> 
> For me, that experience reinforced the importance of aligning on shared goals when dealing with conflict, discussing based on trade-off to address concerns directly, and using data to guide decisions rather than arguing on who’s right.

[Follow Up]
- Java 8 vs Java 17
- Even using new library is still low risk, that's why I suggestion to run **Performance Testing** to validate.
- It's more about push ourselves to out of the **Comfort Zone** a little bit.

#### Tell me about a time how you worked under pressure?
- Rather than reply YES or NO, tell a STORY!
- Hit Key Points:
  - How you have used your skills to handle the situation?
  - How you are taking responsibilities and not blaming?
  - What are the result?

[Google Version]
> One situation was during a product release I had to take over responsibility to deliver features under a very tight deadline. About a week before release meeting, the engineer responsible for that became unavailable because of COVID, ~~and his development was a release blocker~~. I was asked to step in, even though I hadn’t worked on that area before.
>
> For me, the pressure came from the tight deadline and uncertainties, so I focused on how I can reduce uncertainties and execute works efficiently. First, I reviewed the requirement scope and prioritized the critical tasks that would impact the release. Then, I parallelized the work by delegating well-defined subtasks to junior engineers, while I handled the more complex scenarios myself. I also set up daily check-ins with my Tech Lead to make sure we stayed aligned and could adjust quickly if required.
>
> In the end, we completed all the tasks on time, the release went out as scheduled, and we didn’t see any critical issues afterward.
> 
> That experience reinforced for me that under pressure, it’s really about structured planning, leveraging resource smartly, and clear communication.

[Long Version]
> Yes! One example was during a product release when I had to take over the responsibility on very short notice and a very tight timeline. About a week before release, the engineer responsible for a production requirement unavailable due to COVID. This would become a release blocker, and the release would have been delayed. I was asked to take ownership, even though I had 0 preparation for that requirement, and I had only one week to complete the requirement tasks.
>
> For me, pressure comes from the tight deadline and uncertainties. To handle the pressure, I focused on reducing uncertainty and maximizing execution speed: 
> First, I quickly broke down the requirement scope into smaller tasks and identified dependencies. I prioritized critical tasks that would impact the release decision.
> Second, I parallelized execution by delegating well-defined, lower-risk tasks, like preparing verification testing, to junior engineers, while I focused on the more complex scenarios.
> Third, I set up a daily sync with my Tech Lead to validate progress and adjust priorities quickly if needed.
> Since I was new to the area, I also proactively consulted domain experts to avoid misinterpretation of requirements.
>
> We completed all required tasks on schedule, and the product was released without delay. Importantly, we did not encounter any critical issues post-release, which validated the quality of the implementation and testing effort. This experience reinforced for me that under the pressure, success comes from structured planning, smart delegation, and tight communication, not just working harder.

#### Tell me about a time you be asked to go above and beyond.
This question is not just about working hard or doing overtime. Interviewers are trying to evaluate:
- **Ownership Mindset**: Do you treat work like “just your job,” or do you step up when something important is at stake? 
- **Customer / business Focus**: Did you go the extra mile because it mattered (e.g., user impact, revenue, deadlines), or just because someone told you to?
- **Judgment & Prioritization**: Did you choose wisely when to go above and beyond (vs. burning out on everything)?
- **Influence & Initiative**: Did you passively accept extra work, or did you actively drive a better outcome?
In short: **They want to see if you act like an owner, not just an executor.**

[Google Version]
> When I first joined Innova, I was asked to step in as a Scrum Master for two sprints to help the team meet a critical deadline. I didn’t have prior experience in that role, and the team was newly formed, so there was a real risk around the delivery.
>
> Instead of just running ceremonies, I focused on what would actually ensure success. I organized regular meetings with architects to validate our implementation early to avoid rework. I also proactively aligned with another team working on the same feature to reduce dependencies and delays. On top of that, I took a hands-on approach to quickly identify and remove any potential blockers.
> 
> As a result, we not only delivered on time but also improved our team’s collaboration, especially benefit for a newly formed team.
>
> This experience taught me that the main point of going above and beyond is willing to take ownership and step into gaps to enable the team to succeed.

[Long Version]
> When I first joined Innova, I was unexpectedly asked to take on the Scrum Master role for two sprints to help the team meet a critical delivery deadline. At that time, I didn’t have prior experience as a Scrum Master, and the team itself was newly formed, so there was a real risk around alignment and execution. Instead of just facilitating ceremonies, I focused on what would actually unblock the team and ensure delivery.
> 
> First, I noticed there was ambiguity in technical direction, so I proactively organized regular syncs with architects to validate our design decisions early. This helped us avoid rework and kept implementation aligned with long-term architecture. Second, I saw dependencies with another team working on the same feature, so I stepped in to establish direct communication channels and align priorities. This significantly reduced cross-team friction and delays. Third, I took a hands-on approach to identify and remove blockers—whether that meant helping debug issues, clarifying requirements, or reprioritizing tasks to keep progress moving.
>
> As a result, we not only delivered on time, but also improved team velocity and established a strong collaboration model early on. For a team that had only been together for three months, this helped us integrate much faster into the organization.
>
> This experience taught me that going above and beyond isn’t about doing more tasks—it’s about stepping into gaps, identifying what truly drives success, and proactively enabling the team to deliver.



### Questions (Stakeholder Focus):
#### Tell me about a time you had to explain a complex technical concept to non-technical stakeholders. / Tell me about a time stakeholders didn’t understand your solution—what did you do?

[Google Version]
> One example was when I needed to explain our OTA update system, based on Mender, to non-technical stakeholders including product managers. The challenge was that the system architecture was quite complex — built on microservices, and Kubernetes — and if I explained that directly, it could easily overwhelm or confuse non-technical audiences. However, it was important that they understood the value and workflow clearly to present it to our clients. 
> 
> So I approached this from three aspects:
> First, I chose a proper abstraction level to simplify the architecture into a high-level diagram, focus only on key components and the flow of update, instead of internal details.
> Second, I used an analogy, for example I compared the OTA system to a “Pneumatic Tube System,” where carriers are transferred via tube rather than delivered by human, which helped to make the concept more intuitive.
> Third, I prepared a live demo, so they could see the update process end-to-end, which made the system more tangible.
>
> As a result, the PMs were able to present the system to clients confidently. It also helped us to align expectations early, which made later discussions much smoother.
>
> This experience let me understand that effective stakeholder communication is not about explaining everything — it’s about translating complexity into critical points and using the right level of abstraction for the audience.

[Short Version]
> In my current role, I had to explain a complex OTA update system to non-technical stakeholders like PMs who would present it to clients to promote our product. The challenge was that the system used microservices and Kubernetes, which could easily overwhelm a non-technical audience.
>
> So I focused on simplifying the message: I created a high-level architecture diagram, used analogies like a pneumatic tube system to explain the workflow, and prepared a live demo to make it tangible. I also checked for understanding throughout.
>
> As a result, the PMs were able to confidently present to clients with fewer follow-up questions.
>
> This taught me that effective communication is about tailoring the level of abstraction to your audience, not just explaining the technology.

#### Tell me about a time requirements were unclear or constantly changing?
[Google Version]
> One example was when we were asked to handle a technical issue, the requirements were unclear because the problem was highly technical and the product manager didn’t have enough context to define it precisely.
>
> To move forward, I first identified the right stakeholder who had direct insight into the issue to clarify the requirement. Before meeting, I reviewed relevant documentation and prepared questions to make the discussion efficient. During the conversation, I clarified the actual problems, key use cases, and success criteria. After the meeting, I built a prototype to validate my understanding and shared it early to get feedback. I also documented the assumptions and investigation process so everyone could stay aligned and reuse the knowledge later.
>
> As a result, we were able to quickly converge on the correct requirements to avoid rework, and deliver a solution that addressed the real problem.
>
> This experience reinforced for me that when requirements are unclear, it’s important to drive alignment proactively, validate assumptions early, and iterate quickly.

[Long Version]
> One example was when we had a technical issue that needed improvement, but the requirements were quite unclear because it was highly technical, and the product manager didn’t have enough context to describe it precisely. My goal was to clarify the actual problem and deliver a solution without wasting time on incorrect assumptions.
>
> To do that, I first identified the right stakeholders. I asked the PM who had the most context and was directed to a Field Customer Engineer who had direct exposure to the issue. Before meeting them, I reviewed relevant documentation and prepared specific questions to make the discussion more effective. During the meeting, I focused on understanding the real-world scenario, the expected outcome, and the success criteria. After that, I built a simple prototype to validate my understanding and shared it with the Field Customer Engineer for quick feedback. I also documented the discussions, assumptions, and investigation process so the team could reuse this knowledge in the future.
>
> As a result, we were able to clarify the requirements quickly, avoid unnecessary rework, and deliver a solution that matched the actual customer need.
>
> This experience reinforced for me that when requirements are unclear, it’s critical to find the right stakeholders, validate assumptions early through prototyping, and document learnings for scalability.”

#### Tell me about a time when different stakeholders had conflicting requirements. / Describe how you handled competing priorities from product, business, and engineering.
> While I haven’t encountered this directly, here’s how I would approach it...
> [How would you handle conflicting priorities from stakeholders?](#how-would-you-handle-conflicting-priorities-from-stakeholders)



----------
## Hypothetical/Situational Questions
Hypothetical questions evaluate how candidates would handle a challenge they may not have encountered yet. Questions often begin with “Imagine that ..."

- **CFAS** approach:
  - **Clarify**
  - **Framework**
  - **Assumption**
  - **Solution**

- Some common clarifying questions:
  - **Scope**: 
    - it's only for an individual, for a group or across the entire organisation?
    - We are talking about a small, medium, or large company?
  - **Timeline**: It has been always this way or triggered by some special event (COVID)?
  - **Perspective**: It's only my personal opinion or do other people feel this way as well?
  - **Impact**: There is already any impact related by this?

### Questions:
#### What would you do if requirements are unclear?/ How To Deal with Ambiguous Requirements?
The response is supposed to shows:
- Structured thinking
- Ownership
- Communication
- Iterative mindset

> When requirements are unclear, what I would do first is to clarify the problem rather than jumping into implementation.
> 
> But how to do that? The first key point is **asking the right person**. I start by identifying the key stakeholder and asking specific questions to understand the goal and criteria. ~~During the communication I would rephrase the response in both positive and negative ways, to understand the problem from different angles.~~
>
> If there are still gaps, I try to propose a few assumptions — this helps us move the conversation forward instead of waiting for 100% clarity. I also like to break the problem into smaller deliverables or prototypes, so we can get early feedback and reduce the risk of building the wrong thing.
> 
> Throughout the process, I document all the discussion and assumptions, and keep stakeholders aligned with it, so even if requirements evolve later, the team has a shared understanding and can adapt quickly.

#### How Do You Handle Conflict With Teammate? / What would you do if your manager disagrees with your technical decision?
Both questions are fundamentally about:
- Communication skills
- Handling disagreement professionally
- Influence & persuasion
- Emotional intelligence (ego vs outcome)
- Collaboration under tension

> When I handle conflict with coworkers, I try to approach it in a structured and collaborative way.
> 
> First, I make sure I fully understand their perspective before reacting. In many cases, conflicts come from different priorities, so I focus on clarifying the underlying goals. Second, I reframe the discussion around that shared goal, like what’s best for the user or the system, rather than whose approach is correct. Then, I bring concrete examples to make the discussion more objective — such as trade-offs or performance metrics.
>
> If the decision is still unclear, I’m comfortable involving a tech lead or manager to get additional input. Finally, once a decision is made, I fully support it and focus on execution, even if it wasn’t my initial preference.
> 
> Overall, my goal is to resolve conflicts in a way that strengthens collaboration and leads to better technical or business outcomes.

----------
#### How would you handle conflicting priorities from stakeholders?**
This question is less about the exact situation and more about how you think and operate under pressure. They want to evaluate:
- **Prioritization skills**: Can you distinguish urgent vs important, and business impact vs effort?
- **Stakeholder management**: Do you push back professionally, align people, and avoid just saying “yes” to everyone?
- **Communication clarity**: Can you explain trade-offs in a way non-technical stakeholders understand?
- **Decision-making framework**: Do you have a structured approach, or do you just react? 
- **Ownership & accountability**: Do you take responsibility for driving alignment instead of escalating immediately?

> When I receive conflicting priorities from stakeholders, I first focus on making the situation transparent and structured rather than trying to solve it in isolation. In my experience, what seems like a conflict is sometimes just a lack of shared context.
> 
> I start by clarifying each request — understanding the business goal, urgency, and any constraints. Next, I evaluate the effort and impact of each request. For example, I look at factors like customer impact, system risk, deadlines, and dependencies. This helps me frame the discussion objectively ~~rather than emotionally~~.
>
> If the conflict remains, I bring the stakeholders together and present a few clear options with trade-offs. For instance, I might coordinate that to say: “We can prioritize Feature A for urgent customer impact, but Feature B will be partially delivered MVP for this release”. I’ve found that stakeholders are much more aligned when they see the trade-offs explicitly.
>
> Finally, once a decision is made, I make sure expectations are clearly communicated and documented, so the team can execute with focus and avoid further misalignment. 
> 
> Overall, my goal is to act as a bridge — helping stakeholders make informed decisions while keeping delivery realistic and predictable.

----------
## General or Personality-Based Questions

### Questions:
#### Tell me about yourself.
- Get to the core of why this job is open and understand what the company needs in the person
- Job Description: 
  - The problems the company needs to solve
  - Responsibilities
- What value could you bring
- Template:
  - **Present** Who I am today, what I do, and what I am great at:
  - **Past**: 2 or 3 sentences for how I got there
  - **Future**: From what I know, you and your company are looking for someone who can help with X, Y, Z, and that's exactly why I excited about this role

> I’m a software engineer with about 10 years of experience, primarily focused on backend systems ~~using Java, Golang~~ and cloud technologies. ~~In my current role at Innova Solutions,~~ I’ve been working with Optum for 5 years, where I’ve contributed many services in the healthcare domain, including delivering features and facilitating systems design.
> 
> Over time, I found myself often dealing with problems that require both technical depth and strong collaboration. For example, during an internal system development, I worked closely with architect and product manager to communicate the design and requirement frequently. This experience really improve my skill in translating language between technical terms and client requirements for different stakeholders.
> 
> Now I am looking for a role where I can specialise in Software Architecture based on Client's Requirements. What excites me about this opportunity is the chance to work at scale, to help customers to solve real-world problems while leveraging Google’s technologies. I believe my experience in systems design, combined with the ability to communicate across technical and non-technical audiences, would allow me to contribute effectively in this Customer Solutions Engineer role.

#### What are your strengths?
[Google Version]
> I’d say my biggest strengths are structured problem-solving, ownership, and the ability to quickly ramp up in complex domains.
> 
> First, on problem-solving — I am good at breaking down ambiguous or complex problems into smaller, actionable parts. For example, when I was working with unclear or changing requirements, I didn’t just wait for clarification. I proactively aligned with stakeholders, asked targeted questions, and translated ambiguity into concrete technical tasks. That helped the team move forward without unnecessary delays.
> 
> Second, I take strong ownership of outcomes. In situations where there are disagreements on technical decision, I focused on understanding different perspectives, validating trade-offs, and driving toward a solution that balances technical and business factors. I care not just about writing code, but also delivering the right solution for the business.
> 
> Finally, I’m very good at ramping up in new domains. When I first entered the healthcare industry, I had no prior experience. But I created a structured learning approach and quickly got to a point where I could confidently contribute to complex features. 
> 
> Overall, I’d say I bring a balance of analytical thinking, collaboration, and proactive execution.

[Long Version]
> I’d say my key strengths are structured thinking in ambiguous environments, strong ownership, and the ability to ramp up quickly in complex domains—and they tend to reinforce each other in how I work.
> 
> First, on structured thinking — I’m very comfortable dealing with ambiguity and turning it into clear execution. In one project, requirements were not well-defined and kept evolving. Instead of waiting, I proactively worked with stakeholders to clarify goals, asked targeted questions to uncover constraints, and broke the problem into smaller, manageable pieces. That allowed the team to continue making progress while reducing rework, and we were able to deliver iteratively with better alignment.
> 
> Second, I take strong ownership of outcomes, especially when there are trade-offs involved. For example, I had a situation where a teammate and I disagreed on the technical approach — whether to reuse an existing solution or adopt a newer library. Rather than pushing my idea, I focused on understanding his concerns around stability, and we evaluated the trade-offs together, including scalability, maintainability, and delivery risk. We ended up making a balanced decision that met both immediate and long-term needs, and it strengthened our collaboration as well.
> 
> Finally, I’m very proactive in ramping up in new or complex domains. When I first entered the healthcare space, I didn’t have prior domain knowledge, which made it challenging at the beginning. So I created a structured learning plan — combining documentation, hands-on exploration, and learning from experienced teammates — and within a relatively short time, I was able to contribute to core features and discussions with confidence.
> 
> Overall, I’d say I bring a combination of analytical thinking, ownership, and continuous learning, which allows me to navigate ambiguity, collaborate effectively, and consistently deliver meaningful impact.”

#### What are your weaknesses?
- Don't make that about the personality but about the skill set which can be worked on.
- Issue Of personality usually can't be fixed but skill set can be improved.
- Use CAR (Context-Action-Result) Formula

> One area I’ve been improving is that I don’t naturally recall historical incidents as quickly as some colleagues, especially when discussing the issues from months or years ago.
>
> To address this, I built a structured documentation template to record every issue I've worked on. I document the context, investigation process, root cause, and solution, along with some lessons I learned. Over time, I found that these documentation wasn’t just helping me but also valuable for the team. I started to organise these into shared knowledge bases so that others could reference them easily.
>
> As a result, while I still cannot purely rely on my memory ~~in discussions~~, I can retrieve accurate information and provide well-structured insights quickly. More importantly, these documentation helps team to learn from past incidents and avoid repeated issues.
>
> Currently, I’m working on refining the search mechanism and summarisation so the information can be accessed more easily.

#### Why did you leave your last job?
- Never ever put down the Ex-Boss & Ex-Company
- Don't share stuff that gets very personal and emotional

> I’ve really enjoyed the past five years at Innova Solutions and working with Optum. During that time, I’ve contributed many services in healthcare domain and gained deep experience in that. I also worked with a very collaborative team, where knowledge sharing was strongly encouraged, which helped me grow a lot both technically and professionally.
>
> Over time, I realized I’m energized when I’m closer to real-world problem — especially working closely with client side and translate their needs into technology solutions. But in my current role, those opportunities are a bit limited.
>
> So I feel this is the time to take the next step into a role like Customer Solutions Engineer, where I can combine my backend and system design experience with more customer-facing impact, ~~and help to drive solutions more end-to-end~~.

#### What kind of Work environment do you like?
-  Template
  - **Logline**
    - Talking about who you are as professional
    - Talking about what makes you special and good for the job you apply
  - **Past**: 2 or 3 sentences in prior roles that show what you've done in the past and how it connects to the job you apply
  - **Future**: 1 sentence you are going to explain why this job is the perfect next step for you in your career journey

> I do my best work in environments that are both collaborative and ownership-driven. I enjoy working with cross-functional teams where ideas can be openly discussed, but I am also appreciate having clear ownership so I can take responsibility for delivering results.
> 
> In my experience, I often worked on projects where I needed to collaborate closely with architect and product managers, while also independently driving results. That environment kept me stay productive while still aligning with the broader team goals.
> 
> From what I understand, Google emphasizes both collaboration and individual ownership at scale, which is exactly the kind of environment where I believe I can contribute effectively and continue growing.

----------
#### Why do you want to work here?
Reference: [Why Google](https://www.youtube.com/watch?v=hgFKFu5vNug)
- **Research**: Connect a concept that I really tied to or connected with
- **Purpose**: Consider **Long-Term** purpose
- **Job Description**: `Skill Alignment` & `Product`
- **Structure**:
  - 1 minute long
  - Rule of 3: Just focus on 3 items
  - No praising
  - Set up
  - Practice

> “Why Google?” for me basically comes down to two things: culture and impact.
> 
> First, I really resonate with Google’s culture of continuous learning and innovation. Early in my career, a colleague from IBM shared the idea of investing part of my time to learn something beyond my responsibilities. I’ve followed that ever since and constantly exploring new technologies, pursuing certifications, and continuous improving how I solve problems. When I learned about Google’s 20% culture, it strongly aligned with how I already approach my work.
> 
> Second, I’m excited about the scale of impact. Google operates at a level where solving technical problems will directly influence millions of users and businesses. That’s really meaningful to me as someone who enjoys bridging technology and real customer needs.
> 
> Finally, for the Customer Solutions Engineer role, I believe my experience of working closely with non-technical stakeholder, translating their needs into technology solution, and experience in database technologies make me a strong fit for this position. ~~I’m particularly motivated by roles where I can combine deep technical work with customer-facing problem-solving, which is exactly what this position offers.~~

----------
#### What project are you most proud of?
What project are you most proud of?” is not about the project itself. It’s a proxy to evaluate:
- **Ownership & initiative**: Did you just execute, or did you identify and drive a problem?
- **Impact**: Did it actually improve something measurable? (time, cost, reliability, dev productivity)
- **Engineering judgment**: Did you choose the right tools and make good trade-offs?
- **Sustainability**: Did your solution last and scale—or was it a one-off hack?
- **Communication**: Can you explain a complex project clearly and concisely?

> One project I’m most proud of was improving our local ~~development~~ environment setup when I first joined my team. At the time, onboarding was very slow — setting up the system could take hours or even days because we had many dependencies like microservices, databases, and Pub/Sub. The process originally relied on complex shell scripts which were hard to maintain and quite often broke as the system evolved.
> 
> I took an initiative to simplify and automate the entire process. I migrated the deployment logic into Docker Compose to make the architecture more transparent, and then built an Ansible-based automation pipeline to orchestrate the environment setup — including service startup, data initialization, and system configuration.
> 
> As a result, I reduced setup time from hours to just a single command, hugely improving onboarding speed and developer productivity. More importantly, because we used these scripts every day, they stayed up-to-date and became living documentation.
> 
> I’m particularly proud of this project because it solved a real pain point of the team and had a huge impact on developer experience and efficiency.

----------
#### Where do you see yourself in 5 years?
> In the next 1 to 2 years, my main focus is to ramp up quickly to build strong domain knowledge, especially in GCP, and start contributing to real customer problems. ~~I want to build strong domain knowledge and become someone the team can rely on for both technical depth and execution.~~ 
> 
> Around the 3 to 5 year mark, I hope myself taking on more ownership — not just implementing solutions, but helping design scalable architectures and working closely with customers to translate their needs into effective technical solutions. ~~Based on my experience improving complex systems and onboarding processes, I enjoy solving messy, real-world problems, so I’d like to expand that impact at a larger scale.~~
> 
> Longer term, I’d like to grow into a role where I can mentor others and influence technical direction, while still staying hands-on. Ideally, I become a trusted advisor both internally and for customers.


----------
## Questions to ask
- **What's your favourite part about working in this position?**
- **What's the biggest challenges or most difficult part of this role?**
- **How will the performance be evaluated in this position?**
- **Could you please tell me what the typical day looks like in this position?**

