clang ./arc.c -emit-llvm -S -o ./arc.bc
clang ./objects.c -emit-llvm -S -o ./objects.bc
clang ./systemlib.c -emit-llvm -S -o ./systemlib.bc

#opt ./arc.ll > ./arc.bc
#opt ./objects.ll > ./objects.bc
#opt ./systemlib.ll > ./systemlib.bc

llvm-link ./arc.bc ./objects.bc ./systemlib.bc > ./systemlib_lin.bc
llvm-dis ./systemlib_lin.bc > ./systemlib_lin.ll