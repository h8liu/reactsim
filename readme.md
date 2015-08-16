# REACToR Simulator

This is the network simulator used in the 
[REACToR paper](https://www.usenix.org/system/files/conference/nsdi14/nsdi14-paper-liu_he.pdf)
in NSDI 2014 for simulating REACToR and pure circuit switches (like Mordia).

To install:

```
$ go get -u github.com/h8liu/reactsim/nsdisim
```

To reproduce the simulation result:

```
$ nsdisim -fig10=fig10.dat -fig11=fig11.dat
```

The `nsdisim` will save the results in `fig10.dat` and `fig11.dat`.

A copy of the results are also saved in 
[`nsdisim/result`](https://github.com/h8liu/reactsim/tree/master/nsdisim/result).

Google Spreadsheets with 
[the figures](https://docs.google.com/spreadsheets/d/1BQnT4f96MYPwhZupeTeVwgg0u4QOmjaK8jnaLnGTcEI/edit?usp=sharing).
