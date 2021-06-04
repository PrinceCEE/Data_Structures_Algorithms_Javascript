class Stack {
  constructor() {
    this._top = 0;
    this._dataStore = [];
  }

  get top() {
    return this._top;
  }

  toString() {
    return this._dataStore.slice(0);
  }

  get length() {
    return this._dataStore.length;
  }

  push(element) {
    this._dataStore[this._top++] = element;
  }
  
  pop() {
    this._dataStore.pop();
    --this._top;
  }

  clear() {
    this._top = 0;
    delete this._dataStore;
    this._dataStore = [];
  }

  peek() {
    return this._dataStore[this._top - 1];
  }
}