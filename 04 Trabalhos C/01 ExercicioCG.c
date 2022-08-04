#include <stdio.h>

int main()
{
    struct dados{
        int NumeroProjeto[10];
        float Receita[10];
        float Despesa[10];
    };

    int i, QuantiaProjetos;
    struct dados projeto;
    
    //COLOCA AS RECEITAS E DESPESAS DE CADA PROJETO
    for(i = 0; i < 10; i++){
        //DÁ UM CÓDIGO AO PROJETO
        printf("Qual é o código do projeto na posição %d?\n", i);
        scanf("%d", &projeto.NumeroProjeto[i]);
        QuantiaProjetos =+ 1;
        if(projeto.NumeroProjeto[i] == -1){
            break;
        }

        //DÁ UMA RECEITA AO NÚMERO DO PROJETO
        printf("Qual a receita do projeto código %d?\n", projeto.NumeroProjeto[i]);
        scanf("%f", &projeto.Receita[i]);

        //DÁ UMA DESPESA AO NÚMERO DO PROJETO
        printf("Qual a despesa do projeto código %d?\n", projeto.NumeroProjeto[i]);
        scanf("%f", &projeto.Despesa[i]);
    }

    printf("a quantidade de projetos feitos foram: %d\n", QuantiaProjetos);

    //for(i = 0; i < QuantiaProjetos; i++))


    return 0;
}
