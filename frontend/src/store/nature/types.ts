export interface NatureState {
  natures: Nature[];
  current: Nature;
}

export class Nature {
  readonly name: string;
  readonly description: string;

  constructor(name: string, description: string) {
    this.name = name;
    this.description = description;
  }

  static default(): Nature {
    return new Nature('', '');
  }
}
