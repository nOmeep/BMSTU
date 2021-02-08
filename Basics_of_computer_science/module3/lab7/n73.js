
'use strict'

const alphabet = "qwertyuiop[]asdfghjkl;'zxcvbnm,./`1234567890-=QWERTYUIOP{}|ASDFGHJKL:ZXCVBNM<>?!@#$%^&*()"

function main(nel, stringsLength) {
    for (var i = 0; i < nel; i++) {
        var newString = '';
        for (var j = 0; j < stringsLength; j++) {
            var i1;
            i1 = Math.floor(Math.random() * alphabet.length);
            newString += alphabet[i1];
        }
        console.log(newString);
    }
}

main(process.argv[2], process.argv[3]);