var myArgs = process.argv.slice(2);
var numWords = ["Zero", "One", "Two", "Three", "Four", "Five", "Six", "Seven", "Eight", "Nine"];
var toPrint = "";

for (i = 0; i < myArgs.length-1; i++) {
    for(n = 0; n < myArgs[i].length; n++){
        toPrint += numWords[parseInt(myArgs[i].charAt(n))];
    }
    toPrint += ",";
}

for(n = 0; n < myArgs[myArgs.length-1].length; n++){
    toPrint += numWords[parseInt(myArgs[i].charAt(n))];
}

console.log(toPrint);