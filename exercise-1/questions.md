Exercise 1 - Theory questions
-----------------------------

### Concepts

What is the difference between *concurrency* and *parallelism*?
> In concurrency multiple computations are done at the same time intermittently. Processes are switched in a contextual manner. Parallelism on the other hand involves the performing of multiple computations simulataneously on different Central Processing Units (CPUs). With concurrency a single CPU can perform two or more computations by switching(contextually) between them. The disadvantage is that more time is spent on completing a task as compared to parallelism where multiple cpus can perform two different computations at the same time.

What is the difference between a *race condition* and a *data race*? 
> Race condition is a hazard in programming software where the result of a computation is dependent on an interleaving process where as data race is a hazard that occurs when two or more threads attempt to read or write a variable at the same time. This typically occurs when at least one of the threads attempts to modify the variable(content of the register).
 
*Very* roughly - what does a *scheduler* do, and how does it do it?
> A scheduler primarily manages finite resources in the sense that it allows multiple tasks to be performed concurrently. It does so by either one or a combination of these; avoiding stalling of the cpu, ensuring that threads or processes have an equitable access to the processor for example. An example is allocating a time quota to requests so that no process or thread exceeds its time quota at the same time preventing the cpu from stalling.


### Engineering

Why would we use multiple threads? What kinds of problems do threads solve?
> Multiple threads allow for computations that are independent of each other to be performed while accessing a limited resource such as memory space, processor or network. Threads improve energy efficiency, they also preserve data (prevent hazards)by allowing for independent computations to be carried out which in a way depend on shared resources such as memory, threading also improves performance through a better utilization of processors (this reduces or elimates stalling - wastage ).

Some languages support "fibers" (sometimes called "green threads") or "coroutines"? What are they, and why would we rather use them over threads?
> Fibers unlike threads are a sequence of instructions which are a subset of a process which run independently of a kernel scheduler. Fibers can be said to be a subset of a thread for threads that have multiple fibres running concurrently.
Fibers employ cooperative multitasking as compared to threads which use preemptive scheduling or multitasking. In networking, fibers are used in applications where softblocking is required so as to not interrrupt the CPU. The main thread is able to continue what it is doing and then yield to the fibre when it is done. The efficiency of the CPU is maximised when threads are used.

Does creating concurrent programs make the programmer's life easier? Harder? Maybe both?
> With respect to the need for concurrent programs, it improves performance in general.Concurrent programs(application software) makes the programmers life easier in applications where application softwares like IDEs used in editing or writing programs are able to compile code which are being written. It helps to correct errors by notifying the programmer of any syntax or semantic errors and reduce compilation time.
In terms of writing them it can be difficult due to the lower level of abstraction involved. Concurrent programs ar not able to interact with other tasks well in certain situations (when multiple concurrent threads have been performed and there is the need for synchroniszing them to get a single output) which usually results in deadlocks or racing conditions. These errors make the programmers life difficult.

What do you think is best - *shared variables* or *message passing*?
> Shared variables and message passing have their advantages and disadvantages. One concept could be better over the other depending on the application/ use.
For systems that are local and do not require remote communication, shared variables presents a faster means of communication since it does not involve the use of system calls(kernels)  which are relatively slower when using message passing technique. On the other hand for systems that communicate remotely i.e using the help of a network message passing is a superior technique. Conflicts are usually avoided since the medium used has protocols in place that synchronizes activities performed by the processes communicating.


