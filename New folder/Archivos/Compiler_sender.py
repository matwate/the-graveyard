from openpyxl import Workbook, load_workbook
import calendarg
import workers
import colours
import Finishes
wb = Workbook()
hoja = wb.active


def main(StartingDay, TurnoList, CodigoList, TrabajadorList, isLeap):
    WeekDayDateRow = calendarg.GenerateYear(StartingDay, isLeap)
    MonthRow = calendarg.GenerateMonths()
    WorkerRows1, WorkerRows2 = workers.AddWorkers(5, 4, 0)
    WorkerRows3, WorkerRows4 = workers.AddWorkers(5, 4, 1)
    
    WeekDayDateRow[1], WeekDayDateRow[2] = WeekDayDateRow[2], WeekDayDateRow[1]
    WorkerRows1[2], WorkerRows1[3] = WorkerRows1[3], WorkerRows1[2]
    WorkerRows2[2], WorkerRows2[3] = WorkerRows2[3], WorkerRows2[2]
    WorkerRows3[2], WorkerRows3[3] = WorkerRows3[3], WorkerRows3[2]
    WorkerRows4[2], WorkerRows4[3] = WorkerRows4[3], WorkerRows4[2]
    
    hoja.append(MonthRow)
    for row in WeekDayDateRow:
        hoja.append(row)
    for row in WorkerRows1:
        hoja.append(row)
    hoja.append([])
    for row in WorkerRows2:
        hoja.append(row)
    hoja.append([])
    for row in WorkerRows3:
        hoja.append(row)
    hoja.append([])
    for row in WorkerRows4:
        hoja.append(row)
    
    wb.save('test.xlsx')
    Finishes.Finish(TurnoList, CodigoList, TrabajadorList)
    colours.Color()
