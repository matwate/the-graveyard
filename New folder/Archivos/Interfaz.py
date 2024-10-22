import subprocess
import os
import Compiler_sender
import calendarg
import colours
import sys


 
def main(StartingDay, TurnoList, CodigoList, TrabajadorList, isLeap):
    dirc = os.getcwd()

    
    Compiler_sender.main(StartingDay, TurnoList, CodigoList, TrabajadorList, isLeap)

    subprocess.Popen(f"explorer {dirc}")
