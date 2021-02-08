(define-syntax flatten
  (syntax-rules ()
    ((_ xs) (let flatten1 ((ys xs))
                        (cond ((null? ys) '())
                              ((list? (car ys)) (append (flatten1 (car ys))
                                                        (flatten1 (cdr ys))))
                              (else (cons (car ys)
                                          (flatten1 (cdr ys)))))))))

(flatten '(1 (1 (1 (1 (1 1 1))))))
(flatten '(1 (1 (2 3)) 2 3 (3 2)))
(flatten '(1 2 3))
(flatten '(1 (2 3) 1 2 3))