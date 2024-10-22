import customtkinter
import sys
import os

# getting the name of the directory
# where the this file is present.
current = os.path.dirname(os.path.realpath(__file__))

# Getting the parent directory name
# where the current directory is present.
parent = os.path.dirname(current)

# adding the parent directory to
# the sys.path.
sys.path.append(parent)

import Interfaz
Days = ["Dom", "Lun", "Mar", "Mie", "Jue", "Vie", "Sab"]
app = customtkinter.CTk()
app.title("Cronograma")
Turnos = []
Codigos = []
Trabajadores = []

TurnoV = []
CodigoV = []
TrabajadorV = []


def defaultFunc(v):
    pass


def TurnIntoValues(isLeap):
    print("Se Ejecuto")
    for Turno, Codigo, Trabajador in zip(Turnos, Codigos, Trabajadores):
        TurnoV.append(Turno.get())
        CodigoV.append(Codigo.get())
        TrabajadorV.append(Trabajador.get())
    Day = DaySelector.get()
    Interfaz.main(Day,TurnoV, CodigoV, TrabajadorV,isLeap)


for i in range(9):
    label = customtkinter.CTkLabel(
        app, text=f"Trabajador {i+1}" if i != 8 else "Seleccione Día", fg_color="transparent")
    label.grid(row=i, column=0, padx=20, pady=10)
# Entradas de texto
for i in range(8):
    Turnos.append(customtkinter.CTkEntry(app, placeholder_text="Turno"))
    Codigos.append(customtkinter.CTkEntry(app, placeholder_text="Codigo"))
    Trabajadores.append(customtkinter.CTkEntry(
        app, placeholder_text="Nombre Trabajador"))
for Turno, Codigo, Trabajador, i in zip(Turnos, Codigos, Trabajadores, range(8)):
    Turno.grid(row=i, column=1, padx=20, pady=10)
    Codigo.grid(row=i, column=2, padx=20, pady=10)
    Trabajador.grid(row=i, column=3, padx=20, pady=10)

DaySelector = customtkinter.CTkComboBox(app, values=Days, command=defaultFunc)
DaySelector.grid(row=8, column=1, padx=20, pady=20)

LeapSelector = customtkinter.CTkCheckBox(app, text="Es Año Bisiesto?")
LeapSelector.grid(row=8, column=0, padx=20, pady=20)

# Botón De generacion de calendario
GeneratorButton = customtkinter.CTkButton(
    app, text="Generar Calendario", command= lambda: TurnIntoValues(LeapSelector.get()))
GeneratorButton.grid(row=8, column=2, padx=20, pady=20,columnspan=2, sticky="we")


app.mainloop()
