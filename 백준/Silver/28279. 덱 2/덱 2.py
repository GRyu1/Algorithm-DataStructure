import sys

class Deque:
    def __init__(self, capacity):
        self.deque = [0] * capacity
        self.size = 0
        self.head = 0
        self.tail = 0
        self.capacity = capacity

    def add_first(self, value):
        if self.size == 0:
            self.deque[0] = value
            self.head = 0
            self.tail = 0
        else:
            self.head = self.decrease(self.head)
            self.deque[self.head] = value
        self.size += 1

    def add_last(self, value):
        if self.size == 0:
            self.deque[0] = value
            self.head = 0
            self.tail = 0
        else:
            self.tail = self.increase(self.tail)
            self.deque[self.tail] = value
        self.size += 1

    def peek_first(self):
        if self.size > 0:
            return self.deque[self.head]
        return -1

    def peek_last(self):
        if self.size > 0:
            return self.deque[self.tail]
        return -1

    def poll_first(self):
        if self.size > 0:
            value = self.deque[self.head]
            self.head = self.increase(self.head)
            self.size -= 1
            return value
        return -1

    def poll_last(self):
        if self.size > 0:
            value = self.deque[self.tail]
            self.tail = self.decrease(self.tail)
            self.size -= 1
            return value
        return -1

    def get_size(self):
        return self.size

    def increase(self, pointer):
        if pointer == self.capacity - 1:
            return 0
        return pointer + 1

    def decrease(self, pointer):
        if pointer == 0:
            return self.capacity - 1
        return pointer - 1

input = sys.stdin.read
data = input().splitlines()

count = int(data[0])
deque = Deque(count * 2)
output = []

for i in range(1, count + 1):
    tokens = data[i].split()
    command = int(tokens[0])

    if command == 1:
        value = int(tokens[1])
        deque.add_first(value)
    elif command == 2:
        value = int(tokens[1])
        deque.add_last(value)
    elif command == 3:
        output.append(deque.poll_first())
    elif command == 4:
        output.append(deque.poll_last())
    elif command == 5:
        output.append(deque.get_size())
    elif command == 6:
        output.append(1 if deque.get_size() == 0 else 0)
    elif command == 7:
        output.append(deque.peek_first())
    elif command == 8:
        output.append(deque.peek_last())

sys.stdout.write("\n".join(map(str, output)) + "\n")
