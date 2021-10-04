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
console.log(Hash(["2055d229-ce36-495c-96c1-9ecd19585507", "50biz11x012a5jyji3k", now]))
