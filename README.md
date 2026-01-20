

- go mod init example.com


##### Function type
##### Parameter Vs Argument | First Order Function Vs Higher Order Function
1: standard/named func: which func have name

2: init func: which fucn call computer, can not call this func

3: anonymouse fucn: 

4: Immediately invoked fuction expression (IIFE)

5: function expression

6: first order func: working simple task 

7: Higher Order Function/first class func(variable assign any data, str, int, func etc..): 1=> parameter as function, 
2=>function return, 3=> both

8: Receiver function
9: variadic function
10: Defer function - last in first out
11: closure - close over


####  Internal memory | Code Segment | Data Segment | Stack | Heap | GC
[internal memory](https://www.youtube.com/watch?v=J8G--D1tXqU&list=PLpCqPSEm2Xe8sEY2haMDUVgwbkIs5NCJI&index=24)

1: code Segment (all functions)
2: data segment (global memory/variable) => init, main, etc.. func
3: stack (in ececution time like init,main func allocate into stack which is called stack frame) after output its removed
4: heap :  escape analysis: => some outer variable need to be store or use another function then its need to be store into heap, like wheen one heirorder func return a func but this return func have one varible which already diclard into outer or main func, so this variable need to to store into heap (basically from stack ouoer func will be pop ) and also return fucn store to heap besaouse its will call any time and heap just store referance.
5: Garbage Collector =>GC


#### Go Have 2 phases
1: compilation phase
2: Execution phase



#### Closure

- escape analysis: =>


#### Struct

 
#### Context switching | PCB | Concurrency

- controll unit
- instruction register
- proram counter
- General Purpose Rgister
- process controll Block (PCB) (OS) store state info (register set) all meta data
- Concurrency: context switching and concurrently working as like as same time.
- parallelism: without context swtching multiple process run with multiple logical cpu 
- 

