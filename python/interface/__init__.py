from abc import ABC, abstractmethod

class Meat(ABC):
    @abstractmethod
    def get_name(self) -> str:
        pass

    @abstractmethod
    def get_processing_seconds(self) -> int:
        pass

class Beef(Meat):
    def get_name(self) -> str:
        return '牛肉'

    def get_processing_seconds(self) -> int:
        return 1
    
class Pork(Meat):
    def get_name(self) -> str:
        return '豬肉'
    
    def get_processing_seconds(self) -> int:
        return 2
    
class Chicken(Meat):
    def get_name(self) -> str:
        return '雞肉'
    
    def get_processing_seconds(self) -> int:
        return 3
