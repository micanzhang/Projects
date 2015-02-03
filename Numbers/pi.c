#include <stdio.h>
#include <math.h>

int main(void) 
{
	//read from command line 
	/*
	* there are three ways to read from or write to command line 
	* puts and gets : read or write string
	* getchar and putchar : read or write char
	* scanf and printf : read or write by format provided
	*/
	int precision;
	printf("Enter number of digits to round PI to: ");
	scanf("%d", &precision);
	//pi/4 = 4 * arctan(1/5) - arctan(1/239)
	//http://en.wikipedia.org/wiki/Pi 
	// Machin-like formulae, http://en.wikipedia.org/wiki/Machin-like_formula
	long double pi = 4 * (4 * atan(1.0 / 5.0) - atan(1.0/239.0));
	printf("%.*Lf\n", precision, pi);
	return 0;
}