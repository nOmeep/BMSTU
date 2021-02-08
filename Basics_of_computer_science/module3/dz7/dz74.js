'use strict'

const fs = require('fs');

let dict = fs.readFileSync(process.argv[2]).toString().split('\n'); // словарь 
let data = fs.readFileSync(process.argv[3]).toString(); // текст с ошибками 


const parse = (data) => {
    let res = [];
    let strings = String(data).split('\n');
    for (let i = 0; i < strings.length; i++){
        let str = strings[i];
        let words = str.replace(/[^a-z 0-9]/gmi, '').split(' ');
        for (let j = 0; j < words.length; j++){
            res.push([words[j], str.indexOf(words[j]) + 1, i + 1]);
        }
    }
    return res;
};

let tokens = parse(data);
tokens.forEach(element => {
    if (dict.indexOf(element[0]) < 0) {
        console.log(`${element[2]}, ${element[1]} ${element[0]}`)
    }
});