# AI Usage Log – Lorekeeper

## 1. Introduction

This document explains how AI tools were used during the development of Lorekeeper.

AI tools were used as development assistants to support research, problem-solving, code development and documentation.

The purpose of this document is to explain:

- Which AI tools were used.
- What they were used for.
- Examples of useful and unhelpful suggestions.
- How AI-generated suggestions and code were verified.
- What I learned from using AI during development.

AI was used as an aid to development, not as a replacement for understanding the application or its underlying
technologies.

---

## 2. AI Tools Used

List the AI tools used during the challenge and briefly explain their purpose.

| AI Tool | Purpose                                        |
|---------|------------------------------------------------|
| ChatGPT | Research assistance, architectural discussions |

---

## 3. How I Used AI During Development

### 3.1 Planning and Architecture

**How I used AI:**

I used ChatGPT to brainstorm possible ideas for the project before any development had begun. Initially these
suggestions were fairly generic DevOps/AI/Automation projects. I knew I wanted something that was relatively unique to
my personal interests and so prompted for some Dungeons & Dragons themes, which led to the idea of Lorekeeper.
Generally, I am firmly against the use of AI within my D&D campaigns, so I was hesitant to implement any of the
suggestions from ChatGPT. However, my gripes with AI in D&D is that it removes the creative aspect of the game which is
fundamental to its premise. The idea of a session notes lookup tool does not encroach upon any of those creative
sections of the game and so it seemed like an appropriate fit.

I took this concept and then began researching some of the proposed architecture that the model had suggested. This
independent research is what solidified this as the project I would choose as it included technologies and concepts
potentially relevant to the role that I had not looked at before. ChatGPT recommended a WebLLM and WebGPU approach

**My Evaluation**

As mentioned, the initial project ideas were very generic (e.g, Server/Docker monitoring system, AI automation
assistant, Home Network troubleshooter). These initially were not particularly helpful, especially given the context
that this project was for a job interview: a scenario in which I want to stand out from the rest of the applicants.

After I had refined the requirements of the project in my prompts I was able to get some fairly reasonable and workable
solutions.

### 3.2 Research and Learning

**How I used AI:**

Personally, when it comes to research and learning, I prefer to read directly from the source material rather than
relying on AI summarisations. Throughout my academic courses, I have found that AI models tend to mischaracterise the
tone and meaning of more formal pieces of writing like research papers. Something I have extensive experience in is
limiting conclusions that can be drawn from results. My master's thesis relied heavily on this idea; I could have
described the results in a very positive manner but this would have been disingenuous without also considering the
limitations in my research methods and the reliability of them. My experience with AI is that is fails to grasp this
nuance and tends to describe in absolutes, which from a research perspective can be very detrimental.

**My evaluation:**

I did not use AI for the Research and Learning sections of the project.

### 3.3 Code Development

**How I used AI:**

[Explain whether AI helped generate, improve or explain sections of code.]

**My evaluation:**

[Explain how you reviewed generated code, tested it and ensured you understood how it worked.]

### 3.4 Debugging and Problem Solving

**How I used AI:**

[Explain how AI helped investigate errors or unexpected behaviour.]

**My evaluation:**

[Explain whether the suggestions solved the problems and how you verified the solutions.]

### 3.5 Documentation and Testing

**How I used AI:**

ChatGPT was used to develop templates and content for documentation of the project like updating the README.md
throughout development to update the changes in the current branch before creating a PR. Templates for RESEARCH.md,
AI.MD and the GitHub PR template were also created by ChatGPT.

**My evaluation:**

All documentation edits were reviewed manually and edited, if needed, to ensure a consistency and accuracy throughout.
AI-Generated documentation tends to be filled with unnecessary examples and definitions and is generally very explicit
in its writing. This is not inherently a bad thing but a lot of this content needed to be removed as it made the writing
overly verbose.

---

## 4. Incorrect or Unhelpful AI Suggestions

This section documents instances where AI suggestions were incorrect, unsuitable or didn't work as expected.

| Suggestion            | Why it was unsuitable or incorrect | How I identified the problem | What I did instead     |
|-----------------------|------------------------------------|------------------------------|------------------------|
| [Describe suggestion] | [Explain issue]                    | [Explain verification]       | [Describe alternative] |

---

## 5. How I Verified AI-Generated Code and Suggestions

I did not assume that AI-generated code or technical explanations were automatically correct.

The following approaches were used to verify suggestions:

- **Official documentation:** Checking technical information against the documentation for relevant libraries and
  technologies.
- **Code review:** Reading generated code and ensuring I understood its purpose and behaviour.
- **Automated testing:** Running unit tests to verify expected functionality.
- **Manual testing:** Testing application features through the browser using different inputs and scenarios.
- **Error investigation:** Using browser developer tools and terminal output to investigate problems.
- **Independent research:** Comparing AI suggestions with alternative sources where appropriate.

---

## 6. Limitations and Considerations

[Describe limitations and considerations]

---

## 7. Final Reflection

**What was AI most useful for?**

[Describe where AI provided the greatest benefit during development.]

**Where was AI least useful?**

[Describe situations where independent research or experimentation was more effective.]

**What did I learn about working alongside AI?**

[Reflect on how AI affected your development process and understanding of the technologies involved.]

**How did I ensure I understood the final application?**

[Explain how you reviewed, tested and learned the code rather than simply accepting generated solutions.]

---

*AI usage log maintained throughout the development of Lorekeeper.*

*Project: Enterprise Automation Services Technical Challenge – November 2026.*