# Demo 5

Demonstrates the purpose and usage of the tool [gocyclo](https://github.com/fzipp/gocyclo).

* Calculates cyclomatic complexities of functions
* 1 is the base complexity of a function
* +1 for each 'if', 'for', 'case', '&&' or '||'
* Can provide set acceptable boundaries
* Typical example: `if`/`else if`/`else` branches