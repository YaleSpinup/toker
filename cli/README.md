# cli

---
The files in this package are exclusively cobra cli command hooks and are the entry points to the various commands to be
executed in the command line.

## Files List & Descriptions
Below are the current cobra command files used in the application and their descriptions
```
compare.go         Compare a password (UUID) with a bcrypt hash
hash.go            Encrypts the password using bcrypt
new.go             Generate a list of tokens and optionally their encrypted values
root.go            Responsible for setting up the viper configuration file for toker
version.go         Display toker version information
```