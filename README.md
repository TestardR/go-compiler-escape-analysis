**When does the Go compiler place variables on the heap and the stack ?**

- Sharing gown typically stays on the stack 
  - Passing references down as function parameters
- Sharing up typically escapes to the heap
  - Returning references (pointers) from functions 
- "typically": only the compiler knows.

**When are values constructed on the heap ?**
1. When a value could possibly be referenced after the function that constructed the values returns.
2. When the compiler determines a value is too large to fit on the stack.
3. When the compiler does not know the size of a value at compile time.

**Points to remember**
- Optimize for correctness, not performance
- Go only puts function variables on the stack if it can prove a variable is not used after the function returns
- Sharing down typically stays on the stack
- Sharing up typically escapes to the heap
- You want to know ? Ask the compiler to find out using the tooling