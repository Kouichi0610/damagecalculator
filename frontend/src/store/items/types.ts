export interface ItemState {
  items: Item[]
  current: Item;
}

export class Item {
  readonly name: string;
  readonly description: string;

  constructor(name: string, description: string) {
    this.name = name;
    this.description = description;
  }

  enable(): boolean {
    return this.name.length > 0;
  }

  static default(): Item {
    return new Item('', '');
  }
}
