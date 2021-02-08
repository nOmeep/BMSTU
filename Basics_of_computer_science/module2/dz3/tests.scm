(define (run-test the-test)
  (let ((expr (car the-test)))
    (write expr)
    (let* ((result (eval expr (interaction-environment)))
           (status (equal? (cadr the-test) result)))
      (if status
          (display " ok")
          (display " FAIL"))
      (newline)
      (if (not status)
          (begin (display " Expected: ") (write (cadr the-test)) (newline)
                 (display " Returned: ") (write result) (newline)))
      status)))

(define (run-tests the-tests)
  (define (really x xs)
    (if (null? xs)
        x
        (really (and x (car xs)) (cdr xs))))
  (really #t (map run-test the-tests)))

(define-syntax test
  (syntax-rules ()
    ((_ expr expected-result) (list (quote expr) expected-result))))
