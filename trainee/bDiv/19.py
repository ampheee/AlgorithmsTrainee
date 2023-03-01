class Heap:
    def __init__(self):
        self.heap = []

    def insert(self, value):
        self.heap.append(int(value))
        self.sift_up()

    def extract(self):
        max_value = self.heap[0]  # store max value
        self.heap[0] = self.heap[-1]
        self.heap.pop()
        self.sift_down()
        return max_value

    def sift_up(self):
        index = len(self.heap) - 1
        index_parent = (index - 1) // 2  # for index 'i' of a parent, children indexes
        while self.heap[index_parent] < self.heap[index]:  # works also for only root heap
            self.heap[index_parent], self.heap[index] = self.heap[index], self.heap[index_parent]
            index = index_parent
            index_parent = max((index - 1) // 2, 0)  # for 2-element heap index_parent becomes -1

    def sift_down(self):
        index_parent = 0
        while index_parent * 2 + 1 < len(self.heap):  # check whether any children exist
            if 2 * index_parent + 2 == len(self.heap):  # only one child
                max_index = 2 * index_parent + 1
            elif self.heap[2 * index_parent + 1] > self.heap[2 * index_parent + 2]:  # left child is greater
                max_index = 2 * index_parent + 1
            else:  # right child is greater or equal
                max_index = 2 * index_parent + 2

            if self.heap[index_parent] < self.heap[max_index]:  # swap only if child is greater than parent
                self.heap[index_parent], self.heap[max_index] = self.heap[max_index], self.heap[index_parent]
            index_parent = max_index  # update index

total = int(input())
heap = Heap()
for i in range(total):
    command = input().split()
    match command[0]:
        case '1':
            print(heap.extract())
        case '0':
            heap.insert(command[1])