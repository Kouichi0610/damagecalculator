export interface NatureState {
  target: string;
  natures: Nature[];
  current: Nature;
}

export class Nature {
  readonly name: string;
  readonly description: string;

  enable(): boolean {
    return this.name.length > 0;
  }

  constructor(name: string, description: string) {
    this.name = name;
    this.description = description;
  }

  static default(): Nature {
    return new Nature('', '');
  }
}
