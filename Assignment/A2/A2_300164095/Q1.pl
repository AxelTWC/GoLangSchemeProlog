%Given
parent(john, mary).
parent(john, tom).
parent(mary, ann).
parent(mary, fred).
parent(tom, liz).
parent(fred,able).
male(john).
male(tom).
male(fred).
female(mary).
female(ann).
female(liz).
%Self-Function
%parent(X,Y) :- 
    %	parent(X,Y).
%Sibling
sibling(X, Y) :- 
    parent(Z, X), parent(Z, Y), X \= Y.
%GrandParent
grandparent(X, Y) :- 
    parent(X, Z), parent(Z, Y).

%Great GrandParent
ancestor(X, Y) :- 
    parent(X, Z), parent(Z, A), parent(A, Y).