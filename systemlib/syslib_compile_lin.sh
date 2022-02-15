clang ./arc.c -emit-llvm -S -o ./arc.ll
clang ./objects.c -emit-llvm -S -o ./objects.ll
clang ./systemlib.c -emit-llvm -S -o ./systemlib.ll

opt ./arc.ll > ./arc.bc
opt ./objects.ll > ./objects.bc
opt ./systemlib.ll > ./systemlib.bc

llvm-link ./arc.bc ./objects.bc ./systemlib.bc > ./systemlib_lin.bc