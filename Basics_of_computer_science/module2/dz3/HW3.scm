(load "tests.scm")

;Само задание
(define (cut-first xs x) ; -top
  (if (null? x)
      xs
      (cut-first (append xs (cons (car x) '()))
                 (cdr x))))
  (define (sum-build a b) ; представление суммы
    (cond ((equal? a 0) b)
          ((equal? b 0) a)
          (else (list '+ a b))))

  (define (difference-build a b) ; представление разности
    (cond ((equal? a 0) (list '* -1 b))
          ((equal? b 0) a)
          (else (list '- a b))))

  (define (prod-build a b) ;  представление произведения
    (cond ((or (equal? a 0) (equal? b 0)) 0)
          ((equal? a 1) b)
          ((equal? b 1) a)
          (else (list '* a b))))

  (define (sum? exp) ;проверка на сумму
    (eq? (car exp) '+))

  (define (prod? exp);проверка на проиизведение
    (and (eq? (car exp) '*) (equal? (length exp) 3)))

  (define (multiprod? exp) ; множ произведение
    (and (eq? (car exp) '*) (> (length exp) 3)))

  (define (difference? exp) ;разность?
    (and (equal? (car exp) '-)))

  (define (var? exp) ;проверка на символ
    (symbol? exp))

  (define (longexpr? exp) 
    (and (not (symbol? exp)) (not (< (length exp) 3)) (number? (cadr exp)) (number? (caddr exp))))

  (define (sin? exp) ;синус?
    (eq? (car exp) 'sin))

  (define (cos? exp);косинус?
    (eq? (car exp) 'cos))

  (define (expt? ex) ;степень с основанием х?
    (and (eq? (car ex) 'expt) (var? (cadr ex)) (not (eq? (cadr ex) 'exp))))

  (define (expot? ex);число в степени икс?
    (and (eq? (car ex) 'expt) (not (eq? (cadr ex) 'exp)) (var? (caddr ex)))) ;!!!

  (define (exp? ex);экспонента
    (and (eq? (car ex) 'e)))

  (define (ln? exp);логарифм
    (eq? (car exp) 'ln))
;2 случая + воссоздание
  (define (segmentation-1? xs) 
    (and (eq? (car xs) '/) (var? (caddr xs))))

  (define (segmentation-2? xs)
    (and (eq? (car xs) '/) (not (var? (caddr xs))) (not (number? (caddr xs)))))

  (define (reforge xs x)
    (cond ((null? x) xs)
          ((number? (car x)) (reforge (append xs (cons (list '/ 1 (car x)) '())) (cdr x)))
          ((symbol? (car x)) (reforge (append xs (cons (car x) '())) (cdr x)))
          ((eq? (car (car x)) 'expt) (reforge (append xs (cons (list 'expt (cadr (car x)) (* -1 (caddr (car x)))) '())) (cdr x)))))

(define (derivative a) ; само взятие поизводной (derivative)
  (cond ((number? a) 0)
        ((longexpr? a) 0)
        ((var? a) (if (equal? (car (string->list (symbol->string a))) #\-)
                      -1
                      1))
        ((sum? a) (sum-build (derivative (cadr a))
                             (derivative (caddr a))))
        ((prod? a) (sum-build (prod-build (derivative (cadr a))
                                          (caddr a))
                              (prod-build (cadr a)
                                          (derivative (caddr a)))))
        ((difference? a) (difference-build (derivative (cadr a))
                                           (derivative (caddr a))))
        ((sin? a) (prod-build (derivative (cadr a))
                              (list 'cos (cadr a))))
        ((cos? a) (prod-build (derivative (cadr a))
                              (list '* -1 (list 'sin (cadr a)))))
        ((expt? a) (prod-build (derivative (cadr a))
                               (list '* (caddr a) (list 'expt (cadr a) (- (caddr a) 1)))))
        ((expot? a) (prod-build (derivative (caddr a))
                                (list '* a (list 'ln (cadr a)))))
        ((exp? a) (prod-build (derivative (cadr a)) a))
        ((ln? a) (prod-build (derivative (cadr a)) (list '/ 1 (cadr a))))
        ((multiprod? a) (sum-build (prod-build (derivative (cadr a)) (cut-first (list '*) (cdr (cdr a))))
                                    (prod-build (cadr a) (derivative (cut-first (list '*) (cdr (cdr a)))))))
        ((segmentation-1? a) (if (equal? (cadr a) 1)
                                 (list '* -1 (list 'expt (caddr a) -1))
                                 (list '* (cadr a) (list '* -1 (list 'expt (caddr a) -1)))))
        ((segmentation-2? a) (list '* (cadr a) (derivative (reforge '() (caddr a)))))
        ))
                              
    
;тесты для программы
(define the-tests (list
                   (test (derivative 2) 0)
                   (test (derivative 'x) 1)
                   (test (derivative '-x) -1)
                   (test (derivative '(* 1 x)) 1)
                   (test (derivative '(* -1 x)) -1)
                   (test (derivative '(* -4 x)) -4)
                   (test (derivative '(* 10 x)) 10)
                   (test (derivative '(- (* 2 x) 3)) 2)
                   (test (derivative '(expt x 10)) '(* 10 (expt x 9)))
                   (test (derivative '(* 2 (expt x 5))) '(* 2 (* 5 (expt x 4))))
                   (test (derivative '(expt x -2)) '(* -2 (expt x -3)))
                   (test (derivative '(expt 5 x)) '(* (expt 5 x) (ln 5)))
                   (test (derivative '(cos x)) '(* -1 (sin x)))
                   (test (derivative '(sin x)) '(cos x))
                   (test (derivative '(e x)) '(e x))
                   (test (derivative '(* 2 (e x))) '(* 2 (e x)))
                   (test (derivative '(* 2 (e (* 2 x)))) '(* 2 (* 2 (e (* 2 x)))))
                   (test (derivative '(ln x)) '(/ 1 x))
                   (test (derivative '(* 3 (ln x))) '(* 3 (/ 1 x)))
                   (test (derivative '(+ (expt x 3) (expt x 2))) '(+ (* 3 (expt x 2)) (* 2 (expt x 1))))
                   (test (derivative '(- (* 2 (expt x 3)) (* 2 (expt x 2)))) '(- (* 2 (* 3 (expt x 2))) (* 2 (* 2 (expt x 1)))))
                   (test (derivative '(/ 3 x)) '(* 3 (* -1 (expt x -1))))
                   (test (derivative '(/ 3 (* 2 (expt x 2)))) '(* 3 (* (/ 1 2) (* -2 (expt x -3)))))
                   (test (derivative '(* 2 (sin x) (cos x))) '(* 2 (+ (* (cos x) (cos x)) (* (sin x) (* -1 (sin x))))))
                   (test (derivative '(* 2 (e x) (sin x) (cos x))) '(* 2 (+ (* (e x) (* (sin x) (cos x))) (* (e x) (+ (* (cos x) (cos x)) (* (sin x) (* -1 (sin x))))))))
                   (test (derivative '(sin (* 2 x))) '(* 2 (cos (* 2 x))))
                   (test (derivative '(sin (ln (expt x 2)))) '(* (* (* 2 (expt x 1)) (/ 1 (expt x 2))) (cos (ln (expt x 2)))))
                   (test (derivative '(cos (* 2 (expt x 2)))) '(* (* 2 (* 2 (expt x 1))) (* -1 (sin (* 2 (expt x 2))))))
                   (test (derivative '(+ (sin (* 2 x)) (cos (* 2 (expt x 2))))) '(+ (* 2 (cos (* 2 x))) (* (* 2 (* 2 (expt x 1))) (* -1 (sin (* 2 (expt x 2)))))))
                   (test (derivative '(* (sin (* 2 x)) (cos (* 2 (expt x 2))))) '(+ (* (* 2 (cos (* 2 x))) (cos (* 2 (expt x 2)))) (* (sin (* 2 x)) (* (* 2 (* 2 (expt x 1))) (* -1 (sin (* 2 (expt x 2))))))))))
(run-tests the-tests)