#include <stdio.h>

int is_prime(int n);

int main(void) {
	int n;
	printf("Input number to checkout is prime or not ?");
	scanf("%d", &n);
	int i = 2;
	int step = n;
	for(; i<=step; i++) {
		while( step % i == 0) {
			if(is_prime(i)) {
				if (i != n) {
					printf("%d\t", i);
				}	
				step = step / i;
			}
		}
	}
	printf("\n");
}

int is_prime(int n)
{
	if (n < 2) {
		return 0;
	}
	int i = 2;
	for(;i<n; i++)
	{
		if ((n % i) == 0) {
			return 0;
		} 
	}
	return 1;
}