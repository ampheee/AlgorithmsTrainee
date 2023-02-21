class Deque:
    def __init__(self):
        self.deque = []

    def __str__(self):
        return f"{self.deque}"

    def push_front(self, value):
        self.deque.insert(0, value)
        return "ok"

    def push_back(self, value):
        self.deque.append(value)
        return "ok"

    def pop_front(self):
        if len(self.deque) != 0:
            return self.deque.pop(0)
        else:
            return "error"
            
    def size(self):
        return len(self.deque)

    def clear(self):
        self.deque.clear()
        return "ok"

    def pop_back(self):
        if len(self.deque) != 0:
            return self.deque.pop()
        else:
            return "error"

    def front(self):
        if len(self.deque) != 0:
            return self.deque[0]
        else:
            return "error"

    def back(self):
        if len(self.deque) != 0:
            return self.deque[-1]
        else:
            return "error"


deque = Deque()

while True:
    command = input().split()
    match command[0]:
        case "push_back":
            print(deque.push_back(command[1]))
        case "push_front":
            print(deque.push_front(command[1]))
        case "pop_back":
            print(deque.pop_back())
        case "pop_front":
            print(deque.pop_front())
        case "back":
            print(deque.back())
        case "front":
            print(deque.front())
        case "size":
            print(deque.size())
        case "clear":
            print(deque.clear())
        case "exit":
            print("bye")
            break

