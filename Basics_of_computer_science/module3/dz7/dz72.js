'use strict'
const fs = require("fs");
const readline = require('readline');
//const { argv } = require("process");
//const { finished } = require("stream");


let n = false;
let m = false;
let i = false;
var count_m
var data = fs.readFileSync(process.argv[process.argv.length - 1]).toString().split('\n');
var string_to_find = process.argv[process.argv.length - 2];

//try {//Проверка на ключи  
for (let j = 2; j<process.argv.length; j++) {
    if (process.argv[j] == '-n')
        n = true;
    else if (process.argv[j] == '-i')
        i = true;
    else if (process.argv[j] == '-m'){
        m = true;
        count_m = process.argv[j+1];
    }
}
//}
//catch {
//    console.log("Error");
//}
//Проверка на ключи закончилась

if (i) {
    for (let i = 0; i<data.length;i++) {
        data[i] = data[i].toLowerCase();
    }
    string_to_find = string_to_find.toLowerCase();
}
if (m) {
    for (let i = 0; i<count_m; i++) {
        if (data[i]!==undefined && data[i].indexOf(string_to_find) > -1) 
            if (n)
                console.log(`${1+i} : ${data[i]}`);
            else
                console.log(data[i]); 
    }
} 
else {
    for (let i = 0; i<data.length; i++) {
        if (data[i]!==undefined && data[i].indexOf(string_to_find) > -1) 
            if (n)
                console.log(`${1+i} : ${data[i]}`);
            else
                console.log(data[i]); 
    }
}