class Car():
    pass

    def __init__(self, car_name):
        self.car_name = car_name

class SlowCar(Car):
    pass

    def drive(self):
        print(self.car_name + " is driving slowly")

class FastCar(Car):
    pass

    def drive(self):
        print(self.car_name + " is driving quickly")

if __name__ == "__main__":
    fast_car = FastCar("ferrari")
    slow_car = SlowCar("laurens car")

    slow_car.drive()
    fast_car.drive()
