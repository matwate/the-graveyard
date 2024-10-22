import random


def generar_partidos(equipo_local: list[int], equipo_visita: list[int], goles_local: list[int], goles_visita: list[int], num_equipos: int):
    for i in range(90):
        local_idx = i % num_equipos
        visita_idx = (i + 1) % num_equipos
        equipo_local[i] = local_idx
        equipo_visita[i] = visita_idx
        goles_local[i] = random.randint(0, 5)
        goles_visita[i] = random.randint(0, 5)
        
def calcular_puntos(equipo_local: list[int], equipo_visita: list[int], goles_local: list[int], goles_visita: list[int], total_puntos: list[int], num_partidos: int):
   for i, v in enumerate(equipo_local):
       if goles_local[i] > goles_visita[i]:
           total_puntos[v] += 3
       elif goles_local[i] == goles_visita[i]:
           total_puntos[v] += 1
           total_puntos[equipo_visita[i]] += 1
       else:
           total_puntos[equipo_visita[i]] += 3
           
def calcular_campeon(total_puntos: list[int], num_equipos: int) -> int:
    max_puntos = 0
    campeon = 0
    for i in range(num_equipos):
        if total_puntos[i] > max_puntos:
            max_puntos = total_puntos[i]
            campeon = i
    return campeon    

def print_puntos_totales(equipos: list[str], total_puntos: list[int], num_equipos: int):
    print("Tabla de Puntos: \n")
    for i in range(num_equipos):
        print(f"{equipos[i]}:{total_puntos[i]}")

def main():
    random.seed(1234)
    equipos = ["Junior", "Millonarios", "Santafé", "Nacional", "Medellín", "Bucaramanga", "Pasto" ,"Alianza", "Huila", "Tolima"]
    num_equipos = 10
    num_partidos = 90
    equipo_local = [0]* num_partidos  
    equipo_visita = [0]* num_partidos 
    goles_local = [0]* num_partidos 
    goles_visita = [0]* num_partidos  
    total_puntos = [0]* 10 
        
    generar_partidos(equipo_local, equipo_visita, goles_local, goles_visita, num_equipos)
    calcular_puntos(equipo_local, equipo_visita, goles_local, goles_visita, total_puntos, num_partidos)
    campeon = calcular_campeon(total_puntos, num_equipos)
        
    print_puntos_totales(equipos, total_puntos, num_equipos)
    print("El ganador del campeonato es",equipos[campeon],"con",total_puntos[campeon],"puntos.")

if __name__ == "__main__":
    main()

