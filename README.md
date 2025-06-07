# Usage

* Help infomation
  ```
  $ ./test-coredump --help
  Usage of ./test-coredump:
  -delay int
          delay[seconds] before triggering the panic (default 10)
  -size int
          size[MB] of the heap memory (default 16)
  ```

* Generate a 1G core file (wait 3 seconds).
  ```
  $ GOTRACEBACK=crash ./test-coredump --size 1024 --delay 3
  ```

* Look look, then export it.
  ```
  $ sudo coredumpctl 
  $ sudo coredumpctl dump ${pid} -o test-coredump.core
  ```
