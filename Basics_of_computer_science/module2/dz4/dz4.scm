;1

(define (memoized-factorial n)
  (let ((memo '()))
    (let ((memoized (assoc n memo)))
      (if (not (equal? memoized #f))
          (cadr memoized)
          (let ((new
                 (if (< n 1)
                     1
                     (* n (memoized-factorial (- n 1))))))
            (set! memo (cons (list n new) memo))
            new)))))
;TESTS
(begin
  (display (memoized-factorial 10)) (newline)
  (display (memoized-factorial 50)) (newline))
(newline)

;2 ------------------------------------------------------------ 

(define-syntax lazy-cons
  (syntax-rules ()
    ((_ a b)
     (cons a (delay b)))))

(define (lazy-car xs)
  (car xs))

(define (lazy-cdr xs)
  (force (cdr xs)))

(define (lazy-head xs k)
  (if (= 0 k)
      '()
      (cons (lazy-car xs) (lazy-head (lazy-cdr xs) (- k 1)))))

(define (naturals start)
  (lazy-cons start (naturals (+ start 1))))

(define (factorial n)
  (lazy-cons (fact n) (factorial (+ n 1))))

(define (fact a)
  (let ! ((n a))
    (if (= 0 n)
        1
        (* n (! (- n 1))))))  

(define (factorial1 n)
  (lazy-head (factorial 0) (+ n 1)))

(define (lazy-factorial n)
  (lazy-car (reverse (factorial1 n))))

;TESTS

;(display (lazy-head (naturals 10) 12)) (newline) 
;(10 11 12 13 14 15 16 17 18 19 20 21)
;(newline)

;(begin
  ;(display (lazy-factorial 10)) (newline)
  ;(display (lazy-factorial 50)) (newline))

;(newline)

;3 ------------------------------------------------------------

(define (read-words)
  (define (find words word ch)
    (begin
      (cond ((and (eof-object? ch)
                  (not (null? word))) (find (cons (list->string (reverse word)) words) '() (read-char)))
             ((and (eof-object? ch)
                  (not (null? words))) (reverse words))
            ((eof-object? ch) (reverse words))
            ((and (or (equal? ch #\tab)
                      (equal? ch #\newline)
                      (equal? ch #\space))
                  (null? word)) (find words word (read-char)))
            ((or (equal? ch #\tab)
                 (equal? ch #\newline)
                 (equal? ch #\space)) (find (cons (list->string (reverse word)) words) '() (read-char)))
            (else (find words (cons ch word) (read-char))))))
  (find '() '() (read-char)))

;TEST
(read-words)

;5 ------------------------------------------------------------------

(define (strings-conc . strings)
  (if (symbol? (apply string-append strings))
      (symbol->string (apply string-append strings))
      (string->symbol (apply string-append strings))))

(define-syntax struct-create-1
  (syntax-rules ()
    ((_ type exp ...) (begin
                        (eval '(define type (list 'exp ...))
                              (interaction-environment))
                        (eval (list 'define (list (strings-conc (symbol->string 'type) "?") 'x)
                                    '(and (list? x) (member (car x) type) (list? x)))
                              (interaction-environment))))))

(define-syntax struct-create-2
  (syntax-rules ()
    ((_ type exp ...) (eval '(define (type exp ...) (list 'type exp ...))
                            (interaction-environment)))))

(define-syntax define-data
  (syntax-rules ()
    ((_ type ((name expr ...) ...))
     (begin
       (struct-create-1 type name ...)
       (struct-create-2 name expr ...) ...))))
; Определяем тип
;
(define-data figure ((square a)
                     (rectangle a b)
                     (triangle a b c)
                     (circle r)))

; Определяем значения типа
;
(define s (square 10))
(define r (rectangle 10 20))
(define t (triangle 10 20 30))
(define c (circle 10))

; Пусть определение алгебраического типа вводит
; не только конструкторы, но и предикат этого типа:

(and (figure? s)
     (figure? r)
     (figure? t)
     (figure? c)) ;⇒ #t
(newline)


(define (find-match xs ys result)
  (if (null? xs)
      result
      (if (symbol? (car xs))
          (if (equal? (car xs) (car ys))
              (find-match (cdr xs) (cdr ys) result)
              #f)
          (find-match (cdr xs) (cdr ys) (cons (list (car ys) (car xs)) result)))))

(define (make-match exp pr)
  (let loop ((x pr))
    (if (not (null? x))
        (let ((x1 (caar x)) (ss (cadar x)))
          (let ((result (find-match exp x1 '())))
            (if result
                (eval `(let, result, ss) (interaction-environment))
                (loop (cdr x))))))))

(define-syntax match
  (syntax-rules ()
    ((_ exp (name act) ...)
     (make-match exp '((name act) ...)))))

(define pi (acos -1)) ; Для окружности
  
(define (perim f)
  (match f 
    ((square a)       (* 4 a))
    ((rectangle a b)  (* 2 (+ a b)))
    ((triangle a b c) (+ a b c))
    ((circle r)       (* 2 pi r))))
  
(perim s) ;⇒ 40
(perim r) ;⇒ 60
(perim t) ;⇒ 60
(newline)

;АЧИВКА

;(((call-with-current-continuation
;   (lambda (c) c)) (lambda (x) x))
; 'hello)

