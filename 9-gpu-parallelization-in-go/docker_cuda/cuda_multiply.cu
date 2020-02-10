#include <cmath>
#include <cstdlib>
#include <iostream>

const int ELEMENTS = 1 << 20;

__global__ void multiply(int j, float * a, float * b, float * c) {
  int index = threadIdx.x * blockDim.x + threadIdx.x;
  int stride = blockDim.x * gridDim.x;

  for (int i = index; i < j; i += stride)
    c[i] = a[i] * b[i];
}

extern "C" {
  int cuda_multiply(void) {
    float * a, * b, * c;

    cudaMallocManaged( & a, ELEMENTS * sizeof(float));
    cudaMallocManaged( & b, ELEMENTS * sizeof(float));
    cudaMallocManaged( & c, ELEMENTS * sizeof(float));

    for (int i = 0; i < ELEMENTS; i++) {
      a[i] = rand() % 10;
      b[i] = rand() % 10;

    }

    int blockSize = 256;
    int numBlocks = (ELEMENTS + blockSize - 1) / blockSize;
    multiply << < numBlocks, blockSize >>> (ELEMENTS, a, b, c);

    // Wait for GPU to finish before accessing on host
    cudaDeviceSynchronize();

    //for (int k = 0; k < ELEMENTS; k++) {
      //std::cout << k << ":" << a[k] << "*" << b[k] << "=" << c[k] << "\n";
    //}

    cudaFree(a);
    cudaFree(b);
    cudaFree(c);
    return 0;

  }
}
