#include <stdio.h>
#include <stdlib.h>
#include <math.h>

int soma(int num1, int num2){
    int somado;

    somado = num1 + num2;

    return somado;
};

int main(){

    struct informacoes{
    int a, b;
    int resp;
    };

    struct informacoes dados;

    dados.a = 5;
    dados.b = 6;
    dados.resp = soma(dados.a, dados.b);

    printf("o número somado deu em: %d\n", dados.resp);

    return 0;
};