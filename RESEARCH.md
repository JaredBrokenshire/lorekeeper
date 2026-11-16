# Research Log – Lorekeeper

## 1. Introduction

This document records the research and learning undertaken during the development of Lorekeeper as part of the
Enterprise Automation Services Technical Challenge.

Lorekeeper is a web-based AI assistant designed to help Dungeons & Dragons players search their campaign notes and
retrieve information about characters, locations, events and quests.

The purpose of this document is to record:

- New technologies and concepts I needed to understand.
- Research sources I used and how I assessed them.
- Technical problems and approaches I encountered.
- Solutions I implemented and why.
- What I learned throughout development.

Zotero was used for reference management throughout the development process.

---

## 2. Research Log

### Research Topic: Go PDF Extraction

**What did I need to learn?**

I needed to find a way to extract text from PDF files, including the text from each page and its corresponding page
number.

**Why did I need to learn this?**

This is part of the session-note ingestion process. The extracted text is later divided into overlapping chunks so
that Lorekeeper can search campaign notes and identify relevant information.

**Where did I research this?**

| Source         | Link                                              | Why I used it                                                                                                                                                                 |
|----------------|---------------------------------------------------|-------------------------------------------------------------------------------------------------------------------------------------------------------------------------------|
| pdftotext-go   | https://pkg.go.dev/github.com/heussd/pdftotext-go | This was the first option I found. Its documentation led me to a dependency on Poppler utilities, but I could not work out how to install and configure the required package. |
| ledongthuc/pdf | https://github.com/ledongthuc/pdf                 | I found this Go package while looking for an alternative that could extract text from PDF files without the setup difficulties I had encountered.                             |

**What did I discover?**

PDF text extraction can be handled by an existing Go package rather than requiring me to implement the extraction
process myself. I also learned that choosing a library involves considering its dependencies and how practical it is
to configure in the project environment, not just whether it offers the required functionality.

**What approaches did I consider?**

I initially investigated `pdftotext-go`. After encountering difficulties with its Poppler utilities dependency, I
looked for another Go PDF extraction package and found `ledongthuc/pdf`.

**What problems did I encounter?**

I could not understand the setup requirements for the Poppler utilities dependency used by the initial approach. I
was unable to get that option running, so continuing with it would have taken time away from the rest of the project.

**How did I resolve the problem?**

I continued researching alternatives and found the `ledongthuc/pdf` package. I used this package for PDF text
extraction in Lorekeeper.

**What did I decide and why?**

I chose `ledongthuc/pdf` because it provided an alternative that I could use without having to resolve the Poppler
utilities setup issue.

**What did I learn?**

I learned to consider the practical setup and dependencies of a library as part of evaluating it. When an approach
becomes blocked by an unfamiliar dependency, it can be more productive to investigate a suitable alternative rather
than spend too much time trying to force the original choice to work.

---

### Research Topic: Ollama and Docker Compose

**What did I need to learn?**

I needed to understand how to run Ollama in Docker as part of Lorekeeper's containerised application, and how the
API service could communicate with it.

**Why did I need to learn this?**

Lorekeeper uses a local language model. Running Ollama in a container allows it to be included alongside the other
application services in the Docker Compose setup.

**Where did I research this?**

| Source                                                             | Link                                                                                                                | Why I used it                                                                                        |
|--------------------------------------------------------------------|---------------------------------------------------------------------------------------------------------------------|------------------------------------------------------------------------------------------------------|
| Deploying Ollama with Docker Compose: A Simple Guide to Local LLMs | https://medium.com/@prasanta.mohanty/deploying-ollama-with-docker-compose-a-simple-guide-to-local-llms-610db2991581 | I used this guide to research how Ollama could be deployed with Docker Compose.                      |
| How to create an ollama model using docker-compose?                | https://stackoverflow.com/q/78438394                                                                                | I consulted this question for practical guidance on configuring an Ollama model with Docker Compose. |

**What did I discover?**

Ollama can be run as a service in the Docker Compose environment, alongside the database and API. The model name
can be supplied through an environment variable, which means it can be changed through configuration rather than
being fixed in the application.

**What approaches did I consider?**

I wanted the model configuration to be flexible, so I configured the model name through an environment variable.
I chose a CPU-compatible model because I could not assume that the recruitment team would have access to a GPU
when running the project.

**What problems did I encounter?**

I encountered configuration issues when trying to pull the intended model using the value supplied through
environment variables. I also needed to ensure that the API container waited for both the database and Ollama
containers to be up and running successfully before it started.

**How did I resolve the problem?**

I adjusted the Docker Compose and environment-variable configuration to include a one-time init contianer for the ollama 
container which pulled the appropriate model. This meant that the intended model could be selected through configuration
and the API service's startup depended on the database and Ollama services being ready.

**What did I decide and why?**

The model currently configured for the project is `qwen2.5:3b`. I wanted a CPU-only option so that the project would not
depend on the hardware available to whoever reviews it. I kept the model name configurable through an environment
variable so that a different model can be tried later without changing the application code.

**What did I learn?**

