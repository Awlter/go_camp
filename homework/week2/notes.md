Error over exception
- errors.New returns a pointer to avoid equal check when error content (text) are the same, not when two error are the same
- panic is not the same as exception since it's a fatal error, you should not expect anyone to rescue a panic, but recovery can give up this process
  - 越界
- use Error can ensure what need to rollback for the logic that has already processed at any specific point

Error Type
- Errors are value
- Sentinel Error (predefined error)
  -  not good because
    - when you want to wrap more info to the Error, the place that do the error checking will definitely fail, except if you do a pattern matching
    - you have to publiclize / expose the error, which can make a framework expose more internal logic and also make two packages couple too much
- Error types
  - customized error struct, but it shares the same problem with Sentinel Error, which can make a framework fragile
- Opaque Errors
  - you only know the error result but not the internal info
  - assert errors for behavior not type
  - it's good because
    - it checks the type and behavior?? (couldn't grasp here well)

Error Handling
- use struct to save the error as a value
- only handle error once
- record a log when debugging and when something fails
- Wrap errors