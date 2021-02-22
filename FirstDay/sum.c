/*/OneYearProject/FirstDay/sum.c*/
/*terminal上で入力した数を加算したものを出力するプログラム*/
#include <stdio.h>
#include <stdlib.h>
#include <ctype.h>

int main(int argc, char *argv[])
{
    /*argcは入力数,argv[]は入力されたもの*/
    int sum = 0;

    for (int i = 1; i < argc; i++)
    {
        int value = 0;
        value = atoi(argv[i]);
        sum += value;
    }
    printf("出力結果は%dです\n", sum);
    return 0;
}