export LD_LIBRARY_PATH := /usr/local/cuda-10.2/lib64
export PATH := /usr/local/cuda-10.2/bin:/usr/local/cuda-10.2/NsightCompute-2019.1:$(PATH)

target:
	nvcc -o libmultiply.so --shared -Xcompiler -fPIC cuda_multiply.cu
go:
	go run cuda_multiply.go
clean:
	rm *.so