I learned that running multiple services with Docker Compose involves more than defining the containers: their
configuration and startup dependencies also need to be considered. In particular, the API needs its database and
Ollama dependencies to be available before it can operate correctly.

---

### Research Topic: BM25 and Lexical Search

**What did I need to learn?**

I needed to understand how to search the text extracted from session notes and identify chunks that may contain
information relevant to a user's question.

**Why did I need to learn this?**

Players need to be able to ask questions about their campaign notes without manually searching through each
uploaded document. Lorekeeper therefore needs a way to compare a question with the text chunks created from those
documents.

**Where did I research this?**

| Source                                                           | Link                                                                                                           | Why I used it                                                                                                                                            |
|------------------------------------------------------------------|----------------------------------------------------------------------------------------------------------------|----------------------------------------------------------------------------------------------------------------------------------------------------------|
| Understanding TF-IDF (Term Frequency-Inverse Document Frequency) | https://www.geeksforgeeks.org/machine-learning/understanding-tf-idf-term-frequency-inverse-document-frequency/ | I used this as an introduction to term frequency and how the importance of words can be considered when searching text.                                  |
| Understanding the BM25 full text search algorithm                | https://emschwartz.me/understanding-the-bm25-full-text-search-algorithm/                                       | I used this to learn how the BM25 ranking approach works for full-text search.                                                                           |
| What is BM25 (Best Matching 25) Algorithm                        | https://www.geeksforgeeks.org/nlp/what-is-bm25-best-matching-25-algorithm/                                     | I used this as another explanation of BM25 and its use for identifying relevant text.                                                                    |
| BM25 Go Package                                                  | https://pkg.go.dev/github.com/iwilltry42/bm25-go/bm25                                                          | This seemed like an obvious choice for my intentions, but ultimately I did not end up using this due the specificity I needed with my particular models. |

**What did I discover?**

BM25 is a lexical search method: it looks for matches between the terms in a query and the terms present in a
document. In Lorekeeper, it is used to identify the frequency of key terms from a user's question within the
overlapping chunks of extracted session-note text, and use those matches to find relevant chunks.

This differs from a search approach that tries to identify similarity in meaning without relying only on matching
the same words. The current implementation uses BM25 as a standalone search method.

**What approaches did I consider?**

The project could use more advanced methods to identify relevance, potentially combining lexical matching with
other approaches. Due to time constraints, I kept the current implementation to BM25 as a standalone method rather
than expanding the search system further.

**What problems did I encounter?**

The current search can treat unhelpful words in a natural-language question as search terms. For example, question
words such as “what”, “when” and “where”, or filler and connecting words such as “and”, “but” and “so”, may be
included as key terms even when they do not help identify the information being requested.

Differences in spelling and grammar are another area that the current search could handle better.

**How did I resolve the problem?**

I kept BM25 as the search method for this version of the project. I have not yet implemented query preprocessing
to remove unhelpful terms or account for spelling and grammatical differences.

**What did I decide and why?**

I chose BM25 as the current standalone lexical search method because it provides a way to compare the words in a
question with the text in the extracted note chunks. With the time available, I prioritised implementing a basic
search approach rather than attempting a more advanced relevance system.

**What did I learn?**

I learned about lexical search and how BM25 can be used to rank text based on the terms shared by a query and a
document. I also identified that matching terms alone does not necessarily capture the user's intent. Preparing
the query before searching and improving how relevant chunks are identified are worthwhile areas for further work.

---

## 3. Technical Problems and Failed Approaches

This section records significant problems encountered during development, including approaches that did not work.

| Problem                    | Approach attempted                                                                       | Outcome                                                                                                    | What I learned                                                                                      |
|----------------------------|------------------------------------------------------------------------------------------|------------------------------------------------------------------------------------------------------------|-----------------------------------------------------------------------------------------------------|
| PDF extraction dependency  | Initially investigated `pdftotext-go`, which required Poppler utilities.                 | I could not get the dependency configured, so I investigated another Go package and used `ledongthuc/pdf`. | Library dependencies and setup requirements are important when deciding whether a tool is suitable. |
| Ollama model configuration | Tried to pull the configured model using a value supplied through environment variables. | Needed to adjust the configuration so the intended model could be selected correctly.                      | Configuration passed between services needs to be checked carefully.                                |
| Container startup order    | Needed the API to wait for the database and Ollama services.                             | Updated the Docker Compose setup to account for both dependencies being ready.                             | Starting a container does not necessarily mean the service it provides is ready for use.            |
| Search query terms         | Current BM25 search uses terms from the natural-language question.                       | Filler and question words can be included as search terms; preprocessing has not yet been implemented.     | Query preparation may improve lexical search relevance.                                             |

---

## 4. Technical Decisions

This section records significant architectural and implementation decisions made during development.

