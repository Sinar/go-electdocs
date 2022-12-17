# Processing Data

For PRU-15; it is slightly more trickier.

Run the command (after adding command below to the raw data)

```javascript
// console.log(JSON.stringify(dunJSON));
console.log(JSON.stringify(parlimenJSON));
// console.log(JSON.stringify(parlimenProcessed));
// console.log(JSON.stringify(statesJSON));
// console.log(JSON.stringify(partiJSON));


```

Extract out to buffer

```shell
$ node ./0-data.js | pbcopy
```

Paste in ; and download the results