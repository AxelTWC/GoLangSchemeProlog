(define sudoku1 '((2 1 4 3) (4 3 2 1) (1 2 3 4) (3 4 1 2)))
(define sudoku2 '((2 1 4 3) (4 3 2 1) (1 2 3 3) (3 4 1 2)))

(define (different lst)
  (cond ((null? lst) #t) 
        ((member (car lst) (cdr lst)) #f)
        (else (different (cdr lst)))))
(different '(1 3 6 4 8 0))
(different '(1 3 6 4 1 8 0))

(define (extract4Columns sudoku)
  (list (list(car(car sudoku)) (car(car(cdr sudoku))) (car(car(cdr(cdr sudoku)))) (car(car(cdr(cdr(cdr sudoku))))))
        (list (car(cdr(car sudoku))) (car(cdr(car(cdr sudoku)))) (car(cdr(car(cdr(cdr sudoku))))) (car(cdr(car(cdr(cdr(cdr sudoku)))))))
        (list (car(cdr(cdr(car sudoku)))) (car(cdr(cdr(car(cdr sudoku))))) (car(cdr(cdr(car(cdr(cdr sudoku)))))) (car(cdr(cdr(car(cdr(cdr(cdr sudoku))))))))
        (list (car(cdr(cdr(cdr(car sudoku))))) (car(cdr(cdr(cdr(car(cdr sudoku)))))) (car(cdr(cdr(cdr(car(cdr(cdr sudoku))))))) (car(cdr(cdr(cdr(car(cdr(cdr(cdr sudoku)))))))))))

(extract4Columns sudoku1)

(define (extract4Quadrants sudoku)
  (list
  (list (car(car sudoku)) (car(cdr(car sudoku))) (car(car(cdr sudoku))) (car(cdr(car(cdr sudoku)))))
  (list (car(cdr(cdr(car sudoku)))) (car(cdr(cdr(cdr(car sudoku))))) (car(cdr(cdr(car(cdr sudoku))))) (car(cdr(cdr(cdr(car(cdr sudoku)))))))
  (list (car(car(cdr(cdr sudoku)))) (car(cdr(car(cdr(cdr sudoku))))) (car(car(cdr(cdr(cdr sudoku))))) (car(cdr(car(cdr(cdr(cdr sudoku)))))))
  (list (car(cdr(cdr(car(cdr(cdr sudoku)))))) (car(cdr(cdr(cdr(car(cdr(cdr sudoku))))))) (car(cdr(cdr(car(cdr(cdr(cdr sudoku))))))) (car(cdr(cdr(cdr(car(cdr(cdr(cdr sudoku)))))))))))
(extract4Quadrants sudoku1)

(define (merge3 lst lst2 lst3)
  (append lst lst2 lst3))
(merge3 '(1 3 6) '(5 4) '(1 2 3))

(define (allTrue lst)
  (cond ((null? lst) #t) 
        ((not (boolean? (car lst))) #f)
        ((not (car lst)) #f) 
        (else (allTrue (cdr lst)))))
  
(define (checkSudoku sudoku)
  (allTrue (map different (merge3 sudoku (extract4Columns sudoku) (extract4Quadrants sudoku))))
  )
(checkSudoku sudoku1)
(checkSudoku sudoku2)