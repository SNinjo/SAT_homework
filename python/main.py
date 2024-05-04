import random
from typing import List
from queue import Queue
from employee import Employee, await_processing_meat
from interface import Meat, Beef, Pork, Chicken

def main(employee_ids: List[str], meat_list: List[Meat]):
    meat_queue = Queue[Meat]()
    for meat in meat_list:
        meat_queue.put(meat)

    employees = [Employee(employee_id, meat_queue) for employee_id in employee_ids]
    await_processing_meat(employees)

if __name__ == '__main__':
    meat_list: List[Meat] = []
    meat_list += [Beef()] * 10
    meat_list += [Pork()] * 7
    meat_list += [Chicken()] * 5
    # 透過打亂肉品的順序來達到隨機選取的效果
    random.shuffle(meat_list)
    main(['A', 'B', 'C', 'D', 'E'], meat_list)
