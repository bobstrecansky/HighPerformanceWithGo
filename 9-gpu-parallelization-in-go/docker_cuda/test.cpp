#include <cmath>
#include <cstdlib>
#include <iostream>

const long int ELEMENTS = 1<<20;

void multiply(int j, float *a, float *b, float *c)
{
 for (int i = 0; i < j; i++)
      c[i] = a[i] * b[i];
}

int main(void)
{

float *a = new float[ELEMENTS];
float *b = new float[ELEMENTS];
float *c = new float[ELEMENTS];

for(int i = 0; i < ELEMENTS; i++)
{
    a[i] = rand() % 10;
    b[i] = rand() % 10;

}

multiply(ELEMENTS, a, b, c);


//for (int k = 0; k < ELEMENTS; k++)
//{
//    std::cout << k << ":" << a[k] << "*" << b[k] << "=" << c[k] << "\n";
//}

delete [] a;
delete [] b;
delete [] c;
return 0;

}
