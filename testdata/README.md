# Java reference fixture

`java-hq4x.bin.gz` was generated with `Generate.java` and the unmodified
[hqx-java source](https://github.com/Arcnor/hqx-java) at commit
`f15538c55f495b0d6c552145ef239f4c801730ad`.
It contains all 256 binary neighborhood patterns and 512 seeded images.
The latter cover small dimensions, alpha values, and both wrapping options.
Every expected output pixel comes from Java. Each image uses alpha values from the same half of the 0–255 range. Java
treats alpha as signed when comparing pixels, so comparisons across 127/128
are tested separately.

To regenerate it, compile the upstream `hqx` package and `Generate.java`.
Run `java -cp <class-directory> Generate testdata/java-hq4x.bin.gz`.
