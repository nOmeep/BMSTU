"use strict"
// modules
const fs = require('fs');
const path = require('path');

//keys
var D = false;
var O = [false, ""];


for (let i = 0; i < process.argv.length; i++){
    if (process.argv[i] == "-d") {
        D = true;
    }
    if (process.argv[i] == "-o") {
        O[0] = true;
        O[1] = process.argv[i + 1];
        if (O[1] == undefined) {
            console.error("Error!");
            return;
        }
        fs.writeFileSync(O[1], "");
    }
}
var tree = (directory, pref, counts) => {
    if (!D){
        let filePaths = fs.readdirSync(directory);
        filePaths.forEach((filePath, index) => {
            if (filePath.charAt(0) == ".") {
                return
            };
            let absolute = path.join(directory, filePath);
            let isFile = fs.lstatSync(absolute).isFile();
            if (isFile) {
                counts["files"]++;
            }
            else {
                counts["directories"]++;
            }
            if (index == filePaths.length - 1) {
                if (!O[0])
                    console.log(`${pref}└── ${filePath}`);
                else
                    fs.appendFileSync(O[1], `\n${pref}└── ${filePath}`)
                if (!(isFile))
                    tree(absolute, `${pref} `,counts);
            } 
            else {
                if (!O[0]) {
                    console.log(`${pref}├── ${filePath}`);
                }
                else {
                    fs.appendFileSync(O[1], `\n${pref}├── ${filePath}`);
                }
                if (!(isFile)) {
                    tree(absolute, `${pref}│ `, counts);
                }
            }
        });
    } 
    else {
        let filePaths = fs.readdirSync(directory);
        filePaths.forEach((filePath, index) => {
            let absolute = path.join(directory, filePath);
            let isDir = fs.lstatSync(absolute).isDirectory();
            if (filePath.charAt(0) == "." || !(isDir)) {
                return;
            }
            counts["directories"]++;
            if (index == filePaths.length - 1) {
                if (!O[0]) {
                    console.log(`${pref}└── ${filePath}`);
                }
                else {
                    fs.appendFileSync(O[1], `\n${pref}└── ${filePath}`);
                }
                tree(absolute, `${pref} `, counts);
            } 
            else {
                if (!O[0]){
                    console.log(`${pref}└── ${filePath}`);
                } 
                else {
                    fs.appendFileSync(O[1], `\n${pref}├── ${filePath}`)
                }
                tree(absolute, `${pref} `, counts);
            }
        });
    }
};

let directory = (process.argv[2] == "-d" || "-o" ? "." : process.argv[2]);

const counts = {
    directories: 0,
    files: 0
};

if (!O[0]) {
    console.log(directory);
} 
else {
    fs.appendFileSync(O[1], directory);
}
tree(directory, "", counts);

if (!D) {
    if (!O[0]) {
        console.log(`\n ${counts.directories} directories, ${counts.files} files`);
    }
    else {
        fs.appendFileSync(O[1], `\n ${counts.directories} directories, ${counts.files} files`);
    }
}
else {
    if (!O[0])
        console.log(`\n ${counts.directories} directories`);
    else 
        fs.appendFileSync (O[1], `\n ${counts.directories} directories`);
}