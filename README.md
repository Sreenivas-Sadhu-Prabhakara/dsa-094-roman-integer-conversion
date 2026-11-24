# Roman ↔ Integer Conversion

> Striver SDE Sheet 094 · Day 15 String — parse by the smaller-before-larger rule; render greedily from a 13-entry table that includes the subtractive pairs. Go, stdlib only, tested.

## Problem

Convert Roman numerals to integers and back (the classic range 1–3999). `MCMXCIV` ↔ 1994.

## Approach

**Parsing:** the entire subtractive notation collapses to one local rule — a symbol *smaller than its right neighbor* subtracts, everything else adds. No special-casing IV/IX/XL/…: the rule derives them. **Rendering:** greedy over a value table, largest first — valid only because the table lists the six subtractive pairs (900, 400, 90, 40, 9, 4) as first-class units; with them included, Roman numerals form a canonical system where greedy is exact. The round-trip test over all 3999 values is cheap and catches any asymmetry between the two directions — stronger evidence than a handful of fixtures.

**Complexity:** both directions O(length) ≈ O(1) for the bounded range.

## Run

    go test ./...   # verify
    go run .        # demo

---

One of 191 repos solving Striver's SDE Sheet in Go — stdlib only, one problem per repo, every solution tested. Browse the set: [`dsa-*` repos](https://github.com/Sreenivas-Sadhu-Prabhakara?tab=repositories&q=dsa-&type=source).
