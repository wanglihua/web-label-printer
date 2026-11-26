function distinctConcat(a, b) {
    let arr = a.concat(b)
    let result = []
    let obj = {}

    for (let i of arr) {
        if (!obj[i.id]) {
            result.push(i)
            obj[i.id] = 1
        }
    }

    return result
}

export default distinctConcat;
