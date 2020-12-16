class Counter {
    constructor (initialValue) {
        this.__value = initialValue || 0
    }
    increment () {
        this.value++
        return this
    }
    decrement () {
        this.value--
        return this
    }
    reset () {
        this.value = 0
        return this
    }
    step (increment) {
        this.value += increment
        return this
    }
    set value (val) {
        // this is a very long comment so we only expect the first 80 characters and nothing else!
        if (typeof val !== "number") throw new Error("Value must be a number")
        this.__value = val
    }
    get value () {
        return this.__value
    }
}

// create new counter
const counter = new Counter()

// run counter commands
counter
    .decrement()
    .reset()
    .increment()
    .step("four")