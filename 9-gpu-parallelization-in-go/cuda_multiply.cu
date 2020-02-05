#include <cmath>
#include <cstdlib>
#include <iostream>

const int ELEMENTS = 1<<25;

__global__
void multiply(int j, float *a, float *b, float *c)
{
 for (int i = 0; i < j; i++)
      c[i] = a[i] * b[i];
}

int main(void)
{

float *a, *b, *c;

cudaMallocManaged(&a, ELEMENTS*sizeof(float));
cudaMallocManaged(&b, ELEMENTS*sizeof(float));
cudaMallocManaged(&c, ELEMENTS*sizeof(float));

for(int i = 0; i < ELEMENTS; i++)
{
    a[i] = 1.0; //rand() % 10;
    b[i] = 2.0; //rand() % 10;

}

multiply<<<1, 1>>>(ELEMENTS, a, b, c);


// Wait for GPU to finish before accessing on host
cudaDeviceSynchronize();

//for (int k = 0; k < ELEMENTS; k++)
//{
//    std::cout << k << ":" << a[k] << "*" << b[k] << "=" << c[k] << "\n";
//}

cudaFree(a);
cudaFree(b);
cudaFree(c);

return 0;

}
