

## Test Package

使用方式:

```
go test [build/test flags] [packages] [build/test flags & test binary flags]
```

### build/test flags


```
-args
    Pass the remainder of the command line (everything after -args)
    to the test binary, uninterpreted and unchanged.
    Because this flag consumes the remainder of the command line,
    the package list (if present) must appear before this flag.

-c
    Compile the test binary to pkg.test but do not run it
    (where pkg is the last element of the package's import path).
    The file name can be changed with the -o flag.

-exec xprog
    Run the test binary using xprog. The behavior is the same as
    in 'go run'. See 'go help run' for details.

-i
    Install packages that are dependencies of the test.
    Do not run the test.

-json
    Convert test output to JSON suitable for automated processing.
    See 'go doc test2json' for the encoding details.

-o file
    Compile the test binary to the named file.
    The test still runs (unless -c or -i is specified).
```

## testing flags

```
-bench regexp
    Run only those benchmarks matching a regular expression.
    By default, no benchmarks are run.
    To run all benchmarks, use '-bench .' or '-bench=.'.
    The regular expression is split by unbracketed slash (/)
    characters into a sequence of regular expressions, and each
    part of a benchmark's identifier must match the corresponding
    element in the sequence, if any. Possible parents of matches
    are run with b.N=1 to identify sub-benchmarks. For example,
    given -bench=X/Y, top-level benchmarks matching X are run
    with b.N=1 to find any sub-benchmarks matching Y, which are
    then run in full.

-benchtime t
    Run enough iterations of each benchmark to take t, specified
    as a time.Duration (for example, -benchtime 1h30s).
    The default is 1 second (1s).
    The special syntax Nx means to run the benchmark N times
    (for example, -benchtime 100x).

-count n
    Run each test and benchmark n times (default 1).
    If -cpu is set, run n times for each GOMAXPROCS value.
    Examples are always run once.

-cover
    Enable coverage analysis.
    Note that because coverage works by annotating the source
    code before compilation, compilation and test failures with
    coverage enabled may report line numbers that don't correspond
    to the original sources.

-covermode set,count,atomic
    Set the mode for coverage analysis for the package[s]
    being tested. The default is "set" unless -race is enabled,
    in which case it is "atomic".
    The values:
	set: bool: does this statement run?
	count: int: how many times does this statement run?
	atomic: int: count, but correct in multithreaded tests;
		significantly more expensive.
    Sets -cover.

-coverpkg pattern1,pattern2,pattern3
    Apply coverage analysis in each test to packages matching the patterns.
    The default is for each test to analyze only the package being tested.
    See 'go help packages' for a description of package patterns.
    Sets -cover.

-cpu 1,2,4
    Specify a list of GOMAXPROCS values for which the tests or
    benchmarks should be executed. The default is the current value
    of GOMAXPROCS.

-failfast
    Do not start new tests after the first test failure.

-list regexp
    List tests, benchmarks, or examples matching the regular expression.
    No tests, benchmarks or examples will be run. This will only
    list top-level tests. No subtest or subbenchmarks will be shown.

-parallel n
    Allow parallel execution of test functions that call t.Parallel.
    The value of this flag is the maximum number of tests to run
    simultaneously; by default, it is set to the value of GOMAXPROCS.
    Note that -parallel only applies within a single test binary.
    The 'go test' command may run tests for different packages
    in parallel as well, according to the setting of the -p flag
    (see 'go help build').

-run regexp
    Run only those tests and examples matching the regular expression.
    For tests, the regular expression is split by unbracketed slash (/)
    characters into a sequence of regular expressions, and each part
    of a test's identifier must match the corresponding element in
    the sequence, if any. Note that possible parents of matches are
    run too, so that -run=X/Y matches and runs and reports the result
    of all tests matching X, even those without sub-tests matching Y,
    because it must run them to look for those sub-tests.

-short
    Tell long-running tests to shorten their run time.
    It is off by default but set during all.bash so that installing
    the Go tree can run a sanity check but not spend time running
    exhaustive tests.

-timeout d
    If a test binary runs longer than duration d, panic.
    If d is 0, the timeout is disabled.
    The default is 10 minutes (10m).

-v
    Verbose output: log all tests as they are run. Also print all
    text from Log and Logf calls even if the test succeeds.

-vet list
    Configure the invocation of "go vet" during "go test"
    to use the comma-separated list of vet checks.
    If list is empty, "go test" runs "go vet" with a curated list of
    checks believed to be always worth addressing.
    If list is "off", "go test" does not run "go vet" at all.
```


```
-benchmem
    Print memory allocation statistics for benchmarks.

-blockprofile block.out
    Write a goroutine blocking profile to the specified file
    when all tests are complete.
    Writes test binary as -c would.

-blockprofilerate n
    Control the detail provided in goroutine blocking profiles by
    calling runtime.SetBlockProfileRate with n.
    See 'go doc runtime.SetBlockProfileRate'.
    The profiler aims to sample, on average, one blocking event every
    n nanoseconds the program spends blocked. By default,
    if -test.blockprofile is set without this flag, all blocking events
    are recorded, equivalent to -test.blockprofilerate=1.

-coverprofile cover.out
    Write a coverage profile to the file after all tests have passed.
    Sets -cover.

-cpuprofile cpu.out
    Write a CPU profile to the specified file before exiting.
    Writes test binary as -c would.

-memprofile mem.out
    Write an allocation profile to the file after all tests have passed.
    Writes test binary as -c would.

-memprofilerate n
    Enable more precise (and expensive) memory allocation profiles by
    setting runtime.MemProfileRate. See 'go doc runtime.MemProfileRate'.
    To profile all memory allocations, use -test.memprofilerate=1.

-mutexprofile mutex.out
    Write a mutex contention profile to the specified file
    when all tests are complete.
    Writes test binary as -c would.

-mutexprofilefraction n
    Sample 1 in n stack traces of goroutines holding a
    contended mutex.

-outputdir directory
    Place output files from profiling in the specified directory,
    by default the directory in which "go test" is running.

-trace trace.out
    Write an execution trace to the specified file before exiting.
```