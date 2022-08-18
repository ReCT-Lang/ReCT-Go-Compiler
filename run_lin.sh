go run . -- $1

opt ./out.ll > ./out.bc
opt ./packages/test.ll > ./packages/test.bc
llvm-link ./out.bc ./systemlib/systemlib_lin.bc ./packages/test.bc > ./program.bc
clang -lm -pthread -rdynamic ./program.bc -o ./program

./program