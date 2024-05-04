from time import time
from typing import List
from freezegun import freeze_time
from queue import Queue
from interface import Meat
from employee import Employee, await_processing_meat

def toQueue(meat_list: List[Meat]) -> Queue[Meat]:
    queue = Queue[Meat]()
    for meat in meat_list:
        queue.put(meat)
    return queue

class TestEmployee:
    def test_get_current_time(self):
        time = '1970-01-01 00:00:00'
        with freeze_time(time):
            assert Employee('', toQueue([])).get_current_time() == time
            
        time = '2001-12-31 12:59:59'
        with freeze_time(time):
            assert Employee('', toQueue([])).get_current_time() == time

    @freeze_time('1970-01-01 00:00:00')
    def test_run_output(self, capsys):
        class Meat1(Meat):
            def get_name(self) -> str:
                return '肉1'
            def get_processing_seconds(self) -> int:
                return 0
        class Meat2(Meat):
            def get_name(self) -> str:
                return '肉2'
            def get_processing_seconds(self) -> int:
                return 0

        Employee('', toQueue([])).run()
        assert capsys.readouterr().out == ''

        Employee('ID1', toQueue([Meat1()])).run()
        assert capsys.readouterr().out == (
            'ID1 在 1970-01-01 00:00:00 取得肉1\n'
            'ID1 在 1970-01-01 00:00:00 處理完肉1\n'
        )
        
        Employee('ID2', toQueue([Meat1(), Meat2()])).run()
        assert capsys.readouterr().out == (
            'ID2 在 1970-01-01 00:00:00 取得肉1\n'
            'ID2 在 1970-01-01 00:00:00 處理完肉1\n'
            'ID2 在 1970-01-01 00:00:00 取得肉2\n'
            'ID2 在 1970-01-01 00:00:00 處理完肉2\n'
        )

    def test_run_processing_time(self):
        class Meat1(Meat):
            def get_name(self) -> str:
                return ''
            def get_processing_seconds(self) -> int:
                return 1
        class Meat2(Meat):
            def get_name(self) -> str:
                return ''
            def get_processing_seconds(self) -> int:
                return 2

        start_time = time()
        Employee('', toQueue([])).run()
        assert 0.1 > time() - start_time >= 0
        
        start_time = time()
        Employee('', toQueue([Meat1()])).run()
        assert 1.1 > time() - start_time >= 1
        
        start_time = time()
        Employee('', toQueue([Meat1(), Meat2()])).run()
        assert 3.1 > time() - start_time >= 3

@freeze_time('1970-01-01 00:00:00')
def test_await_processing_meat_output(capsys):
    class Meat1(Meat):
        def get_name(self) -> str:
            return ''
        def get_processing_seconds(self) -> int:
            return 0
    class Meat2(Meat):
        def get_name(self) -> str:
            return ''
        def get_processing_seconds(self) -> int:
            return 0

    await_processing_meat([])
    assert capsys.readouterr().out == ''

    meat_queue = toQueue([Meat1(), Meat2()])
    await_processing_meat([Employee('ID1', meat_queue), Employee('ID2', meat_queue)])
    assert capsys.readouterr().out.count('\n') == 2 * 2

def test_await_processing_meat_processing_time():
    class Meat1(Meat):
        def get_name(self) -> str:
            return ''
        def get_processing_seconds(self) -> int:
            return 1
    class Meat2(Meat):
        def get_name(self) -> str:
            return ''
        def get_processing_seconds(self) -> int:
            return 2

    start_time = time()
    await_processing_meat([])
    assert 0.1 > time() - start_time >= 0
    
    start_time = time()
    meat_queue = toQueue([Meat1(), Meat2()])
    await_processing_meat([Employee('ID1', meat_queue), Employee('ID2', meat_queue)])
    assert 3 > time() - start_time >= 3 / 2
