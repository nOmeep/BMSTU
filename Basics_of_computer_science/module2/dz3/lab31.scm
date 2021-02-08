(define-syntax trace-ex
  (syntax-rules ()
    ((_ arg)
     (begin
       (display 'arg)
       (display " => ")
       (let ((result arg))
       (display result)
       (newline)
       result)))))


