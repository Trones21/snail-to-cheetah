# snail-to-cheetah
or gazelle, or jaguar... ...consider the hare.

<iframe width="560" height="315" src="https://youtu.be/3oJCoPd9wh0?si=JBtnTopmOgbPyH3a&t=47" frameborder="0" allow="accelerometer; autoplay; clipboard-write; encrypted-media; gyroscope; picture-in-picture" allowfullscreen></iframe>

# Background: The Challenge and Motivation

## The Problem with Scale
Many tools and ideas in software engineering don’t need to operate at true scale. When building CRUD apps or simpler CLI tools, performance often takes a backseat to usability, flexibility, or time-to-market.

As a result, opportunities to deeply explore performance optimizations in a meaningful way—thinking critically about I/O, concurrency, and architecture—are rare.  

I feel this tension acutely because while I've learned and used these concepts, many of my past projects haven’t demanded their full application. 

## The Tension Between Specificity and Generality
- Optimizations often depend heavily on context—specific use cases, file types, or workflows.
- Writing highly optimized code for one problem (e.g., front matter parsing) often makes the code less useful for broader applications.
- Yet, general-purpose solutions tend to trade away the kind of performance improvements that come from knowing the problem inside-out.

## Turning This Into an Opportunity
This project is not about solving a singular problem. It’s a sandbox for pushing boundaries:
- **Understanding** how different approaches impact performance.
- **Exploring scalability** across file sizes, directory depths, and processing complexity.
- **Gaining mastery** over Go’s concurrency model, I/O patterns, and profiling tools.

I'm setting out to immerse myself in performance-oriented engineering—taking concepts from "I’ve used them before" to "I deeply understand them."

It's all part of making Go my language of focus for the next few years. 

---

# Framing the Project: A Personal Deep Dive
This isn’t just a tool; it’s a deliberate journey into mastery. 

## My Personal Growth Goals

### Master the Layers of Performance:
- **I/O**: Streaming vs. full reads, memory buffers, system calls.
- **Concurrency**: Worker pools, channels, contention, deadlocks.
- **Scalability**: Profiling bottlenecks, balancing memory vs. CPU usage.

### Mastery in Trade-Offs Leads to Thoughtfulness

- Deep exploration of **performance trade-offs** (e.g., `+` string concatenation vs. `strings.Builder`) solidifies an intuitive grasp of **why certain patterns matter**.

- **Profiling and analyzing real-world data**:
  - Reinforces how small choices, like managing allocations or using efficient data structures, scale over time.
  - Builds a sharper sense of **when an optimization is worth the cost** and when simplicity should prevail.

- **This mastery translates to thoughtful coding**:
  - Writing code with a clear awareness of its impact on **performance** and **maintainability**.
  - Balancing immediate needs with an eye toward **long-term scalability** and **resource efficiency**.


### Build a Portfolio of Knowledge:
- Create a resource that showcases not just the final optimized tool but the entire **process of improvement**.
- Document **experiments, benchmarks, and lessons learned** (youtube series)

---

# Project Scope: Defining the Sandbox
To keep this exploration manageable, the project focuses on **file processing at scale**:

## Core Challenges:
- File discovery and traversal (handling deep directories, globbing).
- Streaming vs. full reads for various file sizes.
- Concurrent processing with Go routines and worker pools.

## Scenarios for Experimentation:
- **Small files in large directories** (high directory traversal overhead).
- **Few large files** (high memory pressure during processing).
- **Mixed workloads** (requires adaptive solutions).

## Output Goals:
- A clear progression of increasingly optimized solutions, each documented with **benchmarks and profiles**.
- Insight into **when and why** different approaches work best.

---

# The Plan So Far - Phases of the Project

## Phase 1: Serial Full Read (The Baseline)
**Objective**: Establish a simple, unoptimized baseline.

**Approach**:
- Collect file paths serially.
- Process each file sequentially.
- Read the entire content of each file into memory.

**Outcome**:
- Highlight bottlenecks in memory usage, I/O, and execution time.
- Set the stage for measurable improvements in subsequent phases.

---

## Phase 2: General Speedups
**Objective**: Explore optimizations for specific parts of the process without introducing orchestration complexities.

### Phase 2.IO – Streaming
**Objective**: Optimize file I/O by reading only what’s necessary.

