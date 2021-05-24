# gbablog

Blog post: https://remyhax.xyz/posts/gba-blog/

TinyGo has a compiler target for ARM7TDMI (GameBoy Advance) but it's documentation is literally filled with "?????" in terms of support.

As my first real Go program, I decided to write a blog about writing Go for a microcontroller I was entirely unfamiliar with.

The result: The blog about building the ROM contains an emulator the runs the ROM and the ROM contains a naviagtable copy of the blog itself. Just click the link above. You'll understand.

tinygo build -o main.gba -target gameboy-advance .\blog.go



