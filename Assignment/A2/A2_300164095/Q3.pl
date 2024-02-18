%Initialize base list
sum_odd_numbers([], 0).
%If Even run again
sum_odd_numbers([Odd|Even], Sum) :-
    Odd mod 2 =:= 0, 
    sum_odd_numbers(Even, Sum).
%If odd add to sum
sum_odd_numbers([Odd|Even], Sum) :-
    Odd mod 2 =:= 1, 
    sum_odd_numbers(Even, Rest),
    Sum is Odd + Rest.