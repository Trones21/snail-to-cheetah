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
