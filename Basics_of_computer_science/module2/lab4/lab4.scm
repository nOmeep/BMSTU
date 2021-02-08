;№1

(define reciever #f)
(define call/cc call-with-current-continuation)
(define (use-assertions)
  (call/cc (lambda (escape) (begin
                              (set! reciever escape)))))
(define-syntax assert
  (syntax-rules ()
    ((_ expr) (and (not expr) (begin
                                (display "FAILED: ")
                                (display (quote expr))
                                (newline)
                                (reciever))))))
;Tests
(use-assertions)
(define (1/x x)
  (assert (not (zero? x))) ; Утверждение: x ДОЛЖЕН БЫТЬ > 0
  (/ 1 x))
(map 1/x '(1 2 3 0 5)) 
(map 1/x '(-2 -1 0 1 2))
(map 1/x '(-2 -2 -2 -2))
(map 1/x '(0 0 0 0 0))
(newline)
;№4

(define-syntax my-if
  (syntax-rules ()
    ((_ condition then else) (let ((then1 (delay then)) (else1 (delay else)))
                             (force (or (and condition
                                             then1)
                                        else1))))))
;Tests
(my-if (> 1 1) 1 -1)
(my-if (= 1 1) (display "equal\n") (display "nono\n"))
(my-if (< 1 2) (+ 2 3) (/ 4 2))
(my-if #t #f (/ 1 0))
(newline)
;№3

(define recieve '())
(define (trib n)
  (let ((mem (assoc n recieve)))
    (if (not (equal? mem #f))
        (cadr mem)
        (let ((new (if (<= n 1)
                       0
                       (if (= n 2)
                           1
                           (+ (trib (- n 3))
                              (trib (- n 2))
                              (trib (- n 1)))))))
          (set! recieve (cons (list n new) recieve))
          new))))

(define (bad-trib n) ;для примера времени работы;
  (cond ((<= n 1) 0)
        ((= n 2) 1)
        (else (+ (bad-trib (- n 3)) (bad-trib (- n 2)) (bad-trib (- n 1))))))
                                               
;Testsа
;(trib 1000)
;(trib 10)
;(trib 1)
;(trib 3)
(newline)
;№6

;A
(define-syntax when
  (syntax-rules ()
    ((_ condition . actions) (if condition
                                 (begin . actions)))))
(define-syntax unless
  (syntax-rules ()
    ((_ condition . actions) (if (not condition)
                                 (begin . actions)))))
;Tests (A)
;(define x 1)
;(when   (> x 0) (display "x > 0")  (newline))
;(unless (= x 0) (display "x != 0") (newline))

;B
(define-syntax for
  (syntax-rules (in as)
    ((_ x in xs . actions)
     (for-each (lambda (x) (begin . actions)) xs))
    ((_ xs as x . actions)
     (for-each (lambda (x) (begin . actions)) xs))))

;Tests (B)
;(for i in '(1 2 3)
; (for j in '(4 5 6)
;  (display (list i j))
; (newline)))
(newline)
;(for '(1 2 3) as i
; (for '(4 5 6) as j
;  (display (list i j))
; (newline)))
(newline)
;(for '(1 2 3 4) as i
; (for j in '(5 5 5)
;  (display (list i j))
; (newline)))
(newline)

;C

(define-syntax while
  (syntax-rules ()
    ((_ condition . actions) (let loop ()
                               (if condition
                                   (begin (begin . actions)
                                          (loop)))))))
;Tests (C)
(let ((p 0)
      (q 0))
  (while (< p 3)
         (set! q 0)
         (while (< q 3)
                (display (list p q))
                (newline)
                (set! q (+ q 1)))
         (set! p (+ p 1))))
(newline)

;D

(define-syntax repeat
  (syntax-rules (until)
    ((_ (exp . actions) until condition)
     (let exit ()
       (begin exp . actions)
       (if (not condition)
           (exit))))))
;TESTS (D)
(let ((i 0)
      (j 0))
  (repeat ((set! j 0)
           (repeat ((display (list i j))
                    (set! j (+ j 1)))
                   until (= j 3))
           (set! i (+ i 1))
           (newline))
          until (= i 3)))
(newline)
;E

(define-syntax cout
  (syntax-rules (<< endl)
    ((_ ) (begin)) 
    ((_ << endl . expr) (begin
                          (newline)
                          (cout . expr)))
    ((_ << exp . expr) (begin
                         (display exp)
                         (cout . expr)))))
    

;Test cout
(cout << "a = " << 1 << endl << "b = " << 2 << endl)

;№5

(define-syntax my-let
  (syntax-rules ()
    ((_ ((val exp) ...) actions)
     ((lambda (val ...) actions)
      exp ...))))

(define-syntax my-let*
  (syntax-rules ()
    ((_ ((val exp) (val1 exp1) ...) actions)
     (begin
       (define val exp)
       (my-let* ((val1 exp1) ...) actions)))
    ((_ () actions) actions)))

(let ((x 2) (y 3) (z 5))
  (* x y z))
(my-let ((x 2) (y 3) (z 5))
        (* x y z))

(let* ((x 2) (y 3) (z 5))
  (+ x y z))
(my-let* ((x 2) (y 3) (z 5))
         (+ x y z))

(newline)

;№2

(define (save-data x file)
  (with-output-to-file file
    (lambda () (write x (current-output-port))
      (newline (current-output-port)))))

(define (load-data file)
  (with-input-from-file file
    (lambda () (let ((x (read)))
                 (write x)
                 (newline)))))

;(save-data "приветствую смотрящих" "file1.txt")
;(load-data "file1.txt")

(define (read-str input-port)
  (let ((ch (read-char input-port)))
    (cond
      ((eof-object? ch) ch)
      ((eq? ch #\newline) '())
      (else (cons ch (read-str input-port))))))

(define (counter file)
  (let ((input (open-input-file file)))
    (define (count strings)
      (let ((expr (read-str input)))
        (if (eof-object? expr)
            (begin
              (close-input-port input)
              strings)
            (if (not (null? expr))
                (count (+ strings 1))
                (count strings)))))
    (count 0)))

;(save-data '(1 2 3 4) "1234.txt")
;(apply * (load-data "1234.txt"))
(counter "lab4.scm")
               













                         

