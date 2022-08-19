const crypto = require("crypto");

function Hash(keys) {
    if(!keys) {
        keys = ""
    }
    if (typeof keys == "object" && Array.isArray(keys)){
        keys = keys.join("")
    }
    const r = crypto.createHash('sha256').update(keys).digest('hex');
    return r
}
let now = Number(Date.now() / 1000).toFixed()
console.log(now)
// console.log(Hash(["2055d229-ce36-495c-96c1-9ecd19585507", "50biz11x012a5jyji3k", now]))
console.log(Hash(["de935b02-8f34-46b9-9976-09b176c2ec00", "dz2mgrk5z2ym", now]))
