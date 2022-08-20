package main
import "fmt"

var a, b, c string 

func main() {
    fmt.Scanf("%s", &a)
    fmt.Scanf("%s", &b)
    fmt.Scanf("%s", &c)
    
    if a == "vertebrado"{
        if b == "ave"{
            if c == "carnivoro"{
                fmt.Printf("aguia\n")
            }
            if c == "onivoro"{
                fmt.Printf("pomba\n")
            }
            }
        if b == "mamifero"{
            if c == "onivoro"{
                fmt.Printf("homem\n")
            }
            if c == "herbivoro"{
                fmt.Printf("vaca\n")
            }
        }
    }

    if a == "invertebrado"{
        if b == "inseto"{
            if c == "hematofago"{
                fmt.Printf("pulga\n")
            }
            if c == "herbivoro"{
                fmt.Printf("lagarta\n")
            }
        }
        if b == "anelideo"{
            if c == "hematofago"{
                fmt.Printf("sanguessuga\n")
            }
            if c == "onivoro"{
                fmt.Printf("minhoca\n")
            }
        }
    }
}