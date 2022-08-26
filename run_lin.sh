./rgoc -llvm -o ./out.ll $1

opt ./out.ll > ./out.bc
opt ./packages/sys.ll > ./packages/sys.bc
llvm-link ./out.bc ./systemlib/systemlib_lin.bc ./packages/sys.bc > ./program.bc
clang -lc++ -lm -pthread -rdynamic ./program.bc -o ./program

./program
