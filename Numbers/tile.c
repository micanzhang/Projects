#include <math.h>
#include <stdio.h>

int main(void) {
	float cost, width, height;
	printf("What's the cost of per sq.feet? ");
	scanf("%f", &cost);
	printf("What's the width of the floor? ");
	scanf("%f", &width);
	printf("What's the height of thr floor? ");
	scanf("%f", &height);
	printf("The total cost is $%.2f for %.2f square feet\n", ceil(width * height) * cost, width * height);
	return 0;
}