| Decision                                                           | Alternatives considered                                           | Reason for decision                                                                                                                                              |
|--------------------------------------------------------------------|-------------------------------------------------------------------|------------------------------------------------------------------------------------------------------------------------------------------------------------------|
| Use `ledongthuc/pdf` for PDF extraction                            | `pdftotext-go` with its Poppler utilities dependency              | I could not get the initial dependency configured, so I selected an alternative Go package that I could use.                                                     |
| Run Ollama in Docker Compose                                       | Running the model service separately from the Compose environment | Keeping Ollama alongside the other services makes it part of the project's containerised setup.                                                                  |
| Make the Ollama model configurable through an environment variable | Hard-coding the model name                                        | This allows the model to be changed later without modifying application code.                                                                                    |
| Use `qwen2.5:3b` as the current model                              | Other models were not compared during the project                 | I wanted a CPU-compatible model because the hardware available to the recruitment team is unknown. The specific choice was not based on a documented comparison. |
| Use BM25 as the current search method                              | More advanced or combined relevance methods                       | BM25 provided a basic lexical search approach that could be implemented within the time available.                                                               |

---

## 5. Further Learning and Future Improvements

If I had more time to develop Lorekeeper, I would investigate the following areas:

- **Natural-language query preprocessing:** Investigate simple ways to turn a user's question into more useful search 
terms before passing it to BM25. This could include removing stop words and question words such as “what”, “when” and 
“where”, as well as filler or connecting words that are unlikely to help find relevant notes.
- **Spelling and grammatical variation:** Explore ways to account for spelling differences and variations in word forms 
so that a search is not overly dependent on the exact wording used in the question and session notes. This is 
particularly important for the context of D&D session notes, as the spellings of people/place names can be fairly
ambiguous.
- **More advanced relevance identification:** Research approaches that could improve on standalone lexical search, 
including whether BM25 could be combined with other relevance or semantic-search methods. I would compare these 
approaches before deciding whether the additional complexity improves results.
- **Model evaluation:** Compare alternative CPU-compatible Ollama models to establish whether another model is more 
suitable for Lorekeeper.

---

## 6. Final Reflection

**What did I learn during this challenge?**

I learned about extracting text from PDF files using Go packages, running local LLMs as part of a Docker Compose
environment, configuring services and their dependencies, and using BM25 for lexical search over text chunks.
I also learned the value of investigating alternatives when an initial technical approach is blocked by setup
difficulties.

**What was the most challenging part of the project?**

Some of the challenges involved getting unfamiliar dependencies and service configuration working. The PDF
extraction research led me to move away from `pdftotext-go` when I could not get its Poppler utilities dependency
running. I also encountered configuration issues with selecting the Ollama model through environment variables
and ensuring the API waited for its dependencies.

**What would I approach differently next time?**

I would allow more time to compare alternatives before settling on particular tools and models. For example, I
would evaluate more than one CPU-compatible Ollama model and investigate query preprocessing before deciding how
to improve search relevance.

**What am I most pleased with?**

I am pleased that I was able to work through unfamiliar tooling and find workable alternatives when my initial
approaches did not succeed. I also kept the Ollama model configurable, leaving room to evaluate a more suitable
model in future. Overall, I am very pleased with the user journey created by lorekeeper, it feels like a natural
addition to a D&D game.

**How has my understanding of software development improved?**

The project reinforced that software development involves researching and evaluating dependencies, understanding
how services interact, and making decisions based on the time and environment available. It also showed me that a
working search implementation is a starting point: its limitations can help identify specific areas for further
learning and improvement.

---

## 7. References

The following sources were saved in Zotero during the research process.

[1] P. K. Mohanty, ‘Deploying Ollama with Docker Compose: A Simple Guide to Local LLMs’, Medium. Accessed: Nov. 16,
2026. [Online]. Available: https://medium.com/@prasanta.mohanty/deploying-ollama-with-docker-compose-a-simple-guide-to-local-llms-610db2991581

[2] aturt13, ‘How to create an ollama model using docker-compose?’, Stack Overflow. Accessed: Nov. 16, 2026.
[Online]. Available: https://stackoverflow.com/q/78438394

[3] ‘Understanding TF-IDF (Term Frequency-Inverse Document Frequency)’, GeeksforGeeks. Accessed: Nov. 10, 2026.
[Online]. Available: https://www.geeksforgeeks.org/machine-learning/understanding-tf-idf-term-frequency-inverse-document-frequency/

[4] Evan Schwartz, ‘Understanding the BM25 full text search algorithm’, Evan Schwartz. Accessed: Nov. 10, 2026.
[Online]. Available: https://emschwartz.me/understanding-the-bm25-full-text-search-algorithm/

[5] ‘What is BM25 (Best Matching 25) Algorithm’, GeeksforGeeks. Accessed: Nov. 10, 2026. [Online]. Available:
https://www.geeksforgeeks.org/nlp/what-is-bm25-best-matching-25-algorithm/

[6] ‘bm25 package - github.com/iwilltry42/bm25-go/bm25 - Go Packages’. Accessed: Nov. 16, 2026. 
[Online]. Available: https://pkg.go.dev/github.com/iwilltry42/bm25-go/bm25


---

*Research log maintained throughout the development of Lorekeeper.*

*Project: Enterprise Automation Services Technical Challenge – November 2026.*
