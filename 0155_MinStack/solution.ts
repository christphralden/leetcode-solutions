class MinStack {
  instance: number[];
  minIndex: number[];
  constructor() {
    this.instance = [];
    this.minIndex = [];
  }

  push(val: number): void {
    if (
      val < (this.instance[0] ?? 0)
        ? this.minIndex.unshift(val)
        : this.minIndex.push(val)
    )
      this.instance.push(val);
  }

  pop(): void {
    this.instance.pop();
  }

  top(): number {
    return this.instance[this.instance.length - 1];
  }

  getMin(): number {
    return this.instance[this.minIndex[0]];
  }
}

/**
 * Your MinStack object will be instantiated and called as such:
 * var obj = new MinStack()
 * obj.push(val)
 * obj.pop()
 * var param_3 = obj.top()
 * var param_4 = obj.getMin()
 */
