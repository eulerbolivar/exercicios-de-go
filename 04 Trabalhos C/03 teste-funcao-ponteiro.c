#include <stdio.h>
#include <stdlib.h>
#include <math.h>

int envia(int codigo, int multiplicador);

int main(){
    int data1, data2;
    int resp;
    
    printf("Qual endereço você quer ver?\n");
    scanf("%d", &data1);

    printf("Qual o valor potenciado você quer enviar?\n");
    scanf("%d", &data2);

    resp = envia(data1, data2);

    printf("\nVocê enviou %d e o endereço selecionado [%d] lhe devolveu: %d\n", data2, data1, resp);

    return 0;
}

int envia(int codigo, int multiplicador){
    int *endereco, i, j, acumulado;
    endereco = (int *) malloc (5 * sizeof(int));

    acumulado = multiplicador;
    endereco[0] = 1;
    //ATRIBUI VALORES POTENCIALIZADOS DENTRO DO ENDEREÇO
    for (i = 1; i < 5; i++){
        for (j = 0; j <= i; j++){
            acumulado *= multiplicador; //AJUSTAR ESSA POTÊNCIA HORRÍVEL AQUI
        }
        endereco[i] = acumulado;
    }
    
    //DEMONSTRA OS VALORES QUE FORAM POTENCIALIZADOS
    for(i = 0; i < 5; i++){
        printf("no endereço %d ficou armazenado: %d\n", i, endereco[i]);
    }

    return endereco[codigo];
}