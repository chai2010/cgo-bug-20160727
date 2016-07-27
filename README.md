# cgo-bug-20160727

```
C:\go\gopkg\src\github.com\chai2010\cgo-bug-20160727>go version
go version go1.7rc3 windows/amd64

C:\go\gopkg\src\github.com\chai2010\cgo-bug-20160727>go build ./a

C:\go\gopkg\src\github.com\chai2010\cgo-bug-20160727>go build ./b

C:\go\gopkg\src\github.com\chai2010\cgo-bug-20160727>go build
# github.com/chai2010/cgo-bug-20160727
C:\go\go1.7rc3.windows-amd64\pkg\tool\windows_amd64\link.exe: running g++ failed: exit status 1
C:\Users\chai\AppData\Local\Temp\go-link-802100695\000001.o: In function `mingw_onexit':
C:/crossdev/src/mingw-w64-v4-git/mingw-w64-crt/crt/atonexit.c:34: multiple definition of `mingw_onexit'
C:\Users\chai\AppData\Local\Temp\go-link-802100695\000000.o:C:/crossdev/src/mingw-w64-v4-git/mingw-w64-crt/crt/atonexit.c:34: first defined here
C:\Users\chai\AppData\Local\Temp\go-link-802100695\000001.o: In function `atexit':
C:/crossdev/src/mingw-w64-v4-git/mingw-w64-crt/crt/atonexit.c:57: multiple definition of `atexit'
C:\Users\chai\AppData\Local\Temp\go-link-802100695\000000.o:C:/crossdev/src/mingw-w64-v4-git/mingw-w64-crt/crt/atonexit.c:57: first defined here
C:\Users\chai\AppData\Local\Temp\go-link-802100695\000001.o: In function `_decode_pointer':
C:/crossdev/src/mingw-w64-v4-git/mingw-w64-crt/crt/mingw_helpers.c:20: multiple definition of `_decode_pointer'
C:\Users\chai\AppData\Local\Temp\go-link-802100695\000000.o:C:/crossdev/src/mingw-w64-v4-git/mingw-w64-crt/crt/mingw_helpers.c:20: first defined here
C:\Users\chai\AppData\Local\Temp\go-link-802100695\000001.o: In function `_encode_pointer':
C:/crossdev/src/mingw-w64-v4-git/mingw-w64-crt/crt/mingw_helpers.c:26: multiple definition of `_encode_pointer'
C:\Users\chai\AppData\Local\Temp\go-link-802100695\000000.o:C:/crossdev/src/mingw-w64-v4-git/mingw-w64-crt/crt/mingw_helpers.c:26: first defined here
C:\Users\chai\AppData\Local\Temp\go-link-802100695\000001.o: In function `__gnu_cxx::new_allocator<char>::deallocate(char*, unsigned long long)':
C:/TDM-GCC-64/lib/gcc/x86_64-w64-mingw32/5.1.0/include/c++/ext/new_allocator.h:110: multiple definition of `mingw_app_type'
C:\Users\chai\AppData\Local\Temp\go-link-802100695\000000.o:C:/TDM-GCC-64/lib/gcc/x86_64-w64-mingw32/5.1.0/include/c++/ext/new_allocator.h:110: first defined here
collect2.exe: error: ld returned 1 exit status
```
