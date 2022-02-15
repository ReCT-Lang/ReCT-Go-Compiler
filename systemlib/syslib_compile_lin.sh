clang ./objects.c -emit-llvm -S -o ./objects.ll
clang ./systemlib.c -emit-llvm -S -o ./systemlib.ll

opt ./objects.ll > ./objects.bc
opt ./systemlib.ll > ./systemlib.bc

llvm-link ./objects.bc ./systemlib.bc > ./systemlib_lin.bc