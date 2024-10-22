from pydantic import BaseModel

def find_mean(data: list[float]) -> float:
    return sum(data) / len(data)
        
def find_median(data: list[float]) -> float:
    data.sort()
    n = len(data)
    if n % 2 == 0:
        return (data[n // 2 - 1] + data[n // 2]) / 2
    return data[n // 2]

def find_mode(data: list[float]) -> float:
    mode = max(set(data), key = data.count)
    return mode

def find_max(data: list[float]) -> float:
    return max(data)

def find_min(data: list[float]) -> float:
    return min(data)

class Person(BaseModel):
   uuid: int
   name: str
   grade_1: float
   grade_2: float
   grade_3: float

people: list[Person] = []
averages: list[list[str,float]] = []
with open("MOCK_DATA.csv", "r") as file:
    data = file.readlines()
    for line in data:
        fields = line.split(",")
        if fields[0] == "id":
            continue
        people.append(
            Person(
                uuid=int(fields[0]), 
                name=fields[1] + " " + fields[2], 
                grade_1=float(fields[3]), 
                grade_2=float(fields[4]), 
                grade_3=float(fields[5])
                )
            )

        
for person in people:
    averages.append([person.name, (person.grade_1 + person.grade_2 + person.grade_3) / 3])

average_grades = [average[1] for average in averages]
mean = find_mean(average_grades)
median = find_median(average_grades)
mode = find_mode(average_grades)
maxi = find_max(average_grades)
mini = find_min(average_grades)

with open("results.txt", "w") as f:
    try:
        f.write(f"Mean: {mean}\n")
        f.write(f"Median: {median}\n")
        f.write(f"Mode: {mode}\n")
        f.write(f"Max: {maxi}\n")
        f.write(f"Min: {mini}\n")
    finally:
        print("Results have been written to results.txt")
        f.close()

        
        