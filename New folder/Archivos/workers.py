import calendarg
WorkerRows = []
StarterDayidx = 0
def AddWorkers(num1, num2, mode):
    WeekRow = calendarg.Rows1to3[0]
    DateRow = calendarg.Rows1to3[1]
    DOTWRow = calendarg.Rows1to3[2]
    
    WorkerRows1 = [["Código", "Nombre"] for _ in range(num1)]
    WorkerRows2 = [["Código", "Nombre"] for _ in range(num2)]
    
    states = ["TD", "DA", "AD", "DT"]
    states2 = ["N", "D", "A", "D"]
    
    maridx = -2
    for i, v in enumerate(DOTWRow):
        if v == "Mar":
            StarterDayidx = i
            break
        else:
            maridx += 1
    
    indices = [idx for idx, value in enumerate(DateRow) if value == 31]
    lastDIdx = int(indices[-1])
    
    def fill_worker_rows(worker_rows, idx_start):
        idx = idx_start
        for row in worker_rows:
            week = 0
            row.extend([""] * maridx)
            for i in WeekRow:
                if i != "":
                    CurrentState = states[(idx + week) % 4]
                    if CurrentState == "TD":
                        row.extend(["N"] * 7)
                    elif CurrentState == "DA":
                        row.extend(["D"] * 7)
                    elif CurrentState == "AD":
                        row.extend(["A"] * 7)
                    elif CurrentState == "DT":
                        row.extend(["D"] * 7)
                    week = (week + 1) % 4
            for i, v in enumerate(row):
                if i > lastDIdx:
                    row.pop(i)
            for i in range(StarterDayidx):
                if i >= 2:
                    row[i] = states2[states.index(states[(idx + week - 1) % 4])]
            idx += 1
        return worker_rows
    
    WorkerRows1 = fill_worker_rows(WorkerRows1, 0)
    WorkerRows2 = fill_worker_rows(WorkerRows2, num1)
    
    WorkerRows1 = AddTCS(mode, WorkerRows1)
    WorkerRows2 = AddTCS(mode, WorkerRows2)
    
    return WorkerRows1, WorkerRows2


    return WorkerRows
def AddTCS(mode,table):
    DOTWRow = calendarg.Rows1to3[2]
    for row in table:
        for i,v in enumerate(DOTWRow):
            try:
                if v == "Sab" or v == "Dom":
                    if row[i] == "N":
                        row[i] = "TC" 
                    if row[i] == "A":
                        row[i] = "TC" if mode == 0 else "TP" 
                    if row[i] == "D":
                        row[i] = "L"
            except:
                pass
    return table


    
# end def
