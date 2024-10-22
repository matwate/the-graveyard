import calendarg
WorkerRows = []
StarterDayidx = 0
def AddWorkers(num,mode):
    WeekRow = calendarg.Rows1to3[0]
    DateRow = calendarg.Rows1to3[1]
    DOTWRow = calendarg.Rows1to3[2]
    WorkerRows = []
    
    for i in range(num):
        WorkerRows.append(["Código","Nombre"])
    states = ["TD","DA","AD","DT"]
    states2 = ["N","D","A","D"]
    idx = 0
    maridx = -2
    for i,v in enumerate(DOTWRow):
        if v == "Mar":
            StarterDayidx = i
            break
        else:
            maridx += 1
    indices = []
    for idx, value in enumerate(DateRow):
        if value == 31:
            indices.append(idx)
    idxs = indices[-1]
    lastDIdx = int(idxs)
    for row in WorkerRows:
        
        week = 0
        row.extend([""] * maridx )
        
        for i in WeekRow:
            
            if i != "":
                CurrentState = states[(idx + week)%4]
                if CurrentState == "TD":
                    row.extend(["N"] * 7)
                    week = (week+1)%4
                if CurrentState == "DA":
                    row.extend(["D"] * 7)
                    week = (week+1)%4
                if CurrentState == "AD":
                    row.extend(["A"] * 7)
                    week = (week+1)%4
                if CurrentState == "DT":
                    row.extend(["D"] * 7)
                    week = (week+1)%4
        for i,v in enumerate(row):
            if i > lastDIdx:
                row.pop(i)
        for i in range(StarterDayidx):
            if i >= 2:
                row[i] = states2[states.index(states[(idx + week - 1)%4])] 
        idx += 1
        
    AddTCS(mode,WorkerRows)


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
