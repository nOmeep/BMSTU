(define (interpret program stack)
  (define (to-end counter)
    (if (eq? (vector-ref program counter) 'end)
        counter
        (to-end (+ counter 1))))
  (define (in-if prog count st res-stack tmp mark)
    (let ((cur (vector-ref progog count)))
      (cond
        ((eq? cur 'if) (in-if prog (+ 1 count) st res-stack tmp (cons 1 mark)))
        ((eq? cur 'endif) (if (= (length mark) 1)
                              (false-if prog
                                        (+ 1 count)
                                        st
                                        res-stack
                                        tmp)
                              (in-if prog
                                     (+ 1 count)
                                     st
                                     res-stack
                                     tmp
                                     (cdr mark))))
        (else (in-if prog (+ 1 count) st res-stack tmp mark)))))
      
  (define (false-if prog count st res-stack tmp)
    (let ((cur (vector-ref prog count)))
      (cond
        ((eq? cur 'if) (in-if prog
                              (+ 1 count)
                              st
                              res-stack
                              tmp
                              '(1)))
        ((eq? cur 'endif) (if (eq? (vector-ref prog (+ 1 count)) 'else)
                              (main prog
                                    (+ 2 count)
                                    st
                                    res-stack
                                    tmp)
                              (main prog
                                    (+ 1 count)
                                    st
                                    res-stack
                                    tmp)))
        (else (false-if prog
                        (+ 1 count)
                        st
                        res-stack
                        tmp)))))
        
  (define (false-while prog count st res-stack tmp)
    (let ((cur (vector-ref prog count)))
      (if (eq? cur 'end-while)
          (main prog
                (+ 1 count)
                st
                res-stack
                tmp)
          (false-while prog
                       (+ 1 count)
                       st
                       res-stack
                       tmp))))
  (define (to-end-if prog count st res-stack tmp)
    (let ((cur (vector-ref prog count)))
      (if (eq? cur 'end-else)
          (main prog
                (+ 1 count)
                st
                res-stack
                tmp)
          (to-end-if prog
                     (+ 1 count)
                     st
                     res-stack
                     tmp))))
  (define (break prog count st res-stack tmp)
    (let ((cur (vector-ref prog count)))
      (cond
        ((or (eq? cur 'end-while)
             (eq? cur 'end-repeat)) (main prog
                                          (+ 1 count)
                                          st
                                          res-stack
                                          tmp))
        ((eq? cur 'loop) (main prog
                               (+ 1 count)
                               st
                               res-stack
                               (cddr tmp)))
        (else (break prog
                     (+ 1 count)
                     st
                     res-stack
                     tmp)))))
  (define (continue prog count st res-stack tmp)
    (let ((cur (vector-ref prog count)))
      (cond
        ((or (eq? cur 'end-while)
             (eq? cur 'end-repeat)) (main prog
                                          (car res-stack)
                                          st
                                          (cdr res-stack)
                                          tmp))
        ((eq? cur 'loop) (main prog
                               (car res-stack)
                               (cdr stack)
                               (cdr res-stack)
                               (cons (+ (car tmp) (car stack)) (cdr tmp))))
        (else (continue prog
                        (+ 1 count)
                        st
                        res-stack
                        tmp)))))
  (define (switch prog count st res-stack tmp key)
    (let ((cur (vector-ref prog count)))
      (cond
        ((eq? cur 'case) (if (= (vector-ref prog (+ 1 count)) key)
                             (main prog
                                   (+ 2 count)
                                   st
                                   res-stack
                                   tmp)
                             (switch prog
                                     (+ 1 count)
                                     st
                                     res-stack
                                     tmp)))
        ((eq? cur 'end-switch) (main prog
                                     (+ 1 count)
                                     st
                                     (cdr res-stack)
                                     tmp))
        (else (switch prog
                      (+ 1 count)
                      st
                      res-stack
                      tmp)))))
  (define (find-sw prog count)
    (let ((cur (vector-ref prog count)))
      (if (eq? cur 'end-switch)
          (+ 1 count)
          (find-sw prog
                   (+ 1 count)))))
  (define (main prog pos st res-stack tmp)
    (if (>= pos (vector-length prog))
        st
        (let ((cur (vector-ref prog pos)))
          (cond
            ((number? cur) (main prog
                                 (+ 1 pos)
                                 (cons cur st)
                                 res-stack
                                 tmp))
            ((eq? cur '+) (main prog
                                (+ 1 pos)
                                (cons (+ (car st) (cadr st)) (cddr st))
                                res-stack
                                tmp))
            ((eq? cur '-) (main prog
                                (+ 1 pos)
                                (cons (- (cadr st) (car st)) (cddr st))
                                res-stack
                                tmp))
            ((eq? cur '*) (main prog
                                (+ 1 pos)
                                (cons (* (cadr st) (car st)) (cddr st))
                                res-stack
                                tmp))
            ((eq? cur '/) (main prog
                                (+ 1 pos)
                                (cons (quotient (cadr st) (car st)) (cddr st))
                                res-stack
                                tmp))
            ((eq? cur 'mod) (main prog
                                  (+ 1 pos)
                                  (cons (remainder (cadr st) (car st)) (cddr st))
                                  res-stack
                                  tmp))
            ((eq? cur 'neg) (main prog
                                  (+ 1 pos)
                                  (cons (* -1 (car st)) (cdr st))
                                  res-stack
                                  tmp))
            ((eq? cur '=) (main prog (+ pos 1)
                                (cons (if (= (cadr st) (car st))
                                          -1
                                          0)
                                      (cddr st))
                                res-stack
                                tmp))
            ((eq? cur '>) (main prog (+ pos 1)
                                (cons (if (> (cadr st) (car st))
                                          -1
                                          0)
                                      (cddr st))
                                res-stack
                                tmp))
            ((eq? cur '<) (main prog (+ pos 1)
                                (cons (if (< (cadr st) (car st))
                                          -1
                                          0)
                                      (cddr st))
                                res-stack
                                tmp))
            ((eq? cur 'not) (main prog (+ pos 1)
                                  (cons (if (= (car st) 0)
                                            -1
                                            0)
                                        (cdr st))
                                  res-stack
                                  tmp))
            ((eq? cur 'and) (main prog (+ pos 1)
                                  (cons (if (and (= (cadr st) (car st))
                                                 (= (car st) -1))
                                            -1
                                            0)
                                        (cddr stack))
                                  res-stack
                                  tmp))
            ((eq? cur 'or) (main prog (+ pos 1) (cons (if (or (= (cadr st) -1)
                                                              (= (car st) -1))
                                                          -1
                                                          0)
                                                      (cddr st))
                                 res-stack
                                 tmp))
            ((eq? cur 'drop) (main prog
                                   (+ 1 pos)
                                   (cdr st)
                                   res-stack
                                   tmp))
            ((eq? cur 'swap) (main prog
                                   (+ 1 pos)
                                   (cons (cadr st) (cons (car st) (cddr st)))
                                   res-stack
                                   tmp))
            ((eq? cur 'dup) (main prog
                                  (+ 1 pos)
                                  (cons (car st) st)
                                  res-stack
                                  tmp))
            ((eq? cur 'over) (main prog
                                   (+ 1 pos)
                                   (cons (cadr st) st)
                                   res-stack
                                   tmp))
            ((eq? cur 'rot) (main prog
                                  (+ 1 pos)
                                  (cons (caddr st) (cons (cadr st) (cons (car st) (cdddr st))))
                                  res-stack
                                  tmp))
            ((eq? cur 'depth) (main prog
                                    (+ 1 pos)
                                    (cons (length st) st)
                                    res-stack
                                    tmp))
            ((eq? cur 'define) (main prog
                                     (+ (to-end pos) 1)
                                     st
                                     res-stack
                                     (cons (list (vector-ref prog (+ pos 1)) (+ pos 2)) tmp)))
            ((or (eq? cur 'end) (eq? cur 'exit)) (main prog
                                                       (car res-stack)
                                                       st
                                                       (cdr res-stack)
                                                       tmp))
            ((eq? cur 'if) (if (= (car st) 0)
                               (false-if prog (+ 1 pos) (cdr st) res-stack tmp)
                               (main prog
                                     (+ 1 pos)
                                     (cdr st)
                                     res-stack
                                     tmp)))
            ((eq? cur 'endif) (main prog
                                    (+ 1 pos)
                                    st
                                    res-stack
                                    tmp))
            ((eq? cur 'while) (if (= (car st) 0)
                                  (false-while prog (+ 1 pos) (cdr st) res-stack tmp)
                                  (main prog
                                        (+ 1 pos)
                                        (cdr st)
                                        (cons pos res-stack)
                                        tmp)))
            ((eq? cur 'end-while) (main prog
                                        (car res-stack)
                                        st
                                        (cdr res-stack)
                                        tmp))
            ((eq? cur 'repeat) (main prog
                                     (+ 1 pos)
                                     st
                                     (cons pos res-stack)
                                     tmp))
            ((eq? cur 'end-repeat) (if (= (car st) 0)
                                       (main prog
                                             (+ 1 pos)
                                             (cdr st)
                                             (cdr res-stack)
                                             tmp)
                                       (main prog
                                             (car res-stack)
                                             (cdr st)
                                             (cdr res-stack)
                                             tmp)))
            ((eq? cur 'do) (main prog
                                 (+ 1 pos)
                                 (cddr st)
                                 (cons (+ 1 pos) res-stack)
                                 (cons (car st) (cons (cadr st) tmp))))
            ((eq? cur 'loop) (if (> (+ (car tmp) (car stack)) (cadr tmp))
                                 (main prog
                                       (+ 1 pos)
                                       (cdr st)
                                       (cdr res-stack)
                                       (cddr tmp))
                                 (main prog
                                       (car res-stack)
                                       (cdr st)
                                       res-stack
                                       (cons (+ (car tmp) (car stack)) (cdr tmp)))))
            ((eq? cur 'else) (to-end-if prog (+ 1 pos) st res-stack tmp))
            ((eq? cur 'end-else) (main prog
                                       (+ 1 pos)
                                       st
                                       res-stack
                                       tmp))
            ((eq? cur 'break) (break prog
                                     (+ 1 pos)
                                     st
                                     (cdr res-stack)
                                     tmp))
            ((eq? cur 'continue) (continue prog
                                           (+ 1 pos)
                                           st
                                           res-stack
                                           tmp))
            ((eq? cur 'switch) (switch prog
                                       (+ 1 pos)
                                       (cdr st)
                                       (cons (find-sw prog pos) res-stack)
                                       tmp
                                       (car st)))
            ((assoc cur tmp) (main prog
                                   (cadr (assoc cur tmp))
                                   st
                                   (cons (+ 1 pos) res-stack)
                                   tmp))))))
  (main program 0 stack '() '()))

;TESTS
(load "tests.scm")
(define the-tests (list
                   (test (interpret #(define abs 
                                       dup 0 < 
                                       if neg endif 
                                       end 
                                       9 abs 
                                       -9 abs      ) (quote ())) '(9 9))
                   (test (interpret #(   define =0? dup 0 = end 
                                          define <0? dup 0 < end 
                                          define signum 
                                          =0? if exit endif 
                                          <0? if drop -1 exit endif 
                                          drop 
                                          1 
                                          end 
                                          0 signum 
                                          -5 signum 
                                          10 signum       ) (quote ())) '(1 -1 0))
                   (test (interpret #(   define -- 1 - end 
                                          define =0? dup 0 = end 
                                          define =1? dup 1 = end 
                                          define factorial 
                                          =0? if drop 1 exit endif 
                                          =1? if drop 1 exit endif 
                                          dup -- 
                                          factorial 
                                          * 
                                          end
                                          0 factorial 
                                          1 factorial 
                                          2 factorial 
                                          3 factorial 
                                          4 factorial     ) (quote ())) '(24 6 2 1 1))
                   (test (interpret #(   define =0? dup 0 = end 
                                          define =1? dup 1 = end 
                                          define -- 1 - end 
                                          define fib 
                                          =0? if drop 0 exit endif 
                                          =1? if drop 1 exit endif 
                                          -- dup 
                                          -- fib 
                                          swap fib 
                                          + 
                                          end 
                                          define make-fib 
                                          dup 0 < if drop exit endif 
                                          dup fib 
                                          swap -- 
                                          make-fib 
                                          end 
                                          10 make-fib     ) (quote ())) '(0 1 1 2 3 5 8 13 21 34 55))
                   (test (interpret #(   define =0? dup 0 = end 
                                          define gcd 
                                          =0? if drop exit endif 
                                          swap over mod 
                                          gcd 
                                          end 
                                          90 99 gcd 
                                          234 8100 gcd    ) '()) '(18 9))
                   (test (interpret #(do loop) '(1 5 1 1 1 1 1 0 0)) '(0 0))
                   (test (interpret #( while + dup end-while) '(1 1 2 -3 1 2 3)) '(0 1 2 3))
                   (test (interpret #( repeat + dup end-repeat) '(1 2 -3 1 2 3)) '(0 1 2 3))
                   (test (interpret #(0 if 2 endif else 7 end-else) '()) '(7))
                   (test (interpret #(1 while 1 1 1 break 1 1 end-while) '()) '(1 1 1))
                   (test (interpret #(1 if 0 if 2 endif endif 1) '()) '(1))))
(run-tests the-tests)
       
