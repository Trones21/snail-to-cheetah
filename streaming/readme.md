### Phase 2.IO – Streaming
**Objective**: Optimize file I/O by reading only what’s necessary.

**Approach**:
- Stream file content using tools like `bufio.Scanner`.
- Stop processing once the needed part (e.g., front matter) is read.

**Outcome**:
- Reduced memory footprint for large files.
- Improved I/O efficiency while keeping the process simple.
