from datetime import datetime
from time import sleep
from threading import Thread
from typing import List
from queue import Queue
from interface import Meat

class Employee(Thread):
    def __init__(self, id: str, meat_queue: Queue[Meat]) -> None:
        Thread.__init__(self)
        self._id = id
        self._meat_queue = meat_queue

    def get_current_time(self) -> str:
        return datetime.now().strftime('%Y-%m-%d %H:%M:%S')

    def run(self) -> None:
        while self._meat_queue.qsize() > 0:
            meat = self._meat_queue.get()
            print(f'{self._id} 在 {self.get_current_time()} 取得{meat.get_name()}')
            sleep(meat.get_processing_seconds())
            print(f'{self._id} 在 {self.get_current_time()} 處理完{meat.get_name()}')

def await_processing_meat(employees: List[Employee]) -> None:
    for employee in employees:
        employee.start()
    for employee in employees:
        employee.join()
