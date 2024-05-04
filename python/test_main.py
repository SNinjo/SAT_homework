from time import time
from typing import List
from interface import Meat, Beef, Pork, Chicken
from main import main

def test_main(capsys):
    employee_ids = ['A', 'B']
    meat_list: List[Meat] = [Beef(), Beef(), Pork(), Chicken()]

    start_time = time()
    main(employee_ids, meat_list)
    execution_time = time() - start_time

    total_processing_seconds = sum([meat.get_processing_seconds() for meat in meat_list])
    assert total_processing_seconds > execution_time >= total_processing_seconds / len(employee_ids)
    assert capsys.readouterr().out.count('\n') == len(meat_list) * 2
