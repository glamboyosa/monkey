# Building an Interpreter and Compiler for the Monkey Language

This project is based on my learnings from Thorsten Ball's books, "Writing An Interpreter in Go" and "Writing A Compiler in Go". It involves creating both an interpreter and a compiler for a simple programming language called Monkey. The interpreter directly executes the source code, while the compiler translates it into an intermediate representation that can be executed by a virtual machine.

## The Interpreter

The interpreter serves as the foundation of this project, focusing on parsing and executing the Monkey language.

### Key Features:

1. **Abstract Syntax Tree (AST):** The interpreter constructs an AST to represent the structure of the source code, which it walks during execution to evaluate expressions and statements.

2. **Lexical Analysis and Parsing:** The lexer and parser components break down the source code into tokens and then organize these tokens into the AST using **Pratt Parsing technique** and a **recursive descent parser**.

3. **Environment Handling:** It includes an environment to manage variables and their values, enabling scoping and closures.

4. **Higher-Order Functions:** The interpreter supports higher-order functions, allowing functions to be passed as arguments, returned from other functions, and stored in variables.

5. **Control Flow Constructs:** Basic control flow features, such as conditionals (`if` statements) and loops (`while` loops), are implemented.

## The Compiler and Virtual Machine

Building on the interpreter, the compiler and virtual machine take the language to the next level by translating the code into an intermediate representation that can be executed more efficiently.

### Key Features:

1. **Bytecode Generation:** The compiler converts the source code into bytecode, an intermediate representation that is optimized for execution by a virtual machine.

2. **Virtual Machine (VM):** The VM executes the generated bytecode, providing a platform-independent layer for running the Monkey programs.

3. **Optimizations:** Basic optimizations are implemented to improve the efficiency of the bytecode, including constant folding and dead code elimination.

4. **Garbage Collection:** The virtual machine includes garbage collection to manage memory automatically, reclaiming space used by objects that are no longer in use.

5. **Extended Language Features:** The compiler may introduce additional language features that were not present in the interpreter, such as more complex data types and advanced control flow mechanisms.

## Future Directions

While the initial implementation is in Go, there is potential for exploring the Monkey language in other programming languages. This exploration could offer valuable insights into different paradigms and language implementation techniques.

---

This README outlines my journey in building an interpreter and compiler for the Monkey language, detailing the features of both components and considering future expansions.
