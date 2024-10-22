Days = ["Dom", "Lun", "Mar", "Mie", "Jue", "Vie", "Sab"]
DOM = {
    "31a": "Ene",
    "28b": "Feb",
    "31c": "Mar",
    "30d": "Abr",
    "31e": "May",
    "30f": "Jun",
    "31g": "Jul",
    "31h": "Ago",
    "30i": "Sep",
    "31j": "Oct",
    "30k": "Nov",
    "31l": "Dec",
}
DOML = {
    "31a": "Ene",
    "29b": "Feb",
    "31c": "Mar",
    "30d": "Abr",
    "31e": "May",
    "30f": "Jun",
    "31g": "Jul",
    "31h": "Ago",
    "30i": "Sep",
    "31j": "Oct",
    "30k": "Nov",
    "31l": "Dec",
}
Year = {}
DatesRow = []
DOTWRow = []
WeeksRow = []
Rows1to3 = []
NTecnicosPorRegion = 4


def GenerateYear(StartingDay, Leap):
    CurrentDOM = DOM if Leap == 0 else DOML
    for NumDays in CurrentDOM:
        NumDayss = NumDays[:-1]
        for i in range(int(NumDayss)):
            DatesRow.append(i+1)
    salt = Days.index(StartingDay)
    for i in range(365 if Leap == 0 else 366):
        DOTWRow.append(Days[(i + salt) % 7])
    for i in range(2):
        DatesRow.insert(0, "")
        DOTWRow.insert(0, "")

    return GenerateWeeks()


def GenerateWeeks():
    weekN = 1
    for Day in DOTWRow:
        if Day == "Mar":
            WeeksRow.append(f"Week {weekN}")
            weekN += 1
        else:
            WeeksRow.append("")
    Rows1to3.append(WeeksRow)
    Rows1to3.append(DatesRow)
    Rows1to3.append(DOTWRow)

    return Rows1to3


def GenerateMonths():
    monthslist = []
    monthsidx = 0
    for i in DatesRow:
        if i == 1:
            monthslist.append(DOM[list(DOM.keys())[monthsidx]])
            monthsidx += 1
        else:
            monthslist.append("")
    return monthslist
