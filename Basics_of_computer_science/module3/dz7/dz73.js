'use strict'

const fs = require("fs");
const readline = require("readline");

const rl = readline.createInterface( {
    input: process.stdin,
    output: process.stdout
})

const word_count = (string) => {
    return string.trim().split(/\s+/).length;
};

const lines_count = (string) => {
    return string.split('\n').length;
}

const mainstring = fs.readFileSync(process.argv[process.argv.length-1], "utf8");
let c = false;
let m = false; 
let w = false;
let l = false;

if (process.argv.length == 2){
    console.error("where is file?!");
    return;
} else {
    if (process.argv.length == 3) {
        let stats = fs.statSync(process.argv[process.argv.length-1]);
        //console.log(stats);
        console.log(stats["size"]);
        console.log(mainstring.length);
        console.log(word_count(mainstring));
        console.log(lines_count(mainstring));
    } else {
        for (let i = 0; i<process.argv.length; i++){
            if (process.argv[i] =="-c")
                c = true;
            if (process.argv[i] == "-m")
                m = true;
            if (process.argv[i] =="-w")
                w =true;
            if (process.argv[i] == "-l")
                l = true;
        }
        if (c) {
            let stats = fs.statSync(process.argv[process.argv.length-1]);
            console.log(stats["size"]);
        }
        if (m) {
            console.log(mainstring.length);
        }
        if (w) {
            console.log(word_count(mainstring));
        }
        if (l) {
            console.log(lines_count(mainstring));
        }
    }
}

// Примеры запуска программы: 
//  * node dz73.js -l wcount.txt
//  * node dz73.js -w wcount.txt
//  * node dz73.js -c wcount.txt 
//  * node dz73.js -m wcount.txt