#include <stdio.h>

#define BIN_DEC 1;
#define DEC_BIN 2;

int main(void) {
	int type;
	int n;
	printf("\t1. Binary to Decimal\n");
	printf("\t2. Decimal to Binary\n");
	printf("Make a choice: ");
	scanf("%d", &type);
	if (BIN_DEC == type) {
		printf("Binary to Decimal: ");
	} else if (DEC_BIN == type ){
		printf("Decimal to Binary: ");
	} else {
		printf("Invalid Option");
	}
	scanf("%d", &n);
	if (BIN_DEC == type) {
		printf("The binary %b Decimal: ");
	} else if (DEC_BIN == type ){
		printf("Decimal to Binary: ");
	}
	return 0;
}