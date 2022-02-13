opt ./out.ll > ./out.bc
llvm-link ./out.bc ./systemlib/systemlib_lin.bc > ./program.bc
clang ./program.bc -o ./program