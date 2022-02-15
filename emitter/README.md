# ReCT Emitter 
\*Slaps roof of abstract theoretical concept\*  
This bad boy can compile lowered and bound ReCT functions to LLVM IR.

## Garbage Collection: The Deranged Mumblings of a Grown man at 3 am
So LLVM garbage collection is quite a huge rabbit hole. I've taken some time to look into possibilities of what we could use and how.
**LLVM does not provide a garbage collector** but it does provide some functionality for interfacing
with GCs and also has a few built in GC strategies you can use.

Our goal is to identify all pointers in the program at rune-time which requires us to check the stack and registers.
LLVM provides support for creating safe-points where GC can happen safely, generating stack maps of object
references the GC may need to update, and creating barriers when storing object references in the heap.

We can use built-in GC strategies like "The Shadow Stack GC" this strategy has high overhead per function call
as it maintains a "shadow stack" which mirrors the machine stack - Note: a faster way would be to compile stack maps 
into the executable.

LLVM IR has a bunch of features to interface with the GC runtime. LLVM seems to rely on "statepoints".

### MMm reddit time
So check out [this reddit post](https://www.reddit.com/r/Compilers/comments/8po0lj/garbage_collection_with_llvm/) from about 4 years ago, dude has the exact same problem,
has a language that can compiler to IR but confused on how GC is put together. Unfortunately, the only decent reply is someone
talking about how plugging in your own GC is hella confusing, and they recommend you use the
[Boehm-Demers-Weiser GC](https://github.com/ivmai/bdwgc) or the [Memory Pool System GC](https://www.ravenbrook.com/project/mps/).
Another reply talks about how they researched the [Boehm-Demers-Weiser GC](https://www.reddit.com/r/Compilers/comments/8po0lj/comment/e0i9dl7/?utm_source=share&utm_medium=web2x&context=3)...


Though I really like the idea of us figuring this out on our own, for the time being it may be easier to use a non-llvm approach.
It seems boehm-demer-weider gc is the easiest to set up, so we should just use that for now.

### Resources
[LLVM Garbage Collection](https://llvm.org/docs/GarbageCollection.html#goals-and-non-goals)
[LLVM Accurate Garbage Collection](https://releases.llvm.org/3.5.2/docs/GarbageCollection.html) 