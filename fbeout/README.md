# Fast Binary Encoding
So when you use the CLI tool to create the stuffz, it generates the standard FBE library under some directory called "fbe".
Then it'll create a directory with whatever you declared in your proto file. 
If you call it `package proto` it'll make a directory called "proto". 
Once that's done, you'll have to replace all of the relative paths that are in the "proto" or whatever you package
directory is called.  Go doesn't really like that too much and it appears the go implimentation was made years ago.

Then you can import it and use it as the instructions say
