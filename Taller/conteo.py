# Conteo
# Cargar Datos
# Validaar Datos
"""
La cedula debe tener o 10 digitos, empezando por 1 o 8-6
"""
# Contar Votos
# Una estadistica


votes = []
facultys = []
percentages = []
cedulas = []
def main():
    with open("./votes.csv", "r") as f:
        data = f.readlines()
        for line in data:
            cedula, vote, faculty = line.split(',')
            
            # all() toma una lista y devuelve True si todos los elementos de la lista son True
            try:
                if validate_cedula(cedula):
                    cedulas.append(int(cedula))
                    if validate_vote(vote) and validate_faculty(faculty):
                        votes.append(int(vote))
                        facultys.append(int(faculty))
                    else:
                        raise ValueError("Datos invalidos")
                else:
                    raise ValueError("Datos invalidos")
            except:
                raise ValueError("Datos invalidos")
    percentages = calculate_percentages(votes,facultys)
    
    #Genera la string para cada valor de la lista percentages usando list comprehension            
    output_strings =  [f"Porcentaje de votos de la facultad {i}: {int(v)}% \n" for i,v in enumerate(percentages)] 
    for i in output_strings:
        print(i)
    
    
    
             
                
                
def validate_cedula(cedula: str):
    is_valid = ((len(cedula) == 10 and cedula[0] == "1") or (len(cedula) in range(7,9))) and (int(cedula) not in cedulas)
    
    return is_valid
    
def validate_vote(vote: str):
    return int(vote) in range(1,4)

def validate_faculty(faculty: str):
    return int(faculty) in range(1,11)

def calculate_percentages(votes: list[int], faculty: list[int]):
    amount_of_voters = len(votes)
    percentages= [0] * 9 
    counts ={ } 
    for f in faculty:
        counts[f] = faculty.count(f)
        
    for i in counts.keys():
        percentages[i-1] = counts[i] / amount_of_voters * 100 
    return percentages
    
        
if __name__ == "__main__":
    main()