zero = f"str(sum([bool([])]))"

one = f"str(sum([int(not bool([]))]))"

number = lambda x: str(sum([int(not bool([])) for i in range(x)]))
getDict = lambda x: {str(x): x}


fromString = lambda x: " + ".join([f"{mapping[i]}" for i in x])

mapping = {}


mapping["d"] = f"str(type({getDict(zero)}))[{int(number(8))}]"
mapping["i"] = f"str(type({getDict(zero)}))[{int(number(9))}]"
mapping["r"] = f"str(type({getDict(zero)[str(zero)]}))[{int(number(10))}]"

#abcedfghjklmnopqstuvwxyz


print(eval(f"{mapping['d']}"))






