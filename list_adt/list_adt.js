class InvalidTypeError extends Error {
  constructor(message) {
    this.name = InvalidTypeError.name;
    super(message);
  }
}

export default class List {
  _pos = 0;
  _dataStore = [];
  
  constructor(...args) {
    let typeOfArgs = typeof args[0];
    for(let element of args) {
      if(typeof element !== typeOfArgs) {
        throw new InvalidTypeError("The types of the elements do not match");
      }
    }
  }

  get currPos() {
    return this._pos;
  }

  get length() {
    return this._dataStore.length;
  }

  clear() {
    delete this._dataStore;
    this._dataStore = [];
    return true;
  }

  toString() {
    return this._dataStore;
  }

  getElement(element) {
    let index = this._dataStore.find(element);
    return index != -1 && this._dataStore[index];
  }

  _elementValid(element) {
    if(typeof element !== typeof this._dataStore[0]) {
      throw new InvalidTypeError("The type of the element does not match the types allowable");
    }
  }

  insert(element, after) {
    let index = this._dataStore.find(after);
    if(index !== -1) {
      this._elementValid(element);
      this._dataStore.splice(index, 0, element);
      return true;
    }
    return false;
  }

  append(element) {
    this._elementValid(element);
    this._dataStore.push(element);
  }

  remove(element) {
    let index = this._dataStore.find(element);
    return index !== -1 && this._dataStore.splice(index, 1);
  }

  front() {
    this._pos = 0;
  }

  end() {
    this._pos = this.length - 1;
  }

  prev() {
    --this._pos;
  }

  next() {
    ++this._pos;
  }

  moveTo(newPos) {
    const acceptRange = newPos > -1 && newPos <= this.length - 1;
    if(acceptRange) {
      this._pos = newPos;
    }
    return acceptRange;
  }
}