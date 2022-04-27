# Profiling Notes

## What is a profiler?
Profiler is a **dynamic performance analysis** tool that provides critical
execution insights in various dimensions which enable resolving
performance issues, locating memory leaks, thread contention and more.

---
## Why use it? What are the advantages?
A profiler can help you optimize your code. It can show you, for example, where the greatest amount of time and resources are spent. 

---
## What types of profilings are there?

### cpuprofile
The Go **CPU profiler** uses a `SIGPROF` signal to **record code execution statistics**. Once the **signal got registered**, it will **deliver** every **specified time interval**. This timer, unlike typical timers, will **increment** when the **CPU** is **executing the process**. The **signal** will **interrupt code execution** and make it possible to **see which code was interrupted**.

### memprofile
The **memory profiler** samples **heap allocations**. It will **show function calls allocations**. Recording all allocation and unwinding thestack trace would be expensive, therefore a sampling technique is used.

The sampling process relies on a **pseudo random number generator** based on **exponential distribution** to sample only a fraction of allocations. The ***generated numbers** define the **distance between samples** in terms of **allocated memory size**. This means that **only allocations** that **cross the next random sampling point** will be sampled.