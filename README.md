# StrategyPattern.Go
Strategy Pattern Implementation and Strategy Choose

Strategy pattern sample
strategy.PetShop implements Strategy choose and instatiation, based on dynamic registrations of strategies

All strategies are at package "animals". Every strategy implements IAnimal interface.
At package initialization, each strategy (dog, cat and bird) call strategy.RegisterAnimal method.
strategy.RegisterAnimal method add strategy type and name to a map indexed by type name.
When PetShop struct receives a call to ChooseAnimal, it takes animalType param and finds at Animals map the strategy type.
After strategy type instantiation, ChooseAnimal method do a type assertion to IAnimal interface and call interface method Speech.
 
