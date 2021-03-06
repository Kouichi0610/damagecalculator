export interface WeatherFieldsState {
  target: string;
  weathers: Weather[]; // 天候一覧
  fields: Field[]; // フィールド一覧
  currentWeather :Weather;
  currentField: Field;
}

export class Weather {
  readonly id: number;
  readonly name: string;
  readonly description: string;
  constructor(id: number, name: string, description: string) {
    this.id = id;
    this.name = name;
    this.description = description;
  }
  enable(): boolean {
    return this.name.length > 0;
  }
  static default(): Weather {
    return new Weather(0, '', '');
  }
}
export class Field {
  readonly id: number;
  readonly name: string;
  readonly description: string;
  constructor(id: number, name: string, description: string) {
    this.id = id;
    this.name = name;
    this.description = description;
  }
  enable(): boolean {
    return this.name.length > 0;
  }
  static default(): Field {
    return new Field(0, '', '');
  }
}
export function toFields(data: any): Field[] {
  let res: Field[] = [];
  for (var i = 0; i < data.fields.length; i++) {
    let d = data.fields[i];
    res.push(new Field(i, d.name, d.description));
  }
  return res;
}
export function toWeathers(data: any): Weather[] {
  let res: Weather[] = [];
  for (var i = 0; i < data.weathers.length; i++) {
    let d = data.weathers[i];
    res.push(new Weather(i, d.name, d.description));
  }
  return res;
}
