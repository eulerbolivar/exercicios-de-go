#include <stdio.h>
#include <stdlib.h>

int main()
{
    struct teste{
        float *receita;
        float *despesa;
        float *saldo;
        char texto[10];
    };

    int i, quantidade;
    struct teste dados;
    
    //RECEBE O TAMANHO DO VETOR;
    printf("Coloque a quantidade de projetos:\n");
    scanf("%d", &quantidade);
    
    //DEFINE O TAMANHO DO VETOR *inteiro1
    dados.receita = (float *) malloc(quantidade * sizeof(float));
    dados.despesa = (float *) malloc(quantidade * sizeof(float));
    dados.saldo = (float *) malloc(quantidade * sizeof(float));
    
    //RECEBE OS VALORES DE RECEITA
    printf("Coloque os %d valores de receita para cada projeto:\n", quantidade);
    for(i = 0; i < quantidade; i++){
        scanf("%f", &dados.receita[i]);
    }
    
    //RECEBE OS VALORES DE DESPESA
    printf("Coloque os %d valores de despesa para cada projeto:\n", quantidade);
    for(i = 0; i < quantidade; i++){
        scanf("%f", &dados.despesa[i]);
    }
    
    int soma(float numero1, float numero2){
        float *saldo;
        saldo = (float *) malloc(quantidade * sizeof(float));

        *saldo = (numero1 - numero2);

        return *saldo;
        };

    for(i = 0; i < quantidade; i++){
        dados.saldo[i] = soma(dados.receita[i], dados.despesa[i]);
    }

    //DEVOLVE OS NÚMEROS ESCOLHIDOS
    printf("\nEssas são as receitas e despesas de cada projeto:\n");
    for(i = 0; i < quantidade; i++){
        if(dados.saldo[i] > 0){
            printf("o Projeto %d teve uma receita de R$%.2f e despesa de R$%.2f, restando R$%.2f se tornando viável para a produção\n", i, dados.receita[i], dados.despesa[i], dados.saldo[i]);
        }
        if(dados.saldo[i] <= 0){
            printf("o Projeto %d teve uma receita de R$%.2f e despesa de R$%.2f, restando R$%.2f se tornando invíavel para produção\n", i, dados.receita[i], dados.despesa[i], dados.saldo[i]);
        }
    }
    free(dados.receita);
    free(dados.despesa);
    free(dados.saldo);

    printf("fazendo uma alteração no git\n");
    
    return 0;
}