**Approach**:
- Stream file content using tools like `bufio.Scanner`.
- Stop processing once the needed part (e.g., front matter) is read.

**Outcome**:
- Reduced memory footprint for large files.
- Improved I/O efficiency while keeping the process simple.

### Phase 2.Parallelization – Concurrency
**Objective**: Introduce concurrency to improve throughput and scalability.

**Approach**:
- Experiment with OS threads (e.g., using C or Python) to understand their limits.
- Use Go routines to leverage Go’s lightweight concurrency model.

**Outcome**:
- Demonstrate the scalability and memory efficiency of Go routines.
- Highlight differences in execution time, resource usage, and ease of implementation.

---

## Phase 3: Combining Phase 2 Techniques
**Objective**: Explore the interaction of streaming and concurrency for maximum performance.

**Approach**:
- Combine streaming (Phase 2.IO) with parallel processing (Phase 2.Parallelization).
- Experiment with concurrency in different parts of the pipeline:
  - File discovery (directory traversal).
  - File processing (reading and analyzing).
- Evaluate resource allocation (e.g., balancing workers across tasks).
- Push the system to limits:
  - Investigate scenarios where processes become **CPU-bound** or **I/O-bound**.
  - Test large-scale workloads (e.g., millions of files or massive file sizes) to observe downstream effects.

**Outcome**:
- Achieve a more sophisticated solution that balances I/O and parallelism.
- Understand how workloads shift between **CPU-bound** and **I/O-bound**, and their downstream effects.
- Identify whether **massive-scale workloads** reveal inherent time or space complexities.

---

## Phase 4: Problem-Specific Optimizations
**Objective**: Address highly specific, real-world problems or edge cases.

**Approach**:
- Identify bottlenecks unique to specific workloads or file structures.
- Experiment with advanced techniques, such as:
  - Dynamic worker allocation based on file size or type.
  - Prioritization of tasks (e.g., processing certain files first).
  - System-level optimizations (e.g., pre-fetching data, caching).
- Analyze improvements at both **micro-scale** (e.g., reducing runtime from 300ms to 100ms) and **macro-scale** (e.g., handling millions of files efficiently).

**Outcome**:
- Showcase the depth of performance tuning for specialized scenarios.
- Provide insights into when such optimizations are worth the effort.
- Explore whether **unknown or unexplored time or space complexities** emerge at large scales.

---

## Where Structural Concurrency Fits
Structural concurrency is relevant when tasks depend on each other, requiring coordination. While this project doesn’t involve dependency chains or orchestration needs, structural concurrency concepts could still apply in specific areas:

### Directory Traversal
- If directory traversal is deep or complex, parallelize it and manage traversal state using channels or worker pools.
- **Example**: Parallelizing directory traversal while ensuring only one thread writes to the shared list of file paths.

### Resource Focus
- Use profiling to decide where resources (e.g., threads or memory) should focus:
  - Should you parallelize file traversal or file processing?
  - How do you prioritize resource-intensive tasks?

### Experimentation
- Even if the problem doesn’t need orchestration now, introduce mock dependency chains to explore structural concurrency techniques.
- **Example**: Simulate a pipeline where a file-read step triggers a dependent processing step, introducing delays or artificial bottlenecks to study Go's handling of dependencies.

---

## Benefits of This Approach
- **Exploratory Depth**: Each phase lets me zoom into a specific optimization area while leaving room for advanced concepts like resource allocation and orchestration.
- **Realism vs. Mastery**: While the current problem is simple, the structure allows for experimentation with more complex scenarios (e.g., dependency chains, structural concurrency) to push my understanding of resource management and Go’s capabilities.
- **Practical + Theoretical Balance**: I’m not just solving the immediate problem but creating a framework for deeper learning and real-world application.
- **Pushing Limits**: By testing at extreme scales, I aim to understand:
  - The effects of workloads becoming CPU-bound or I/O-bound.
  - The impact of massive workloads on time and space complexity.
  - How optimizations behave at both small and large scales.

---

## What’s Next?
- Refine how Phase 3 handles different parts of the process (e.g., directory traversal vs. file reading) and decide on specific experiments.
- Push the boundaries of the system to uncover previously unknown performance bottlenecks.
- If structural concurrency feels unnecessary now, keep it as an optional exploration in Phase 4, tied to mock problems or future expansions.


