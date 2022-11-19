#clang ./arc.c -emit-llvm -S -o ./arc.bc
clang ./objects.c -emit-llvm -S -o ./objects.bc
clang ./exceptions.c -emit-llvm -S -o ./exceptions.bc

#opt ./arc.ll > ./arc.bc
#opt ./objects.ll > ./objects.bc
#opt ./systemlib.ll > ./systemlib.bc

llvm-link ./objects.bc ./exceptions.bc > ./systemlib_lin.bc
llvm-dis ./systemlib_lin.bc > ./systemlib_lin.ll