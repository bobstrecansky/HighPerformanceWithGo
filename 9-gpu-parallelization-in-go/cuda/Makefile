export LD_LIBRARY_PATH := /usr/local/cuda-10.2/lib64
export PATH := /usr/local/cuda-10.2/bin:/usr/local/cuda-10.2/NsightCompute-2019.1:$(PATH)

target:
	nvcc -o cuda_add.so --shared -Xcompiler -fPIC cuda_add.cu
go:
	go run cuda_add.go
clean:
	rm cuda_add.so
