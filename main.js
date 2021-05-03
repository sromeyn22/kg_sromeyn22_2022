var myArgs = process.argv.slice(2);
// this array of the strings from "Zero" to "Nine" is an easy way to deal with getting the string value of an integer between 0-9
// simply use the integer you want the string of as the index into the array
var numWords = ["Zero", "One", "Two", "Three", "Four", "Five", "Six", "Seven", "Eight", "Nine"];
var toPrint = "";

// loop through command line arguments and get the integer value of each character in each string in myArgs
// Use that integer to append the correct string in the numWords array to the string we will print
for (i = 0; i < myArgs.length-1; i++) {
    for(n = 0; n < myArgs[i].length; n++){
        toPrint += numWords[parseInt(myArgs[i].charAt(n))];
    }
    toPrint += ",";
}
// this extra for loop gets the string value of the last command line argument
// this makes sure the printed string doesn't end in a ','
for(n = 0; n < myArgs[myArgs.length-1].length; n++){
    toPrint += numWords[parseInt(myArgs[i].charAt(n))];
}

console.log(toPrint);