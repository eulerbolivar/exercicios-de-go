#include <stdio.h>
#include <stdlib.h>
#include <math.h>

float CalculoIMC(float peso, float altura);
char ControleIMC(float IMC);

int main(){

    struct pessoal{
        int idade;
        float peso;
        float altura;
    };

    struct pessoal pessoa1;
    float RespIMC;
    char RespControle[50];

    printf("Qual é a idade da pessoa?\n");
    scanf("%d", &pessoa1.idade);

    printf("Qual é o peso da pessoa?\n");
    scanf("%f", &pessoa1.peso);
    
    printf("Qual é a altura da pessoa?\n");
    scanf("%f", &pessoa1.altura);

    printf("\na pessoa possui %d anos, %.2fkg e %.2f de altura\n", pessoa1.idade, pessoa1.peso, pessoa1.altura);

    RespIMC = CalculoIMC(pessoa1.peso, pessoa1.altura);

    printf("\no IMC dessa pessoa é: %.2f\n", RespIMC);
    ControleIMC(RespIMC);

    return 0;
}

    float CalculoIMC(float peso, float altura){
            float imc;

            imc = peso/(altura * altura);

            return imc;
        };

    char ControleIMC(float IMC){
        char *condicao;
        condicao = (char *) malloc(50 * sizeof(char));
        
        if(IMC < 18.5){
            condicao = "A pessoa está abaixo do peso";
        };

        if(18.6 < IMC && IMC < 24.9){
            condicao = "A pessoa está com o peso saudável";
        };

        if(25 < IMC && IMC < 29.9){
            condicao = "A pessoa está com sobrepeso";
        };

        if(30 <  IMC && IMC < 34.9){
            condicao = "A pessoa está obesa";
        };

        printf("\n%s\n", condicao);
        return 0;
    }