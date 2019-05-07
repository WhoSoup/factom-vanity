# Factom Vanity Address Generator

This is a simple FCT Address generator that looks for specified prefixes in the human readable address. It is intended more as a proof of concept and didactic tool for Factomize's blog on Addresses.

While running, it outputs a list of private keys and addresses to the console where the address contains the prefix.

## Compiling

You'll need dependencies `FactomProject/btcutil/base58` and `FactomProject/ed25519` to compile. Otherwise there are no requirements and you can just `go build` or `go install`.

## Usage

Run with `factom-vanity [-threads 2] [-sleep 50] -file <input file>`.

* threads: The number of goroutines to start. More goroutines means more threads will be started and more cpu will be used.
* sleep: The amount of time in microseconds that each thread sleeps between generating the next address pair. The default value is suitable to run in the background for longer periods of time without tying up a significant amount of cpu resources. Lower it to calculate more hashes per second.
* file: The input file is a plain text file with every index you are looking for on a new line. The prefixes are case insensitive. All lowercase characters are represented in Base58 but not all uppercase.

Stop the application with `Ctrl+C`.

I recommend piping all output to a file so that you can review it later. 

Example of an input text file:
```
test
factom
vanity
```

## What to expect

It's fairly easy to generate address with short prefixes of 4 or fewer characters and you can expect those in the order of minutes with the default `sleep` setting. Five characters take in the order of several hours and six characters in the order of days. If you want a very specific spelling (for example all lower caps, or upper case first character) it will likely take a lot longer unless you get lucky.

I am also not aware whether or not all letter combinations are possible. They may not be, depending on the SHA256 algorithm, which has not been proven to be surjective. As such, I recommend quantity over quality when it comes to prefixes.

For example, the address `FA2VaNiTy3WwHhDkA7XjA1KH6fpWcUXgsiUTHA72EK9pD9mkY4JZ` took 109 hours. 

## Security

All addresses are generated with a random number generator that uses golang's cryptographically secure generator (the same as inside the factom protocol). If you want to try this yourself, I recommend strongly against any sort of sequential generation.
