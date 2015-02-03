#include <stdio.h>

int fibonacci(int n);

void fibonaccis(int n);

int main(void) {
	int n;
	printf("How many numbers do you need? ");
	scanf("%d", &n);
	fibonaccis(n);
	return 0;
}

int fibonacci(int n) {
	//http://en.wikipedia.org/wiki/Fibonacci_number
	// The first two number choiced be 0 and 1 
	if (0 == n) {
		return 0;
	} else if (1 == n) {
		return 1;
	} else {
		return fibonacci(n-1) + fibonacci(n-2);
	}
}

void fibonaccis(int n)
{
	int i = 0;
	while(i < n) {
	 	printf("%d\t", fibonacci(i++));
	}
	printf("\n");
}