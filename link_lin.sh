opt ./out.ll > ./out.bc
llvm-link ./out.bc ./systemlib/systemlib_lin.bc > ./program.bc
clang -lm -rdynamic ./program.bc -o ./program