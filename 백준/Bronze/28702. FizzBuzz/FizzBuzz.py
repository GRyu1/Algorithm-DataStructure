index_value = 0
index = 0

for i in range(3):
    data = input().strip()
    try:
        value = int(data)
        index_value = value
        index = i
    except ValueError:
        pass

index_value += 4 - (index + 1)

if index_value % 3 == 0 and index_value % 5 == 0:
    print("FizzBuzz")
elif index_value % 3 == 0:
    print("Fizz")
elif index_value % 5 == 0:
    print("Buzz")
else:
    print(index_value)