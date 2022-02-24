opt ./out.ll > ./out.bc
llvm-link ./out.bc ./systemlib/systemlib_lin.bc > ./program.bc
clang -lm ./program.bc -o ./program