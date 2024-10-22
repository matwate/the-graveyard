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
    WorkerRows = workers.AddWorkers(4, 0)
    WorkerRows2 = workers.AddWorkers(4, 1)
    WeekDayDateRow[1], WeekDayDateRow[2] = WeekDayDateRow[2], WeekDayDateRow[1]
    WorkerRows[2], WorkerRows[3] = WorkerRows[3], WorkerRows[2]
    WorkerRows2[2], WorkerRows2[3] = WorkerRows2[3], WorkerRows2[2]
    hoja.append(MonthRow)
    for row in WeekDayDateRow:
        hoja.append(row)
    for row in WorkerRows:
        hoja.append(row)
    hoja.append([])
    for row in WorkerRows2:
        hoja.append(row)

    wb.save('test.xlsx')
    Finishes.Finish(TurnoList, CodigoList, TrabajadorList)
    colours.Color()
