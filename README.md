# An English -> Japanese language API 

- Currently uses jmdict but will eventually have two way support, more dictionaries, pitch accent information, and Japanese definitions for Japanese words

## Performance
- Currently this API is quite performant however there may be improvements in the future
- Using a Trie data structure in conjunction with a variation of Levenshteins edit distance algorithm, almost 300,000 entries are compared at lightning speeds (nanoseconds, microseconds!)