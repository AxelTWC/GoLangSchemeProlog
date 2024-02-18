%Given
pet(fido, dog, 3).
pet(spot, dog, 5).
pet(mittens, cat, 2).
pet(tweety, bird, 1).
male(fido).
male(spot).
female(mittens).
%Self Function
%pet(Name, Species, Age) :-
	%pet(Name, Species, Age).
%species(Species,Count)
species(Species,Count) :- 
	findall(Name, pet(Name, Species, _), List), length(List, Count).
%age_range(MinAge,MaxAge,Count)
age_range(MinAge,MaxAge,Count) :- 
	findall(Name, (pet(Name, _, Age), Age >= MinAge, Age =< MaxAge), List), length(List, Count).
