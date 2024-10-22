# Ingresar 
# 1. Cedula
# 2. Si, No , Blanco
# 3. Facultad (1,9)

def main():
    print("Software de votacion")
    print("--------------------")
    cedula = int(input("Ingrese numero de cedula: "))
    vote = int(input(" Ingrese voto segun lo que indique \n1. Si\n2. No\n3. Blanco"))                   
    faculty = int(input("Ingrese facultad (1-9): "))
    
    with open('./votes.csv', 'w') as f:
        f.write(f"{cedula},{faculty},{vote}")
    
    print("Los datos han sido guardados de manera exitosa")
    
    


if __name__  == "__main__":
    main()