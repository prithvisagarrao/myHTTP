# myHTTP
###The coding test repository for Adjust

- The tool prints md5sum of the response of an HTTP request of a URL. It takes the entire response object and calculates its md5sum value
- The list of URLs are passed as command-line arguments to the tool.
- There are two flags defined in the tool
  - limit:- This flag is to limit concurrency by limiting the number of threads running at a given time.
    - Usage:- -limit=7
    - Default value=10
  - timeout:- This is an extra flag provided to set the timeout value for http requests in seconds.
    - Usage:- timeout=5
    - Default value=10
  - The logs of the tool are maintained in myHTTP.log file.
    
#### Example
```
go build -o myHTTP main.go
./myHTTP -limit=7 www.adjust.com http://google.com http://facebook.com http://yahoo.com http://yandex.com http://twitter.com http://reddit.com/r/funny http://reddit.com/r/notfunny http://baroquemusiclibrary.com http://wikipedia.com http://nokia.com http://hotstar.com 

http://hotstar.com f8a60819ad0cc5407215139cce1cb4b3
http://google.com a1288f669010bd33c580da80a2fcedbc
http://reddit.com/r/notfunny 7564d1b6c4b125c8d73764e24a3e778e
http://reddit.com/r/funny 84b669b93e371299e36299c83b63a712
http://facebook.com 6c307ca1fd0b573a1b51a53057bc1ce4
http://baroquemusiclibrary.com 2816634ad990be18fac5753b7257e769
http://wikipedia.com ebee9eb66f01c8747ace1f3a54c40b75
http://www.adjust.com 08a9fa50149bd1906080479ffd21c776
http://nokia.com 099107f604269b3e5c2ef0b8efdb2ee1
http://yandex.com d5e0c9119c9df058f42453a20bf10281
http://yahoo.com 0463a37121c48ff842645002be2f5f53
http://twitter.com bca420925c0cdb262a76482c8a0de87c
```
