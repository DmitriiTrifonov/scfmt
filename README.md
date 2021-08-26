# scfmt
`scfmt` is a tiny util for C to produce semicolons `;`

## How to use 

get the util by:

```go get github.com/DmitriiTrifonov/scfmt```

build the project and just use it at project root:

```./bin/scfmt example/hello_world.c```

you can add the binary to `$PATH` to use it across the system

### Before
```c
#include "stdio.h"

int main() {
     printf("hello, world!")
     printf("you can test the util with this file")
}
```

### After
```c
#include "stdio.h"

int main() {
     printf("hello, world!");
     printf("you can test the util with this file");
}
```

![👍](https://bloximages.chicago2.vip.townnews.com/celebretainment.com/content/tncms/assets/v3/editorial/0/d7/0d74573f-8eee-582f-9f36-dd76ca27bbdf/5a145d702fd73.image.jpg)

